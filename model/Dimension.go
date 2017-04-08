package model

type Dimension struct {
	Metric   string `xml:"metric"`
	Imperial string `xml:"imperial"`
	Mixed    string `xml:"mixed"`
}
