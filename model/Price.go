package model

import (
	"github.com/jinzhu/gorm"
)

type Price struct {
	gorm.Model
	PropertyID uint
	Qualifier  string       `xml:"qualifier,attr"`
	Currency   string       `xml:"currency,attr"`
	Display    string       `xml:"display,attr"`
	Rent       string       `xml:"rent,attr"`
	Value      SanitizedInt `xml:",chardata"`
}
