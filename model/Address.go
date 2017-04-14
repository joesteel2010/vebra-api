package model

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	PropertyID     int
	Name           string `xml:"name"`
	Street         string `xml:"street"`
	Locality       string `xml:"locality"`
	Town           string `xml:"town"`
	County         string `xml:"county"`
	Postcode       string `xml:"postcode"`
	CustomLocation string `xml:"custom_locatiom"`
	Display        string `xml:"display"`
}
