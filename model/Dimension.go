package model

import "github.com/jinzhu/gorm"

type Dimension struct {
	gorm.Model
	Metric   string `xml:"metric"`
	Imperial string `xml:"imperial"`
	Mixed    string `xml:"mixed"`
}
