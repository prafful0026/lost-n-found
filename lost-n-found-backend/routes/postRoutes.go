package routes

func (routes *Routes) postRoutes() {
	routes.Gin.POST("/post", routes.Controllers.WithAuth(), routes.Controllers.CreatePost)
	routes.Gin.GET("/post/:slug", routes.Controllers.WithAuth(), routes.Controllers.GetPosts)
	routes.Gin.GET("/post", routes.Controllers.WithAuth(), routes.Controllers.GetPosts)
}
