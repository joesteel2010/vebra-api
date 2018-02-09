package model

import (
	"time"
	"strings"
	"regexp"
	"fmt"
	"strconv"
)

const (
	BaseURL                 = "http://webservices.vebra.com/"
	URLGetBranches          = BaseURL + "export/{datafeedid}/v10/branch"
	URLGetBranch            = BaseURL + "export/{datafeedid}/v10/branch/{clientid}"
	URLGetChangedFiles      = BaseURL + "export/{datafeedid}/v10/files/{yyyy}/{MM}/{dd}/{HH}/{mm}/{ss}"
	URLGetProperties        = BaseURL + "export/{datafeedid}/v10/branch/{clientid}/property"
	URLGetProperty          = BaseURL + "export/{datafeedid}/v10/branch/{clientid}/property/{prop_id}"
	URLGetChangedProperties = BaseURL + "export/{datafeedid}/v10/property/{yyyy}/{MM}/{dd}/{HH}/{mm}/{ss}"
)

type URLBuilder interface {
	Build() string
}

type vebraURLBuilder struct {
	url        string
	dataFeedID string
}

type URLGetBranchesBuilder struct {
	vebraURLBuilder
}

func (b *URLGetBranchesBuilder) SetDataFeedID(dataFeedID string) *URLGetBranchesBuilder {
	b.dataFeedID = dataFeedID
	return b
}

func (b *URLGetBranchesBuilder) Build() string {
	b.url = URLGetBranches
	return strings.Replace(b.url, "{datafeedid}", b.dataFeedID, 1)
}

type URLGetBranchBuilder struct {
	vebraURLBuilder
	clientID string
}

func (b *URLGetBranchBuilder) SetDataFeedID(dataFeedID string) *URLGetBranchBuilder {
	b.dataFeedID = dataFeedID
	return b
}

func (b *URLGetBranchBuilder) SetClientID(clientID string) *URLGetBranchBuilder {
	b.clientID = clientID
	return b
}

func (b *URLGetBranchBuilder) Build() string {
	b.url = URLGetBranch
	b.url = strings.Replace(b.url, "{datafeedid}", b.dataFeedID, 1)
	b.url = strings.Replace(b.url, "{clientid}", b.clientID, 1)
	return b.url
}

type URLGetChangedFilesBuilder struct {
	vebraURLBuilder
	since time.Time
}

func (b *URLGetChangedFilesBuilder) SetDataFeedID(dataFeedID string) *URLGetChangedFilesBuilder {
	b.dataFeedID = dataFeedID
	return b
}

func (b *URLGetChangedFilesBuilder) SetSince(time time.Time) *URLGetChangedFilesBuilder {
	b.since = time
	return b
}

func (b *URLGetChangedFilesBuilder) Build() string {
	b.url = URLGetChangedFiles
	b.url = strings.Replace(b.url, "{datafeedid}", b.dataFeedID, 1)
	b.url = strings.Replace(b.url, "{yyyy}", fmt.Sprintf("%04d", b.since.Year()), 1)
	b.url = strings.Replace(b.url, "{MM}", fmt.Sprintf("%02d", int(b.since.Month())), 1)
	b.url = strings.Replace(b.url, "{dd}", fmt.Sprintf("%02d", int(b.since.Day())), 1)
	b.url = strings.Replace(b.url, "{HH}", fmt.Sprintf("%02d", int(b.since.Hour())), 1)
	b.url = strings.Replace(b.url, "{mm}", fmt.Sprintf("%02d", int(b.since.Minute())), 1)
	b.url = strings.Replace(b.url, "{ss}", fmt.Sprintf("%02d", int(b.since.Second())), 1)
	return b.url
}

type URLGetPropertiesBuilder struct {
	vebraURLBuilder
	clientID string
}

func (b *URLGetPropertiesBuilder) SetDataFeedID(dataFeedID string) *URLGetPropertiesBuilder {
	b.dataFeedID = dataFeedID
	return b
}

