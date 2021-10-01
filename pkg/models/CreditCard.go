package models

type CreditCard struct {
	Id     uint    `json:"credit_card_id"`
	Firstname string
	Lastname string
	Mobile string
	Number string  `json:"number"`
	Money  float64 `json:"money"`
}

type CreditCard_item struct {
	Id uint
	ProductId uint
	Product *Product `gorm:"foreignKey:Id"`
	CreditCardId uint
	CreditCard *CreditCard `gorm:"foreignKey:Id"`
}
