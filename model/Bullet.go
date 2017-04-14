package model

type Bullet struct {
	PropertyID uint
	ID         SanitizedInt `xml:"id,attr"`
	Value      string       `xml:",chardata"`
}
