package model

type Donation struct {
	Amount           float64
	ReceiptRequested bool
	DonorName        string
	Status           Status
	Account          Account
}

type Status string

const (
	StatusInProcess   Status = "IN_PROCESS"
	StatusTransferred Status = "TRANSFERRED"
)
