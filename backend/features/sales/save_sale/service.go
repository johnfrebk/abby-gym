package save_sale

import (
	"POS/backend/database/models"
	"POS/backend/database/sqlite"
	"POS/backend/domain"
	"POS/backend/domain/services"
	"POS/backend/infrastructure"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SaveSaleService struct{}

func NewSaveSaleService() *SaveSaleService {
    return &SaveSaleService{}
}

func (s *SaveSaleService) Save(req SaveSaleRequest) error {
	sale := MapSaleRequestToDomain(req)
	details := MapDetailsRequestToDomain(req)
	services.CalculateSale(sale, details)

	return sqlite.DB.Transaction(func(tx *gorm.DB) error {
		txProductRepo := infrastructure.NewProductRepository(tx)
		txSaleRepo := infrastructure.NewSaleRepository(tx)

		for _, detail := range details {
			if err := txProductRepo.ValidateStock(detail.ProductID, detail.Quantity); err != nil {
				return errors.New("stock insuficiente para el producto: " + err.Error())
			}
		}

		saleID, err := txSaleRepo.SaveSale(sale)
		if err != nil {
			return fmt.Errorf("no se pudo guardar la venta: %w", err)
		}

		detailsModel := MapDetailsToModel(saleID, details)
		if err := txSaleRepo.SaveSaleDetails(detailsModel); err != nil {
			return err
		}

		for _, detail := range details {
			if err := txProductRepo.DecreaseStock(detail.ProductID, detail.Quantity); err != nil {
				return fmt.Errorf("no se pudo guardar la venta: %w", err)
			}
		}

		infrastructure.NewActivityRepository().CreateActivity(models.ActivityLog{
			Entity:   "Sale",
			EntityID: saleID,
			Action:   "Create",
			Summary:  fmt.Sprintf("Venta #%d guardada", saleID),
		})

		return nil
	})
}

func MapSaleRequestToDomain(req SaveSaleRequest) *domain.Sale {
	return &domain.Sale{
		ClientID: uint(req.ClientID),
		Total:    0,
	}
}

func MapDetailsRequestToDomain(req SaveSaleRequest) []domain.SalesDetail {
	var domainDetails []domain.SalesDetail
	for _, detail := range req.Details {
		domainDetails = append(domainDetails, domain.SalesDetail{
			ProductID: uint(detail.ProductID),
			Quantity:  int(detail.Quantity),
			Price:     float64(detail.Price),
		})
	}
	return domainDetails
}

func MapDetailsToModel(sale_id uint, details []domain.SalesDetail) []models.SalesDetail {
	var modelDetails []models.SalesDetail
	for _, detail := range details {
		modelDetails = append(modelDetails, models.SalesDetail{
			SaleID:    sale_id,
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
			Price:     detail.Price,
		})
	}
	return modelDetails
}
