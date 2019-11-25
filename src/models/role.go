package models

type Role struct {
	ID          uint          `gorm:"primary_key" json:"id"`
	Name        string        `gorm:"type:varchar(50);not null;unique_index"json:"name"`
	Description string        `gorm:"type:varchar(200);not null"json:"description"`
	Users       []*User       `gorm:"many2many:user_role;"`
	Permission  []*Permission `gorm:"many2many:role_permission;"`
}
