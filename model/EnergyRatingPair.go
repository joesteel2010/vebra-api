package model

import (
	"time"
)

type EnergyEfficiency struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	PropertyID uint         `gorm:"primary_key" sql:"type:int(10) unsigned"`
	Current    SanitizedInt `xml:"current"`
	Potential  SanitizedInt `xml:"potential"`
}

type EnvironmentalImpact struct {
	EnergyEfficiency
}
