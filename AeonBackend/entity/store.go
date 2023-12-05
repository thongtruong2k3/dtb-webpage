package entity

import "time"

//should be hardcode database for store
type Store struct {
	StoreID     int       `json:"store_id" gorm:"store_id"`
	Name        string    `json:"name" gorm:"name"`
	OpeningDate time.Time `json:"opening_date" gorm:"opening_date"`
	ContactInfo string    `json:"contact_info" gorm:"contact_info"`
	Location    string    `json:"location" gorm:"location"`
}