func (b *URLGetPropertiesBuilder) SetClientID(clientID string) *URLGetPropertiesBuilder {
	b.clientID = clientID
	return b
}

func (b *URLGetPropertiesBuilder) Build() string {
	b.url = URLGetProperties
	b.url = strings.Replace(b.url, "{datafeedid}", b.dataFeedID, 1)
	b.url = strings.Replace(b.url, "{clientid}", b.clientID, 1)
	return b.url
}

type URLGetPropertyBuilder struct {
	URLGetPropertiesBuilder
	propertyID string
}

func (b *URLGetPropertyBuilder) SetDataFeedID(dataFeedID string) *URLGetPropertyBuilder {
	b.dataFeedID = dataFeedID
	return b
}

func (b *URLGetPropertyBuilder) SetClientID(clientID string) *URLGetPropertyBuilder {
	b.clientID = clientID
	return b
}

func (b *URLGetPropertyBuilder) SetPropertyID(propertyID string) *URLGetPropertyBuilder {
	b.propertyID = propertyID
	return b
}

func (b *URLGetPropertyBuilder) Build() string {
	b.url = URLGetProperty
	b.url = strings.Replace(b.url, "{datafeedid}", b.dataFeedID, 1)
	b.url = strings.Replace(b.url, "{clientid}", b.clientID, 1)
	b.url = strings.Replace(b.url, "{prop_id}", b.propertyID, 1)
	return b.url
}

type URLGetChangedPropertiesBuilder struct {
	vebraURLBuilder
	since time.Time
}

func (b *URLGetChangedPropertiesBuilder) SetDataFeedID(dataFeedID string) *URLGetChangedPropertiesBuilder {
	b.dataFeedID = dataFeedID
	return b
}

func (b *URLGetChangedPropertiesBuilder) SetSince(time time.Time) *URLGetChangedPropertiesBuilder {
	b.since = time
	return b
}

func (b *URLGetChangedPropertiesBuilder) Build() string {
	b.url = URLGetChangedProperties
	b.url = strings.Replace(b.url, "{datafeedid}", b.dataFeedID, 1)
	b.url = strings.Replace(b.url, "{yyyy}", fmt.Sprintf("%04d", b.since.Year()), 1)
	b.url = strings.Replace(b.url, "{MM}", fmt.Sprintf("%02d", int(b.since.Month())), 1)
	b.url = strings.Replace(b.url, "{dd}", fmt.Sprintf("%02d", int(b.since.Day())), 1)
	b.url = strings.Replace(b.url, "{HH}", fmt.Sprintf("%02d", int(b.since.Hour())), 1)
	b.url = strings.Replace(b.url, "{mm}", fmt.Sprintf("%02d", int(b.since.Minute())), 1)
	b.url = strings.Replace(b.url, "{ss}", fmt.Sprintf("%02d", int(b.since.Second())), 1)
	return b.url
}

// BranchSummaries list of BranchSummary
// Returned by the API call http://webservices.vebra.com/export/{datafeedid}/v10/branch List of branches for a Search Group.
type BranchSummaries struct {
	Branches []BranchSummary `xml:"branch"`
}

// BranchSummary Returned by the API call http://webservices.vebra.com/export/{datafeedid}/v10/branch List of branches for a Search Group.
// Contains:
// Name: 		Branch name
// FirmID: 		PSG Firm Identifier
// BranchID: 	PSG Branch Identifier for this Firm.
// Url:			REST Uri for the full branch details
type BranchSummary struct {
	Name     string `xml:"name"`
	FirmID   int    `xml:"firmid"`
	BranchID int    `xml:"branchid"`
	Url      string `xml:"url"`
}

// GetClientID returns the 4 digit client ID for the branch
func (bs BranchSummary) GetClientID() (int, error) {
	index := strings.LastIndex(bs.Url, "/")
	return strconv.Atoi(bs.Url[(index + 1):])
}

