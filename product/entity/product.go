package entity

// Database Design
type Product struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	Name        string `gorm:"type:VARCHAR(255);NOT NULL"`
	Description string `gorm:"type:VARCHAR(255)"`
	CategoryID  uint   `gorm:"type:INT;NOT NULL;INDEX"`
	Price       int    `gorm:"type:VARCHAR(30);NOT NULL;default:0"`
	Stock       int    `gorm:"type:VARCHAR(30);NOT NULL;default:0"`
	IsActive    bool   `gorm:"type:BOOLEAN;default:true"`
}
