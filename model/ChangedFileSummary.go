package model

type ChangedFilesSummaries struct {
	Files []ChangedFileSummary `xml:"file"`
}

type ChangedFileSummary struct {
	FileID      int    `xml:"file_id"`
	FilePropId  int    `xml:"file_propid"`
	LastChanged string `xml:"updated"`
	IsDeleted   bool   `xml:"deleted"`
	Url         string `xml:"url"`
	PropUrl     string `xml:"prop_url"`
}
