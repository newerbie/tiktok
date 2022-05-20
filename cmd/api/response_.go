package entity

import (
	"github.com/ipfans/tiktok"
)

type HTTPErrResp struct {
	Meta Meta `json:"metadata"`
}

type HTTPEmptyResp struct {
	Meta Meta `json:"metadata"`
}

type Meta struct {
	Path       string    `json:"path"`
	StatusCode int       `json:"status_code"`
	Status     string    `json:"status"`
	Message    string    `json:"message"`
	Error      *AppError `json:"error,omitempty" swaggertype:"primitive,integer"`
	Timestamp  string    `json:"timestamp"`
}

type ErrorCode uint32

// AppError - Application Error Structure
type AppError struct {
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
	DebugError *string   `json:"debug,omitempty"`
	sys        error
}

// =======================================end ========================================

type HTTPCategoryTreeResp struct {
	Meta Meta              `json:"metadata"`
	Data []tiktok.Category `json:"data"`
}

type HTTBrandListResp struct {
	Meta       Meta          `json:"metadata"`
	Data       BrandListData `json:"data"`
	Pagination *Pagination   `json:"pagination"`
}

type Value struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type InputType struct {
	IsMandatory        bool `json:"is_mandatory"`
	IsMultipleSelected bool `json:"is_multiple_selected"`
	IsCustomized       bool `json:"is_customized"`
}

type Attribute struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Label         string    `json:"label"`
	AttributeType int       `json:"attribute_type"`
	InputType     InputType `json:"input_type"`
	Values        []Value   `json:"values"`
}

type ProductAttrsData struct {
	Attributes []Attribute `json:"attributes"`
}

type HTTPProductAttrsResp struct {
	Meta Meta             `json:"metadata"`
	Data ProductAttrsData `json:"data"`
}

type AddEditProductOpData struct {
	TotalCount      int `json:"total_count"`
	SuccessCount    int `json:"success_count"`
	SuccessRowsData []struct {
		ProductID int64 `json:"product_id"`
	} `json:"success_rows_data"`
	FailRowsData []struct {
		RequestID  string `json:"request_id"`
		SellerSKU  string `json:"seller_sku"`
		FailReason string `json:"fail_reason"`
	} `json:"fail_rows_data"`
}

type HTTPAddEditProductOpResp struct {
	Meta Meta                 `json:"metadata"`
	Data AddEditProductOpData `json:"data"`
}

type HTTPUpdateProductResp struct {
	Meta Meta                 `json:"metadata"`
	Data AddEditProductOpData `json:"data"`
}

type HTTPProductListResp struct {
	Meta       Meta                  `json:"metadata"`
	Data       ProductDetailListData `json:"data"`
	Pagination *Pagination           `json:"pagination"`
}

type SKUItem struct {
	SKUID             string   `json:"sku_id,omitempty"`
	SellerSku         string   `json:"seller_sku,omitempty"`
	Status            string   `json:"status,omitempty"`
	WarehouseID       string   `json:"warehouse_id,omitempty"`
	AvailableStock    int      `json:"available_stock,omitempty"`
	PackageWeight     int      `json:"package_weight,omitempty"`
	PackageHeight     int      `json:"package_height,omitempty"`
	PackageLength     int      `json:"package_length,omitempty"`
	PackageWidth      int      `json:"package_width,omitempty"`
	PackageWeightUnit string   `json:"package_weight_unit,omitempty"`
	Price             int      `json:"price,omitempty"`
	PriceCurrency     string   `json:"price_currency,omitempty"`
	Images            []string `json:"images,omitempty"`
}

type ProductItem struct {
	ProductID     string    `json:"product_id"`
	ProductStatus string    `json:"product_status"`
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description"`
	CategoryID    string    `json:"category_id"`
	BrandID       string    `json:"brand_id"`
	Images        []string  `json:"Images"`
	SKUS          []SKUItem `json:"skus"`
}

type ProductDetailListData struct {
	Products []ProductItem `json:"products"`
}

type BrandListData struct {
	BrandList []tiktok.Brand `json:"brand_list"`
}
