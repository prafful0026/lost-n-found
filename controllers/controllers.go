package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prafful0026/lost-n-found/responses"
	"go.mongodb.org/mongo-driver/mongo"
)

type Controllers struct {
	DB  *mongo.Client
	Gin *gin.Engine
}

func NewControllers(router *gin.Engine, client *mongo.Client) (*Controllers, error) {
	return &Controllers{
		Gin: router,
		DB:  client,
	}, nil
}

func throwInternalServerError(err error, ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusInternalServerError, responses.HttpResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
}

func throwError(httpStatus int, ginCtx *gin.Context, err error, errorMessage string) {

	if errorMessage == "" {
		errorMessage = err.Error()
	}
	ginCtx.JSON(httpStatus, responses.HttpResponse{Status: httpStatus, Message: errorMessage, Data: map[string]interface{}{}})

}
