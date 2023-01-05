package dto

// Database Design
type Item struct {
	ID          uint   `json:"id"`
	OrderID     uint   `json:"order_id"`
	ProductID   uint   `json:"product_id"`
	CategoryID  uint   `json:"category_id"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	IsActive    int    `json:"is_active"`
}
