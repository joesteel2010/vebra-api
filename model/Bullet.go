package model

import (
	"github.com/jinzhu/gorm"
)

type Bullet struct {
	gorm.Model
	PropertyID uint
	BulletID   SanitizedInt `xml:"id,attr" json:"ID"`
	Value      string       `xml:",chardata"`
}
