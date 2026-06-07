package get_clients

import (
	"POS/backend/database/models"
	"POS/backend/database/sqlite"
	"POS/backend/utils"
	"errors"
)

func GetAllClients(params utils.PaginationParams) ([]ClientResponse, error) {
	var clients []models.Client
	result := sqlite.DB.Offset(params.Offset()).Limit(params.Limit()).Find(&clients)

	if result.Error != nil {
		return nil, errors.New("no se encontraron clientes")
	}

	return MapClientModelListToResponse(clients), nil
}

func GetAllClientsPaginated(params utils.PaginationParams) (*utils.PaginatedResult, error) {
	var clients []models.Client
	var total int64

	sqlite.DB.Model(&models.Client{}).Count(&total)
	result := sqlite.DB.Offset(params.Offset()).Limit(params.Limit()).Find(&clients)

	if result.Error != nil {
		return nil, errors.New("no se encontraron clientes")
	}

	data := MapClientModelListToResponse(clients)
	pr := utils.NewPaginatedResult(data, total, params)
	return &pr, nil
}
