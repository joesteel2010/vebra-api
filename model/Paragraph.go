package model

type Paragraph struct {
	ID          uint
	PropertyID  uint         `gorm:"primary_key" sql:"type:int"`
	ParagraphID int          `xml:"id,attr" gorm:"primary_key" sql:"type:int"`
	Type        SanitizedInt `xml:"type,attr" json:"Type"`
	Name        string       `xml:"name"`
	File        string       `xml:"file"`
	Metric      string       `xml:"dimensions>metric"`
	Imperial    string       `xml:"dimensions>imperial"`
	Mixed       string       `xml:"dimensions>mixed"`
	Text        string       `xml:"text" sql:"type:text"`
}
