package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uname    string `gorm:"type:varchar(50);not null;unique_index"`
	Password string `gorm:"type:varchar(200);not null"`
	Mobile   string `gorm:"type:varchar(11);not null"`
	Email    string `gorm:"type:varchar(200);not null"`
	Motto    string `gorm:"type:varchar(50)"`
	Status   string `gorm:"type:enum('normal', 'abnormal');default:'normal'"`
}
