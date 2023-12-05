package entity

type TakingOrdering struct {
	TransactionID int `json:"transaction_id" gorm:"transaction_id"`
	ShipperID     int `json:"shipper_id" gorm:"shipper_id"`
}
