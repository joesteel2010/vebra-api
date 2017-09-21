package model

import "time"

type File struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	PropertyID uint          `gorm:"primary_key" sql:"type:int(10) unsigned"`
	FileID     int           `xml:"id,attr" json:"ID" gorm:"primary_key" sql:"type:int"`
	Type       int           `xml:"type,attr"`
	Name       string        `xml:"name"`
	Url        string        `xml:"url"`
	Updated    SanitizedDate `xml:"updated" sql:"type:datetime"`
}
