package model

import (
	"time"
)

type StreetView struct {
	PropertyID   uint `gorm:"primary_key" sql:"type:int(10) unsigned"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	PovLatitude  float32
	PovLongitude float32
	PovPitch     float32
	PovHeading   float32
	PovZoom      int
}