// Branch Returned by the API call http://webservices.vebra.com/export/{datafeedid}/v10/branch/{clientid} Client address details.
// Contains:
// ClientID:	ClientID used in subsiquent API calls
// FirmID: 		PSG Firm Identifier
// BranchID: 	PSG Branch Identifier for this Firm.
// Name:		Branch name
// Street:		Branch street
// Town:		Branch town
// County:		Branch county
// Postcode:	Branch postcode
// Phone:		Branch telephone number
// Email:		Branch email
// QueriedAt:	Branch street
type Branch struct {
	ClientID int    `xml:"clientid"`
	FirmID   int    `xml:"FirmID"`
	BranchID int    `xml:"BranchID"`
	Name     string `xml:"name"`
	URL      string `xml:"url"`
	Street   string `xml:"street"`
	Town     string `xml:"town"`
	County   string `xml:"county"`
	Postcode string `xml:"postcode"`
	Phone    string `xml:"phone"`
	Email    string `xml:"email"`
}

type PropertySummaries struct {
	Properties []PropertySummary `xml:"property"`
}

// PropertySummary Returned by the API call http://webservices.vebra.com/export/{datafeedid}/v10/branch/{clientid}/property List of properties in a branch.
// Contains:
// PropertyID: Unique identifier for this property.
// LastChanged: Last changed Datetime for this property. ISO Date Time format - YYYY-MM-DDTHH:MI:SS
// Url: REST Uri for the full details of the this property.
type PropertySummary struct {
	PropertyID  uint   `xml:"prop_id" gorm:"primary_key" sql:"type:int"`
	LastChanged string `xml:"lastchanged"`
	Url         string `xml:"url"`
}

const RentalPeriodPattern string = "pw|PW|pcm|PCM|pq|pa"

type RentalPeriod string

type RMType SanitizedInt

const (
	RMTypeNotSpecified                   RMType = iota
	RMTypeTerracedHouse
	RMTypeEndOfTerraceHouse
	RMTypeSemidetachedHouse
	RMTypeDetachedHouse
	RMTypeMewsHouse
	RMTypeClusterHouse
	RMTypeGroundFloorFlat
	RMTypeFlat
	RMTypeStudioFlat
	RMTypeGroundFloorMaisonette
	RMTypeMaisonette
	RMTypeBungalow
	RMTypeTerracedBungalow
	RMTypeSemidetachedBungalow
	RMTypeDetachedBungalow
	RMTypeMobileHome
	RMTypeLandResidential
	RMTypeLinkDetachedHouse
	RMTypeTownHouse
	RMTypeCottage
	RMTypeChalet
	RMTypeCharacterProperty
	RMTypeHouseUnspecified
	RMTypeVilla
	RMTypeApartment
	RMTypePenthouse
	RMTypeFinca
	RMTypeBarnConversion
	RMTypeServicedApartment
	RMTypeParking
	RMTypeShelteredHousing
	RMTypeReteirmentProperty
	RMTypeHouseShare
	RMTypeFlatShare
	RMTypeParkHome
	RMTypeGarages
	RMTypeFarmHouse
	RMTypeEquestrianFacility
	RMTypeDuplex
	RMTypeTriplex
	RMTypeLongere
	RMTypeGite
	RMTypeBarn
	RMTypeTrulli
	RMTypeMill
	RMTypeRuins
	RMTypeRestaurant
	RMTypeCafe
	RMTypeMillII
	RMTypeCastle
	RMTypeVillageHouse
	RMTypeCaveHouse
	RMTypeCortijo
	RMTypeFarmLand
	RMTypePlot
	RMTypeCountryHouse
	RMTypeStoneHouse
	RMTypeCaravan
	RMTypeLodge
	RMTypeLogCabin
	RMTypeManorHouse
	RMTypeStatelyHome
	RMTypeOffPlan
	RMTypeSemidetachedVilla
	RMTypeDetachedVilla
	RMTypeBarNightclub
	RMTypeShop
	RMTypeRiad
	RMTypeHouseBoat
	RMTypeHotelRoom
	RMTypeBlockOfApartments
	RMTypePrivateHalls
	RMTypeOffice
	RMTypeBusinessPark
	RMTypeServicedOffice
	RMTypeRetailPropertyHighStreet
	RMTypeRetailPropertyOutOfTown
	RMTypeConvenienceStore
	RMTypeGaragesII
	RMTypeHairdresserBarberShop
	RMTypeHotel
	RMTypePetrolStation
	RMTypePostOffice
	RMTypePub
	RMTypeWorkshopAndRetailSpace
	RMTypeDistributionWarehouse
	RMTypeFactory
	RMTypeHeavyIndustrial
	RMTypeIndustrialPark
	RMTypeLightIndustrial
	RMTypeStorage
	RMTypeShowroom
	RMTypeWarehouse
	RMTypeLandCommercial
	RMTypeCommercialDevelopment
	RMTypeIndustrialDevelopment
	RMTypeResidentialDevelopment
	RMTypeCommercialProperty
	RMTypeDataCentre
	RMTypeFarm
	RMTypeHealthcareFacility
	RMTypeMarineProperty
	RMTypeMixedUse
	RMTypeResearchAndDevelopmentFacility
	RMTypeSciencePark
	RMTypeGuestHouse
	RMTypeHospitality
	RMTypeLeisureFacility
	RMTypeTakeaway
	RMTypeChildcareFacility
	RMTypeSmallholding
	RMTypePlaceOfWorship
	RMTypeTradeCounter
	RMTypeCoachHouse
	RMTypeHouseOfMultipleOccupation
	RMTypeSportsFacilities
	RMTypeSpa
	RMTypeCampsiteAndHolidayVillage
)

