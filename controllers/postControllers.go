package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prafful0026/lost-n-found/configs"
	"github.com/prafful0026/lost-n-found/models"
	"github.com/prafful0026/lost-n-found/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (cnt *Controllers) CreatePost(ginCtx *gin.Context) {
	var validate = validator.New()
	var postCollection *mongo.Collection = configs.GetCollections(cnt.DB, "posts")
	userId := (ginCtx.MustGet("id").(string))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var post models.Post
	defer cancel()

	if err := ginCtx.BindJSON(&post); err != nil {
		ginCtx.JSON(http.StatusBadRequest, responses.HttpResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&post); validationErr != nil {
		ginCtx.JSON(http.StatusBadRequest, responses.HttpResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	uID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}

	newPost := models.Post{
		Title:       post.Title,
		Address:     post.Address,
		Description: post.Description,
		User:        uID,
		Email:       post.Email,
		PhoneNumber: post.PhoneNumber,
		ImageUrls:   post.ImageUrls,
		Status:      post.Status,
	}
	newPost.Claims = []primitive.ObjectID{}

	newPost.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	result, err := postCollection.InsertOne(ctx, newPost)
	if err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}
	ginCtx.JSON(http.StatusCreated, responses.HttpResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

}
func (cnt *Controllers) GetPosts(ginCtx *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var postCollection *mongo.Collection = configs.GetCollections(cnt.DB, "posts")
	slug := ginCtx.Param("slug")
	var key string
	var value interface{}

	userId := (ginCtx.MustGet("id").(string))
	uID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}

	if slug == "" {
		value = uID
		key = "user"

	} else if slug != "LOST" && slug != "FOUND" {
		postId, err := primitive.ObjectIDFromHex(slug)
		if err != nil {
			throwInternalServerError(err, ginCtx)
			return
		}
		value = postId
		key = "_id"
	} else {
		value = slug
		key = "status"
	}

	matchStage := bson.D{{"$match", bson.D{{key, value}}}}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "users"}, {"localField", "user"}, {"foreignField", "_id"}, {"as", "userDetails"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$userDetails"}, {"preserveNullAndEmptyArrays", false}}}}

	cursor, err := postCollection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage})
	if err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}

	var posts []responses.PostResponse

	if err = cursor.All(ctx, &posts); err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}

	ginCtx.JSON(http.StatusOK, responses.HttpResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": posts}})

}
