package entity

type Employee struct {
	EmployeeID int     `json:"employee_id" gorm:"employee_id"`
	FirstName  string  `json:"first_name" gorm:"first_name"`
	Salary     float64 `json:"salary" gorm:"salary"`
	LastName   string  `json:"last_name" gorm:"last_name"`
	Address    string  `json:"address" gorm:"address"`
	StoreID    int     `json:"store_id" gorm:"store_id"`
}

type EPhone struct {
	EmployeeID int    `json:"employee_id" gorm:"employee_id"`
	EPhone     string `json:"ephone" gorm:"ephone"`
}
type EPhoneCreation struct {
	EmployeeID int    `json:"employee_id" gorm:"employee_id"`
	EPhone     string `json:"ephone" gorm:"ephone"`
}

func (EPhoneCreation) TableName() string { return "ephone" }

type EmployeeCreation struct {
	EmployeeID int     `json:"employee_id" gorm:"employee_id"`
	FirstName  string  `json:"first_name" gorm:"first_name"`
	Salary     float64 `json:"salary" gorm:"salary"`
	LastName   string  `json:"last_name" gorm:"last_name"`
	Address    string  `json:"address" gorm:"address"`
	StoreID    int     `json:"store_id" gorm:"store_id"`
}

func (EmployeeCreation) TableName() string { return "employee" }

type EmployeeUpdate struct {
	EmployeeID int     `json:"employee_id" gorm:"employee_id"`
	FirstName  string  `json:"first_name" gorm:"first_name"`
	Salary     float64 `json:"salary" gorm:"salary"`
	LastName   string  `json:"last_name" gorm:"last_name"`
	Address    string  `json:"address" gorm:"address"`
	StoreID    int     `json:"store_id" gorm:"store_id"`
}

func (Employee) TableName() string { return "employee" }