type RMTypeFurnished SanitizedInt

const (
	RMTypeFurnishedFurnished            RMTypeFurnished = iota
	RMTypeFurnishedPartFurnished
	RMTypeFurnishedUnFurnished
	RMTypeFurnishedNotSpecified
	RMTypeFurnishedFurnishedUnFurnished
	RMTypeFurnishedNotUsed              RMTypeFurnished = iota + 3
	RMTypeFurnishedNotUsedII            RMTypeFurnished = iota + 6
	RMTypeFurnishedNotUsedIII           RMTypeFurnished = iota + 16
	RMTypeFurnishedNotUsedIIII          RMTypeFurnished = iota + 18
)

type RMTypeLetType SanitizedInt

const (
	RMLetTypeNotSpecified RMTypeLetType = iota
	RMLetTypeLongTerm
	RMLetTypeShortTerm
	RMLetTypeStudent
	RMLetTypeCommercial
)

type RMQualifier SanitizedInt

const (
	RMQualifierDefault            RMQualifier = iota
	RMQualifierPriceOnApplication
	RMQualifierGuidePrice
	RMQualifierFixedPrice
	RMQualifierOffersInExcessOf
	RMQualifierOffersInRegionOf
	RMQualifierSaleByTender
	RMQualifierFrom
	RMQualifierNotUsed
	RMQualifierSharedOwnership
	RMQualifierOffersOver
	RMQualifierPartTimeBuyRent
	RMQualifierSharedEquality
	RMQualifierComingSoon         RMQualifier = iota + 3
)

type PropertyStatus SanitizedInt

// Sale or let types
const (
	ForSaleOrToLet                            PropertyStatus = iota
	ForSaleOrToLetUnderOfferOrLet
	ForSaleOrToLetSoldOrUnderOffer
	ForSaleOrToLetSSTCOrReserved
	ForSaleOrToLetForSaleByAuctionOrLetAgreed
	ForSaleOrToLetReserved
	ForSaleOrToLetNewInstruction
	ForSaleOrToLetJustOnMarket
	ForSaleOrToLetPriceReduction
	ForSaleOrToLetKeenToSell
	ForSaleOrToLetNoChain
	ForSaleOrToLetVendorWillPayStampDuty
	ForSaleOrToLetOffersInRegionOf
	ForSaleOrToLetGuidePrice
)

