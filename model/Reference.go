package model

import (
	"time"
)

type Reference struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	PropertyID uint   `gorm:"primary_key" sql:"type:int"`
	Agents     string `xml:"agents"`
	Software   string `xml:"software"`
}
