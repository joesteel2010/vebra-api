package model

import (
	"time"
)

type Address struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	PropertyID     uint   `gorm:"primary_key" sql:"type:int(10) unsigned"`
	Name           string `xml:"name"`
	Street         string `xml:"street"`
	Locality       string `xml:"locality"`
	Town           string `xml:"town"`
	County         string `xml:"county"`
	Postcode       string `xml:"postcode"`
	CustomLocation string `xml:"custom_locatiom"`
	Display        string `xml:"display"`
}
