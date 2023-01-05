package entity

// Database Design
type Status struct {
	ID               uint   `gorm:"primaryKey;autoIncrement:true"`
	Name             string `gorm:"type:VARCHAR(50);NOT NULL"`
	FrontDescription string `gorm:"type:VARCHAR(100);NOT NULL"`
	BackDescription  string `gorm:"type:VARCHAR(100);NOT NULL"`
	IsActive         bool   `gorm:"type:BOOLEAN;default:true"`
}
