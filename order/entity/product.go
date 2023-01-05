package entity

// Request Message Design

type ProductAllResponse struct {
	SuccessCode int       `json:"SuccessCode"`
	Data        []Product `json:"Data"`
}

type ProductResponse struct {
	SuccessCode int     `json:"SuccessCode"`
	Data        Product `json:"Data"`
}

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	IsActive    bool   `json:"is_active"`
}
