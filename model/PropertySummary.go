package model

import (
	"github.com/jinzhu/gorm"
)

type PropertySummary struct {
	gorm.Model
	PropertyID  int    `xml:"prop_id"`
	LastChanged string `xml:"lastchanged"`
	Url         string `xml:"url"`
}

type PropertySummaries struct {
	Properties []PropertySummary `xml:"property"`
}
