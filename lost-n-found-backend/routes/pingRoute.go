package routes

func (routes *Routes) pingRoutes() {
	routes.Gin.GET("/ping", routes.Controllers.HandlePing)
}
