package model

type LandArea struct {
	Area `xml:"landarea"`
}

type LandAreas struct {
	Landareas []Area `xml:"landarea"`
}
