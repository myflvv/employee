package schema

import "github.com/jinzhu/gorm"

type Site struct {
	gorm.Model
	Name string `gorm:"type:varchar(200);"`
	UniqueId string `gorm:"type:varchar(200)"`
}