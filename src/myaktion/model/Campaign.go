package model

import "gorm.io/gorm"

type Campaign struct {
	gorm.Model

	Name               string     `json:"name" gorm:"notNull;size:30"`
	DonationMinimum    float64    `json:"donationMinimum" gorm:"notNull;check:donation_minimum>=1.0"`
	TargetAmount       float64    `json:"targetAmount" gorm:"notNull;check:target_amount>=10.0"`
	Account            Account    `json:"account" gorm:"embedded;embeddedPrefix:account_"`
	Organizer          string     `json:"organizer" gorm:"notNull"`
	Donations          []Donation `json:"donations" gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AmountDonatedSoFar float64    `gorm:"-"`
}
