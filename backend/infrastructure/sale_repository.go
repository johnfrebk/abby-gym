package infrastructure

import (
	"POS/backend/database/models"
	"POS/backend/domain"
	"gorm.io/gorm"
)

type SaleRepository struct {
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) *SaleRepository {
	return &SaleRepository{db: db}
}

func (r *SaleRepository) UseDB(db *gorm.DB) {
	r.db = db
}

func (r *SaleRepository) SaveSale(sale *domain.Sale) (uint, error) {
	var saleORM models.Sale
	saleORM.ClientID = uint(sale.ClientID)

	if err := r.db.Create(&saleORM).Error; err != nil {
		return 0, err
	}
	saleORM.Total = sale.Total

	if err := r.db.Save(&saleORM).Error; err != nil {
		return 0, err
	}

	return saleORM.ID, nil
}

func (r *SaleRepository) SaveSaleDetails(details []models.SalesDetail) error {
	return r.db.Create(details).Error
}

func (r *SaleRepository) UpdateSale(sale *models.Sale) error {
	return r.db.Save(sale).Error
}

func (r *SaleRepository) UpdateSaleDetails(details []models.SalesDetail) error {
	return r.db.Create(details).Error
}

func (r *SaleRepository) GetSaleDetails(saleID uint) ([]models.SalesDetail, error) {
	var details []models.SalesDetail
	if err := r.db.Where("sale_id = ?", saleID).Find(&details).Error; err != nil {
		return nil, err
	}
	return details, nil
}

func (r *SaleRepository) DeleteSaleDetails(saleID uint) error {
	return r.db.Where("sale_id = ?", saleID).Delete(&models.SalesDetail{}).Error
}
