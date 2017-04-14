package model

import "time"

type Reference struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	PropertyID int
	Agents     string `xml:"agents"`
	Software   string `xml:"software"`
}
