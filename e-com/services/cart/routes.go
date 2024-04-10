package cart

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/matizaj/go-app/e-com/services/auth"
	"github.com/matizaj/go-app/e-com/types"
	"github.com/matizaj/go-app/e-com/utils"
	"net/http"
)

type Handler struct {
	orderRepo   types.OrderRepository
	productRepo types.ProductRepository
	userRepo    types.UserRepository
}

func NewHandler(orepo types.OrderRepository, prepo types.ProductRepository, urepo types.UserRepository) *Handler {
	return &Handler{
		orderRepo:   orepo,
		productRepo: prepo,
		userRepo:    urepo,
	}
}

func (h *Handler) RegisterHandler(router *http.ServeMux) {
	router.HandleFunc("POST /cart/checkout", auth.WithJwtAuth(h.Checkout, h.userRepo))
}

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	var payload types.CartCheckoutPayload
	err := utils.ParseJson(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.Validate.Struct(payload)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// get products
	productsIds, err := getCardItemsIds(payload.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	products, err := h.productRepo.GetProductsByIds(productsIds)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	userId := auth.GetUserIdFromCtx(r.Context())
	orderId, totalPrice, err := h.CreateOrder(products, payload.Items, userId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]any{
		"orderId": orderId,
		"price":   totalPrice,
	})
}
