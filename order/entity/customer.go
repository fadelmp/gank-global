package entity

// Request Message Design

type CustomerAllResponse struct {
	SuccessCode int        `json:"SuccessCode"`
	Data        []Customer `json:"Data"`
}

type CustomerResponse struct {
	SuccessCode int      `json:"SuccessCode"`
	Data        Customer `json:"Data"`
}

type Customer struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	IsActive bool      `json:"is_active"`
	Address  []Address `json:"address"`
}

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
