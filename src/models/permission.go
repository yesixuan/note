package models

type Permission struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	Name        string  `gorm:"type:varchar(50);not null;unique_index"json:"name"`
	Description string  `gorm:"type:varchar(200);not null"json:"description"`
	Role        []*Role `gorm:"many2many:role_permission;"`
}
