package model

type Bullet struct {
	ID    SanitizedInt `xml:"id,attr"`
	Value string       `xml:",chardata"`
}

type Bullets struct {
	Bulllets []Bullet `xml:"bullet"`
}
