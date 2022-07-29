package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prafful0026/lost-n-found/configs"
	"github.com/prafful0026/lost-n-found/routes"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	client := configs.ConnectDB()
	defer configs.DisconnectDB(client)

	newRoutes, _ := routes.NewRoutes(router, client)
	newRoutes.RegisterRoutes()

	router.Run(":3000")

}
