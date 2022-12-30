package entity

// Database Design
type Address struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	CustomerID  uint   `gorm:"type:INT;NOT NULL;INDEX"`
	AddressLine string `gorm:"type:VARCHAR(255);NOT NULL"`
	City        string `gorm:"type:VARCHAR(30);NOT NULL"`
	Province    string `gorm:"type:VARCHAR(30);NOT NULL"`
	Country     string `gorm:"type:VARCHAR(30);NOT NULL"`
	PostalCode  string `gorm:"type:VARCHAR(10);NOT NULL"`
	IsActive    bool   `gorm:"type:BOOLEAN;default:true"`
}
