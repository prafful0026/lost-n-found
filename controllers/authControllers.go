package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prafful0026/lost-n-found/configs"
	"github.com/prafful0026/lost-n-found/models"
	"github.com/prafful0026/lost-n-found/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (cnt *Controllers) RegisterUser(ginCtx *gin.Context) {

	var validate = validator.New()
	var userCollection *mongo.Collection = configs.GetCollections(cnt.DB, configs.UsersCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	if err := ginCtx.BindJSON(&user); err != nil {
		throwError(http.StatusBadRequest, ginCtx, err, "")
		return
	}

	if err := validate.Struct(&user); err != nil {
		throwError(http.StatusBadRequest, ginCtx, err, "")
		return
	}
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Err()

	if err == nil {
		throwError(http.StatusConflict, ginCtx, nil, "Email address already registered")
		return
	}
	if err != nil && err != mongo.ErrNoDocuments {
		throwError(http.StatusInternalServerError, ginCtx, err, "")
		return
	}

	hashByte, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		throwError(http.StatusInternalServerError, ginCtx, err, "")
		return
	}

	hash := string(hashByte)

	newUser := models.User{
		Id:        primitive.NewObjectID(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hash,
	}

	result, err := userCollection.InsertOne(ctx, newUser)

	if err != nil {
		throwError(http.StatusInternalServerError, ginCtx, err, "")
		return
	}
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"id":  result.InsertedID,
	// 	"exp": time.Now().Add(time.Hour * 24 * 15).Unix(),
	// })
	// var (
	// 	mySigningKey = []byte(configs.GetEnvVar("JWT_PASSKEY"))
	// )
	// tokenStr, err := token.SignedString(mySigningKey)

	// if err != nil {
	// 	throwError(http.StatusInternalServerError, ginCtx, err, "")
	// 	return
	// }

	ginCtx.JSON(http.StatusCreated, responses.HttpResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}

func (cnt *Controllers) WithAuth() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		tokenStr := ginCtx.GetHeader("authorization")
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		var (
			mySigningKey = []byte(configs.GetEnvVar("JWT_PASSKEY"))
		)
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})

		if err != nil || !token.Valid {
			ginCtx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ginCtx.Set("id", claims["id"])
			ginCtx.Next()
		} else {
			ginCtx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
		}
	}
}

func (cnt *Controllers) LoginUser(ginCtx *gin.Context) {

	var userCollection *mongo.Collection = configs.GetCollections(cnt.DB, configs.UsersCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	type input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var i input

	err := ginCtx.BindJSON(&i)
	if err != nil {
		throwError(http.StatusBadRequest, ginCtx, err, "")
		return
	}

	err = userCollection.FindOne(ctx, bson.M{"email": i.Email}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			throwError(http.StatusBadRequest, ginCtx, nil, "Invalid email address")

		} else {
			throwError(http.StatusInternalServerError, ginCtx, err, "")
		}
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 15).Unix(),
	})

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(i.Password))

	if err != nil {
		throwError(http.StatusUnauthorized, ginCtx, nil, "Invalid email address or password")
		return
	}

	var (
		mySigningKey = []byte(configs.GetEnvVar("JWT_PASSKEY"))
	)
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		throwError(http.StatusInternalServerError, ginCtx, err, "")
		return
	}
	var userRes responses.UserResponse
	userRes.Id = user.Id
	userRes.Email = user.Email
	userRes.FirstName = user.FirstName
	userRes.LastName = user.LastName
	userRes.AuthToken = tokenStr

	ginCtx.JSON(http.StatusOK, responses.HttpResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": userRes}})

}
