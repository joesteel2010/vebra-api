package model

import "time"

type Area struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	PropertyID uint    `gorm:"primary_key" sql:"type:int(10) unsigned"`
	Unit       string  `xml:"unit,attr"`
	Min        float64 `xml:"min"`
	Max        float64 `xml:"max"`
}
