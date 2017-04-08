package model

type Branch struct {
	ClientID  int           `xml:"clientid"`
	FirmID    int           `xml:"FirmID"`
	BranchID  int           `xml:"BranchID"`
	Name      string        `xml:"name"`
	URL       string        `xml:"url"`
	Street    string        `xml:"street"`
	Town      string        `xml:"town"`
	County    string        `xml:"county"`
	Postcode  string        `xml:"postcode"`
	Phone     string        `xml:"phone"`
	Email     string        `xml:"email"`
	QueriedAt SanitizedDate `xml:"queriedat"`
}

type Branches struct {
	Branches []Branch `xml:"branch"`
}
