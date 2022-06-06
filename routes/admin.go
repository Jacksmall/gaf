package routes

func AdminRoute() {
	adminAPI := router.Group("admin")

	// 获取商品列表数据
	{
		adminAPI.GET("/api/v1/getAll", adminProductController.GetProducts)
		adminAPI.GET("/api/v1/getOne/:id", adminProductController.GetProduct)
		// adminAPI.POST("/api/v1/add", product.CreateProduct)
		// adminAPI.POST("/api/v1/update", product.UpdateProduct)
		// adminAPI.DELETE("/api/v1/delete/:id", product.DeleteProduct)
	}
}
