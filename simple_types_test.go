package main

import (
	"testing"
	"io/ioutil"
	"encoding/xml"
)

func PropertyReaderHelper(t *testing.T) *Property {
	file, err := ioutil.ReadFile("test_assets/api/branch/3741/property/26858499.xml")
	if err != nil {
		t.Fatalf("Error opening file: %s", err.Error())
	}
	prop := &Property{}
	if err := xml.Unmarshal(file, prop); err != nil {
		t.Fatalf("Error unmarshaling file: %s", err.Error())
	}
	return prop
}

func TestUnmarshalXMLUploadedValue(t *testing.T) {
	prop := PropertyReaderHelper(t)
	expected := "21/03/2017"
	actual   := prop.Uploaded.TimeValue().Format("02/01/2006")

	if actual != expected {
		t.Errorf("Parsed date %s does not match the expected date %s", actual, expected)
		t.Fail()
	}
}

func TestUnmarshalXMLFileUploadedValue(t *testing.T) {
	prop := PropertyReaderHelper(t)
	expected := "2017-03-21T13:39:33"
	actual   := prop.Files[0].Updated.TimeValue().Format("2006-01-02T15:04:05")

	if actual != expected {
		t.Errorf("Parsed date %s does not match the expected date %s", actual, expected)
		t.Fail()
	}
}

func TestUnmarshalXMLFileInstructedValue(t *testing.T) {
	prop := PropertyReaderHelper(t)
	expected := "2016-05-31"
	actual   := prop.Instructed.TimeValue().Format(`2006-01-02`)

	if actual != expected {
		t.Errorf("Parsed date %s does not match the expected date %s", actual, expected)
		t.Fail()
	}
}

func TestUnmarshalXMLRMTypeValue(t *testing.T) {
	prop := PropertyReaderHelper(t)
	expected := RMTypeFlat
	actual   := prop.RmType

	if actual != expected {
		t.Errorf("Parsed date %s does not match the expected date %s", actual, expected)
		t.Fail()
	}
}

func TestUnmarshalXMLRMFurnishedValue(t *testing.T) {
	prop := PropertyReaderHelper(t)
	expected := RMTypeFurnishedPartFurnished
	actual   := prop.Furnished

	if actual != expected {
		t.Errorf("Parsed date %s does not match the expected date %s", actual, expected)
		t.Fail()
	}
}

func TestUnmarshalXMLRMQualifierValue(t *testing.T) {
	prop := PropertyReaderHelper(t)
	expected := RMQualifierPriceOnApplication
	actual   := prop.RmQualifier

	if actual != expected {
		t.Errorf("Parsed date %s does not match the expected date %s", actual, expected)
		t.Fail()
	}
}

func TestUnmarshalXMLRMLetTypeValue(t *testing.T) {
	prop := PropertyReaderHelper(t)
	expected := RMLetTypeStudent
	actual   := prop.RmLetTypeID

	if actual != expected {
		t.Errorf("Parsed date %t does not match the expected date %t", actual, expected)
		t.Fail()
	}
}

func TestUnmarshalXMLRMWebStatus(t *testing.T) {
	prop := PropertyReaderHelper(t)
	expected := ForSaleOrToLetSSTCOrReserved
	actual   := prop.WebStatus

	if actual != expected {
		t.Errorf("Parsed date %t does not match the expected date %t", actual, expected)
		t.Fail()
	}
}
