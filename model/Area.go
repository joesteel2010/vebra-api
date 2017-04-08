package model

type Area struct {
	Unit string  `xml:"unit,attr"`
	Min  float64 `xml:"min"`
	Max  float64 `xml:"max"`
}

type Areas struct {
	Areas []Area `xml:"area"`
}
