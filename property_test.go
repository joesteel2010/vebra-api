// + build unit

package api

import (
	"testing"
	"time"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"encoding/xml"
	"log"
	"github.com/jinzhu/gorm"
	"os"
)

func TestURLGetBranchesBuilder(t *testing.T) {
	const expected = "http://webservices.vebra.com/export/ABCDEFG/v10/branch"

	branchesBuilder := new(URLGetBranchesBuilder)
	url := branchesBuilder.SetDataFeedID("ABCDEFG").Build()

	if url != expected {
		t.Errorf("Expected [%s] but found [%s]", expected, url)
	}
}

func TestURLGetBranchBuilder(t *testing.T) {
	const expected = "http://webservices.vebra.com/export/ABCDEFG/v10/branch/1234567"

	branchBuilder := new(URLGetBranchBuilder)
	url := branchBuilder.SetDataFeedID("ABCDEFG").SetClientID("1234567").Build()

	if url != expected {
		t.Errorf("Expected [%s] but found [%s]", expected, url)
	}
}

func TestURLGetChangedFiles(t *testing.T) {
	const expected = "http://webservices.vebra.com/export/ABCDEFG/v10/files/1989/05/25/00/00/00"
	dtime, err := time.Parse("02/01/2006", "25/05/1989")
	if err != nil {
		t.Error(err.Error())
	}
	changedFilesBuilder := new(URLGetChangedFilesBuilder)
	url := changedFilesBuilder.SetDataFeedID("ABCDEFG").SetSince(dtime).Build()

	if url != expected {
		t.Errorf("Expected [%s] but found [%s]", expected, url)
	}
}

func TestURLGetPropertiesBuilder(t *testing.T) {
	const expected = "http://webservices.vebra.com/export/ABCDEFG/v10/branch/1234567/property"

	propertiesBuilder := new(URLGetPropertiesBuilder)
	url := propertiesBuilder.SetDataFeedID("ABCDEFG").SetClientID("1234567").Build()

	if url != expected {
		t.Errorf("Expected [%s] but found [%s]", expected, url)
	}
}

func TestURLGetPropertyBuilder(t *testing.T) {
	const expected = "http://webservices.vebra.com/export/ABCDEFG/v10/branch/1234567/property/7654321"

	propertyBuilder := new(URLGetPropertyBuilder)
	url := propertyBuilder.SetDataFeedID("ABCDEFG").SetClientID("1234567").SetPropertyID("7654321").Build()

	if url != expected {
		t.Errorf("Expected [%s] but found [%s]", expected, url)
	}
}

func TestURLGetChangedProperties(t *testing.T) {
	const expected = "http://webservices.vebra.com/export/ABCDEFG/v10/property/1989/05/25/00/00/00"
	dtime, err := time.Parse("02/01/2006", "25/05/1989")
	if err != nil {
		t.Error(err.Error())
	}
	changedPropertiesBuilder := new(URLGetChangedPropertiesBuilder)
	url := changedPropertiesBuilder.SetDataFeedID("ABCDEFG").SetSince(dtime).Build()

	if url != expected {
		t.Errorf("Expected [%s] but found [%s]", expected, url)
	}
}

func ReadPropertiesHelper(t *testing.T) []Property {
	files, err := ioutil.ReadDir("test_assets/api/branch/3741/property/")
	if err != nil {
		t.Error(err.Error())
	}

	properties := make([]Property, len(files))
	for _, file := range (files) {
		file.Name()
		file, err := ioutil.ReadFile("test_assets/api/branch/3741/property/" + file.Name())
		if err != nil {
			t.Fatalf("Error opening file: %s", err.Error())
		}
		prop := &Property{}
		if err := xml.Unmarshal(file, prop); err != nil {
			t.Fatalf("Error unmarshaling file: %s", err.Error())
		}
		properties = append(properties, *prop)
	}
	return properties
}

func getDBConnectionHelper() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	log.Printf("Connecting to [%s]", dsn)

	db, err := gorm.Open("mysql", dsn)
	db = db.Debug()

	if err != nil {
		panic(err)
	}

	return db
}

