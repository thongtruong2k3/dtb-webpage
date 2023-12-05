package entity

type Customer struct {
	CustomerID int    `json:"customer_id" gorm:"customer_id;primaryKey"`
	CAddress   string `json:"c_address" gorm:"c_address"`
	CFName     string `json:"c_fname" gorm:"c_fname"`
	CLName     string `json:"c_lname" gorm:"c_lname"`
	CPhone     string `json:"c_phone" gorm:"c_phone"`
}

type CustomerCreation struct {
	CustomerID int    `json:"customer_id" gorm:"customer_id;primaryKey"`
	CAddress   string `json:"c_address" gorm:"c_address"`
	CFName     string `json:"c_fname" gorm:"c_fname"`
	CLName     string `json:"c_lname" gorm:"c_lname"`
	CPhone     string `json:"c_phone" gorm:"c_phone"`
}

func (CustomerCreation) TableName() string { return "customer" }

type CustomerUpdate struct {
	CustomerID int    `json:"customer_id" gorm:"customer_id;primaryKey"`
	CAddress   string `json:"c_address" gorm:"c_address"`
	CFName     string `json:"c_fname" gorm:"c_fname"`
	CLName     string `json:"c_lname" gorm:"c_lname"`
	CPhone     string `json:"c_phone" gorm:"c_phone"`
}

func (CustomerUpdate) TableName() string { return "customer" }
