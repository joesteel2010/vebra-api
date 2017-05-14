package model

import (
	"time"
)

type Dimension struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	PropertyID uint   `gorm:"primary_key" sql:"type:int"`
	Metric     string `xml:"metric"`
	Imperial   string `xml:"imperial"`
	Mixed      string `xml:"mixed"`
}
