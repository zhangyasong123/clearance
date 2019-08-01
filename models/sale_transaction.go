package models

type SaleTransaction struct {
	Id             int64   `json:"id" query:"id"`
	OrderId        int64   `json:"orderId" query:"orderId"`
	StoreId        int64   `json:"storeId" query:"storeId" xorm:"index default 0" validate:"required"`
	TotalSalePrice float64 `json:"totalSalePrice" query:"totalSalePrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	Quantity       int64   `json:"quantity" query:"quantity" xorm:"notnull" validate:"required"`
	SalePrice      float64 `json:"salePrice" query:"salePrice" xorm:"DECIMAL(18,2) default 0" validate:"gte=0"`
	SkuId          int64   `json:"skuId" query:"skuId" xorm:"index notnull" validate:"gte=0"`
}