//Let types
const (
	LetingsToLet      PropertyStatus = iota + 100
	LetingsLet
	LetingsUnderOffer
	LetingsReserved
	LettingsLetAgreed
)

// Hidden Properties
const (
	NotMarketed                      PropertyStatus = iota + 200
	NotMarketedUnderOffer
	NotMarketedSold
	NotMarketedSoldSubjectToContract
	NotMarketedLet                   PropertyStatus = iota + 200 + 10
	NotMarketedII                    PropertyStatus = iota + 200 + 50
)

// Properties is a collection of the type Property
type Properties struct {
	properties []Property `xml:"property"`
}

// Property is the main data type returned from the web service. It is the parent
// type for all other types listed in this file
// Contains:
// ID: Our Property Identifier
// System: Deprecated.
// Firmid: Our firm (company) identifer (FirmId).
// Branchid: Our branch identifer ( BranchId).
// Database: See dbids.xsd
// Featured: The property 'is featured' flag.
// AgentReference: See AgentReference
// Address: See Address
// Price: See Price
// RentalFees: This field is to comply with CAP legislation.
// LettingsFee: This field is to comply with CAP legislation.
// RmQualifier: RightMove qualifier. See rightmovetypes.xsd.
// Available: Date the property is available in format “dd/mm/yyyy”. Used mainly for rental properties. If value held is NULL, is populated with 01/01/1900.
// Uploaded: Date the property was uploaded in format “dd/mm/yyyy”.
// Longitude: The exact longitude of the property.
// Latitude: The exact latitude of the property.
// Easting: Easting (Cartesian coordinates)
// Northing: Northing (Cartesian coordinates)
// StreetView: See StreetView
// WebStatus: The status of the property. See propertyrelatedtypes.xsd.
// CustomStatus: Relates to Commercial properties.
// CommRent: Commercial rental amount. Listed rental price plus a qualifier which is configurable by the agent, e.g. per sq ft. Now DEPRECATED - please use price node.
// Premium: Premium for property, if supplied. Relates to Commerical properties.
// ServiceCharge: The service charge for the property, if supplied.
// RateableValue: The rateable value for the property, if supplied.
// Type: Apartment, House, (Not Specified) or the agent's own type. May not always contain data.
// Furnished: A value giving the furnished status of the property. See propertyrelatedtypes.xsd.
// RmType: This is property_type as per the Rightmove Real-time specification.
// LetBond: Lettings Bond (deposit).
// RmLetTypeID: RightMove Type for property. See rightmovetypes.xsd.
// Bedrooms: No of Bedrooms for the property.
// Receptions: No of Reception Rooms for the property.
// Bathrooms: No of Bathrooms for the property. NB: we advise if this value is 0 to regard this as non-live data.
// UserField1: User defined field used by vebra agents searching.
// UserField2: User defined field used by vebra agents searching.
// SoldDate: date the property was sold in format "yyyy-mm-dd"
// LeaseEnd: lease end date in format "yyyy-mm-dd"
// Instructed: Date the instruction was received in format “yyyy-mm-dd”.
// SoldPrice: Sold price
// Garden: Garden
// Parking: parking
// NewBuild: New build
// GroundRent: Ground Rent
// Commission: Commission as defined by the agent.
// Area: See Area
// LandArea: See LandArea
// Description: The description of the property.
// EnergyEfficiency: See EnergyEfficiency
// EnvironmentalImpact: See EnvironmentalImpact
// Paragraphs: See Paragraph
// Bullets: See Bullet
// Files: See File
type Property struct {
	ID                  uint                       `xml:"id,attr" gorm:"primary_key" sql:"type:int(10) unsigned`
	System              string                     `xml:"system,attr"`
	Firmid              string                     `xml:"firmid,attr"`
	Branchid            int                        `xml:"branchid,attr"`
	Database            int                        `xml:"database,attr"`
	Featured            int                        `xml:"featured,attr"`
	AgentReference      Reference                  `xml:"reference"`
	Address             Address                    `xml:"address"`
	Price               Price                      `xml:"price"`
	RentalFees          string                     `xml:"rentalfees"`
	LettingsFee         string                     `xml:"lettingsfee"`
	RmQualifier         RMQualifier                `xml:"rm_qualifier"`
	Available           *SanitizedDateUKDateFormat `xml:"available"`
	Uploaded            *SanitizedDateUKDateFormat `xml:"uploaded" gorm:"type:Datetime"`
	Longitude           float32                    `xml:"longitude"`
	Latitude            float32                    `xml:"latitude"`
	Easting             SanitizedInt               `xml:"easting"`
	Northing            SanitizedInt               `xml:"northing"`
	StreetView          StreetView                 `xml:"streetview"`
	WebStatus           PropertyStatus             `xml:"web_status"`
	CustomStatus        string                     `xml:"custom_status"`
	CommRent            string                     `xml:"comm_rent"`
	Premium             string                     `xml:"premium"`
	ServiceCharge       string                     `xml:"service_charge"`
	RateableValue       string                     `xml:"rateable_value"`
	Type                string                     `xml:"type"`
	Furnished           RMTypeFurnished            `xml:"furnished"`
	RmType              RMType                     `xml:"rm_type"`
	LetBond             SanitizedInt               `xml:"let_bond"`
	RmLetTypeID         RMTypeLetType              `xml:"rm_let_type_id"`
	Bedrooms            SanitizedInt               `xml:"bedrooms"`
	Receptions          SanitizedInt               `xml:"receptions"`
	Bathrooms           SanitizedInt               `xml:"bathrooms"`
	UserField1          string                     `xml:"userfield1"`
	UserField2          SanitizedInt               `xml:"userfield2"`
	SoldDate            *SanitizedDateISODate      `xml:"solddate" gorm:"type:Datetime"`
	LeaseEnd            *SanitizedDateISODate      `xml:"leaseend" gorm:"type:Datetime"`
	Instructed          *SanitizedDateISODate      `xml:"instructed" gorm:"type:Datetime"`
	SoldPrice           SanitizedInt               `xml:"soldprice"`
	Garden              SanitizedBool              `xml:"garden"`
	Parking             SanitizedBool              `xml:"parking"`
	NewBuild            SanitizedBool              `xml:"newbuild"`
	GroundRent          string                     `xml:"groundrent"`
	Commission          string                     `xml:"commission"`
	Area                []Area                     `xml:"area"`
	LandArea            LandArea                   `xml:"landarea"`
	Description         string                     `xml:"description" gorm:"type:varchar(2056)"`
	EnergyEfficiency    EnergyEfficiency           `xml:"hip>energy_performance>energy_efficiency"`
	EnvironmentalImpact EnvironmentalImpact        `xml:"hip>energy_performance>environmental_impact"`
	Paragraphs          []Paragraph                `xml:"paragraphs>paragraph"`
	Bullets             []Bullet                   `xml:"bullets>bullet"`
	Files               []File                     `xml:"files>file"`
}

