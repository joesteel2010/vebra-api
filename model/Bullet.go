package model

import (
	"github.com/jinzhu/gorm"
)

type Bullet struct {
	gorm.Model
	PropertyID uint         `gorm:"primary_key" sql:"type:int"`
	BulletID   SanitizedInt `xml:"id,attr" json:"ID" gorm:"primary_key" sql:"type:int"`
	Value      string       `xml:",chardata"`
}
