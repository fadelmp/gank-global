package dto

// Database Design
type Address struct {
	ID          uint   `json:"id"`
	CustomerID  uint   `json:"customer_id"`
	AddressLine string `json:"address_line"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	PostalCode  string `json:"postal_code"`
	IsActive    bool   `json:"is_active"`
}
