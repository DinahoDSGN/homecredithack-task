package models

type Category struct {
	Id uint
	ShopId uint
	Shop *Shop `gorm:"foreignKey:ShopId"`
	Title string
}

type Product_Category struct {
	Id uint
	ProductId uint
	Product []Product `gorm:"foreignKey:Id"`
	CategoryId uint
	Category *Category `gorm:"foreignKey:Id"`
}

type Product struct {
	Id uint `json:"product_id"`
	ShopId uint
	Shop *Shop `gorm:"foreignKey:ShopId"`
	Name string
	Description string
	Price float64
}

type Transaction struct {
	Id uint
	OrderId uint
	Order *Order `gorm:"foreignKey:Id"`
	Type string
	Code string
	Status bool
}
