package update_sale

type UpdateSaleHandler struct {
	Service *UpdateSaleService
}

func NewUpdateSaleHandler() *UpdateSaleHandler {
	return &UpdateSaleHandler{
		Service: NewUpdateSaleService(),
	}
}

func (h *UpdateSaleHandler) Handle(req UpdateSaleRequest) error {
	return h.Service.UpdateSale(req)
}