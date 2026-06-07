package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"POS/backend/features/clients/delete_client"
	"POS/backend/features/clients/get_client_by_id"
	"POS/backend/features/clients/get_clients"
	"POS/backend/features/clients/save_client"
	"POS/backend/features/clients/update_client"
	"POS/backend/features/dashboard/get_activities"
	"POS/backend/features/dashboard/get_dashboard"
	"POS/backend/features/memberships/delete_membership"
	"POS/backend/features/memberships/get_memberships"
	"POS/backend/features/memberships/save_membership"
	"POS/backend/features/memberships/update_membership"
	"POS/backend/features/products/delete_product"
	"POS/backend/features/products/get_products"
	"POS/backend/features/products/save_product"
	"POS/backend/features/products/update_product"
	"POS/backend/features/sales/delete_sale"
	"POS/backend/features/sales/get_sales"
	"POS/backend/features/sales/save_sale"
	"POS/backend/features/sales/update_sale"
	"POS/backend/features/subscriptions/delete_subscription"
	"POS/backend/features/subscriptions/get_subscriptions"
	"POS/backend/features/subscriptions/save_subscription"
	"POS/backend/features/subscriptions/update_subscription"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(CORSMiddleware)

	// Auth
	r.Post("/api/auth/register", RegisterHandler)
	r.Post("/api/auth/login", LoginHandler)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)

		// Clients
		r.Post("/api/clients", handleSaveClient)
		r.Get("/api/clients", handleGetAllClients)
		r.Get("/api/clients/paginated", handleGetClientsPaginated)
		r.Get("/api/clients/{id}", handleGetClientByID)
		r.Put("/api/clients/{id}", handleUpdateClient)
		r.Delete("/api/clients/{id}", handleDeleteClient)

		// Products
		r.Post("/api/products", handleSaveProduct)
		r.Get("/api/products", handleGetAllProducts)
		r.Put("/api/products/{id}", handleUpdateProduct)
		r.Delete("/api/products/{id}", handleDeleteProduct)

		// Memberships
		r.Post("/api/memberships", handleSaveMembership)
		r.Get("/api/memberships", handleGetAllMemberships)
		r.Put("/api/memberships/{id}", handleUpdateMembership)
		r.Delete("/api/memberships/{id}", handleDeleteMembership)

		// Subscriptions
		r.Post("/api/subscriptions", handleSaveSubscription)
		r.Get("/api/subscriptions", handleGetAllSubscriptions)
		r.Put("/api/subscriptions/{id}", handleUpdateSubscription)
		r.Delete("/api/subscriptions/{id}", handleDeleteSubscription)

		// Sales
		r.Post("/api/sales", handleSaveSale)
		r.Get("/api/sales", handleGetAllSales)
		r.Put("/api/sales/{id}", handleUpdateSale)
		r.Delete("/api/sales/{id}", handleDeleteSale)

		// Dashboard
		r.Get("/api/dashboard", handleGetDashboard)
		r.Get("/api/activities", handleGetActivities)
	})

	return r
}

// --- Clients ---

func handleSaveClient(w http.ResponseWriter, r *http.Request) {
	var req save_client.SaveClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	if err := save_client.NewSaveClientHandler().Handle(req); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusCreated, map[string]string{"message": "cliente creado"})
}

func handleGetAllClients(w http.ResponseWriter, r *http.Request) {
	clients, err := get_clients.NewGetClientsHandler().Handle()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, clients)
}

func handleGetClientsPaginated(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	result, err := get_clients.NewGetClientsHandler().HandlePaginated(page, pageSize)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, result)
}

func handleGetClientByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "id invalido")
		return
	}
	client, err := get_client_by_id.NewGetClientByIDHandler().Handle(get_client_by_id.GetClientByIDQuery{ID: uint(id)})
	if err != nil {
		errorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, client)
}

func handleUpdateClient(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "id invalido")
		return
	}
	var req update_client.UpdateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	req.ID = uint(id)
	if err := update_client.NewUpdateClientHandler().Handle(req); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]string{"message": "cliente actualizado"})
}

func handleDeleteClient(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "id invalido")
		return
	}
	if err := delete_client.NewDeleteClientHandler().Handle(delete_client.DeleteClientRequest{ID: uint(id)}); err != nil {
		errorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]string{"message": "cliente eliminado"})
}

