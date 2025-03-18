package models

type Product struct {
	Id          int    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Price       int    `gorm:"type:int" json:"price"`
	Description string `gorm:"type:text" json:"description"`
}
