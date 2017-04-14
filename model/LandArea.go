package model

import "github.com/jinzhu/gorm"

type LandArea struct {
	gorm.Model
	Area `xml:"landarea"`
}
