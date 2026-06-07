package save_sale

type SaveSaleHandler struct {
	Service *SaveSaleService
}

func NewSaveSaleHandler() *SaveSaleHandler {
	return &SaveSaleHandler{
		Service: NewSaveSaleService(),
	}
}

func (h *SaveSaleHandler) Handle(req SaveSaleRequest) error {
    return h.Service.Save(req)
}
