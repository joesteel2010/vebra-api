package model

import "github.com/jinzhu/gorm"

type EnergyEfficiency struct {
	gorm.Model
	Current   SanitizedInt `xml:"current"`
	Potential SanitizedInt `xml:"potential"`
}

type EnvironmentalImpact struct {
	gorm.Model
	EnergyEfficiency
}
