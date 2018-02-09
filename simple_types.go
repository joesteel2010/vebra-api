package model

import (
	"time"
	"strings"
	"fmt"
	"strconv"
	"encoding/xml"
	"database/sql/driver"
)

type SanitizedInt int

var MySQLNullDate, _ = time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:01")
var VebraNullDate, _ = time.Parse("2006-01-02 15:04:05", "1900-01-01 00:00:00")

func (u SanitizedInt) Value() (driver.Value, error) { return int64(u), nil }

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

func (si *SanitizedInt) AsInt() string {
	n := int(*si)

	in := strconv.FormatInt(int64(n), 10)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
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

type SanitizedBool bool

func (u SanitizedBool) Value() (driver.Value, error) {
	if bool(u) {
		return int64(1), nil
	}
	return int64(0), nil
}

func (u *SanitizedBool) Scan(value interface{}) error {
	var val int
	var err error

	switch value.(type) {
	case []uint8:
		s := value.([]uint8)
		val, err = strconv.Atoi(string(s))
	case int64:
		s := value.(int64)
		val = int(s)
	}

	if err != nil {
		return fmt.Errorf("Error converting to sanitized bool: %s", err)
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

const (
	UKDateFormat string = `02/01/2006`
	ISODate      string = `2006-01-02`
	ISODateTime  string = `2006-01-02T15:04:05`
)

type SanitizedDate interface {
	Scan(value interface{}) error
	Value() (driver.Value, error)
	TimeValue() time.Time
	UnmarshalXML(d *xml.Decoder, start xml.StartElement)
}

type SanitizedDateType struct {
	Datetime time.Time
	DbLayout string "02/01/2006"
	Layout   string "2006-01-02 15:04:05"
}

func (u *SanitizedDateType) Scan(value interface{}) error {
	switch value.(type) {
	case time.Time:
		u.Datetime = value.(time.Time)
		return nil
	case []uint8:
		str := string(value.([]uint8))
		t, err := time.Parse(u.DbLayout, str)

		if err != nil {
			fmt.Errorf("Error scanning SanitizedDateTimeType: %s", err)
		}
		u.Datetime = t
		return nil
	default:
		fmt.Errorf("Error: cant handle type %T in SanitizedDateTimeType.Scan", value)
		return nil
	}
}

func (u *SanitizedDateType) Value() (driver.Value, error) {
	if u == nil {
		return nil, nil
	}
	if u.Datetime.Before(MySQLNullDate) {
		return nil, nil
	}
	return u.Datetime, nil
}

func (u *SanitizedDateType) TimeValue() time.Time { return u.Datetime }

type SanitizedDateTimeType struct {
	Datetime time.Time
	DbLayout string "02/01/2006"
	Layout   string "2006-01-02 15:04:05"
}

func (u *SanitizedDateTimeType) Scan(value interface{}) error {
	switch value.(type) {
	case time.Time:
		u.Datetime = value.(time.Time)
		return nil
	case []uint8:
		str := string(value.([]uint8))
		t, err := time.Parse(u.DbLayout, str)

		if err != nil {
			fmt.Errorf("Error scanning SanitizedDateTimeType: %s", err)
		}
		u.Datetime = t
		return nil
	default:
		fmt.Errorf("Error: cant handle type %T in SanitizedDateTimeType.Scan", value)
		return nil
	}
}

func (u *SanitizedDateTimeType) Value() (driver.Value, error) {
	if u == nil {
		return nil, nil
	}
	if u.Datetime.Before(MySQLNullDate) {
		return nil, nil
	}
	return u.Datetime, nil
}

func (u *SanitizedDateTimeType) TimeValue() time.Time { return u.Datetime }

type SanitizedDateUKDateFormat struct {
	SanitizedDateTimeType
}

func (u *SanitizedDateUKDateFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var value string
	var tmpTime time.Time
	d.DecodeElement(&value, &start)

	if value == "" {
		u.Datetime = time.Now()
		return nil
	}

	if tmpTime, err = time.Parse(UKDateFormat, value); err == nil {
		u.Datetime = tmpTime
		return nil
	}

	u.Datetime = time.Now()
	return nil
}

func (u *SanitizedDateUKDateFormat) Scan(value interface{}) error {
	switch value.(type) {
	case time.Time:
		u.Datetime = value.(time.Time)
		return nil
	case []uint8:
		str := string(value.([]uint8))
		t, err := time.Parse(u.DbLayout, str)

		if err != nil {
			fmt.Errorf("Error scanning SanitizedDateTimeType: %s", err)
		}
		u.Datetime = t
		return nil
	default:
		fmt.Errorf("Error: cant handle type %T in SanitizedDateTimeType.Scan", value)
		return nil
	}
}

func (u *SanitizedDateUKDateFormat) Value() (driver.Value, error) {
	if u == nil {
		return nil, nil
	}
	if u.Datetime.Before(MySQLNullDate) {
		return nil, nil
	}
	return u.Datetime, nil
}

func (u *SanitizedDateUKDateFormat) TimeValue() time.Time { return u.Datetime }

type SanitizedDateISODate struct {
	SanitizedDateType
}

func (u *SanitizedDateISODate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var value string
	var tmpTime time.Time
	d.DecodeElement(&value, &start)

	if value == "" {
		u.Datetime = time.Now()
		return nil
	}

	if tmpTime, err = time.Parse(ISODate, value); err == nil {
		u.Datetime = tmpTime
		return nil
	}

	u.Datetime = time.Now()
	return nil
}

func (u *SanitizedDateISODate) Value() (driver.Value, error) {
	if u == nil {
		return nil, nil
	}
	if u.Datetime.Before(MySQLNullDate) {
		return nil, nil
	}
	return u.Datetime, nil
}

type SanitizedDateISODateTime struct {
	SanitizedDateTimeType
}

func (u *SanitizedDateISODateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var value string
	var tmpTime time.Time
	d.DecodeElement(&value, &start)

	if value == "" {
		u.Datetime = time.Now()
		return nil
	}

	if tmpTime, err = time.Parse(ISODateTime, value); err == nil {
		u.Datetime = tmpTime
		return nil
	}

	u.Datetime = time.Now()
	return nil
}

func (u *SanitizedDateISODateTime) Scan(value interface{}) error {
	switch value.(type) {
	case time.Time:
		u.Datetime = value.(time.Time)
		return nil
	case []uint8:
		str := string(value.([]uint8))
		t, err := time.Parse(u.DbLayout, str)

		if err != nil {
			fmt.Errorf("Error scanning SanitizedDateTimeType: %s", err)
		}
		u.Datetime = t
		return nil
	default:
		fmt.Errorf("Error: cant handle type %T in SanitizedDateTimeType.Scan", value)
		return nil
	}
}

func (u *SanitizedDateISODateTime) Value() (driver.Value, error) {
	if u == nil {
		return nil, nil
	}
	if u.Datetime.Before(MySQLNullDate) {
		return nil, nil
	}
	return u.Datetime, nil
}

func (u *SanitizedDateISODateTime) TimeValue() time.Time { return u.Datetime }