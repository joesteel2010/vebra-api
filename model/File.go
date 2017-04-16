package model

type File struct {
	PropertyID int
	FileID     int           `xml:"id,attr" sql:"type=integer" json:"ID"`
	Type       int           `xml:"type,attr"`
	Name       string        `xml:"name"`
	Url        string        `xml:"url"`
	Updated    SanitizedDate `xml:"updated" sql:"type=datetime"`
}
