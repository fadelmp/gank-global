package entity

import "time"

// Database Design
type Order struct {
	ID          uint      `gorm:"primaryKey;autoIncrement:true"`
	Number      string    `gorm:"type:VARCHAR(255);NOT NULL"`
	CustomerID  uint      `gorm:"type:INT;NOT NULL;INDEX"`
	StatusID    uint      `gorm:"type:INT;NOT NULL;INDEX"`
	AddressID   uint      `gorm:"type:INT;NOT NULL;INDEX"`
	Total       int       `gorm:"type:INT;NOT NULL;default:0"`
	Date        time.Time `gorm:"type:DATETIME;NOT NULL;default:now()"`
	Description string    `gorm:"type:VARCHAR(255)"`
	Status      Status    `gorm:"foreignKey:StatusID"`
	Items       []Item
}
