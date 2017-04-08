package model

type Price struct {
	Qualifier string       `xml:"qualifier,attr"`
	Currency  string       `xml:"currency,attr"`
	Display   string       `xml:"display,attr"`
	Rent      string       `xml:"rent,attr"`
	Value     SanitizedInt `xml:",chardata"`
}
