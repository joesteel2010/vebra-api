package model

import "github.com/jinzhu/gorm"

type Reference struct {
	gorm.Model
	PropertyID int
	Agents     string `xml:"agents"`
	Software   string `xml:"software"`
}
