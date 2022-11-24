package main

import (
	"GinMES/database"
	"GinMES/middleware"
	"GinMES/models"
	"GinMES/routes"
	"GinMES/views"
)

func main() {
	route := routes.InitRouter()
	// 迁移到数据库
	database.Db.AutoMigrate(&models.ProductInfo{}, &models.ProductUnitInfo{},
		&models.LineInfo{}, &models.WorkProcessInfo{},
		&models.WorkCraftInfo{}, &models.WorkStationInfo{}, &models.Users{})
	// 产品
	//route.GET()
	apiV1 := route.Group("/api")
	apiV1.POST("/login", views.Login)
	apiV1.Use(middleware.JWTAuth())
	apiV1.POST("/product/create", views.ProductCreate)
	apiV1.GET("/product/list", views.GetProduct)
	apiV1.POST("/product/update", views.ProductUpdate)
	apiV1.POST("/product/delete", views.ProductDelete)
	apiV1.GET("/product/detail/:product_id", views.ProductDetail)
	// 单位
	apiV1.POST("/unit/create", views.UnitCreate)
	apiV1.POST("/unit/update", views.UnitUpdate)
	apiV1.GET("/unit/get", views.UnitGet)
	apiV1.POST("/unit/delete", views.UnitDelete)
	// 工序
	apiV1.POST("/workprocess", views.WorkProcessCreate)
	apiV1.PUT("/workprocess", views.WorkProcessupdate)
	apiV1.GET("/workprocess", views.WorkProcessGet)
	apiV1.DELETE("/workprocess", views.WorkProcessDelete)
	// 工艺
	apiV1.POST("/workcraft", views.WorkCraftCreate)
	apiV1.GET("/workcraft", views.WorkCraftGET)
	apiV1.PUT("/workcraft", views.WorkCraftupdate)
	apiV1.DELETE("/workcraft", views.WorkCraftDelete)
	// 工位

	route.Run("0.0.0.0:5000")
}
