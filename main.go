package main

import (
	"GinMES/routes"
	"net/http"
)

func main() {
	route := routes.InitRouter()
	// 迁移到数据库
	//database.Db.AutoMigrate(&models.ProductUnitInfo{}, &models.ProductInfo{},
	//	&models.LineInfo{}, &models.WorkProcessInfo{},
	//	&models.WorkCraftInfo{}, &models.WorkStationInfo{}, &models.Users{})
	// 产品
	//route.GET()
	//apiV1 := route.Group("/api")
	//apiV1.Use(cors.Default())
	//apiV1.POST("/login", views.Login)
	//apiV1.Use(middleware.JWTAuth())
	//apiV1.POST("/product/create", views.ProductCreate)
	//apiV1.GET("/product/list", views.GetProduct)
	//apiV1.POST("/product/update", views.ProductUpdate)
	//apiV1.POST("/product/delete", views.ProductDelete)
	//apiV1.GET("/product/detail/:product_id", views.ProductDetail)
	//// 单位
	//apiV1.POST("/unit/create", views.UnitCreate)
	//apiV1.POST("/unit/update", views.UnitUpdate)
	//apiV1.GET("/unit/get", views.UnitGet)
	//apiV1.POST("/unit/delete", views.UnitDelete)
	//// 工序
	//apiV1.POST("/workprocess", views.WorkProcessCreate)
	//apiV1.PUT("/workprocess", views.WorkProcessupdate)
	//apiV1.GET("/workprocess", views.WorkProcessGet)
	//apiV1.DELETE("/workprocess", views.WorkProcessDelete)
	//// 工艺
	//apiV1.POST("/workcraft", views.WorkCraftCreate)
	//apiV1.GET("/workcraft", views.WorkCraftGET)
	//apiV1.PUT("/workcraft", views.WorkCraftupdate)
	//apiV1.DELETE("/workcraft", views.WorkCraftDelete)
	////产线
	//apiV1.POST("/line", views.LineCreate)
	//apiV1.GET("/line", views.LineGet)
	//apiV1.PUT("/line", views.LineUpdate)
	//apiV1.DELETE("/line", views.LineDelete)
	//// 工位
	//apiV1.POST("/station", views.StationCreate)
	//apiV1.GET("/station", views.StationGet)
	//apiV1.PUT("/station", views.StationPut)
	//apiV1.DELETE("/station", views.StationDelete)
	//route.Run("0.0.0.0:5000")
	server := &http.Server{Addr: "0.0.0.0:5000", Handler: route}

	server.ListenAndServe()
}
