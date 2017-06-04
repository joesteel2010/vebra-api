package model

import "time"

type Bullet struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	PropertyID uint         `gorm:"primary_key" sql:"type:int(10) unsigned"`
	BulletID   SanitizedInt `xml:"id,attr" json:"ID" gorm:"primary_key" sql:"type:int"`
	Value      string       `xml:",chardata"`
}
