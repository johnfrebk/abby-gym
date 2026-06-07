package get_clients

import "POS/backend/utils"

type GetClientsHandler struct{}

func NewGetClientsHandler() *GetClientsHandler {
	return &GetClientsHandler{}
}

func (h *GetClientsHandler) Handle() ([]ClientResponse, error) {
	return GetAllClients(utils.NewPaginationParams(1, 1000))
}

func (h *GetClientsHandler) HandlePaginated(page, pageSize int) (*utils.PaginatedResult, error) {
	return GetAllClientsPaginated(utils.NewPaginationParams(page, pageSize))
}
