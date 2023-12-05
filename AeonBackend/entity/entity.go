package entity

import "time"

type Bill struct {
	TransactionID int       `json:"transaction_id" gorm:"transaction_id"`
	PaymentMethod string    `json:"payment_method" gorm:"payment_method"`
	DateAndTime   time.Time `json:"date_and_time" gorm:"date_and_time"`
	CustomerID    int       `json:"customer_id" gorm:"customer_id"`
	CashierID     int       `json:"cashier_id" gorm:"cashier_id"`
	StoreID       int       `json:"store_id" gorm:"store_id"`
}

type BillPromotion struct {
	ApplyPrice  int `json:"apply_price" gorm:"apply_price"`
	PromotionID int `json:"promotion_id" gorm:"promotion_id"`
}

type ProductPromotion struct {
	PromotionID int `json:"promotion_id" gorm:"promotion_id"`
}

type PromoteBill struct {
	PromotionID   int `json:"promotion_id" gorm:"promotion_id"`
	TransactionID int `json:"transaction_id" gorm:"transaction_id"`
}

type PromoteProduct struct {
	PromotionID int `json:"promotion_id" gorm:"promotion_id"`
	ProductID   int `json:"product_id" gorm:"product_id"`
}


type StoreManager struct {
	EmployeeID int `json:"employee_id" gorm:"employee_id"`
	StoreID    int `json:"store_id" gorm:"store_id"`
}

type Cashier struct {
	EmployeeID int    `json:"employee_id" gorm:"employee_id"`
	Shift      string `json:"shift" gorm:"shift"`
}

type At struct {
	ProductID     int `json:"product_id" gorm:"product_id"`
	StoreID       int `json:"store_id" gorm:"store_id"`
	NumberAtStore int `json:"number_at_store" gorm:"number_at_store"`
}

type Include struct {
	TransactionID         int `json:"transaction_id" gorm:"transaction_id"`
	ProductID             int `json:"product_id" gorm:"product_id"`
	NumberOfProductInBill int `json:"number_of_product_in_bill" gorm:"number_of_product_in_bill"`
}
