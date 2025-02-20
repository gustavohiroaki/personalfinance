package models

type Fees struct {
	DBBase
	TransactionId int     `json:"transaction_id"`
	Settlement    float64 `json:"settlement"`
	Emolument     float64 `json:"emolument"`
	Brokerage     float64 `json:"brokerage"`
	Iss           float64 `json:"iss"`
}
