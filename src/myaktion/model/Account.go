package model

type Account struct {
	Number   string `json:"number" gorm:"notNull;size:60"`
	Name     string `json:"name" gorm:"notNull;size:40"`
	BankName string `json:"bankName" gorm:"notNull;size:20"`
}
