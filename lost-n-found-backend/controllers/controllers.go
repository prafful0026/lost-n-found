package controllers

import (
	"github.com/gin-gonic/gin"
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
