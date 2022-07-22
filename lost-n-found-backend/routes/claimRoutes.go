package routes

func (routes *Routes) claimRoutes() {
	routes.Gin.POST("/claim", routes.Controllers.WithAuth(), routes.Controllers.CreateClaim)
	routes.Gin.GET("/claim/:postId/:claimId", routes.Controllers.WithAuth(), routes.Controllers.GetClaims)
	routes.Gin.GET("/claim/:postId", routes.Controllers.WithAuth(), routes.Controllers.GetClaims)
}
