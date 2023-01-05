package entity

// Database Design
type Item struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	OrderID     uint   `gorm:"type:INT;NOT NULL;INDEX"`
	ProductID   uint   `gorm:"type:INT;NOT NULL;INDEX"`
	Quantity    int    `gorm:"type:INT;NOT NULL"`
	Description string `gorm:"type:VARCHAR(255)"`
	IsActive    int    `gorm:"type:BOOLEAN;default:true"`
}
