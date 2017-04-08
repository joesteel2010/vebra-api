package model

type PropertySummary struct {
	PropID      int    `xml:"prop_id"`
	LastChanged string `xml:"lastchanged"`
	Url         string `xml:"url"`
}

type PropertySummaries struct {
	Properties []PropertySummary `xml:"property"`
}
