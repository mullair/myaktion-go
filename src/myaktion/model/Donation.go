package model

import "gorm.io/gorm"

type Donation struct {
	gorm.Model
	Amount           float64 `gorm:"notNull;check:amount >= 1.0"`
	ReceiptRequested bool    `gorm:"notNull"`
	DonorName        string  `gorm:"notNull;size:40"`
	Status           Status  `gorm:"notNull;type:ENUM('TRANSFERRED','IN_PROCESS')"`
	Account          Account `gorm:"embedded;embeddedPrefix:account_"`
	CampaignID       uint
}

type Status string

const (
	StatusInProcess   Status = "IN_PROCESS"
	StatusTransferred Status = "TRANSFERRED"
)
