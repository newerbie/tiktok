package entity

import (
	"github.com/ipfans/tiktok"
)

type Pagination struct {
	CurrentPage     int64    `json:"current_page" example:"1"`
	CurrentElements int64    `json:"current_elements" example:"10"`
	TotalPages      int64    `json:"total_pages" example:"5"`
	TotalVariants   *int64   `json:"total_variants,omitempty" example:"50"`
	TotalElements   int64    `json:"total_elements" example:"50"`
	SortBy          []string `json:"sort_by"`
	CursorStart     *string  `json:"cursor_start,omitempty"`
	CursorEnd       *string  `json:"cursor_end,omitempty"`
	HasNextPage     bool     `json:"has_next_page,omitempty"`
	NextOffset      int      `json:"next_offset,omitempty"`
}

type StandardSKU struct {
	IsPrimary     bool             `json:"is_primary"`
	SellerSku     int              `json:"seller_sku"`
	Price         int              `json:"price"`
	PriceCurrency string           `json:"price_currency"`
	StockInfo     tiktok.StockInfo `json:"stock_info"`
	Weight        int              `json:"weight"`
	ImageURLList  []string         `json:"image_url_list"`
	Attributes    []Attribute      `json:"attribute_list"`
	PackageLength int              `json:"package_length"`
	PackageWidth  int              `json:"package_width"`
	PackageHeight int              `json:"package_height"`
	PackageWeight int              `json:"package_weight"`
	WeightUnit    string           `json:"weight_unit"`
}

type StandardShipperOption struct {
	SizeID     int  `json:"size_id"`
	ShipFee    int  `json:"ship_fee"`
	Enable     bool `json:"enabled"`
	LogisticID int  `json:"logistic_id"`
	IsFreeShip bool `json:"is_free_ship"`
}

type WarrantPolicy struct {
	WarrantType        int    `json:"warrant_type" example:"-1"` //
	WarrantDescription string `json:"warrant_description" example:"warrant_type: -1 lifetime,0 No-warrant policy 1 limited 2 full, WarrantTime: day unit "`
	WarrantTime        int    `json:"warrant_time"`
}

type AddEditProductReq struct {
	// ProductID              string                  `json:"product_id,omitempty" example:"when call EditProduct API"`
	Name                   string                  `json:"name"`
	Description            string                  `json:"description" example:"usually rich text,contains media file, like image, video content."`
	ShortDescription       string                  `json:"short_description"`
	CategoryID             string                  `json:"category_id"`
	BrandID                string                  `json:"brand_id"`
	ImagesURLList          []string                `json:"images_url_list"`
	IsSingleProduct        bool                    `json:"is_single_product"`
	IsDangerous            bool                    `json:"is_dangerous"`
	Price                  int                     `json:"price"`
	PriceCurrency          string                  `json:"price_currency"`
	ProductStatus          string                  `json:"product_status"`
	ProductCondition       string                  `json:"product_condition"`
	MinOrder               int                     `json:"min_order"`
	ProductWeight          int                     `json:"product_weight"`
	WeightUnit             string                  `json:"weight_unit"`
	WarrantPolicy          WarrantPolicy           `json:"warrant_policy"`
	StandardSKUS           []StandardSKU           `json:"standard_sku_list"`
	StandardShipperOptions []StandardShipperOption `json:"standard_shipper_option_list"`
}

type SKUInfo struct {
	ID int `json:"id"`
	//SKUNo       string `json:"sku_no"`
	SellerSKUNo string `json:"seller_sku_no"`
	Name        string `json:"name"`
}

type SKUPriceItem struct {
	SKUID         string `json:"sku_id"`
	OriginalPrice int    `json:"original_price"` // example:"default Rp"
	PriceCurrency int    `json:"price_currency"`
}

type UpdateProductPriceRequest struct {
	ProductID string         `json:"product_id" `
	SKUPrices []SKUPriceItem `json:"prices" validate:"required,min=1"`
}

type SKUStockItem struct {
	SKUID      string             `json:"sku_id"`
	StockInfos []tiktok.StockInfo `json:"stock_infos"`
}

type UpdateProductStockRequest struct {
	ProductID string         `json:"product_id"`
	SKUS      []SKUStockItem `json:"skus"`
}

type ProductIDSRequest struct {
	ProductIDS []string `json:"product_ids"`
}

type ActivateProductRequest struct {
	SpuID string    `json:"spu_id,omitempty"`
	Skus  []SKUInfo `json:"skus" validate:"required"`
}
