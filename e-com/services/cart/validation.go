package cart

import (
	"fmt"
	"github.com/matizaj/go-app/e-com/types"
)

func getCardItemsIds(items []types.CartItem) ([]int, error) {
	productsIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid qty for products")
		}
		productsIds[i] = item.ProductId
	}
	return productsIds, nil
}

func (h *Handler) CreateOrder(ps []types.Product, items []types.CartItem, uid int) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.Id] = product
	}
	// products are in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, nil
	}

	// calculate total price
	totalPrice := calculateTotalPrice(items, productMap)

	// reduce qty of products stock
	for _, item := range items {
		product := productMap[item.ProductId]
		product.Quantity -= item.Quantity
		// h.productRepo.UpdateProduct(product) -> TODO
	}

	// create order
	orderId, err := h.orderRepo.CreateOrder(types.Order{
		UserId:  uid,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address",
	})
	if err != nil {
		return 0, 0, fmt.Errorf("cant create order: %s", err)
	}

	// create order item
	for _, item := range items {
		h.orderRepo.CreateOrderItem(types.OrderItem{
			OrderId:   orderId,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductId].Price,
		})
	}
	return orderId, totalPrice, nil
}

func calculateTotalPrice(items []types.CartItem, productMap map[int]types.Product) float64 {
	var total float64
	for _, item := range items {
		product := productMap[item.ProductId]
		total += product.Price * float64(item.Quantity)
	}
	return total
}

func checkIfCartIsInStock(items []types.CartItem, productMap map[int]types.Product) interface{} {
	if len(items) == 0 {
		return fmt.Errorf("cart is empty")
	}
	for _, item := range items {
		product, ok := productMap[item.ProductId]
		if !ok {
			return fmt.Errorf("product %s is not available in stock", item.ProductId)
		}
		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in stock in this quantity %s", item.ProductId, item.Quantity)
		}
	}
	return nil
}