type ChangedPropertySummaries struct {
	PropertySummaries []ChangedPropertySummary `xml:"property"`
}

// ChangedPropertySummary
// Returned by the API call
// http://webservices.vebra.com/export/{datafeedid}/v10/property/{yyyy}/{MM}/{dd}/{HH}/{mm}/{ss}
// List of properties updated in a search group.
// Contains:
// PropertyID: Unique identifier for this property
// LastChanged: Last changed Datetime for this property. ISO Date Time format - YYYY-MM-DDTHH:MI:SS
// LastAction Action field to indicate whether this property was updated or deleted. Possible values: ["updated", "deleted"]
type ChangedPropertySummary struct {
	PropertyID  uint   `xml:"propid"`
	LastChanged string `xml:"lastchanged"`
	Url         string `xml:"url"`
	LastAction  string `xml:"action"`
}

func (cps *ChangedPropertySummary) SetLastAction(lastAction string) {
	cps.LastAction = lastAction
}

func (cps *ChangedPropertySummary) GetLastAction() string {
	return cps.LastAction
}

// GetClientID returns the 4 digit client ID for the branch
func (cps *ChangedPropertySummary) GetClientID() (int, error) {
	re := regexp.MustCompile(`branch\/(\d+)/`)
	matches := re.FindStringSubmatch(cps.Url)

	if matches == nil || len(matches) < 2 {
		return 0, fmt.Errorf("couldnt match client ID in URL: [%s]", cps.Url)
	}

	return strconv.Atoi(matches[1])
}

