package routes

func ApiRoute() {
	apiApi := router.Group("api")

	{
		apiApi.POST("/v1/cart/add", apiCartController.Add)
	}
}
