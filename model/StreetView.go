package model

import "github.com/jinzhu/gorm"

type StreetView struct {
	gorm.Model
	PropertyID   uint
	PovLatitude  float32
	PovLongitude float32
	PovPitch     float32
	PovHeading   float32
	PovZoom      int
}
