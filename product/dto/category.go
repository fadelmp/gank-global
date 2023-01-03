package dto

// Database Design
type Category struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	Products    []Product `json:"product"`
}
