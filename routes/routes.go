package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prafful0026/lost-n-found/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

type Routes struct {
	DB          *mongo.Client
	Gin         *gin.RouterGroup
	Controllers *controllers.Controllers
}

func NewRoutes(router *gin.Engine, client *mongo.Client) (*Routes, error) {
	v1 := router.Group("/api/v1")
	newControllers, _ := controllers.NewControllers(router, client)

	return &Routes{
		Gin:         v1,
		DB:          client,
		Controllers: newControllers,
	}, nil
}

func (routes *Routes) RegisterRoutes() {
	routes.pingRoutes()
	routes.authRoutes()
	routes.postRoutes()
	routes.claimRoutes()
}
