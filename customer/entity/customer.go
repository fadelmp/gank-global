package entity

// Database Design
type Customer struct {
	ID        uint   `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"type:VARCHAR(50);NOT NULL"`
	Phone     string `gorm:"type:VARCHAR(15);NOT NULL"`
	Email     string `gorm:"type:VARCHAR(50)"`
	IsActive  bool   `gorm:"type:BOOLEAN;default:true"`
	Addresses []Address
}
