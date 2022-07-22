package controllers

import (
	"context"
	"log"
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

func (cnt *Controllers) CreateClaim(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var validate = validator.New()
	var claimCollection *mongo.Collection = configs.GetCollections(cnt.DB, "claims")
	var postCollection *mongo.Collection = configs.GetCollections(cnt.DB, "posts")
	userId := (ginCtx.MustGet("id").(string))

	var claim models.Claim
	claim.Status = "PENDING"
	if err := ginCtx.BindJSON(&claim); err != nil {
		ginCtx.JSON(http.StatusBadRequest, responses.HttpResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&claim); validationErr != nil {
		ginCtx.JSON(http.StatusBadRequest, responses.HttpResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	uID, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}

	newClaim := models.Claim{
		Title:       claim.Title,
		Address:     claim.Address,
		Description: claim.Description,
		User:        uID,
		Post:        claim.Post,
		Email:       claim.Email,
		PhoneNumber: claim.PhoneNumber,
		ImageUrls:   claim.ImageUrls,
		Status:      claim.Status,
	}

	newClaim.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	var post models.Post
	err = postCollection.FindOne(ctx, bson.M{"_id": claim.Post}).Decode(&post)
	if err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}
	if post.User == uID {
		ginCtx.JSON(http.StatusBadRequest, responses.HttpResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "you cant claim on your own post"}})
		return
	}

	err = claimCollection.FindOne(ctx, bson.M{"post": claim.Post, "user": uID}).Err()
	if err != nil && err != mongo.ErrNoDocuments {
		throwInternalServerError(err, ginCtx)
		return
	}
	if err == nil {
		ginCtx.JSON(http.StatusBadRequest, responses.HttpResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "you cant claim twice on one post"}})
		return
	}

	result, err := claimCollection.InsertOne(ctx, newClaim)

	if err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}

	match := bson.M{"_id": claim.Post}
	change := bson.M{"$push": bson.M{"claims": uID}}

	updateResult, err := postCollection.UpdateOne(ctx, match, change)
	if err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}
	log.Println(updateResult)
	ginCtx.JSON(http.StatusCreated, responses.HttpResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result.InsertedID}})

}

func (cnt *Controllers) GetClaims(ginCtx *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// var validate = validator.New()
	var claimCollection *mongo.Collection = configs.GetCollections(cnt.DB, "claims")
	var postCollection *mongo.Collection = configs.GetCollections(cnt.DB, "posts")

	userId := (ginCtx.MustGet("id").(string))
	uID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}

	cId := ginCtx.Param("claimId")
	pId := ginCtx.Param("postId")
	postId, err := primitive.ObjectIDFromHex(pId)
	if err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}
	claimId, err := primitive.ObjectIDFromHex(cId)
	if err != nil {
		throwInternalServerError(err, ginCtx)
		return
	}
	// var key string
	// var value interface{}

	if cId == "" {
		// key = "post"
		// value = postId
		var post models.Post
		err = postCollection.FindOne(ctx, bson.M{"_id": postId}).Decode(&post)
		if err != nil {
			throwInternalServerError(err, ginCtx)
			return
		}
		if post.User != uID {
			var claim []responses.ClaimResponse
			matchStage := bson.D{{"$match", bson.D{{"post", postId}, {"user", uID}}}}
			lookupStage := bson.D{{"$lookup", bson.D{{"from", "users"}, {"localField", "user"}, {"foreignField", "_id"}, {"as", "userDetails"}}}}
			unwindStage := bson.D{{"$unwind", bson.D{{"path", "$userDetails"}, {"preserveNullAndEmptyArrays", false}}}}
			cursor, err := claimCollection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage})

			if err != nil {
				throwInternalServerError(err, ginCtx)
				return
			}

			if err = cursor.All(ctx, &claim); err != nil {
				throwInternalServerError(err, ginCtx)
				return
			}
			ginCtx.JSON(http.StatusOK, responses.HttpResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": claim}})
			return
		}

		var claims []responses.ClaimResponse
		matchStage := bson.D{{"$match", bson.D{{"post", postId}}}}
		lookupStage := bson.D{{"$lookup", bson.D{{"from", "users"}, {"localField", "user"}, {"foreignField", "_id"}, {"as", "userDetails"}}}}
		unwindStage := bson.D{{"$unwind", bson.D{{"path", "$userDetails"}, {"preserveNullAndEmptyArrays", false}}}}

		cursor, err := claimCollection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage})

		if err != nil {
			throwInternalServerError(err, ginCtx)
			return
		}

		if err = cursor.All(ctx, &claims); err != nil {
			throwInternalServerError(err, ginCtx)
			return
		}
		ginCtx.JSON(http.StatusOK, responses.HttpResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": claims}})

	} else {
		var post models.Post
		err = postCollection.FindOne(ctx, bson.M{"_id": postId}).Decode(&post)
		if err != nil {
			throwInternalServerError(err, ginCtx)
			return
		}

		var claim models.Claim
		err = claimCollection.FindOne(ctx, bson.M{"_id": claimId}).Decode(&claim)
		if err != nil {
			throwInternalServerError(err, ginCtx)
			return
		}

		// if(claim.Post!=post.Id){

		// 	throwInternalServerError(err, ginCtx)
		// 	return
		// }
		if post.User == uID || claim.User == uID {
			var claim []responses.ClaimResponse
			matchStage := bson.D{{"$match", bson.D{{"_id", claimId}}}}
			lookupStage := bson.D{{"$lookup", bson.D{{"from", "users"}, {"localField", "user"}, {"foreignField", "_id"}, {"as", "userDetails"}}}}
			unwindStage := bson.D{{"$unwind", bson.D{{"path", "$userDetails"}, {"preserveNullAndEmptyArrays", false}}}}
			cursor, err := claimCollection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage})

			if err != nil {
				ginCtx.JSON(http.StatusBadRequest, responses.HttpResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "claimId and postId dont match"}})
				return
			}

			if err = cursor.All(ctx, &claim); err != nil {
				throwInternalServerError(err, ginCtx)
				return
			}
			ginCtx.JSON(http.StatusOK, responses.HttpResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": claim}})
			return
		}
		ginCtx.JSON(http.StatusBadRequest, responses.HttpResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "this is not your post "}})

	}
}
