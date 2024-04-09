package product

import (
	"github.com/matizaj/go-app/e-com/types"
	"github.com/matizaj/go-app/e-com/utils"
	"net/http"
)

type Handler struct {
	repository types.ProductRepository
}

func NewHandler(repo types.ProductRepository) *Handler {
	return &Handler{repository: repo}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /products", h.GetProducts)
	router.HandleFunc("GET /products-id", h.GetProductsByIds)
	router.HandleFunc("POST /products", h.CreateProduct)
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.repository.GetAllProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, products)
}
func (h *Handler) GetProductsByIds(w http.ResponseWriter, r *http.Request) {
	products, err := h.repository.GetProductsByIds([]int{1, 2})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, products)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.Product
	err := utils.ParseJson(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err = h.repository.CreateProduct(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, nil)
}
