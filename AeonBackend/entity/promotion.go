package entity

import "time"

type Promotion struct {
	PromotionID int       `json:"promotion_id" gorm:"promotion_id"`
	Discount    float64   `json:"discount" gorm:"discount"`
	Name        string    `json:"name" gorm:"name"`
	Description string    `json:"description" gorm:"description"`
	StartDay    time.Time `json:"start_day" gorm:"start_day"`
	EndDay      time.Time `json:"end_day" gorm:"end_day"`
}
