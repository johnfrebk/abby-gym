package update_sale

import (
	"POS/backend/database/models"
	"POS/backend/database/sqlite"
	"POS/backend/domain"
	"POS/backend/domain/services"
	"POS/backend/infrastructure"
	"fmt"

	"gorm.io/gorm"
)

type UpdateSaleService struct{}

func NewUpdateSaleService() *UpdateSaleService {
	return &UpdateSaleService{}
}

func (s *UpdateSaleService) UpdateSale(req UpdateSaleRequest) error {
	return sqlite.DB.Transaction(func(tx *gorm.DB) error {
		txSaleRepo := infrastructure.NewSaleRepository(tx)
		txProductRepo := infrastructure.NewProductRepository(tx)

		currentDetails, err := txSaleRepo.GetSaleDetails(req.ID)
		if err != nil {
			return fmt.Errorf("error al obtener detalles actuales: %w", err)
		}

		currentMap := make(map[uint]models.SalesDetail)
		for _, detail := range currentDetails {
			currentMap[detail.ID] = detail
		}

		newMap := make(map[uint]ProductItem)
		for _, detail := range req.Details {
			if detail.ID > 0 {
				newMap[detail.ID] = detail
			}
		}

		for id, current := range currentMap {
			if _, exists := newMap[id]; !exists {
				txProductRepo.IncreaseStock(current.ProductID, current.Quantity)
			}
		}

		var newDetails []models.SalesDetail
		for _, detail := range req.Details {
			if detail.ID > 0 {
				if current, exists := currentMap[detail.ID]; exists {
					adjustStockForEdit(txProductRepo, current, detail)
				}
			} else {
				txProductRepo.DecreaseStock(detail.ProductID, detail.Quantity)
			}

			newDetails = append(newDetails, models.SalesDetail{
				ID:        detail.ID,
				SaleID:    req.ID,
				ProductID: detail.ProductID,
				Quantity:  detail.Quantity,
				Price:     detail.Price,
			})
		}

		if err := txSaleRepo.DeleteSaleDetails(req.ID); err != nil {
			return fmt.Errorf("error al eliminar detalles: %w", err)
		}

		sale := &models.Sale{
			ID:       req.ID,
			ClientID: req.ClientID,
			Total:    0,
		}

		detailsDomain := make([]domain.SalesDetail, 0)
		for _, detail := range newDetails {
			detailsDomain = append(detailsDomain, domain.SalesDetail{
				ID:        detail.ID,
				SaleID:    detail.SaleID,
				ProductID: detail.ProductID,
				Quantity:  detail.Quantity,
				Price:     detail.Price,
			})
		}

		saleDomain := &domain.Sale{
			ID:       sale.ID,
			ClientID: sale.ClientID,
			Total:    sale.Total,
		}

		services.CalculateSale(saleDomain, detailsDomain)
		sale.Total = saleDomain.Total

		if err := txSaleRepo.UpdateSale(sale); err != nil {
			return fmt.Errorf("error al actualizar venta: %w", err)
		}
		if err := txSaleRepo.UpdateSaleDetails(newDetails); err != nil {
			return fmt.Errorf("error al guardar detalles: %w", err)
		}

		infrastructure.NewActivityRepository().CreateActivity(models.ActivityLog{
			Entity:   "Sale",
			EntityID: req.ID,
			Action:   "Update",
			Summary:  fmt.Sprintf("Venta #%d actualizada", req.ID),
		})

		return nil
	})
}

func adjustStockForEdit(repo *infrastructure.ProductRepository, current models.SalesDetail, new ProductItem) {
	if current.ProductID != new.ProductID {
		repo.IncreaseStock(current.ProductID, current.Quantity)
		repo.DecreaseStock(new.ProductID, new.Quantity)
	} else if current.Quantity != new.Quantity {
		diff := new.Quantity - current.Quantity
		if diff > 0 {
			repo.DecreaseStock(current.ProductID, diff)
		} else {
			repo.IncreaseStock(current.ProductID, -diff)
		}
	}
}
