package entity

// Database Design
type Category struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	Name        string `gorm:"type:VARCHAR(50);NOT NULL"`
	Description string `gorm:"type:VARCHAR(100);NOT NULL"`
	IsActive    bool   `gorm:"type:BOOLEAN;default:true"`
	Products    []Product
}
