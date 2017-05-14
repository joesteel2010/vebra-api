package model

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
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

func (u SanitizedInt) Value() (driver.Value, error)  { return int64(u), nil }
func (u SanitizedDate) Value() (driver.Value, error) { return time.Time(u), nil }
func (u SanitizedBool) Value() (driver.Value, error) {
	if bool(u) {
		return int64(1), nil
	}
	return int64(0), nil
}

func (u *SanitizedInt) Scan(value interface{}) error {
	switch value.(type) {
	case []uint8:
		s := value.([]uint8)
		val, err := strconv.Atoi(string(s))

		if err != nil {
			return fmt.Errorf("Error converting to sanitized int: %s", err)
		}

		*u = SanitizedInt(val)
	case int64:
		*u = SanitizedInt(value.(int64))
	}
	return nil
}

func (u *SanitizedDate) Scan(value interface{}) error {
	switch value.(type) {
	case time.Time:
		*u = SanitizedDate(value.(time.Time))
		return nil
	case []uint8:
		layout := "2006-01-02 15:04:05"
		str := string(value.([]uint8))
		t, err := time.Parse(layout, str)

		if err != nil {
			fmt.Errorf("Error scanning SanitizedDate: %s", err)
		}
		*u = SanitizedDate(t)
		return nil
	default:
		fmt.Errorf("Error: cant handle type %T in SanitizedDate.Scan", value)
		return nil
	}
}

func (u *SanitizedBool) Scan(value interface{}) error {
	s := value.([]uint8)
	val, err := strconv.Atoi(string(s))

	if err != nil {
		return fmt.Errorf("Error converting to sanitized int: %s", err)
	}
	if val == 1 {
		*u = SanitizedBool(true)
		return nil
	}

	*u = SanitizedBool(false)
	return nil
}

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
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time
	ID                  uint                `xml:"id,attr"`
	PropertyID          int                 `xml:"propertyid,attr"`
	System              string              `xml:"system,attr"`
	Firmid              string              `xml:"firmid,attr"`
	Branchid            string              `xml:"branchid,attr"`
	Database            string              `xml:"database,attr"`
	Featured            string              `xml:"featured,attr"`
	AgentReference      Reference           `xml:"reference"` // to implement
	Address             Address             `xml:"address"`
	Price               Price               `xml:"price"`
	RentalFees          string              `xml:"rentalfees"`
	LettingsFee         string              `xml:"lettingsfee"`
	RmQualifier         SanitizedInt        `xml:"rm_qualifier"`
	Available           string              `xml:"available"`
	Uploaded            SanitizedDate       `xml:"uploaded" gorm:"type:datetime"`
	Longitude           float32             `xml:"longitude"`
	Latitude            float32             `xml:"latitude"`
	Easting             SanitizedInt        `xml:"easting"`
	Northing            SanitizedInt        `xml:"northing"`
	StreetView          StreetView          `xml:"streetview"`
	WebStatus           SanitizedInt        `xml:"web_status"`
	CustomStatus        string              `xml:"custom_status"`
	CommRent            string              `xml:"comm_rent"`
	Premium             string              `xml:"premium"`
	ServiceCharge       string              `xml:"service_charge"`
	RateableValue       string              `xml:"rateable_value"`
	Type                string              `xml:"type"`
	Furnished           string              `xml:"furnished"`
	RmType              SanitizedInt        `xml:"rm_type"`
	LetBond             SanitizedInt        `xml:"let_bond"`
	RmLetTypeID         SanitizedInt        `xml:"rm_let_type_id"`
	Bedrooms            SanitizedInt        `xml:"bedrooms"`
	Receptions          SanitizedInt        `xml:"receptions"`
	Bathrooms           SanitizedInt        `xml:"bathrooms"`
	UserField1          string              `xml:"userfield1"`
	UserField2          SanitizedInt        `xml:"userfield2"`
	SoldDate            SanitizedDate       `xml:"solddate" gorm:"type:datetime"`
	LeaseEnd            SanitizedDate       `xml:"leaseend" gorm:"type:datetime"`
	Instructed          SanitizedDate       `xml:"instructed" gorm:"type:datetime"`
	SoldPrice           SanitizedInt        `xml:"soldprice"`
	Garden              SanitizedBool       `xml:"garden"`
	Parking             SanitizedBool       `xml:"parking"`
	NewBuild            SanitizedBool       `xml:"newbuild"`
	GroundRent          string              `xml:"groundrent"`
	Commission          string              `xml:"commission"`
	Area                Area                `xml:"area"`
	LandArea            LandArea            `xml:"landarea"`
	Description         string              `xml:"description" gorm:"type:varchar(2056)"`
	EnergyEfficiency    EnergyEfficiency    `xml:"hip>energy_performance>energy_efficiency"`
	EnvironmentalImpact EnvironmentalImpact `xml:"hip>energy_performance>environmental_impact"`
	Paragraphs          []Paragraph         `xml:"paragraphs>paragraph"`
	Bullets             []Bullet            `xml:"bullets>bullet"`
	Files               []File              `xml:"files>file"`
	QueriedAt           SanitizedDate       `xml:"queriedat" gorm:"type:datetime"` // To implement - date-time retreived.
	LocalFiles          []File              // Path used for local storage
}
