package model

type Paragraph struct {
	PropertyID uint
	ID         SanitizedInt `xml:"id,attr"`
	Type       SanitizedInt `xml:"type,attr"`
	Name       string       `xml:"name"`
	File       string       `xml:"file"`
	Dimension  Dimension    `xml:"dimensions"`
	Text       string       `xml:"text"`
}
