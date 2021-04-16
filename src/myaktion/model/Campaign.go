package model

type Campaign struct {
	ID              uint       `json:"ID"`
	Name            string     `json:"name"`
	DonationMinimum float64    `json:"donationMinimum"`
	TargetAmount    float64    `json:"targetAmount"`
	Account         Account    `json:"account"`
	Organizer       Organizer  `json:"organizer"`
	Donations       []Donation `json:"donations"`
}