// Reference This is the agents reference and can be displayed on an agent's search.
// Rightmove use this as part of property reference.
type Reference struct {
	PropertyID uint `gorm:"primary_key" sql:"type:int(10) unsigned"`
	Agents     int  `xml:"agents"`
	Software   int  `xml:"software"`
}

// Address represents the property address
// Contains:
// PropertyID: Unique ID for the property
// Name: Property house No / name.
// Street: Property street, not including house No or  name.
// Locality: Property locality.
// Town: Property town
// County: Property county.
// Postcode: Property postcode (part or full) if available.
// CustomLocation: Agent Location
// Display: The address to display on the website. If this is not supplied we will display street, town, county
//			from the address above.
type Address struct {
	PropertyID     uint   `gorm:"primary_key" sql:"type:int(10) unsigned"`
	Name           string `xml:"name"`
	Street         string `xml:"street"`
	Locality       string `xml:"locality"`
	Town           string `xml:"town"`
	County         string `xml:"county"`
	Postcode       string `xml:"postcode"`
	CustomLocation string `xml:"custom_locatiom"`
	Display        string `xml:"display"`
}

// Price represents the price of the property
// Contains:
// Currency: The currency that the property price is denominated in. See currency.xsd
// Qualifier: ???
// Display: If display="no" the price should not be displayed on the website (i.e. this property is Price On Application or similar)
// Rent: Rental Period, e.g. PCM (per calendar month). Present if property is rented.
// 		 Possible values are: pw|PW|pcm|PCM|pq|pa (Is this a per week(pw), per month(pcm), per quarter (pq) or per annum (pa) rental)
// Value: Property value (in given unit of currency)
type Price struct {
	PropertyID uint         `gorm:"primary_key" sql:"type:int(10) unsigned"`
	Qualifier  string       `xml:"qualifier,attr"`
	Currency   string       `xml:"currency,attr"`
	Display    string       `xml:"display,attr"`
	Rent       string       `xml:"rent,attr"`
	Value      SanitizedInt `xml:",chardata"`
}

// StreetView describes the longitude, latitude, yaw, pitch and zoom for the property using Google StreetView.
// Contains:
// PovLatitude: The latitude for the google StreetView camera
// PovLongitude: The longitude for the google StreetView camera
// PovPitch: The pitch for the google StreetView camera
// PovHeading: The heading for the google StreetView camera
// PovZoom: The zoom level for the google StreetView camera
type StreetView struct {
	PropertyID   uint `gorm:"primary_key" sql:"type:int(10) unsigned"`
	PovLatitude  float32
	PovLongitude float32
	PovPitch     float32
	PovHeading   float32
	PovZoom      int
}

