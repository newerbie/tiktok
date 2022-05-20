package main

import (
	"github.com/ipfans/tiktok/cmd/api"
	_ "github.com/ipfans/tiktok/cmd/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title          Swagger Product API
// @version        v0.0.1
// @description    This is product api swagger doc.
// @contact.name   yan.jinbin
// @contact.email  yan.jinbin@shipper.id
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @host           127.0.0.1:10000
// @BasePath       /v1
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		v1.GET("/categories", entity.HandleGetCategoryTree)
		v1.GET("/brands", entity.HandleGetBrands)
		v1.GET("/category/:category_id/attributes", entity.HandleGetCategoryAttrs)
		v1.POST("/products", entity.HandleCreateProducts)
		v1.PATCH("/products/:product_id", entity.HandleEditProduct)
		v1.PUT("/products/:product_id/prices", entity.HandleUpdateProductsPrices)
		v1.PUT("/products/:product_id/stocks", entity.HandleUpdateProductsStocks)
		v1.DELETE("/products", entity.HandleDeleteProducts)
		v1.GET("/products", entity.HandleGetProducts)
		v1.PUT("/products/active", entity.HandleActiveProducts)
		v1.PUT("/products/inactive", entity.HandleInactiveProducts)
	}
	r.Run(":10000")
}
