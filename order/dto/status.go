package dto

// Database Design
type Status struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	FrontDescription string `json:"front_description"`
	BackDescription  string `json:"back_description"`
	IsActive         bool   `json:"is_active"`
}
