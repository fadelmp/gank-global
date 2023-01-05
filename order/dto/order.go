package dto

import "time"

// Database Design
type Order struct {
	ID            uint      `json:"id"`
	Number        string    `json:"number"`
	StatusID      uint      `json:"status_id"`
	CustomerName  string    `json:"name"`
	CustomerPhone string    `json:"phone"`
	CustomerEmail string    `json:"email"`
	AddressLine   string    `json:"address_line"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	Country       string    `json:"country"`
	PostalCode    string    `json:"postal_code"`
	Total         int       `json:"total"`
	Date          time.Time `json:"date"`
	Description   string    `json:"description"`
	Item          []Item    `json:"item"`
}
