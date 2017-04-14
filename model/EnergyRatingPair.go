package model

import "github.com/jinzhu/gorm"

type EnergyRatingPair struct {
	gorm.Model
	Current   SanitizedInt
	Potential SanitizedInt
}
