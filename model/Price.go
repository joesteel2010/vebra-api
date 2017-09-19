package model

import (
	"time"
)

type Price struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	PropertyID uint         `gorm:"primary_key" sql:"type:int(10) unsigned"`
	Qualifier  string       `xml:"qualifier,attr"`
	Currency   string       `xml:"currency,attr"`
	Display    string       `xml:"display,attr"`
	Rent       string       `xml:"rent,attr"`
	Value      SanitizedInt `xml:",chardata"`
}
