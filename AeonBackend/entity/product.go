package entity

type Product struct {
	ProductID   int     `json:"product_id" gorm:"product_id"`
	Category    string  `json:"category" gorm:"category"`
	Description string  `json:"description" gorm:"description"`
	PName       string  `json:"p_name" gorm:"p_name"`
	Price       float64 `json:"price" gorm:"price"`
	Weight      float64 `json:"weight" gorm:"weight"`
}
