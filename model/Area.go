package model

type Area struct {
	PropertyID int
	Unit       string  `xml:"unit,attr"`
	Min        float64 `xml:"min"`
	Max        float64 `xml:"max"`
}
