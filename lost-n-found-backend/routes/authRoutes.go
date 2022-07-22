package routes

func (routes *Routes) authRoutes() {
	routes.Gin.POST("/user", routes.Controllers.RegisterUser)
	routes.Gin.POST("/login", routes.Controllers.LoginUser)
}
