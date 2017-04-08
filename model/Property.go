package model

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

type Properties struct {
	properties []Property `xml:"property"`
}

type SanitizedInt int
type SanitizedDate time.Time

func (si *SanitizedInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	// Read tag content into value
	d.DecodeElement(&value, &start)
	if value == "" {
		*si = (SanitizedInt)(0)
		return nil
	}

	i, err := strconv.ParseInt(strings.Replace(value, `""`, "", -1), 0, 64)
	if err != nil {
		return err
	}
	// Cast int64 to SanitizedInt
	*si = (SanitizedInt)(i)
	return nil
}

type SanitizedBool bool

func (si *SanitizedBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	// Read tag content into value
	d.DecodeElement(&value, &start)
	if value == "" {
		*si = (SanitizedBool)(false)
		return nil
	}

	i, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	// Cast int64 to SanitizedInt
	*si = (SanitizedBool)(i)
	return nil
}

func (si *SanitizedDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var value string
	var tmpTime time.Time
	d.DecodeElement(&value, &start)

	if value == "" {
		*si = (SanitizedDate)(time.Now())
		return nil
	}

	if len(value) == len("2006-01-02T15:04:05") {
		if tmpTime, err = time.Parse(`2006-01-02T15:04:05`, value); err == nil {
			*si = (SanitizedDate)(tmpTime)
		}
	}

	if strings.Contains(value, "/") {
		if tmpTime, err = time.Parse(`02/01/2006`, value); err == nil {
			*si = (SanitizedDate)(tmpTime)
		}
	}

	if strings.Contains(value, "-") && len(value) != len("2006-01-02T15:04:05") {
		if tmpTime, err = time.Parse(`2006-01-02`, value); err == nil {
			*si = (SanitizedDate)(tmpTime)
		}
	}

	if time.Time(*si).Before(time.Now().Add(-(time.Duration(time.Now().Unix()) * time.Second))) {
		*si = (SanitizedDate)(time.Now())
		return nil
	}

	*si = (SanitizedDate)(time.Now())
	return nil
}

type Property struct {
	ID             int           `xml:"id,attr"`
	PropertyID     string        `xml:"propertyid,attr"`
	System         string        `xml:"system,attr"`
	Firmid         string        `xml:"firmid,attr"`
	Branchid       string        `xml:"branchid,attr"`
	Database       string        `xml:"database,attr"`
	Featured       string        `xml:"featured,attr"`
	AgentReference Reference     `xml:"reference"` // to implement
	Address        Address       `xml:"address"`
	Price          Price         `xml:"price"`
	RentalFees     string        `xml:"rentalfees"`
	LettingsFee    string        `xml:"lettingsfee"`
	RmQualifier    SanitizedInt  `xml:"rm_qualifier"`
	Available      string        `xml:"available"`
	Uploaded       SanitizedDate `xml:"uploaded"`
	Longitude      float32       `xml:"longitude"`
	Latitude       float32       `xml:"latitude"`
	Easting        SanitizedInt  `xml:"easting"`
	Northing       SanitizedInt  `xml:"northing"`
	StreetView     StreetView    `xml:"streetview"`
	WebStatus      SanitizedInt  `xml:"web_status"`
	CustomStatus   string        `xml:"custom_status"`
	CommRent       string        `xml:"comm_rent"`
	Premium        string        `xml:"premium"`
	ServiceCharge  string        `xml:"service_charge"`
	RateableValue  string        `xml:"rateable_value"`
	Type           string        `xml:"type"`
	Furnished      string        `xml:"furnished"`
	RmType         SanitizedInt  `xml:"rm_type"`
	LetBond        SanitizedInt  `xml:"let_bond"`
	RmLetTypeID    SanitizedInt  `xml:"rm_let_type_id"`
	Bedrooms       SanitizedInt  `xml:"bedrooms"`
	Receptions     SanitizedInt  `xml:"receptions"`
	Bathrooms      SanitizedInt  `xml:"bathrooms"`
	UserField1     string        `xml:"userfield1"`
	UserField2     SanitizedInt  `xml:"userfield2"`
	SoldDate       SanitizedDate `xml:"solddate"`
	LeaseEnd       SanitizedDate `xml:"leaseend"`
	Instructed     SanitizedDate `xml:"instructed"`
	SoldPrice      SanitizedInt  `xml:"soldprice"`
	Garden         SanitizedBool `xml:"garden"`
	Parking        SanitizedBool `xml:"parking"`
	NewBuild       SanitizedBool `xml:"newbuild"`
	GroundRent     string        `xml:"groundrent"`
	Commission     string        `xml:"commission"`
	Area           Area          `xml:"area"`
	LandArea       LandArea      `xml:"landarea"`
	Description    string        `xml:"description"`
	Hip            Hip           `xml:"hip"`
	Paragraphs     Paragraphs    `xml:"paragraphs"`
	Bullets        Bullets       `xml:"bullets"`
	Files          Files         `xml:"files"`
	QueriedAt      SanitizedDate `xml:"queriedat"` // To implement - date-time retreived.
	LocalFiles     Files         // Path used for local storage
}
