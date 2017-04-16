package model

type Paragraph struct {
	ID          uint `gorm:"primary_key"`
	PropertyID  uint
	ParagraphID int          `xml:"id,attr"`
	Type        SanitizedInt `xml:"type,attr" json:"Type"`
	Name        string       `xml:"name"`
	File        string       `xml:"file"`
	Metric      string       `xml:"dimensions>metric"`
	Imperial    string       `xml:"dimensions>imperial"`
	Mixed       string       `xml:"dimensions>mixed"`
	Text        string       `xml:"text" sql:"type:text"`
}
