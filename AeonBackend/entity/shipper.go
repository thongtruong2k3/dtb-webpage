package entity

type Shipper struct {
	ShipperID       int    `json:"shipper_id" gorm:"shipper_id"`
	VehicleCapacity int    `json:"vehicle_capacity" gorm:"vehicle_capacity"`
	Area            string `json:"area" gorm:"area"`
	Status          string `json:"status" gorm:"status"`
	SAddress        string `json:"s_address" gorm:"s_address"`
	SFName          string `json:"s_fname" gorm:"s_fname"`
	SLName          string `json:"s_lname" gorm:"s_lname"`
}

type ShipperPhoneNo struct {
	ShipperID int    `json:"shipper_id" gorm:"shipper_id"`
	PhoneNo   string `json:"phone_no" gorm:"phone_no"`
}

type AreaShip struct {
	ShipperID int    `json:"shipper_id" gorm:"shipper_id"`
	Area      string `json:"area" gorm:"area"`
}
