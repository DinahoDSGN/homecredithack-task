package models

type Shop struct {
	Id       uint       `json:"shop_id"`
	Name     string     `json:"name"`
	Product  []Product  `gorm:"foreignKey:Id"`
	Category []Category `gorm:"foreignKey:Id"`
	Order    []Order    `gorm:"foreignKey:Id"`
}
