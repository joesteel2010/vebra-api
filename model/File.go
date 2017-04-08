package model

type File struct {
	ID      int           `xml:"id,attr"`
	Type    int           `xml:"type,attr"`
	Name    string        `xml:"name"`
	Url     string        `xml:"url"`
	Updated SanitizedDate `xml:"updated"`
}

type Files struct {
	Files []File `xml:"file"`
}
