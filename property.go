package api

import (
	"time"
	"strings"
	"regexp"
	"fmt"
	"strconv"
	"encoding/json"
	"database/sql/driver"
	"reflect"
	"encoding/xml"
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

type ChangedPropertyURLBuilder struct {
	url string
}

func (builder *ChangedPropertyURLBuilder) SetURL(url string) {
	builder.url = url
}

func (builder *ChangedPropertyURLBuilder) Build() string {
	return builder.url
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
	Branches []BranchSummary `json:"branch" xml:"branch"`
}

// BranchSummary Returned by the API call http://webservices.vebra.com/export/{datafeedid}/v10/branch List of branches for a Search Group.
// Contains:
// Name: 		Branch name
// FirmID: 		PSG Firm Identifier
// BranchID: 	PSG Branch Identifier for this Firm.
// Url:			REST Uri for the full branch details
type BranchSummary struct {
	Name     string `json:"name" xml:"name"`
	FirmID   int    `json:"firmid" xml:"firmid"`
	BranchID int    `json:"branchid" xml:"branchid"`
	Url      string `json:"url" xml:"url"`
}

// GetClientID returns the 4 digit client ID for the branch
func (bs BranchSummary) GetClientID() (int, error) {
	index := strings.LastIndex(bs.Url, "/")
	return strconv.Atoi(bs.Url[(index + 1):])
}

func (bs BranchSummary) GetClientIDString() string {
	index := strings.LastIndex(bs.Url, "/")
	return bs.Url[(index + 1):]
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
	ClientID int    `json:"clientid" xml:"clientid"`
	FirmID   int    `json:"firmID" xml:"FirmID"`
	BranchID int    `json:"branchID" xml:"BranchID"`
	Name     string `json:"name" xml:"name"`
	URL      string `json:"url" xml:"url"`
	Street   string `json:"street" xml:"street"`
	Town     string `json:"town" xml:"town"`
	County   string `json:"county" xml:"county"`
	Postcode string `json:"postcode" xml:"postcode"`
	Phone    string `json:"phone" xml:"phone"`
	Email    string `json:"email" xml:"email"`
}

type PropertySummaries struct {
	Properties []PropertySummary `json:"property" xml:"property"`
}

// PropertySummary Returned by the API call http://webservices.vebra.com/export/{datafeedid}/v10/branch/{clientid}/property List of Properties in a branch.
// Contains:
// PropertyID: Unique identifier for this property.
// Updated: Last changed Datetime for this property. ISO Date Time format - YYYY-MM-DDTHH:MI:SS
// Url: REST Uri for the full details of the this property.
type PropertySummary struct {
	PropertyID  uint                      `json:"propId" xml:"prop_id" gorm:"primary_key" sql:"type:int"`
	LastChanged *SanitizedDateISODateTime `json:"lastchanged" xml:"lastchanged"`
	Url         string                    `json:"url" xml:"url"`
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
	Properties []Property `json:"property" xml:"property"`
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
// Available: Date the property is available in format “dd/mm/yyyy”. Used mainly for rental Properties. If value held is NULL, is populated with 01/01/1900.
// Uploaded: Date the property was uploaded in format “dd/mm/yyyy”.
// Longitude: The exact longitude of the property.
// Latitude: The exact latitude of the property.
// Easting: Easting (Cartesian coordinates)
// Northing: Northing (Cartesian coordinates)
// StreetView: See StreetView
// WebStatus: The status of the property. See propertyrelatedtypes.xsd.
// CustomStatus: Relates to Commercial Properties.
// CommRent: Commercial rental amount. Listed rental price plus a qualifier which is configurable by the agent, e.g. per sq ft. Now DEPRECATED - please use price node.
// Premium: Premium for property, if supplied. Relates to Commerical Properties.
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
	ID                  uint                       `json:"id" xml:"id,attr" gorm:"primary_key" sql:"type:int(10) unsigned"`
	System              string                     `json:"system" xml:"system,attr"`
	Firmid              string                     `json:"firmid" xml:"firmid,attr"`
	Branchid            int                        `json:"branchid" xml:"branchid,attr"`
	Database            int                        `json:"database" xml:"database,attr"`
	Featured            int                        `json:"featured" xml:"featured,attr"`
	AgentReference      Reference                  `json:"reference" xml:"reference"`
	Address             Address                    `json:"address" xml:"address"`
	Price               Price                      `json:"price" xml:"price"`
	RentalFees          string                     `json:"rentalfees" xml:"rentalfees"`
	LettingsFee         string                     `json:"lettingsfee" xml:"lettingsfee"`
	RmQualifier         RMQualifier                `json:"rmQualifier" xml:"rm_qualifier"`
	Available           *SanitizedDateUKDateFormat `json:"available" xml:"available"`
	Uploaded            *SanitizedDateUKDateFormat `json:"uploaded" xml:"uploaded" gorm:"type:Datetime"`
	Longitude           float32                    `json:"longitude" xml:"longitude"`
	Latitude            float32                    `json:"latitude" xml:"latitude"`
	Easting             SanitizedInt               `json:"easting" xml:"easting"`
	Northing            SanitizedInt               `json:"northing" xml:"northing"`
	StreetView          StreetView                 `json:"streetview" xml:"streetview"`
	WebStatus           PropertyStatus             `json:"webStatus" xml:"web_status"`
	CustomStatus        string                     `json:"customStatus" xml:"custom_status"`
	CommRent            string                     `json:"commRent" xml:"comm_rent"`
	Premium             string                     `json:"premium" xml:"premium"`
	ServiceCharge       string                     `json:"serviceCharge" xml:"service_charge"`
	RateableValue       string                     `json:"rateableValue" xml:"rateable_value"`
	Type                string                     `json:"type" xml:"type"`
	Furnished           RMTypeFurnished            `json:"furnished" xml:"furnished"`
	RmType              RMType                     `json:"rmType" xml:"rm_type"`
	LetBond             SanitizedInt               `json:"letBond" xml:"let_bond"`
	RmLetTypeID         RMTypeLetType              `json:"rmLetTypeId" xml:"rm_let_type_id"`
	Bedrooms            SanitizedInt               `json:"bedrooms" xml:"bedrooms"`
	Receptions          SanitizedInt               `json:"receptions" xml:"receptions"`
	Bathrooms           SanitizedInt               `json:"bathrooms" xml:"bathrooms"`
	UserField1          string                     `json:"userfield1" xml:"userfield1"`
	UserField2          SanitizedInt               `json:"userfield2" xml:"userfield2"`
	SoldDate            *SanitizedDateISODate      `json:"solddate" xml:"solddate" gorm:"type:Datetime"`
	LeaseEnd            *SanitizedDateISODate      `json:"leaseend" xml:"leaseend" gorm:"type:Datetime"`
	Instructed          *SanitizedDateISODate      `json:"instructed" xml:"instructed" gorm:"type:Datetime"`
	SoldPrice           SanitizedInt               `json:"soldprice" xml:"soldprice"`
	Garden              SanitizedBool              `json:"garden" xml:"garden"`
	Parking             SanitizedBool              `json:"parking" xml:"parking"`
	NewBuild            SanitizedBool              `json:"newbuild" xml:"newbuild"`
	GroundRent          string                     `json:"groundrent" xml:"groundrent"`
	Commission          string                     `json:"commission" xml:"commission"`
	Area                []Area                     `json:"area" xml:"area"`
	LandArea            LandArea                   `json:"landarea" xml:"landarea"`
	Description         string                     `json:"description" xml:"description" gorm:"type:varchar(2056)"`
	EnergyEfficiency    EnergyEfficiency           `json:"energyEfficiency" xml:"hip>energy_performance>energy_efficiency"`
	EnvironmentalImpact EnvironmentalImpact        `json:"environmentalImpact" xml:"hip>energy_performance>environmental_impact"`
	Paragraphs          []Paragraph                `json:"paragraphs" xml:"paragraphs>paragraph"`
	Bullets             []Bullet                   `json:"bullets" xml:"bullets>bullet"`
	Files               []File                     `json:"files" xml:"files>file"`
}

const (
	Updated = "updated"
	Deleted = "deleted"
)

type ChangedPropertySummaries struct {
	PropertySummaries []ChangedPropertySummary `json:"property" xml:"property"`
}

// ChangedPropertySummary
// Returned by the API call
// http://webservices.vebra.com/export/{datafeedid}/v10/property/{yyyy}/{MM}/{dd}/{HH}/{mm}/{ss}
// List of Properties updated in a search group.
// Contains:
// PropertyID: Unique identifier for this property
// Updated: Last changed Datetime for this property. ISO Date Time format - YYYY-MM-DDTHH:MI:SS
// LastAction Action field to indicate whether this property was updated or deleted. Possible values: ["updated", "deleted"]
type ChangedPropertySummary struct {
	PropertyID  uint                      `xml:"propid" json:"-"`
	LastChanged *SanitizedDateISODateTime `json:"lastchanged" xml:"lastchanged"`
	Url         string                    `json:"url" xml:"url"`
	LastAction  string                    `json:"action" xml:"action"`
}

func (cps *ChangedPropertySummary) SetLastAction(lastAction string) {
	cps.LastAction = lastAction
}

func (cps *ChangedPropertySummary) GetLastAction() string {
	return cps.LastAction
}

// GetClientID returns the 4 digit client ID for the branch
func (cps *ChangedPropertySummary) GetClientID() (int, error) {
	re := regexp.MustCompile(`branch/(\d+)/`)
	matches := re.FindStringSubmatch(cps.Url)

	if matches == nil || len(matches) < 2 {
		return 0, fmt.Errorf("couldnt match client ID in URL: [%s]", cps.Url)
	}

	return strconv.Atoi(matches[1])
}

// Reference This is the agents reference and can be displayed on an agent's search.
// Rightmove use this as part of property reference.
type Reference struct {
	PropertyID uint `gorm:"primary_key" sql:"type:int(10) unsigned" json:"-"`
	Agents     int  `json:"agents" xml:"agents"`
	Software   int  `json:"software" xml:"software"`
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
	PropertyID     uint   `json:"-" gorm:"primary_key" sql:"type:int(10) unsigned"`
	Name           string `json:"name" xml:"name"`
	Street         string `json:"street" xml:"street"`
	Locality       string `json:"locality" xml:"locality"`
	Town           string `json:"town" xml:"town"`
	County         string `json:"county" xml:"county"`
	Postcode       string `json:"postcode" xml:"postcode"`
	CustomLocation string `json:"customLocatiom" xml:"custom_locatiom"`
	Display        string `json:"display" xml:"display"`
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
	PropertyID uint         `json:"-" gorm:"primary_key" sql:"type:int(10) unsigned"`
	Qualifier  string       `json:"qualifier" xml:"qualifier,attr"`
	Currency   string       `json:"currency" xml:"currency,attr"`
	Display    string       `json:"display" xml:"display,attr"`
	Rent       string       `json:"rent" xml:"rent,attr"`
	Value      SanitizedInt `json:"value" xml:",chardata"`
}

// StreetView describes the longitude, latitude, yaw, pitch and zoom for the property using Google StreetView.
// Contains:
// PovLatitude: The latitude for the google StreetView camera
// PovLongitude: The longitude for the google StreetView camera
// PovPitch: The pitch for the google StreetView camera
// PovHeading: The heading for the google StreetView camera
// PovZoom: The zoom level for the google StreetView camera
type StreetView struct {
	PropertyID   uint    `json:"-" gorm:"primary_key" sql:"type:int(10) unsigned"`
	PovLatitude  float32 `json:"povLatitude"`
	PovLongitude float32 `json:"povLongitude"`
	PovPitch     float32 `json:"povPitch"`
	PovHeading   float32 `json:"povHeading"`
	PovZoom      int     `json:"povZoom"`
}

// Area The minimum / maximum internal area for the property, if supplied.
// May be in Imperial, Metric or both. Used for Commercial Properties.
// Contains:
// Unit: "sqft", "sqm", "acre", "hectare"
// Min:
// Max:
type Area struct {
	PropertyID uint    `json:"-" gorm:"primary_key" sql:"type:int(10) unsigned"`
	Unit       string  `json:"unit" xml:"unit,attr"`
	Min        float64 `json:"min" xml:"min"`
	Max        float64 `json:"max" xml:"max"`
}

// LandArea The land / external area for the property, if supplied.
// May be in Imperial, Metric. Used for Commercial Properties.
// Contains:
// Unit: "sqft", "sqm", "acre", "hectare"
// Min:
// Max:
type LandArea struct {
	Area `json:"-" xml:"landarea"`
}

// EnergyEfficiency The Environmental Impact value for the property.
// Values are 1-100. Includes Current and Potential values.
type EnergyEfficiency struct {
	PropertyID uint         `json:"-" gorm:"primary_key" sql:"type:int(10) unsigned"`
	Current    SanitizedInt `json:"current" xml:"current"`
	Potential  SanitizedInt `json:"potential" xml:"potential"`
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

type ParagraphFileIndex struct {
	File uint `xml:"ref,attr"`
	value string `json:"-"`
}

func (pfi *ParagraphFileIndex) MarshalJSON() ([]byte, error) {
	if pfi.value == "" {
		return json.Marshal(nil)
	}
	return json.Marshal(pfi.File)
}

func (pfi *ParagraphFileIndex) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for _, attr := range start.Attr {
		if attr.Name.Local != "ref" {
			continue
		}
		if attr.Value == "" {
			return nil
		}
		pfi.value = attr.Value
		var index int
		if index, err = strconv.Atoi(attr.Value); err != nil {
			return err
		}
		pfi.File = uint(index)
	}
	return d.Skip()
}

func (u *ParagraphFileIndex) Scan(value interface{}) error {
	switch value.(type) {
	case int:
		u.File = value.(uint)
		return nil
	case uint:
		u.File = value.(uint)
		return nil
	case int64:
		u.File = value.(uint)
	default:
		return fmt.Errorf("unexpected ParagraphFileIndex type: %s", reflect.TypeOf(value))
	}
	return nil
}

func (u *ParagraphFileIndex) Value() (driver.Value, error) {
	if u.value == "" {
		return nil, nil
	}
	return int64(u.File), nil
}

// Paragraph contains detailed property information
// Contains:
// Name: The heading for the paragraph, for example, room name, e.g. Lounge, Kitchen etc
// File: The index of the image that relates to this paragraph. If no image is referenced, the value will be NULL.
// Text: The description of the room.
// Metric: The dimensions of the room (if supplied).
// Imperial: The dimensions of the room (if supplied).
// Mixed: The dimensions of the room (if supplied).
type Paragraph struct {
	PropertyID  uint                `json:"-" gorm:"primary_key" sql:"type:int(10) unsigned"`
	ParagraphID int                 `json:"id" xml:"id,attr" gorm:"primary_key" sql:"type:int"`
	Type        ParagraphType       `json:"type" xml:"type,attr" json:"Type"`
	Name        string              `json:"name" xml:"name"`
	File        *ParagraphFileIndex `json:"file" xml:"file"`
	Metric      string              `json:"metric" xml:"dimensions>metric"`
	Imperial    string              `json:"imperial" xml:"dimensions>imperial"`
	Mixed       string              `json:"mixed" xml:"dimensions>mixed"`
	Text        string              `json:"text" xml:"text" sql:"type:text"`
}

// Bullet Bullet points, if supplied.
// Contains:
// PropertyID: ID of the parent property
// BulletID: ID of the Bullet
type Bullet struct {
	PropertyID uint         `json:"-" gorm:"primary_key" sql:"type:int(10) unsigned"`
	BulletID   SanitizedInt `json:"id" xml:"id,attr" json:"ID" gorm:"primary_key" sql:"type:int"`
	Value      string       `json:"value" xml:",chardata"`
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
	PropertyID uint                      `json:"-" gorm:"primary_key" sql:"type:int(10) unsigned"`
	FileID     SanitizedInt              `json:"id" xml:"id,attr" gorm:"primary_key" sql:"type:int"`
	Type       FileURLType               `json:"type" xml:"type,attr"`
	Name       string                    `json:"name" xml:"name"`
	Url        string                    `json:"url" xml:"url"`
	Updated    *SanitizedDateISODateTime `json:"updated" xml:"updated" sql:"type:Datetime"`
}

type ChangedFilesSummaries struct {
	Files []ChangedFileSummary `json:"file" xml:"file"`
}

type ChangedFileSummary struct {
	FileID     int    `json:"-" xml:"file_id"`
	FilePropId int    `json:"filePropId" xml:"file_propid"`
	Updated    string `json:"updated" xml:"updated"`
	Deleted    bool   `json:"deleted" xml:"deleted"`
	Url        string `json:"url" xml:"url"`
	PropUrl    string `json:"propUrl" xml:"prop_url"`
}
