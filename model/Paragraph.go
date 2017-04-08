package model

type Paragraph struct {
	Id        SanitizedInt `xml:"id,attr"`
	Type      SanitizedInt `xml:"type,attr"`
	Name      string       `xml:"name"`
	File      string       `xml:"file"`
	Dimension Dimension    `xml:"dimensions"`
	Text      string       `xml:"text"`
}

type Paragraphs struct {
	Paragraphs []Paragraph `xml:"paragraph"`
}
