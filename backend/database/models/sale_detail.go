package models

type SalesDetail struct {
	ID        uint    `gorm:"primaryKey"`
	SaleID    uint    `gorm:"not null;index:idx_sales_detail_sale_id"`
	ProductID uint    `gorm:"not null;index:idx_sales_detail_product_id"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`

	Sale    Sale    `gorm:"foreignKey:SaleID"`
	Product Product `gorm:"foreignKey:ProductID"`
}