func CreateTables(db *gorm.DB) error {
	prop := Property{}
	db.SingularTable(true)

	if err := db.AutoMigrate(&prop).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&Branch{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&prop.Address).AddForeignKey("property_id", "property(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&prop.AgentReference).AddForeignKey("property_id", "property(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&prop.Area).AddForeignKey("property_id", "property(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&Bullet{}).AddForeignKey("property_id", "property(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&File{}).AddForeignKey("property_id", "property(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&EnergyEfficiency{}).AddForeignKey("property_id", "property(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&EnvironmentalImpact{}).AddForeignKey("property_id", "property(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&prop.LandArea).AddForeignKey("property_id", "property(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&Paragraph{}).AddForeignKey("property_id", "property(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&prop.Price).AddForeignKey("property_id", "property(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&prop.StreetView).AddForeignKey("property_id", "property(id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}
	return nil
}

func TestSavePropertyIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	db := getDBConnectionHelper(t)
	CreateTables(db)

	for _, prop := range ReadPropertiesHelper(t) {

		for i := range prop.Files {
			prop.Files[i].FileID += 1
		}

		if err := db.Create(&prop).Error; err != nil {
			t.Error(err.Error())
		}
	}

	for _, prop := range ReadPropertiesHelper(t) {
		outProp := &Property{}
		db.First(outProp, prop.ID).Related(&outProp.AgentReference).
			Related(&outProp.Area).
			Related(&outProp.Price).
			Related(&outProp.Files).
			Related(&outProp.Bullets).
			Related(&outProp.Paragraphs).
			Related(&outProp.Files).
			Related(&outProp.Address).
			Related(&outProp.EnergyEfficiency).
			Related(&outProp.EnvironmentalImpact).
			Related(&outProp.LandArea)

		log.Println(prop.ID)
		errors := compareProperties(prop, *outProp)
		if errors != nil {
			for _, err := range errors {
				log.Println(err.Error())
			}
		}
	}
}

func compareProperties(a Property, b Property) []error {

	errors := make([]error, 0)

	if a.ID != b.ID {
		errors = append(errors, fmt.Errorf("property ID difference: [%s] [%s]", a.ID, b.ID))
	}

	if a.System != b.System {
		errors = append(errors, fmt.Errorf("property System difference: [%s] [%s]", a.System, b.System))
	}

	if a.Firmid != b.Firmid {
		errors = append(errors, fmt.Errorf("property Firmid difference: [%s] [%s]", a.Firmid, b.Firmid))
	}

	if a.Branchid != b.Branchid {
		errors = append(errors, fmt.Errorf("property Branchid difference: [%s] [%s]", a.Branchid, b.Branchid))
	}

	if a.Database != b.Database {
		errors = append(errors, fmt.Errorf("property Database difference: [%s] [%s]", a.Database, b.Database))
	}

	if a.Featured != b.Featured {
		errors = append(errors, fmt.Errorf("property Featured difference: [%s] [%s]", a.Featured, b.Featured))
	}

	if a.AgentReference.Agents != b.AgentReference.Agents {
		errors = append(errors, fmt.Errorf("property AgentReference Agents difference: [%s] [%s]", a.AgentReference.Agents, b.AgentReference.Agents))
	}

	if a.AgentReference.Software != b.AgentReference.Software {
		errors = append(errors, fmt.Errorf("property AgentReference Software difference: [%s] [%s]", a.AgentReference.Software, b.AgentReference.Software))
	}

	if a.Address.Name != b.Address.Name {
		errors = append(errors, fmt.Errorf("property Address Name difference: [%s] [%s]", a.Address.Name, b.Address.Name))
	}

	if a.Address.Street != b.Address.Street {
		errors = append(errors, fmt.Errorf("property Address Street difference: [%s] [%s]", a.Address.Street, b.Address.Street))
	}

	if a.Address.Locality != b.Address.Locality {
		errors = append(errors, fmt.Errorf("property Address Locality difference: [%s] [%s]", a.Address.Locality, b.Address.Locality))
	}

	if a.Address.Town != b.Address.Town {
		errors = append(errors, fmt.Errorf("property Address Town difference: [%s] [%s]", a.Address.Town, b.Address.Town))
	}

	if a.Address.County != b.Address.County {
		errors = append(errors, fmt.Errorf("property Address County difference: [%s] [%s]", a.Address.County, b.Address.County))
	}

	if a.Address.Postcode != b.Address.Postcode {
		errors = append(errors, fmt.Errorf("property Address Postcode difference: [%s] [%s]", a.Address.Postcode, b.Address.Postcode))
	}

	if a.Address.CustomLocation != b.Address.CustomLocation {
		errors = append(errors, fmt.Errorf("property Address CustomLocation difference: [%s] [%s]", a.Address.CustomLocation, b.Address.CustomLocation))
	}

	if a.Address.Display != b.Address.Display {
		errors = append(errors, fmt.Errorf("property Address Display difference: [%s] [%s]", a.Address.Display, b.Address.Display))
	}

	if a.Price.Qualifier != b.Price.Qualifier {
		errors = append(errors, fmt.Errorf("property Price Qualifier difference: [%s] [%s]", a.Price.Qualifier, b.Price.Qualifier))
	}

	if a.Price.Currency != b.Price.Currency {
		errors = append(errors, fmt.Errorf("property Price Currency difference: [%s] [%s]", a.Price.Currency, b.Price.Currency))
	}

	if a.Price.Display != b.Price.Display {
		errors = append(errors, fmt.Errorf("property Price Display difference: [%s] [%s]", a.Price.Display, b.Price.Display))
	}

	if a.Price.Rent != b.Price.Rent {
		errors = append(errors, fmt.Errorf("property Price Rent difference: [%s] [%s]", a.Price.Rent, b.Price.Rent))
	}

	if a.Price.Value != b.Price.Value {
		errors = append(errors, fmt.Errorf("property Price Value difference: [%s] [%s]", a.Price.Value, b.Price.Value))
	}

	if a.RentalFees != b.RentalFees {
		errors = append(errors, fmt.Errorf("property RentalFees difference: [%s] [%s]", a.RentalFees, b.RentalFees))
	}

	if a.LettingsFee != b.LettingsFee {
		errors = append(errors, fmt.Errorf("property LettingsFee difference: [%s] [%s]", a.LettingsFee, b.LettingsFee))
	}

	if a.RmQualifier != b.RmQualifier {
		errors = append(errors, fmt.Errorf("property RmQualifier difference: [%s] [%s]", a.RmQualifier, b.RmQualifier))
	}

	if a.Available == nil && b.Available != nil {
		if (! a.Available.Datetime.Equal(MySQLNullDate) && b.Available == nil ) ||
			(!b.Available.Datetime.Equal(MySQLNullDate) && a.Available == nil) ||
			(!a.Available.Datetime.Equal(VebraNullDate) && b.Available == nil) ||
			(!b.Available.Datetime.Equal(VebraNullDate) && a.Available == nil) {
			errors = append(errors, fmt.Errorf("property Available difference: [%s] [%s]", a.Available, b.Available))
		}
	}

	if a.Available != nil && b.Available == nil {
		if (! a.Available.Datetime.Equal(MySQLNullDate) && b.Available == nil ) ||
			(!b.Available.Datetime.Equal(MySQLNullDate) && a.Available == nil) ||
			(!a.Available.Datetime.Equal(VebraNullDate) && b.Available == nil) ||
			(!b.Available.Datetime.Equal(VebraNullDate) && a.Available == nil) {
			errors = append(errors, fmt.Errorf("property Available difference: [%s] [%s]", a.Available, b.Available))
		}
	}

	if a.Uploaded != b.Uploaded {
		errors = append(errors, fmt.Errorf("property Uploaded difference: [%s] [%s]", a.Uploaded, b.Uploaded))
	}

	if a.Longitude != b.Longitude {
		errors = append(errors, fmt.Errorf("property Longitude difference: [%s] [%s]", a.Longitude, b.Longitude))
	}

	if a.Latitude != b.Latitude {
		errors = append(errors, fmt.Errorf("property Latitude difference: [%s] [%s]", a.Latitude, b.Latitude))
	}

	if a.Easting != b.Easting {
		errors = append(errors, fmt.Errorf("property Easting difference: [%s] [%s]", a.Easting, b.Easting))
	}

	if a.Northing != b.Northing {
		errors = append(errors, fmt.Errorf("property Northing difference: [%s] [%s]", a.Northing, b.Northing))
	}

	if a.StreetView != b.StreetView {
		errors = append(errors, fmt.Errorf("property StreetView difference: [%s] [%s]", a.StreetView, b.StreetView))
	}

	if a.WebStatus != b.WebStatus {
		errors = append(errors, fmt.Errorf("property WebStatus difference: [%s] [%s]", a.WebStatus, b.WebStatus))
	}

	if a.CustomStatus != b.CustomStatus {
		errors = append(errors, fmt.Errorf("property CustomStatus difference: [%s] [%s]", a.CustomStatus, b.CustomStatus))
	}

	if a.CommRent != b.CommRent {
		errors = append(errors, fmt.Errorf("property CommRent difference: [%s] [%s]", a.CommRent, b.CommRent))
	}

	if a.Premium != b.Premium {
		errors = append(errors, fmt.Errorf("property Premium difference: [%s] [%s]", a.Premium, b.Premium))
	}

	if a.ServiceCharge != b.ServiceCharge {
		errors = append(errors, fmt.Errorf("property ServiceCharge difference: [%s] [%s]", a.ServiceCharge, b.ServiceCharge))
	}

	if a.RateableValue != b.RateableValue {
		errors = append(errors, fmt.Errorf("property RateableValue difference: [%s] [%s]", a.RateableValue, b.RateableValue))
	}

	if a.Type != b.Type {
		errors = append(errors, fmt.Errorf("property Type difference: [%s] [%s]", a.Type, b.Type))
	}

	if a.Furnished != b.Furnished {
		errors = append(errors, fmt.Errorf("property Furnished difference: [%s] [%s]", a.Furnished, b.Furnished))
	}

	if a.RmType != b.RmType {
		errors = append(errors, fmt.Errorf("property RmType difference: [%s] [%s]", a.RmType, b.RmType))
	}

	if a.LetBond != b.LetBond {
		errors = append(errors, fmt.Errorf("property LetBond difference: [%s] [%s]", a.LetBond, b.LetBond))
	}

	if a.RmLetTypeID != b.RmLetTypeID {
		errors = append(errors, fmt.Errorf("property RmLetTypeID difference: [%s] [%s]", a.RmLetTypeID, b.RmLetTypeID))
	}

	if a.Bedrooms != b.Bedrooms {
		errors = append(errors, fmt.Errorf("property Bedrooms difference: [%s] [%s]", a.Bedrooms, b.Bedrooms))
	}

	if a.Receptions != b.Receptions {
		errors = append(errors, fmt.Errorf("property Receptions difference: [%s] [%s]", a.Receptions, b.Receptions))
	}

	if a.Bathrooms != b.Bathrooms {
		errors = append(errors, fmt.Errorf("property Bathrooms difference: [%s] [%s]", a.Bathrooms, b.Bathrooms))
	}

	if a.UserField1 != b.UserField1 {
		errors = append(errors, fmt.Errorf("property UserField1 difference: [%s] [%s]", a.UserField1, b.UserField1))
	}

	if a.UserField2 != b.UserField2 {
		errors = append(errors, fmt.Errorf("property UserField2 difference: [%s] [%s]", a.UserField2, b.UserField2))
	}

	if a.SoldDate != b.SoldDate {
		errors = append(errors, fmt.Errorf("property SoldDate difference: [%s] [%s]", a.SoldDate, b.SoldDate))
	}

	if a.LeaseEnd != b.LeaseEnd {
		errors = append(errors, fmt.Errorf("property LeaseEnd difference: [%s] [%s]", a.LeaseEnd, b.LeaseEnd))
	}

	if a.Instructed != b.Instructed {
		errors = append(errors, fmt.Errorf("property Instructed difference: [%s] [%s]", a.Instructed, b.Instructed))
	}

	if a.SoldPrice != b.SoldPrice {
		errors = append(errors, fmt.Errorf("property SoldPrice difference: [%s] [%s]", a.SoldPrice, b.SoldPrice))
	}

	if a.Garden != b.Garden {
		errors = append(errors, fmt.Errorf("property Garden difference: [%s] [%s]", a.Garden, b.Garden))
	}

	if a.Parking != b.Parking {
		errors = append(errors, fmt.Errorf("property Parking difference: [%s] [%s]", a.Parking, b.Parking))
	}

	if a.NewBuild != b.NewBuild {
		errors = append(errors, fmt.Errorf("property NewBuild difference: [%s] [%s]", a.NewBuild, b.NewBuild))
	}

	if a.GroundRent != b.GroundRent {
		errors = append(errors, fmt.Errorf("property GroundRent difference: [%s] [%s]", a.GroundRent, b.GroundRent))
	}

	if a.Commission != b.Commission {
		errors = append(errors, fmt.Errorf("property Commission difference: [%s] [%s]", a.Commission, b.Commission))
	}

	if len(a.Area) != len(b.Area) {
		errors = append(errors, fmt.Errorf("property Area length difference: [%d] [%d]", len(a.Area), len(b.Area)))
	}

	if len(a.Area) == len(b.Area) {
		for index := range (a.Area) {
			if a.Area[index].Unit != b.Area[index].Unit {
				errors = append(errors, fmt.Errorf("property Area Unit difference: [%s] [%s]", a.Area[index].Unit, b.Area[index].Unit))
			}

			if a.Area[index].Min != b.Area[index].Min {
				errors = append(errors, fmt.Errorf("property Area Min difference: [%s] [%s]", a.Area[index].Min, b.Area[index].Min))
			}

			if a.Area[index].Max != b.Area[index].Max {
				errors = append(errors, fmt.Errorf("property Area Max difference: [%s] [%s]", a.Area[index].Max, b.Area[index].Max))
			}

		}
	}

	if a.LandArea.Unit != b.LandArea.Unit {
		errors = append(errors, fmt.Errorf("property LandArea Unit difference: [%s] [%s]", a.LandArea.Unit, b.LandArea.Unit))
	}

	if a.LandArea.Min != b.LandArea.Min {
		errors = append(errors, fmt.Errorf("property LandArea Min difference: [%s] [%s]", a.LandArea.Min, b.LandArea.Min))
	}

	if a.LandArea.Max != b.LandArea.Max {
		errors = append(errors, fmt.Errorf("property LandArea Max difference: [%s] [%s]", a.LandArea.Max, b.LandArea.Max))
	}

	if a.Description != b.Description {
		errors = append(errors, fmt.Errorf("property Description difference: [%s] [%s]", a.Description, b.Description))
	}

	if a.EnergyEfficiency.Current != b.EnergyEfficiency.Current {
		errors = append(errors, fmt.Errorf("property EnergyEfficiency Current difference: [%s] [%s]", a.EnergyEfficiency.Current, b.EnergyEfficiency.Current))
	}

	if a.EnergyEfficiency.Potential != b.EnergyEfficiency.Potential {
		errors = append(errors, fmt.Errorf("property EnergyEfficiency Potential difference: [%s] [%s]", a.EnergyEfficiency.Potential, b.EnergyEfficiency.Potential))
	}

	if a.EnvironmentalImpact.Current != b.EnvironmentalImpact.Current {
		errors = append(errors, fmt.Errorf("property EnvironmentalImpact Current difference: [%s] [%s]", a.EnvironmentalImpact.Current, b.EnvironmentalImpact.Current))
	}

	if a.EnvironmentalImpact.Potential != b.EnvironmentalImpact.Potential {
		errors = append(errors, fmt.Errorf("property EnvironmentalImpact Potential difference: [%s] [%s]", a.EnvironmentalImpact.Potential, b.EnvironmentalImpact.Potential))
	}

	if len(a.Paragraphs) != len(b.Paragraphs) {
		errors = append(errors, fmt.Errorf("property Paragraphs length difference: [%d] [%d]", len(a.Paragraphs), len(b.Paragraphs)))
	}

	if len(a.Paragraphs) == len(b.Paragraphs) {
		for index := range a.Paragraphs {
			if a.Paragraphs[index].ParagraphID != b.Paragraphs[index].ParagraphID {
				errors = append(errors, fmt.Errorf("property Paragraphs ParagraphID difference at index[%d]: [%s] [%s]", index, a.Paragraphs[index].ParagraphID, b.Paragraphs[index].ParagraphID))
			}

			if a.Paragraphs[index].Type != b.Paragraphs[index].Type {
				errors = append(errors, fmt.Errorf("property Paragraphs Type difference at index[%d]: [%s] [%s]", index, a.Paragraphs[index].Type, b.Paragraphs[index].Type))
			}

			if a.Paragraphs[index].Name != b.Paragraphs[index].Name {
				errors = append(errors, fmt.Errorf("property Paragraphs Name difference at index[%d]: [%s] [%s]", index, a.Paragraphs[index].Name, b.Paragraphs[index].Name))
			}

			if a.Paragraphs[index].File != b.Paragraphs[index].File {
				errors = append(errors, fmt.Errorf("property Paragraphs File difference at index[%d]: [%s] [%s]", index, a.Paragraphs[index].File, b.Paragraphs[index].File))
			}

			if a.Paragraphs[index].Metric != b.Paragraphs[index].Metric {
				errors = append(errors, fmt.Errorf("property Paragraphs Metric difference at index[%d]: [%s] [%s]", index, a.Paragraphs[index].Metric, b.Paragraphs[index].Metric))
			}

			if a.Paragraphs[index].Imperial != b.Paragraphs[index].Imperial {
				errors = append(errors, fmt.Errorf("property Paragraphs Imperial difference at index[%d]: [%s] [%s]", index, a.Paragraphs[index].Imperial, b.Paragraphs[index].Imperial))
			}

			if a.Paragraphs[index].Mixed != b.Paragraphs[index].Mixed {
				errors = append(errors, fmt.Errorf("property Paragraphs Mixed difference at index[%d]: [%s] [%s]", index, a.Paragraphs[index].Mixed, b.Paragraphs[index].Mixed))
			}

			if a.Paragraphs[index].Text != b.Paragraphs[index].Text {
				errors = append(errors, fmt.Errorf("property Paragraphs Text difference at index[%d]: [%s] [%s]", index, a.Paragraphs[index].Text, b.Paragraphs[index].Text))
			}
		}
	}

	if len(a.Bullets) != len(b.Bullets) {
		errors = append(errors, fmt.Errorf("property Bullets length difference: [%d] [%d]", len(a.Bullets), len(b.Bullets)))
	}

	if len(a.Bullets) == len(b.Bullets) {
		for index := range (a.Bullets) {
			if a.Bullets[index].BulletID != b.Bullets[index].BulletID {
				errors = append(errors, fmt.Errorf("property Bullets BulletID difference at index[%d]: [%s] [%s]", index, a.Bullets[index].BulletID, b.Bullets[index].BulletID))
			}

			if a.Bullets[index].Value != b.Bullets[index].Value {
				errors = append(errors, fmt.Errorf("property Bullets Value difference at index[%d]: [%s] [%s]", index, a.Bullets[index].Value, b.Bullets[index].Value))
			}
		}
	}

	if len(a.Files) != len(b.Files) {
		errors = append(errors, fmt.Errorf("property Files length difference: [%d] [%d]", len(a.Files), len(b.Files)))
	}

	if len(a.Files) == len(b.Files) {
		for index := range (a.Files) {
			if a.Files[index].FileID != b.Files[index].FileID {
				errors = append(errors, fmt.Errorf("property Files FileID difference at index[%d]: [%s] [%s]", index, a.Files[index].FileID, b.Files[index].FileID))
			}

			if a.Files[index].Type != b.Files[index].Type {
				errors = append(errors, fmt.Errorf("property Files Type difference at index[%d]: [%s] [%s]", index, a.Files[index].Type, b.Files[index].Type))
			}

			if a.Files[index].Name != b.Files[index].Name {
				errors = append(errors, fmt.Errorf("property Files Name difference at index[%d]: [%s] [%s]", index, a.Files[index].Name, b.Files[index].Name))
			}

			if a.Files[index].Url != b.Files[index].Url {
				errors = append(errors, fmt.Errorf("property Files Url difference at index[%d]: [%s] [%s]", index, a.Files[index].Url, b.Files[index].Url))
			}

			if a.Files[index].Updated != b.Files[index].Updated {
				errors = append(errors, fmt.Errorf("property Files Updated difference at index[%d]: [%s] [%s]", index, a.Files[index].Updated, b.Files[index].Updated))
			}
		}
	}
	return errors
}
