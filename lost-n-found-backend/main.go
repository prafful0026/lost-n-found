package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prafful0026/lost-n-found/configs"
	"github.com/prafful0026/lost-n-found/routes"
)

func main() {
	router := gin.Default()

	client := configs.ConnectDB()
	defer configs.DisconnectDB(client)

	newRoutes, _ := routes.NewRoutes(router, client)
	routes.RegisterRoutes(newRoutes)

	router.Run(":3000")

}
