package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prafful0026/lost-n-found/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

type Routes struct {
	DB          *mongo.Client
	Gin         *gin.Engine
	Controllers *controllers.Controllers
}

func NewRoutes(router *gin.Engine, client *mongo.Client) (*Routes, error) {

	newControllers, _ := controllers.NewControllers(router, client)

	return &Routes{
		Gin:         router,
		DB:          client,
		Controllers: newControllers,
	}, nil
}

func RegisterRoutes(routes *Routes) {
	routes.pingRoutes()
}
