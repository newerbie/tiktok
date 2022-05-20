package entity

import (
	"github.com/gin-gonic/gin"
)

// HandleGetBrands godoc
// HandleGetBrands is get brand list from marketplace
// @Summary      get brand list from marketplace
// @Description  Marketplace open api doc link. Lazada: https://bit.ly/3l8loTF shopee: https://bit.ly/3L7JMQ5 Tokopedia: Tiktok: https://bit.ly/3FO6Pym
// @Description  e.g. get brand list like Apple Google Facebook Airbnb
// @Tags         brand
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  false  "Authorization signature"
// @Param        x-test-id      header    string  false  "test-id Key via Header"  Enums(mock-positive, mock-negative)
// @Param        shop_id        query     string  true   "common query params: shop id"
// @Param        mp             query     int     true   "common query params: marketplace type"
// @Param        page_no        query     int     true   "common query params: page number"
// @Param        page_size      query     int     true   "common query params: page size"
// @Success      200            {object}  entity.HTTBrandListResp
// @Failure      400            {object}  entity.HTTPErrResp
// @Failure      401            {object}  entity.HTTPErrResp
// @Failure      500            {object}  entity.HTTPErrResp
// @Router       /v1/brands [GET]
func HandleGetBrands(ctx *gin.Context) {}

// HandleGetCategoryTree godoc
// HandleGetCategoryTree is get category tree or dictionary from marketplace
// @Summary      Get category tree or dictionary from marketplace
// @Description  Marketplace open api doc link. Lazada: https://bit.ly/3wbFq64, shopee: https://bit.ly/3l7cjKT, Tokopedia: https://bit.ly/39fr2Rj, Tiktok: https://bit.ly/3FI0HaM
// @Description  e.g. get category tree, like telecommunication->mobile phone-> iphone -> iphone 13 series
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  false  "Authorization signature"
// @Param        x-test-id      header    string  false  "test-id Key via Header"  Enums(mock-positive, mock-negative)
// @Param        shop_id        query     string  true   "common query params: shop id"
// @Param        mp             query     int     true   "common query params: marketplace type"
// @Success      200            {object}  entity.HTTPCategoryTreeResp
// @Failure      400            {object}  entity.HTTPErrResp
// @Failure      401            {object}  entity.HTTPErrResp
// @Failure      500            {object}  entity.HTTPErrResp
// @Router       /v1/categories [GET]
func HandleGetCategoryTree(ctx *gin.Context) {

}

// HandleGetCategoryAttrs godoc
// HandleGetCategoryAttrs is get category attributes from marketplace
// @Summary      get category attributes from marketplace
// @Description  Marketplace open api doc link. Lazada: https://bit.ly/3Ncd0yp  shopee: https://bit.ly/3PxuYOb  Tokopedia: https://bit.ly/3Lwltvj Tiktok: https://bit.ly/3wtQdrm
// @Description  e.g. under category_id 100, product attributes, include mandatory or not, are waterproof fireproof isDangerous R18
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  false  "Authorization signature"
// @Param        x-test-id      header    string  false  "test-id Key via Header"  Enums(mock-positive, mock-negative)
// @Param        shop_id        query     string  true   "common query params: shop id"
// @Param        mp             query     int     true   "common query params: marketplace type"
// @Param        category_id    path      int     true   "category_id"
// @Success      200            {object}  entity.HTTPProductAttrsResp
// @Failure      400            {object}  entity.HTTPErrResp
// @Failure      401            {object}  entity.HTTPErrResp
// @Failure      500            {object}  entity.HTTPErrResp
// @Router       /v1/categories/{category_id}/attributes [GET]
func HandleGetCategoryAttrs(ctx *gin.Context) {

}

// HandleGetProducts godoc
// HandleGetProducts is fetch product id list
// @Summary      fetch product info list
// @Summary      Refer current Atoor: https://dashboard.staging.atoor.com/ocms/product
// @Description  Marketplace open api doc link. Lazada: batch query https://bit.ly/3MfPaBQ single query https://bit. ly/39mmw3i, shopee: https://bit.ly/3syQb0h get_item_list->get_item_base_info->get_item_extra_info Tokopedia: https://bit.ly/3PdbI8e Tiktok: https://bit.ly/3l7AkBy
// @Description  e.g. get Apple iphone series product list, iPhone15  red/green iphone16 black/white
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string   false  "Authorization signature"
// @Param        x-test-id      header    string   false  "test-id Key via Header"  Enums(mock-positive, mock-negative)
// @Param        shop_id        query     string   true   "common query params: shop id"
// @Param        mp             query     integer  true   "common query params: marketplace type"
// @Param        page_no        query     int      true   "common query params: page number"
// @Param        page_size      query     int      true   "common query params: page size"
// @Success      200            {object}  entity.HTTPProductListResp
// @Failure      400            {object}  entity.HTTPErrResp
// @Failure      401            {object}  entity.HTTPErrResp
// @Failure      500            {object}  entity.HTTPErrResp
// @Router       /v1/products [GET]
func HandleGetProducts(ctx *gin.Context) {

}