// Area The minimum / maximum internal area for the property, if supplied.
// May be in Imperial, Metric or both. Used for Commercial properties.
// Contains:
// Unit: "sqft", "sqm", "acre", "hectare"
// Min:
// Max:
type Area struct {
	PropertyID uint    `gorm:"primary_key" sql:"type:int(10) unsigned"`
	Unit       string  `xml:"unit,attr"`
	Min        float64 `xml:"min"`
	Max        float64 `xml:"max"`
}

// LandArea The land / external area for the property, if supplied.
// May be in Imperial, Metric. Used for Commercial properties.
// Contains:
// Unit: "sqft", "sqm", "acre", "hectare"
// Min:
// Max:
type LandArea struct {
	Area `xml:"landarea"`
}

// EnergyEfficiency The Environmental Impact value for the property.
// Values are 1-100. Includes Current and Potential values.
type EnergyEfficiency struct {
	PropertyID uint         `gorm:"primary_key" sql:"type:int(10) unsigned"`
	Current    SanitizedInt `xml:"current"`
	Potential  SanitizedInt `xml:"potential"`
}

// EnvironmentalImpact The Environmental Impact	value for the property.
// Values are 1-100. Includes Current and Potential values.
type EnvironmentalImpact struct {
	EnergyEfficiency
}

type ParagraphType SanitizedInt

const (
	StandardTextParagraph    ParagraphType = iota
	EnergyEfficiencyRatings
	DisclaimerTextForDetails
)

// Paragraph contains detailed property information
// Contains:
// Name: The heading for the paragraph, for example, room name, e.g. Lounge, Kitchen etc
// File: The index of the image that relates to this paragraph. If no image is referenced, the value will be NULL.
// Text: The description of the room.
// Metric: The dimensions of the room (if supplied).
// Imperial: The dimensions of the room (if supplied).
// Mixed: The dimensions of the room (if supplied).
type Paragraph struct {
	PropertyID  uint          `gorm:"primary_key" sql:"type:int(10) unsigned"`
	ParagraphID int           `xml:"id,attr" gorm:"primary_key" sql:"type:int"`
	Type        ParagraphType `xml:"type,attr" json:"Type"`
	Name        string        `xml:"name"`
	File        string        `xml:"file"`
	Metric      string        `xml:"dimensions>metric"`
	Imperial    string        `xml:"dimensions>imperial"`
	Mixed       string        `xml:"dimensions>mixed"`
	Text        string        `xml:"text" sql:"type:text"`
}

// Bullet Bullet points, if supplied.
// Contains:
// PropertyID: ID of the parent property
// BulletID: ID of the Bullet
type Bullet struct {
	PropertyID uint         `gorm:"primary_key" sql:"type:int(10) unsigned"`
	BulletID   SanitizedInt `xml:"id,attr" json:"ID" gorm:"primary_key" sql:"type:int"`
	Value      string       `xml:",chardata"`
}

type FileURLType SanitizedInt

const (
	Image                        FileURLType = iota
	Map
	FloorPlan
	Vebra360Tour
	EHouse
	IPix
	FullDetails
	PDFDetails
	ExternalURL
	EnergyPerformanceCertificate
	HouseInformationPack
	VirtualTour
)

// File holds the location of external files uploaded for a property
// Contains:
// Name: Caption for the file, if uploaded by the agent.
// Url: Link to an external resource.
// Updated: Date and time the file was last updated in format “dd/mm/yyyy hh:mm:ss”.
// Type: The type of file. See propertyrelatedtypes.xsd.
// FileID: ID of the file
// PropertyID: ID of the parent property
type File struct {
	PropertyID uint                      `gorm:"primary_key" sql:"type:int(10) unsigned"`
	FileID     SanitizedInt              `xml:"id,attr" json:"ID" gorm:"primary_key" sql:"type:int"`
	Type       FileURLType               `xml:"type,attr"`
	Name       string                    `xml:"name"`
	Url        string                    `xml:"url"`
	Updated    *SanitizedDateISODateTime `xml:"updated" sql:"type:Datetime"`
}

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
