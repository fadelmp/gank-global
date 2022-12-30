package dto

// Database Design
type Customer struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	Addresses []Address `json:"address"`
}
