package model

import (
	"time"
)

type PropertySummary struct {
	PropertyID  uint `xml:"prop_id" gorm:"primary_key" sql:"type:int"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	LastChanged string `xml:"lastchanged"`
	Url         string `xml:"url"`
}

type PropertySummaries struct {
	Properties []PropertySummary `xml:"property"`
}