// HandleCreateProducts godoc
// HandleCreateProducts is create product on multiple marketplaces
// @Summary      create product on multiple marketplaces, how to create a product refer link: lazada usecase-> https://bit.ly/3yym6BG
// @Description  Marketplace open api doc link. Lazada: https://bit.ly/3PwiIgO shopee: https://bit.ly/3LjQimY Tokopedia: https://bit.ly/3wqKD9c  Tiktok: https://bit.ly/3NaZiMr
// @Description  e.g. create iphone-13 256GB on tiktok shop.
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                    false  "Authorization signature"
// @Param        x-test-id      header    string                    false  "test-id Key via Header"  Enums(mock-positive, mock-negative)
// @Param        shop_id        query     string                    true   "common query params: shop id"
// @Param        mp             query     integer                   true   "common query params: marketplace type"
// @Param        data           body      entity.AddEditProductReq  true   "Create Product metadata fields"
// @Success      200            {object}  entity.HTTPAddEditProductOpResp
// @Failure      400            {object}  entity.HTTPErrResp
// @Failure      401            {object}  entity.HTTPErrResp
// @Failure      500            {object}  entity.HTTPErrResp
// @Router       /v1/products [POST]
func HandleCreateProducts(ctx *gin.Context) {

}

// HandleEditProduct godoc
// HandleEditProduct is edit product attributes
// @Summary      edit product attributes
// @Description  Marketplace open api doc link. lazada: https://bit.ly/3ld8CU5 shopee: https://bit.ly/3lg4vX7 tokopedia: https://bit.ly/3FJFEV4 tiktok: https://bit.ly/3lb5VCs
// @Description  e.g. edie package weight,width, height, is_dangerous
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                    false  "Authorization signature"
// @Param        x-test-id      header    string                    false  "test-id Key via Header"  Enums(mock-positive, mock-negative)
// @Param        shop_id        query     string                    true   "common query params: shop id"
// @Param        mp             query     integer                   true   "common query params: marketplace type"
// @Param        data           body      entity.AddEditProductReq  true   "Update Product attributes"
// @Param        product_id     path      string                    true   "product id"
// @Success      200            {object}  entity.HTTPUpdateProductResp
// @Failure      400            {object}  entity.HTTPErrResp
// @Failure      401            {object}  entity.HTTPErrResp
// @Failure      500            {object}  entity.HTTPErrResp
// @Router       /v1/products/{product_id} [PATCH]
func HandleEditProduct(ctx *gin.Context) {

}

// HandleUpdateProductsPrices godoc
// HandleUpdateProductsPrices is update product sku price
// @Summary      update product price
// @Description  Marketplace open api doc link. lazada: https://bit.ly/3yBQVVV shopee: https://bit.ly/3Pn4HSe tokopedia: https://bit.ly/3PoACBR tiktok: https://bit.ly/3szsTHt
// @Description  e.g. old price is 1000Rp, new price is 668Rp
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                            false  "Authorization signature"
// @Param        x-test-id      header    string                            false  "test-id Key via Header"  Enums(mock-positive, mock-negative)
// @Param        shop_id        query     string                            true   "common query params: shop id"
// @Param        mp             query     integer                           true   "common query params: marketplace type"
// @Param        data           body      entity.UpdateProductPriceRequest  true   "Update Product Price"
// @Param        product_id     path      string                            true   "product id"
// @Success      200            {object}  HTTPEmptyResp
// @Failure      400            {object}  entity.HTTPErrResp
// @Failure      401            {object}  entity.HTTPErrResp
// @Failure      500            {object}  entity.HTTPErrResp
// @Router       /v1/products/{product_id}/price [PUT]
func HandleUpdateProductsPrices(ctx *gin.Context) {

}