// --- Products ---

func handleSaveProduct(w http.ResponseWriter, r *http.Request) {
	var req save_product.SaveProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	if err := save_product.NewSaveProductHandler().Handle(req); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusCreated, map[string]string{"message": "producto creado"})
}

func handleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := get_products.NewGetAllProductsHandler().Handle()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, products)
}

func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "id invalido")
		return
	}
	var req update_product.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	req.ID = id
	if err := update_product.NewUpdateProductHandler().Handle(req); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]string{"message": "producto actualizado"})
}

func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "id invalido")
		return
	}
	if err := delete_product.NewDeleteProductHandler().Handle(delete_product.DeleteProductRequest{ID: uint(id)}); err != nil {
		errorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]string{"message": "producto eliminado"})
}

// --- Memberships ---

func handleSaveMembership(w http.ResponseWriter, r *http.Request) {
	var req save_membership.SaveMembershipRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	if err := save_membership.NewSaveMembershipHandler().Handle(req); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusCreated, map[string]string{"message": "membresia creada"})
}

func handleGetAllMemberships(w http.ResponseWriter, r *http.Request) {
	memberships, err := get_memberships.NewGetAllMembershipsHandler().Handle()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, memberships)
}

func handleUpdateMembership(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "id invalido")
		return
	}
	var req update_membership.UpdateMembershipRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	req.ID = id
	if err := update_membership.NewUpdateMembershipHandler().Handle(req); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]string{"message": "membresia actualizada"})
}

func handleDeleteMembership(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "id invalido")
		return
	}
	if err := delete_membership.NewDeleteMembershipHandler().Handle(delete_membership.DeleteMembershipRequest{ID: uint(id)}); err != nil {
		errorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]string{"message": "membresia eliminada"})
}

// --- Subscriptions ---

func handleSaveSubscription(w http.ResponseWriter, r *http.Request) {
	var req save_subscription.SaveSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	if err := save_subscription.NewSaveSubscriptionHandler().Handle(req); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusCreated, map[string]string{"message": "suscripcion creada"})
}

func handleGetAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	subscriptions, err := get_subscriptions.NewGetAllSubscriptionsHandler().Handle()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, subscriptions)
}

func handleUpdateSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "id invalido")
		return
	}
	var req update_subscription.UpdateSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	req.ID = uint(id)
	if err := update_subscription.NewUpdateSubscriptionHandler().Handle(req); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]string{"message": "suscripcion actualizada"})
}

func handleDeleteSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "id invalido")
		return
	}
	if err := delete_subscription.NewDeleteSubscriptionHandler().Handle(delete_subscription.DeleteSubscriptionRequest{ID: uint(id)}); err != nil {
		errorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]string{"message": "suscripcion eliminada"})
}

// --- Sales ---

func handleSaveSale(w http.ResponseWriter, r *http.Request) {
	var req save_sale.SaveSaleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	if err := save_sale.NewSaveSaleHandler().Handle(req); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusCreated, map[string]string{"message": "venta registrada"})
}

func handleGetAllSales(w http.ResponseWriter, r *http.Request) {
	sales, err := get_sales.NewGetSalesHandler().Handle()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, sales)
}

func handleUpdateSale(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "id invalido")
		return
	}
	var req update_sale.UpdateSaleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	req.ID = uint(id)
	if err := update_sale.NewUpdateSaleHandler().Handle(req); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]string{"message": "venta actualizada"})
}

func handleDeleteSale(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "id invalido")
		return
	}
	if err := delete_sale.NewDeleteSaleHandler().Handle(delete_sale.DeleteSaleRequest{ID: uint(id)}); err != nil {
		errorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, map[string]string{"message": "venta eliminada"})
}

// --- Dashboard ---

func handleGetDashboard(w http.ResponseWriter, r *http.Request) {
	dashboard, err := get_dashboard.NewGetDashboardHandler().Handle()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, dashboard)
}

func handleGetActivities(w http.ResponseWriter, r *http.Request) {
	activities, err := get_activities.NewGetActivitiesHandler().Handle()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonResponse(w, http.StatusOK, activities)
}
