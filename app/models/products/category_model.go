package product

type Category struct {
	ID          int `gorm:"primary_key;"`
	Name        string
	Description string
}
