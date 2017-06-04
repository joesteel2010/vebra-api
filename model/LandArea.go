package model

type LandArea struct {
	PropertyID uint `gorm:"primary_key" sql:"type:int(10) unsigned"`
	Area       `xml:"landarea"`
}