// HandleUpdateProductsStocks godoc
// HandleUpdateProductsStocks is update product stock at specific warehouse
// @Summary      update product stock at specific warehouse
// @Description  Marketplace open api doc link. lazada: https://bit.ly/3yBQVVV shopee: https://bit.ly/3wrusIL tokopedia: https://bit.ly/3NcVf1Z tiktok: https://bit.ly/3sAyf57
// @Description  update product stock at specific warehouse, e.g. update stock from 1 to 100 at specific warehouse
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                            false  "Authorization signature"
// @Param        x-test-id      header    string                            false  "test-id Key via Header"  Enums(mock-positive, mock-negative)
// @Param        shop_id        query     string                            true   "common query params: shop id"
// @Param        mp             query     int                               true   "common query params: marketplace type"
// @Param        data           body      entity.UpdateProductStockRequest  true   "Update Product Stock"
// @Param        product_id     path      string                            true   "product id"
// @Success      200            {object}  HTTPEmptyResp
// @Failure      400            {object}  entity.HTTPErrResp
// @Failure      401            {object}  entity.HTTPErrResp
// @Failure      500            {object}  entity.HTTPErrResp
// @Router       /v1/products/{product_id}/stock [PUT]
func HandleUpdateProductsStocks(ctx *gin.Context) {

}

// HandleDeleteProducts godoc
// HandleDeleteProducts is batch delete product from product list
// @Summary      batch delete product from product list
// @Description  Marketplace open api doc link. tokopedia: https://bit.ly/38fpR4n, lazada: https://bit.ly/3yCt7l4 shopee: https://bit.ly/3wtMMAY tiktok: https://bit.ly/38qiuHe
// @Description  batch delete product from product list, e.g. delete iphone 13 red/black 256GB from shop product list
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                    false  "Authorization signature"
// @Param        x-test-id      header    string                    false  "test-id Key via Header"  Enums(mock-positive, mock-negative)
// @Param        shop_id        query     string                    true   "common query params: shop id"
// @Param        mp             query     integer                   true   "common query params: marketplace type"
// @Param        data           body      entity.ProductIDSRequest  true   "Delete Product From Shelf"
// @Success      200            {object}  HTTPEmptyResp
// @Failure      400            {object}  entity.HTTPErrResp
// @Failure      401            {object}  entity.HTTPErrResp
// @Failure      500            {object}  entity.HTTPErrResp
// @Router       /v1/products [delete]
func HandleDeleteProducts(ctx *gin.Context) {

}

// HandleActiveProducts godoc
// HandleActiveProducts is batch activate products on shelf
// @Summary      batch activate products on shelf
// @Description  Marketplace open api doc link. Lazada: https://bit.ly/37Ix6kR shopee: https://bit.ly/3wbm6G6 Tokopedia: https://bit.ly/3wkDsiH Tiktok: https://bit.ly/3sxkOTu
// @Description  e.g. iphone 13 red 256GB is top seller goods. so we often activate it
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                    false  "Authorization signature"
// @Param        x-test-id      header    string                    false  "test-id Key via Header"  Enums(mock-positive, mock-negative)
// @Param        shop_id        query     integer                   true   "common query params: shop id"
// @Param        mp             query     integer                   true   "common query params: marketplace type"
// @Param        data           body      entity.ProductIDSRequest  true   "activate Product From Shelf"
// @Success      200            {object}  HTTPEmptyResp
// @Failure      400            {object}  entity.HTTPErrResp
// @Failure      401            {object}  entity.HTTPErrResp
// @Failure      500            {object}  entity.HTTPErrResp
// @Router       /v1/products/active [POST]
func HandleActiveProducts(ctx *gin.Context) {

}

// HandleInactiveProducts godoc
// HandleInactiveProducts is batch deactivate products from shelf
// @Summary      batch deactivate products from shelf
// @Description  Marketplace open api doc link. Lazada: https://bit.ly/37Ix6kR shopee: https://bit.ly/3wbm6G6 Tokopedia: https://bit.ly/3wkDsiH Tiktok: https://bit.ly/3sxkOTu
// @Description  e.g. iphone 13 red 256GB is top seller goods. so we often deactivate it for the sake of shortage
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                    false  "Authorization signature"
// @Param        x-test-id      header    string                    false  "test-id Key via Header"  Enums(mock-positive, mock-negative)
// @Param        shop_id        query     integer                   true   "common query params: shop id"
// @Param        mp             query     integer                   true   "common query params: marketplace type"
// @Param        data           body      entity.ProductIDSRequest  true   "deactivate Product From Shelf"
// @Success      200            {object}  HTTPEmptyResp
// @Failure      400            {object}  entity.HTTPErrResp
// @Failure      401            {object}  entity.HTTPErrResp
// @Failure      500            {object}  entity.HTTPErrResp
// @Router       /v1/products/inactive [POST]
func HandleInactiveProducts(ctx *gin.Context) {

}
