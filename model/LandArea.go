package model

type LandArea struct {
	PropertyID uint `gorm:"primary_key" sql:"type:int"`
	Area       `xml:"landarea"`
}
