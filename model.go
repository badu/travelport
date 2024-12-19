package travelport

import (
	"regexp"
	"time"
)

// acceptedCreditCardProcessingCountryPattern is the validation pattern for AcceptedCreditCard.ProcessingCountry
var acceptedCreditCardProcessingCountryPattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// acceptedCreditCardValuePattern is the validation pattern for AcceptedCreditCard.Value
var acceptedCreditCardValuePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// AcceptedCreditCard is an object. Credit card code
type AcceptedCreditCard struct {
	ProcessingCountry string `json:"processingCountry,omitempty"` // ProcessingCountry: Country Code ISO
	Value             string `json:"value,omitempty"`             // Value:
}

// Accounting is an object.
type Accounting struct {
	Type          string          `json:"@type"`                   // Type:
	AccountingRef string          `json:"AccountingRef,omitempty"` // AccountingRef: Accounting reference
	Identifier    *Identifier     `json:"Identifier,omitempty"`    // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NameValuePair []NameValuePair `json:"NameValuePair,omitempty"` // NameValuePair:
	DataType      string          `json:"dataType,omitempty"`      // DataType: Accounting data type
	Id            string          `json:"id,omitempty"`            // Id:
	Template      string          `json:"template,omitempty"`      // Template: Accounting template
}

// AccountingID is an object.
type AccountingID struct {
	Type          string      `json:"@type"`                   // Type:
	AccountingRef string      `json:"AccountingRef,omitempty"` // AccountingRef: Accounting reference
	Identifier    *Identifier `json:"Identifier,omitempty"`    // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id            string      `json:"id,omitempty"`            // Id:
}

// additionalBrandAttributeGroupCodePattern is the validation pattern for AdditionalBrandAttribute.GroupCode
var additionalBrandAttributeGroupCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// additionalBrandAttributeSubCodePattern is the validation pattern for AdditionalBrandAttribute.SubCode
var additionalBrandAttributeSubCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// additionalBrandAttributeSubGroupCodePattern is the validation pattern for AdditionalBrandAttribute.SubGroupCode
var additionalBrandAttributeSubGroupCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// AdditionalBrandAttribute is an object.
type AdditionalBrandAttribute struct {
	Type                string             `json:"@type"`                         // Type:
	Classification      string             `json:"Classification"`                // Classification: The Travelport classification used for categories of ancillaries
	ImageURL            []string           `json:"ImageURL,omitempty"`            // ImageURL:
	Inclusion           BrandInclusionEnum `json:"Inclusion"`                     // Inclusion: The indicator as to how the brand and the product are associated.
	GroupCode           string             `json:"groupCode,omitempty"`           // GroupCode:
	MatchedAttributeInd bool               `json:"matchedAttributeInd,omitempty"` // MatchedAttributeInd: If true, this is a matched attribute
	SubCode             string             `json:"subCode,omitempty"`             // SubCode:
	SubGroupCode        string             `json:"subGroupCode,omitempty"`        // SubGroupCode:
}

// additionalDetailDetailTypePattern is the validation pattern for AdditionalDetail.DetailType
var additionalDetailDetailTypePattern = regexp.MustCompile(`[0-9A-Z]{1,3}(\.[A-Z]{3}(\.X){0,1}){0,1}`)

// AdditionalDetail is an object.
type AdditionalDetail struct {
	Type        string                   `json:"@type,omitempty"`       // Type:
	Amount      *CurrencyAmount          `json:"Amount,omitempty"`      // Amount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	Description *TextTitleAndDescription `json:"Description,omitempty"` // Description: Descriptive text
	Code        string                   `json:"code,omitempty"`        // Code: Partner code
	DetailType  string                   `json:"detailType,omitempty"`  // DetailType: OTA Code
}

// AdditionalDetails is an object.
type AdditionalDetails struct {
	Type             string             `json:"@type,omitempty"`  // Type:
	AdditionalDetail []AdditionalDetail `json:"AdditionalDetail"` // AdditionalDetail:
}

// Address is an object.
type Address struct {
	Type        string               `json:"@type"`                 // Type:
	AddressLine []string             `json:"AddressLine,omitempty"` // AddressLine: Additional address line details
	BldgRoom    *AddressBldgRoom     `json:"BldgRoom,omitempty"`    // BldgRoom: Address with building and room number
	City        string               `json:"City"`                  // City: City (e.g., Dublin), town, or postal station (i.e., a postal service territory, often used in a military address).
	Country     *Country             `json:"Country,omitempty"`     // Country: ISO 3166 code for a country with optional name
	County      string               `json:"County,omitempty"`      // County: County or Region Name (e.g., Fairfax).
	Number      *AddressStreetNumber `json:"Number,omitempty"`      // Number: The street number alone is the numerical number that precedes the street name in the address
	PostalCode  string               `json:"PostalCode,omitempty"`  // PostalCode: Post Office Code number.
	StateProv   *StateProv           `json:"StateProv,omitempty"`   // StateProv: The standard code or abbreviation for the state, province, or region with optional name
	Street      string               `json:"Street,omitempty"`      // Street: May contain the street number when the Street number element is missing.
	Id          string               `json:"id,omitempty"`          // Id: Internally referenced id
	Role        EnumAddressRole      `json:"role,omitempty"`        // Role:
}

// AddressBldgRoom is an object. Address with building and room number
type AddressBldgRoom struct {
	BuldingInd bool   `json:"buldingInd,omitempty"` // BuldingInd: When true, the information is a building name. When false, it is an apartment or room #
	Value      string `json:"value,omitempty"`      // Value:
}

// addressStreetNumberRuralRouteNmbrPattern is the validation pattern for AddressStreetNumber.RuralRouteNmbr
var addressStreetNumberRuralRouteNmbrPattern = regexp.MustCompile(`[0-9]{1,5}`)

// AddressStreetNumber is an object. The street number alone is the numerical number that precedes the street name in the address
type AddressStreetNumber struct {
	PoBox            string `json:"po_Box,omitempty"`           // PoBox: PO Box Number
	RuralRouteNmbr   string `json:"ruralRouteNmbr,omitempty"`   // RuralRouteNmbr: RuralRoute Number
	StreetDirection  string `json:"streetDirection,omitempty"`  // StreetDirection: Dircetion of the Street
	StreetNmbrSuffix string `json:"streetNmbrSuffix,omitempty"` // StreetNmbrSuffix: Street Number Suffix
	Value            string `json:"value,omitempty"`            // Value:
}

// AgeQualifying is an object.
type AgeQualifying struct {
	Type              string `json:"@type,omitempty"`             // Type:
	AgeBucket         string `json:"ageBucket,omitempty"`         // AgeBucket: The age bucket
	AgeQualifyingCode string `json:"ageQualifyingCode,omitempty"` // AgeQualifyingCode: Enter 10 for an adult or 08 for a child
	Count             int32  `json:"count,omitempty"`             // Count: The number of age qualifying
	MaxAge            int32  `json:"maxAge,omitempty"`            // MaxAge: Max Age: The maximum age to qualify for AgeQualifyingCode.
	MinAge            int32  `json:"minAge,omitempty"`            // MinAge: MinAge: The minimum age to qualify for AgeQualifyingCode.
}

// agencyInfoCodePattern is the validation pattern for AgencyInfo.Code
var agencyInfoCodePattern = regexp.MustCompile(`([0-9]{8})`)

// agencyInfoTicketedDatePattern is the validation pattern for AgencyInfo.TicketedDate
var agencyInfoTicketedDatePattern = regexp.MustCompile(`(\d{4}-\d{2}-\d{2})`)

// agencyInfoTicketingPCCPattern is the validation pattern for AgencyInfo.TicketingPCC
var agencyInfoTicketingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// AgencyInfo is an object. Detail of the travel agency that issues the ticket
type AgencyInfo struct {
	Code             string `json:"code,omitempty"`         // Code: Agency code
	Name             string `json:"name"`                   // Name: Name of the Agency
	Place            string `json:"place"`                  // Place: Place of the agency
	SalesType        string `json:"salesType,omitempty"`    // SalesType: Sales type
	TicketedDate     string `json:"ticketedDate,omitempty"` // TicketedDate: Ticketed date
	TicketingCity    string `json:"ticketingCity"`          // TicketingCity: Ticketing city
	TicketingCountry string `json:"ticketingCountry"`       // TicketingCountry: Ticketing country
	TicketingPCC     string `json:"ticketingPCC,omitempty"` // TicketingPCC: Ticketing PCC
}

// AgencyServiceFee is an object.
type AgencyServiceFee struct {
	Type                  string          `json:"@type"`                           // Type:
	Amount                CurrencyAmount  `json:"Amount"`                          // Amount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	Description           string          `json:"Description,omitempty"`           // Description: The description of the service fee
	ExpiryDate            time.Time       `json:"ExpiryDate,omitempty"`            // ExpiryDate: The service fee expiry date. Once expiry date has been reached, the service fee information will only be stored in the ReservationReceipt
	OfferRef              string          `json:"OfferRef,omitempty"`              // OfferRef: Reference to an Offer within the Reservation that this service fee applies to
	RelatedDocumentNumber *DocumentNumber `json:"RelatedDocumentNumber,omitempty"` // RelatedDocumentNumber:
	Tax                   []Tax           `json:"Tax,omitempty"`                   // Tax:
	TravelerRef           string          `json:"TravelerRef,omitempty"`           // TravelerRef: Reference to a Traveler within the Reservation that this service fee applies to
	Id                    string          `json:"id,omitempty"`                    // Id: Unique id for this object within a message
}

// AgencyServiceFeeID is an object.
type AgencyServiceFeeID struct {
	Type string `json:"@type"`        // Type:
	Id   string `json:"id,omitempty"` // Id: Unique id for this object within a message
}

// AgencyServiceFeeIdentifier is an object.
type AgencyServiceFeeIdentifier struct {
	Id string `json:"id,omitempty"` // Id: Unique id for this object within a message
}

// AlternateContact is an object.
type AlternateContact struct {
	Type         string      `json:"@type,omitempty"`        // Type:
	Address      []Address   `json:"Address,omitempty"`      // Address:
	Email        []Email     `json:"Email,omitempty"`        // Email:
	PersonName   PersonName  `json:"PersonName"`             // PersonName:
	Telephone    []Telephone `json:"Telephone,omitempty"`    // Telephone:
	ContactType  string      `json:"contactType,omitempty"`  // ContactType: Contact type value
	DefaultInd   bool        `json:"defaultInd,omitempty"`   // DefaultInd: This is the default contact
	EmergencyInd bool        `json:"emergencyInd,omitempty"` // EmergencyInd: This is the contact in case of an emergency
	Id           string      `json:"id,omitempty"`           // Id:
	Relation     string      `json:"relation,omitempty"`     // Relation: Relation value
}

// AmenitySurcharges is an object.
type AmenitySurcharges struct {
	Type            string  `json:"@type"`                     // Type:
	TotalSurcharges float32 `json:"TotalSurcharges,omitempty"` // TotalSurcharges: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	ApproximateInd  bool    `json:"approximateInd,omitempty"`  // ApproximateInd: if true, the surcharge amounts are approximate
}

// AmenitySurchargesDetail is an object.
type AmenitySurchargesDetail struct {
	Type            string      `json:"@type"`                     // Type:
	Surcharge       []Surcharge `json:"Surcharge"`                 // Surcharge:
	TotalSurcharges float32     `json:"TotalSurcharges,omitempty"` // TotalSurcharges: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	ApproximateInd  bool        `json:"approximateInd,omitempty"`  // ApproximateInd: if true, the surcharge amounts are approximate
}

// Amount is an object.
type Amount struct {
	Type           string             `json:"@type,omitempty"`          // Type:
	Base           float32            `json:"Base,omitempty"`           // Base: The price prior to all applicable taxes of a product such as the rate for a room or fare for a flight.
	CurrencyCode   *CurrencyCode      `json:"CurrencyCode,omitempty"`   // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	Fees           *Fees              `json:"Fees,omitempty"`           // Fees:
	Taxes          *Taxes             `json:"Taxes,omitempty"`          // Taxes:
	Total          float32            `json:"Total,omitempty"`          // Total: Specifies the total price including base + taxes + fees
	ApproximateInd bool               `json:"approximateInd,omitempty"` // ApproximateInd: True if this amount has been converted from the original amount
	CurrencySource CurrencySourceEnum `json:"currencySource,omitempty"` // CurrencySource: The system requesting or returning the currency code specified in the attribute
}

// AmountPercent is an object.
type AmountPercent struct {
	Type        string         `json:"@type"`                 // Type:
	Application CommissionEnum `json:"application,omitempty"` // Application: Type of commission
}

// AmountPercentAmount is an object.
type AmountPercentAmount struct {
	Type        string          `json:"@type"`                 // Type:
	Amount      *CurrencyAmount `json:"Amount,omitempty"`      // Amount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	Application CommissionEnum  `json:"application,omitempty"` // Application: Type of commission
}

// AmountPercentPercent is an object.
type AmountPercentPercent struct {
	Type        string         `json:"@type"`                 // Type:
	Percent     float32        `json:"Percent,omitempty"`     // Percent: Percent amount of commission
	Application CommissionEnum `json:"application,omitempty"` // Application: Type of commission
}

// Ancillary is an object.
type Ancillary struct {
	Type        string                 `json:"@type"`              // Type:
	Description []AncillaryDescription `json:"Description"`        // Description:
	Quantity    int32                  `json:"quantity,omitempty"` // Quantity: The quantity value
}

// AncillaryAir is an object.
type AncillaryAir struct {
	Type        string                 `json:"@type"`               // Type:
	Description []AncillaryDescription `json:"Description"`         // Description:
	FlightRef   []string               `json:"FlightRef,omitempty"` // FlightRef: The list of travel segments the ancillary applies to
	Quantity    int32                  `json:"quantity,omitempty"`  // Quantity: The quantity value
}

// AncillaryAirBaggage is an object.
type AncillaryAirBaggage struct {
	Type        string                 `json:"@type"`               // Type:
	BaggageType BaggageTypeEnum        `json:"BaggageType"`         // BaggageType:
	Description []AncillaryDescription `json:"Description"`         // Description:
	FlightRef   []string               `json:"FlightRef,omitempty"` // FlightRef: The list of travel segments the ancillary applies to
	Measurement []Measurement          `json:"Measurement"`         // Measurement:
	Quantity    int32                  `json:"quantity,omitempty"`  // Quantity: The quantity value
}

// AncillaryAirSeat is an object.
type AncillaryAirSeat struct {
	Type           string                 `json:"@type"`               // Type:
	Description    []AncillaryDescription `json:"Description"`         // Description:
	FlightRef      []string               `json:"FlightRef,omitempty"` // FlightRef: The list of travel segments the ancillary applies to
	SeatAssignment SeatAssignment         `json:"SeatAssignment"`      // SeatAssignment:
	Quantity       int32                  `json:"quantity,omitempty"`  // Quantity: The quantity value
}

// ancillaryDescriptionCodePattern is the validation pattern for AncillaryDescription.Code
var ancillaryDescriptionCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// ancillaryDescriptionSubCodePattern is the validation pattern for AncillaryDescription.SubCode
var ancillaryDescriptionSubCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// AncillaryDescription is an object. A description of the ancillary with two description codes
type AncillaryDescription struct {
	Code        string `json:"code,omitempty"`        // Code: The code value
	CodeContext string `json:"codeContext,omitempty"` // CodeContext: The code Context value
	SubCode     string `json:"subCode,omitempty"`     // SubCode: The subcode value
	Value       string `json:"value,omitempty"`       // Value:
}

// AncillaryOfferingID is an object.
type AncillaryOfferingID struct {
	Type                 string      `json:"@type"`                          // Type:
	AncillaryOfferingRef string      `json:"AncillaryOfferingRef,omitempty"` // AncillaryOfferingRef:
	CatalogOfferingRef   string      `json:"CatalogOfferingRef,omitempty"`   // CatalogOfferingRef: Used to reference another instance of this object in the same message
	Identifier           *Identifier `json:"Identifier,omitempty"`           // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id                   string      `json:"id,omitempty"`                   // Id: Local indentifier within a given message for this object.
}

// AncillaryOfferingIdentifier is an object.
type AncillaryOfferingIdentifier struct {
	AncillaryOfferingRef string      `json:"AncillaryOfferingRef,omitempty"` // AncillaryOfferingRef:
	CatalogOfferingRef   string      `json:"CatalogOfferingRef,omitempty"`   // CatalogOfferingRef: Used to reference another instance of this object in the same message
	Identifier           *Identifier `json:"Identifier,omitempty"`           // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id                   string      `json:"id,omitempty"`                   // Id: Local indentifier within a given message for this object.
}

// AncillaryOfferings is an object.
type AncillaryOfferings struct {
	Type                            string `json:"@type"`                                     // Type:
	IncludeUnsellableAncillariesInd bool   `json:"includeUnsellableAncillariesInd,omitempty"` // IncludeUnsellableAncillariesInd: If true, the response will include unsellable ancillary options
}

// AncillaryOfferingsBuildFromCatalogOfferings is an object.
type AncillaryOfferingsBuildFromCatalogOfferings struct {
	Type                             string                           `json:"@type"`                                     // Type:
	BuildFromCatalogOfferingsRequest BuildFromCatalogOfferingsRequest `json:"BuildFromCatalogOfferingsRequest"`          // BuildFromCatalogOfferingsRequest:
	IncludeUnsellableAncillariesInd  bool                             `json:"includeUnsellableAncillariesInd,omitempty"` // IncludeUnsellableAncillariesInd: If true, the response will include unsellable ancillary options
}

// AncillaryOfferingsBuildFromCatalogProductOfferings is an object.
type AncillaryOfferingsBuildFromCatalogProductOfferings struct {
	Type                                    string                                  `json:"@type"`                                     // Type:
	BuildFromCatalogProductOfferingsRequest BuildFromCatalogProductOfferingsRequest `json:"BuildFromCatalogProductOfferingsRequest"`   // BuildFromCatalogProductOfferingsRequest:
	IncludeUnsellableAncillariesInd         bool                                    `json:"includeUnsellableAncillariesInd,omitempty"` // IncludeUnsellableAncillariesInd: If true, the response will include unsellable ancillary options
}

// AncillaryOfferingsBuildFromOfferList is an object.
type AncillaryOfferingsBuildFromOfferList struct {
	Type                            string              `json:"@type"`                                     // Type:
	BuildFromOfferList              *BuildFromOfferList `json:"BuildFromOfferList,omitempty"`              // BuildFromOfferList:
	IncludeUnsellableAncillariesInd bool                `json:"includeUnsellableAncillariesInd,omitempty"` // IncludeUnsellableAncillariesInd: If true, the response will include unsellable ancillary options
}

// ApplicationLimit is an object.
type ApplicationLimit struct {
	End   string              `json:"end,omitempty"`   // End: The end value
	Start string              `json:"start,omitempty"` // Start: The start value
	Value ApplicableLevelEnum `json:"value,omitempty"` // Value:
}

// AppliesTo is an object.
type AppliesTo struct {
	Type string `json:"@type"` // Type:
}

// AppliesToOffer is an object.
type AppliesToOffer struct {
	Type            string            `json:"@type"`                     // Type:
	OfferIdentifier []OfferIdentifier `json:"OfferIdentifier,omitempty"` // OfferIdentifier:
}

// AppliesToOfferProduct is an object.
type AppliesToOfferProduct struct {
	Type              string              `json:"@type"`                       // Type:
	OfferIdentifier   *OfferIdentifier    `json:"OfferIdentifier,omitempty"`   // OfferIdentifier:
	ProductIdentifier []ProductIdentifier `json:"ProductIdentifier,omitempty"` // ProductIdentifier:
}

// AppliesToOfferProductSegment is an object.
type AppliesToOfferProductSegment struct {
	Type                string             `json:"@type"`                         // Type:
	OfferIdentifier     *OfferIdentifier   `json:"OfferIdentifier,omitempty"`     // OfferIdentifier:
	ProductIdentifier   *ProductIdentifier `json:"ProductIdentifier,omitempty"`   // ProductIdentifier:
	SegmentSequenceList []int32            `json:"SegmentSequenceList,omitempty"` // SegmentSequenceList: Segment Sequence List
}

// ApproximateRate is an object.
type ApproximateRate struct {
	Type                 string          `json:"@type,omitempty"`                // Type:
	BaseRate             *CurrencyAmount `json:"BaseRate,omitempty"`             // BaseRate: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	DropOffCharge        *CurrencyAmount `json:"DropOffCharge,omitempty"`        // DropOffCharge: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	EstimatedTotalAmount *CurrencyAmount `json:"EstimatedTotalAmount,omitempty"` // EstimatedTotalAmount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	ExtraMileageCharge   *CurrencyAmount `json:"ExtraMileageCharge,omitempty"`   // ExtraMileageCharge: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	RateForPeriod        *CurrencyAmount `json:"RateForPeriod,omitempty"`        // RateForPeriod: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
}

// arrivalLocationPattern is the validation pattern for Arrival.Location
var arrivalLocationPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// Arrival is an object.
type Arrival struct {
	Type     string `json:"@type"`    // Type:
	Date     string `json:"date"`     // Date: Local date of for arrival or departure
	Location string `json:"location"` // Location: Location of departure or arrival
	Time     string `json:"time"`     // Time: Local time Date of for arrival or departure
}

// arrivalDetailCountryPattern is the validation pattern for ArrivalDetail.Country
var arrivalDetailCountryPattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// arrivalDetailLocationPattern is the validation pattern for ArrivalDetail.Location
var arrivalDetailLocationPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// arrivalDetailTerminalPattern is the validation pattern for ArrivalDetail.Terminal
var arrivalDetailTerminalPattern = regexp.MustCompile(`([0-9a-zA-Z]+)?`)

// ArrivalDetail is an object.
type ArrivalDetail struct {
	Type     string `json:"@type"`              // Type:
	Country  string `json:"country,omitempty"`  // Country: Country of to departure or arrival
	Date     string `json:"date"`               // Date: Local date of for arrival or departure
	Location string `json:"location"`           // Location: Location of departure or arrival
	Terminal string `json:"terminal,omitempty"` // Terminal: Departure/Arrival terminal
	Time     string `json:"time"`               // Time: Local time Date of for arrival or departure
}

// arrivalDetailsCarrierPattern is the validation pattern for ArrivalDetails.Carrier
var arrivalDetailsCarrierPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,3})`)

// arrivalDetailsFlightNumberPattern is the validation pattern for ArrivalDetails.FlightNumber
var arrivalDetailsFlightNumberPattern = regexp.MustCompile(`[0-9]{1,4}[A-Z]?|OPEN|ARNK`)

// ArrivalDetails is an object.
type ArrivalDetails struct {
	Type               string                     `json:"@type,omitempty"`              // Type:
	Carrier            string                     `json:"Carrier,omitempty"`            // Carrier: Carrier
	FlightNumber       string                     `json:"FlightNumber,omitempty"`       // FlightNumber: Flight Number
	TransportationCode TransportationCategoryEnum `json:"TransportationCode,omitempty"` // TransportationCode: Transportation Category
}

// authenticationVerificationMaskPattern is the validation pattern for AuthenticationVerification.Mask
var authenticationVerificationMaskPattern = regexp.MustCompile(`[0-9a-zA-Z]{1,32}`)

// authenticationVerificationTokenPattern is the validation pattern for AuthenticationVerification.Token
var authenticationVerificationTokenPattern = regexp.MustCompile(`[0-9a-zA-Z]{1,32}`)

// AuthenticationVerification is an object.
type AuthenticationVerification struct {
	Type                 string                      `json:"@type,omitempty"`                // Type:
	ErrorWarning         *ErrorWarning               `json:"ErrorWarning,omitempty"`         // ErrorWarning:
	PlainText            string                      `json:"PlainText,omitempty"`            // PlainText: Don't use this unless it is REALLY ok to not use encryption. Non-secure (plain text) value.
	AuthenticationMethod EncryptionTokenTypeAuthEnum `json:"authenticationMethod,omitempty"` // AuthenticationMethod: Type of Authentication
	EncryptedValue       string                      `json:"encryptedValue,omitempty"`       // EncryptedValue: Encrypted value
	EncryptionKey        string                      `json:"encryptionKey,omitempty"`        // EncryptionKey: Note: This contains a key required to retrieve the full payment instrument details compliant with PCI DSS standards.
	EncryptionKeyMethod  string                      `json:"encryptionKeyMethod,omitempty"`  // EncryptionKeyMethod: Developer: This contains a reference to the key generation method being used - this is NOT the key value.
	EncryptionMethod     string                      `json:"encryptionMethod,omitempty"`     // EncryptionMethod: OpenTravel Best Practice: Encryption Method: When using the OpenTravel Encryption element, it is RECOMMENDED that all trading partners be informed of all encryption methods being used in advance of implementation to ensure message processing compatibility.
	Mask                 string                      `json:"mask,omitempty"`                 // Mask: Masked Value
	Token                string                      `json:"token,omitempty"`                // Token: Token value
	TokenProviderID      string                      `json:"tokenProviderID,omitempty"`      // TokenProviderID: Developer: This contains a provider ID if multiple providers are used for secure information exchange.
}

// AvailabilityStatusENUM is an object.
type AvailabilityStatusENUM struct {
	Value AvailabilityStatusENUMBase `json:"value,omitempty"` // Value:
}

// BaggageAllowance is an object.
type BaggageAllowance struct {
	Type                string          `json:"@type"`                         // Type:
	BaggageItem         []BaggageItem   `json:"BaggageItem,omitempty"`         // BaggageItem:
	ProductRef          []string        `json:"ProductRef,omitempty"`          // ProductRef: A product is any product, service or package of products and services  that can be priced and purchased by a specific supplier.
	SegmentSequenceList []int32         `json:"SegmentSequenceList,omitempty"` // SegmentSequenceList: Segment sequence is only to be used when the baggage allowance differs between segments within a product. If so, then the ProducRef must be specified.
	Text                []string        `json:"Text,omitempty"`                // Text:
	BaggageType         BaggageTypeEnum `json:"baggageType,omitempty"`         // BaggageType:
	PassengerTypeCodes  []string        `json:"passengerTypeCodes,omitempty"`  // PassengerTypeCodes: List of passenger type codes
}

// BaggageAllowanceDetail is an object.
type BaggageAllowanceDetail struct {
	Type                string          `json:"@type"`                         // Type:
	BaggageItem         []BaggageItem   `json:"BaggageItem,omitempty"`         // BaggageItem:
	ProductRef          []string        `json:"ProductRef,omitempty"`          // ProductRef: A product is any product, service or package of products and services  that can be priced and purchased by a specific supplier.
	SegmentSequenceList []int32         `json:"SegmentSequenceList,omitempty"` // SegmentSequenceList: Segment sequence is only to be used when the baggage allowance differs between segments within a product. If so, then the ProducRef must be specified.
	Text                []string        `json:"Text,omitempty"`                // Text:
	BaggageType         BaggageTypeEnum `json:"baggageType,omitempty"`         // BaggageType:
	PassengerTypeCodes  []string        `json:"passengerTypeCodes,omitempty"`  // PassengerTypeCodes: List of passenger type codes
	Url                 string          `json:"url,omitempty"`                 // Url: Url for the airline's baggage information web site
}

// BaggageItem is an object.
type BaggageItem struct {
	Type        string          `json:"@type,omitempty"`       // Type:
	BaggageFee  *CurrencyAmount `json:"BaggageFee,omitempty"`  // BaggageFee: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	Measurement []Measurement   `json:"Measurement,omitempty"` // Measurement:
	Text        string          `json:"Text,omitempty"`        // Text: Text returned from the shop response
	Quantity    int32           `json:"quantity,omitempty"`    // Quantity: Baggage item quantity
}

// baggageRecheckAtPattern is the validation pattern for BaggageRecheck.At
var baggageRecheckAtPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// BaggageRecheck is an object.
type BaggageRecheck struct {
	Type            string   `json:"@type,omitempty"` // Type:
	ArrivalFlight   FlightID `json:"ArrivalFlight"`   // ArrivalFlight:
	At              string   `json:"At"`              // At: The city where the baggage recheck is required
	DepartureFlight FlightID `json:"DepartureFlight"` // DepartureFlight:
}

// BaseResponse is an object.
type BaseResponse struct {
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"` // CurrencyRateConversion:
	Identifier             *Identifier              `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`              // NextSteps:
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`          // ReferenceList:
	Result                 *Result                  `json:"Result,omitempty"`                 // Result:
	TraceId                string                   `json:"traceId,omitempty"`                // TraceId: Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TransactionId          string                   `json:"transactionId,omitempty"`          // TransactionId: Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
}

// BedConfiguration is an object.
type BedConfiguration struct {
	BedType  string `json:"bedType,omitempty"`  // BedType: Configuration of bed(s) in room.
	Quantity int32  `json:"quantity,omitempty"` // Quantity: The number of bed of this type and size in the room
	Size     string `json:"size,omitempty"`     // Size: Size of bed(s) in the room.
}

// BestCombinablePrice is an object.
type BestCombinablePrice struct {
	Type                string               `json:"@type"`                         // Type:
	Base                float32              `json:"Base,omitempty"`                // Base: The total amount not including taxes and\/or fees
	CurrencyCode        *CurrencyCode        `json:"CurrencyCode,omitempty"`        // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	TotalFees           float32              `json:"TotalFees,omitempty"`           // TotalFees: The total of the fees included in the total price
	TotalPrice          float32              `json:"TotalPrice,omitempty"`          // TotalPrice: The total price of the product in the currency indicated
	TotalTaxes          float32              `json:"TotalTaxes,omitempty"`          // TotalTaxes: The total of the taxes included in the total price
	VendorCurrencyTotal *VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"` // VendorCurrencyTotal:
	Id                  string               `json:"id,omitempty"`                  // Id: Internally referenced id
}

// BestCombinablePriceDetail is an object.
type BestCombinablePriceDetail struct {
	Type                string               `json:"@type"`                         // Type:
	Base                float32              `json:"Base,omitempty"`                // Base: The total amount not including taxes and\/or fees
	CurrencyCode        *CurrencyCode        `json:"CurrencyCode,omitempty"`        // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	PriceBreakdown      []PriceBreakdown     `json:"PriceBreakdown,omitempty"`      // PriceBreakdown:
	TotalFees           float32              `json:"TotalFees,omitempty"`           // TotalFees: The total of the fees included in the total price
	TotalPrice          float32              `json:"TotalPrice,omitempty"`          // TotalPrice: The total price of the product in the currency indicated
	TotalTaxes          float32              `json:"TotalTaxes,omitempty"`          // TotalTaxes: The total of the taxes included in the total price
	VendorCurrencyTotal *VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"` // VendorCurrencyTotal:
	Id                  string               `json:"id,omitempty"`                  // Id: Internally referenced id
}

// Brand is an object.
type Brand struct {
	Type                     string                     `json:"@type"`                              // Type:
	AdditionalBrandAttribute []AdditionalBrandAttribute `json:"AdditionalBrandAttribute,omitempty"` // AdditionalBrandAttribute:
	BrandAttribute           []BrandAttribute           `json:"BrandAttribute,omitempty"`           // BrandAttribute:
	BrandRef                 string                     `json:"BrandRef,omitempty"`                 // BrandRef: Used to reference another instance of this object in the same message
	Identifier               *Identifier                `json:"Identifier,omitempty"`               // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ImageURL                 []string                   `json:"ImageURL,omitempty"`                 // ImageURL:
	Id                       string                     `json:"id,omitempty"`                       // Id: Local indentifier within a given message for this object.
	Name                     string                     `json:"name"`                               // Name: Assigned Type: ctbr-1100:BrandName
	ShelfNumbers             []int32                    `json:"shelfNumbers,omitempty"`             // ShelfNumbers: Assigned Type: ctbr-1100:ShelfNumbers
	Tier                     int32                      `json:"tier,omitempty"`                     // Tier: Assigned Type: c-1100:NumberSingleDigit
}

// brandAttributeGroupCodePattern is the validation pattern for BrandAttribute.GroupCode
var brandAttributeGroupCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// brandAttributeSubCodePattern is the validation pattern for BrandAttribute.SubCode
var brandAttributeSubCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// brandAttributeSubGroupCodePattern is the validation pattern for BrandAttribute.SubGroupCode
var brandAttributeSubGroupCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// BrandAttribute is an object.
type BrandAttribute struct {
	Type                string                  `json:"@type"`                         // Type:
	ImageURL            []string                `json:"ImageURL,omitempty"`            // ImageURL:
	Classification      BrandClassificationEnum `json:"classification"`                // Classification: The Travelport classification used for a category of ancillaries such as Seat, Bags, etc. This is an initial list that will be added to.
	GroupCode           string                  `json:"groupCode,omitempty"`           // GroupCode:
	Inclusion           BrandInclusionEnum      `json:"inclusion"`                     // Inclusion: The indicator as to how the brand and the product are associated.
	MatchedAttributeInd bool                    `json:"matchedAttributeInd,omitempty"` // MatchedAttributeInd: if true, this is a matched attribute.
	SubCode             string                  `json:"subCode,omitempty"`             // SubCode:
	SubGroupCode        string                  `json:"subGroupCode,omitempty"`        // SubGroupCode:
}

// BrandAttributeInclusion is an object.
type BrandAttributeInclusion struct {
	Type                     string                    `json:"@type,omitempty"`                    // Type:
	AdditionalClassification []string                  `json:"AdditionalClassification,omitempty"` // AdditionalClassification:
	Classification           []BrandClassificationEnum `json:"Classification,omitempty"`           // Classification:
	LegSequence              []int32                   `json:"legSequence,omitempty"`              // LegSequence: the leg sequence
}

// BrandID is an object.
type BrandID struct {
	Type       string      `json:"@type"`                // Type:
	BrandRef   string      `json:"BrandRef,omitempty"`   // BrandRef: Used to reference another instance of this object in the same message
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Local indentifier within a given message for this object.
}

// BuildAncillaryOffersFromCatalogOfferings is an object.
type BuildAncillaryOffersFromCatalogOfferings struct {
	Type                                    string                     `json:"@type"`                                             // Type:
	CatalogOfferingIdentifier               CatalogOfferingIdentifier  `json:"CatalogOfferingIdentifier"`                         // CatalogOfferingIdentifier:
	CatalogOfferingsAncillaryListIdentifier *Identifier                `json:"CatalogOfferingsAncillaryListIdentifier,omitempty"` // CatalogOfferingsAncillaryListIdentifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	CatalogOfferingsIdentifier              CatalogOfferingsIdentifier `json:"CatalogOfferingsIdentifier"`                        // CatalogOfferingsIdentifier:
	ProductIdentifier                       ProductIdentifier          `json:"ProductIdentifier"`                                 // ProductIdentifier:
	Quantity                                int32                      `json:"Quantity,omitempty"`                                // Quantity: The quantity of ancillaries to be included in the Offer
	TravelerIdentifierRef                   *IdentifierRef             `json:"TravelerIdentifierRef,omitempty"`                   // TravelerIdentifierRef:
	IncludeUnsellableAncillariesInd         bool                       `json:"includeUnsellableAncillariesInd,omitempty"`         // IncludeUnsellableAncillariesInd: If true, the response will include unsellable ancillary options
}

// BuildAncillaryOffersFromCatalogOfferingsAirSeat is an object.
type BuildAncillaryOffersFromCatalogOfferingsAirSeat struct {
	Type                                     string                                     `json:"@type"`                                    // Type:
	BuildAncillaryOffersFromCatalogOfferings []BuildAncillaryOffersFromCatalogOfferings `json:"BuildAncillaryOffersFromCatalogOfferings"` // BuildAncillaryOffersFromCatalogOfferings:
	SeatAssignment                           string                                     `json:"SeatAssignment"`                           // SeatAssignment: The specific seat number to be assigned to a Traveler
}

// BuildFromCatalogOfferingHospitality is an object.
type BuildFromCatalogOfferingHospitality struct {
	Type                      string      `json:"@type,omitempty"`                     // Type:
	CatalogOfferingIdentifier *Identifier `json:"CatalogOfferingIdentifier,omitempty"` // CatalogOfferingIdentifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NumberOfRooms             int32       `json:"NumberOfRooms,omitempty"`             // NumberOfRooms: Number of rooms required.
}

// BuildFromCatalogOfferingsRequest is an object.
type BuildFromCatalogOfferingsRequest struct {
	Type                        string                        `json:"@type"`                                 // Type:
	AncillaryOfferingIdentifier []AncillaryOfferingIdentifier `json:"AncillaryOfferingIdentifier,omitempty"` // AncillaryOfferingIdentifier:
	CatalogOfferingIdentifier   CatalogOfferingIdentifier     `json:"CatalogOfferingIdentifier"`             // CatalogOfferingIdentifier:
	CatalogOfferingsIdentifier  CatalogOfferingsIdentifier    `json:"CatalogOfferingsIdentifier"`            // CatalogOfferingsIdentifier:
	ProductIdentifier           []ProductIdentifier           `json:"ProductIdentifier"`                     // ProductIdentifier:
}

// BuildFromCatalogOfferingsRequestAir is an object.
type BuildFromCatalogOfferingsRequestAir struct {
	Type                        string                        `json:"@type"`                                 // Type:
	AncillaryOfferingIdentifier []AncillaryOfferingIdentifier `json:"AncillaryOfferingIdentifier,omitempty"` // AncillaryOfferingIdentifier:
	CatalogOfferingIdentifier   CatalogOfferingIdentifier     `json:"CatalogOfferingIdentifier"`             // CatalogOfferingIdentifier:
	CatalogOfferingsIdentifier  CatalogOfferingsIdentifier    `json:"CatalogOfferingsIdentifier"`            // CatalogOfferingsIdentifier:
	PricingModifiersAir         *PricingModifiersAir          `json:"PricingModifiersAir,omitempty"`         // PricingModifiersAir:
	ProductIdentifier           []ProductIdentifier           `json:"ProductIdentifier"`                     // ProductIdentifier:
	SegmentSequence             []int32                       `json:"SegmentSequence,omitempty"`             // SegmentSequence: The segmentSequence within the product the action is being requested for. Used when multiple flights exist within a product. Only one product may be selected with this option.
}

// BuildFromCatalogOfferingsRequestAirChange is an object.
type BuildFromCatalogOfferingsRequestAirChange struct {
	Type                        string                        `json:"@type"`                                 // Type:
	AncillaryOfferingIdentifier []AncillaryOfferingIdentifier `json:"AncillaryOfferingIdentifier,omitempty"` // AncillaryOfferingIdentifier:
	CatalogOfferingIdentifier   CatalogOfferingIdentifier     `json:"CatalogOfferingIdentifier"`             // CatalogOfferingIdentifier:
	CatalogOfferingsIdentifier  CatalogOfferingsIdentifier    `json:"CatalogOfferingsIdentifier"`            // CatalogOfferingsIdentifier:
	PricingModifiersAirChange   PricingModifiersAirChange     `json:"PricingModifiersAirChange"`             // PricingModifiersAirChange:
	ProductIdentifier           []ProductIdentifier           `json:"ProductIdentifier"`                     // ProductIdentifier:
}

// BuildFromCatalogProductOfferingsRequest is an object.
type BuildFromCatalogProductOfferingsRequest struct {
	Type                              string                            `json:"@type,omitempty"`                    // Type:
	CatalogProductOfferingSelection   []CatalogProductOfferingSelection `json:"CatalogProductOfferingSelection"`    // CatalogProductOfferingSelection:
	CatalogProductOfferingsIdentifier CatalogProductOfferingsIdentifier `json:"CatalogProductOfferingsIdentifier"`  // CatalogProductOfferingsIdentifier:
	UpsellOfferingIdentifier          []UpsellOfferingIdentifier        `json:"UpsellOfferingIdentifier,omitempty"` // UpsellOfferingIdentifier:
}

// BuildFromCatalogProductOfferingsRequestAir is an object.
type BuildFromCatalogProductOfferingsRequestAir struct {
	Type                              string                            `json:"@type,omitempty"`                      // Type:
	CabinPreference                   *CabinPreference                  `json:"CabinPreference,omitempty"`            // CabinPreference:
	CatalogProductOfferingSelection   []CatalogProductOfferingSelection `json:"CatalogProductOfferingSelection"`      // CatalogProductOfferingSelection:
	CatalogProductOfferingsIdentifier CatalogProductOfferingsIdentifier `json:"CatalogProductOfferingsIdentifier"`    // CatalogProductOfferingsIdentifier:
	CustomResponseModifiersAir        *CustomResponseModifiersAir       `json:"CustomResponseModifiersAir,omitempty"` // CustomResponseModifiersAir: Modifiers to customize the result set
	FareRuleCategory                  []FareRuleCategoryEnum            `json:"FareRuleCategory,omitempty"`           // FareRuleCategory:
	FareRuleType                      FareRuleEnum                      `json:"FareRuleType,omitempty"`               // FareRuleType:
	PricingModifiersAir               *PricingModifiersAir              `json:"PricingModifiersAir,omitempty"`        // PricingModifiersAir:
	UpsellOfferingIdentifier          []UpsellOfferingIdentifier        `json:"UpsellOfferingIdentifier,omitempty"`   // UpsellOfferingIdentifier:
	InhibitBrandContentInd            bool                              `json:"inhibitBrandContentInd,omitempty"`     // InhibitBrandContentInd: If true, Brand content will not be returned with the Offer
	LowFareFinderInd                  bool                              `json:"lowFareFinderInd,omitempty"`           // LowFareFinderInd: If true, the price service will return the lowest fare available for the itinerary requested
	ReCheckInventoryInd               bool                              `json:"reCheckInventoryInd,omitempty"`        // ReCheckInventoryInd: If true, the price service will recheck inventory at the time of pricing the Offer
	ValidateInventoryInd              bool                              `json:"validateInventoryInd,omitempty"`       // ValidateInventoryInd: If true, the flight inventory will be checked during the price step
}

// BuildFromOffer is an object.
type BuildFromOffer struct {
	Type              string              `json:"@type"`                       // Type:
	OfferIdentifier   OfferIdentifier     `json:"OfferIdentifier"`             // OfferIdentifier:
	PaymentCriteria   *PaymentCriteria    `json:"PaymentCriteria,omitempty"`   // PaymentCriteria:
	ProductIdentifier []ProductIdentifier `json:"ProductIdentifier,omitempty"` // ProductIdentifier:
}

// BuildFromOfferAir is an object.
type BuildFromOfferAir struct {
	Type              string              `json:"@type"`                       // Type:
	OfferIdentifier   OfferIdentifier     `json:"OfferIdentifier"`             // OfferIdentifier:
	PaymentCriteria   *PaymentCriteria    `json:"PaymentCriteria,omitempty"`   // PaymentCriteria:
	ProductIdentifier []ProductIdentifier `json:"ProductIdentifier,omitempty"` // ProductIdentifier:
	SegmentSequence   []int32             `json:"SegmentSequence"`             // SegmentSequence: The segmentSequence within the product the action is being requested for. Used when multiple in a product
}

// BuildFromOfferList is an object.
type BuildFromOfferList struct {
	Type                string              `json:"@type,omitempty"`             // Type:
	OfferIdentifier     []OfferIdentifier   `json:"OfferIdentifier"`             // OfferIdentifier:
	OfferListIdentifier string              `json:"OfferListIdentifier"`         // OfferListIdentifier: The OfferListIdentifer (GUID) to retrieve the OfferList from cache
	ProductIdentifier   []ProductIdentifier `json:"ProductIdentifier,omitempty"` // ProductIdentifier:
	SegmentSequence     int32               `json:"SegmentSequence,omitempty"`   // SegmentSequence: The segmentSequence within the product the action is being requested for. Used when multiple exist within a product
}

// BuildFromProductsRequest is an object.
type BuildFromProductsRequest struct {
	Type string `json:"@type"` // Type:
}

// BuildFromProductsRequestAir is an object.
type BuildFromProductsRequestAir struct {
	Type                       string                       `json:"@type"`                                // Type:
	CustomResponseModifiersAir []CustomResponseModifiersAir `json:"CustomResponseModifiersAir,omitempty"` // CustomResponseModifiersAir:
	PassengerCriteria          []PassengerCriteria          `json:"PassengerCriteria"`                    // PassengerCriteria:
	PricingModifiersAir        PricingModifiersAir          `json:"PricingModifiersAir"`                  // PricingModifiersAir:
	ProductCriteriaAir         []ProductCriteriaAir         `json:"ProductCriteriaAir"`                   // ProductCriteriaAir:
}

// BuildFromProductsRequestAirChange is an object.
type BuildFromProductsRequestAirChange struct {
	Type                      string                     `json:"@type"`                               // Type:
	PricingModifiersAirChange *PricingModifiersAirChange `json:"PricingModifiersAirChange,omitempty"` // PricingModifiersAirChange:
	ProductCriteriaAir        []ProductCriteriaAir       `json:"ProductCriteriaAir"`                  // ProductCriteriaAir:
}

// BuildFromReservationWorkbench is an object.
type BuildFromReservationWorkbench struct {
	Type                  string                 `json:"@type,omitempty"`                 // Type:
	OfferIdentifier       *OfferIdentifier       `json:"OfferIdentifier,omitempty"`       // OfferIdentifier:
	ProductIdentifier     []ProductIdentifier    `json:"ProductIdentifier,omitempty"`     // ProductIdentifier:
	ReservationIdentifier *ReservationIdentifier `json:"ReservationIdentifier,omitempty"` // ReservationIdentifier:
	SegmentSequenceList   []int32                `json:"SegmentSequenceList,omitempty"`   // SegmentSequenceList: The segmentSequence within the product the action is being requested for. Used when multiple flights exist within a product. Only one product may be selected with this option.
}

// CabinPreference is an object.
type CabinPreference struct {
	Type           string                  `json:"@type,omitempty"`          // Type:
	Cabins         []CabinAirEnum          `json:"cabins,omitempty"`         // Cabins:
	LegSequence    []int32                 `json:"legSequence,omitempty"`    // LegSequence: Leg sequence
	PreferenceType CabinPreferenceTypeEnum `json:"preferenceType,omitempty"` // PreferenceType:
}

// CalculatedFareAdjustment is an object.
type CalculatedFareAdjustment struct {
	Type string `json:"@type"` // Type:
}

// CalculatedFareAdjustmentDiscount is an object.
type CalculatedFareAdjustmentDiscount struct {
	Type          string         `json:"@type"`                   // Type:
	AmountPercent *AmountPercent `json:"AmountPercent,omitempty"` // AmountPercent:
}

// CalculatedFareAdjustmentIncrease is an object.
type CalculatedFareAdjustmentIncrease struct {
	Type          string        `json:"@type"`         // Type:
	AmountPercent AmountPercent `json:"AmountPercent"` // AmountPercent:
}

// Calculation is an object.
type Calculation struct {
	Applicability *Comment        `json:"Applicability,omitempty"` // Applicability: Textual information.
	MaxQuantity   int32           `json:"MaxQuantity,omitempty"`   // MaxQuantity: The maximum quantity allowed for a charge e.g Baby seat charged at $10 per day for a maximum of 10 days
	Percent       float32         `json:"Percent,omitempty"`       // Percent: Used when the charge is based on a percentage of a TotalAmount
	Quantity      int32           `json:"Quantity,omitempty"`      // Quantity: The quantity used in the calculation of the vehicle charge e.g 2 x $500 per week
	TotalAmount   *CurrencyAmount `json:"TotalAmount,omitempty"`   // TotalAmount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	UnitAmount    *CurrencyAmount `json:"UnitAmount,omitempty"`    // UnitAmount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	UnitName      RatePeriodEnum  `json:"UnitName,omitempty"`      // UnitName: The time period for a rate such as daily, weekly, monthly
}

// Cancel is an object.
type Cancel struct {
	Type            string   `json:"@type"`                     // Type:
	ProductRefs     []string `json:"ProductRefs,omitempty"`     // ProductRefs: The productRefs this change or cancel applies to
	SegmentSequence []int32  `json:"SegmentSequence,omitempty"` // SegmentSequence: The SegmentSequence of the product this change or cancel applies to
}

// CancelIndeterminate is an object.
type CancelIndeterminate struct {
	Type             string   `json:"@type"`                      // Type:
	IndeterminateInd bool     `json:"IndeterminateInd,omitempty"` // IndeterminateInd: structured fare rules could not be determined for this category
	ProductRefs      []string `json:"ProductRefs,omitempty"`      // ProductRefs: The productRefs this change or cancel applies to
	SegmentSequence  []int32  `json:"SegmentSequence,omitempty"`  // SegmentSequence: The SegmentSequence of the product this change or cancel applies to
}

// CancelNotPermitted is an object.
type CancelNotPermitted struct {
	Type            string            `json:"@type"`                     // Type:
	NotPermittedInd bool              `json:"NotPermittedInd,omitempty"` // NotPermittedInd: Changes are not permitted
	ProductRefs     []string          `json:"ProductRefs,omitempty"`     // ProductRefs: The productRefs this change or cancel applies to
	SegmentSequence []int32           `json:"SegmentSequence,omitempty"` // SegmentSequence: The SegmentSequence of the product this change or cancel applies to
	PenaltyTypes    []PenaltyTypeEnum `json:"penaltyTypes,omitempty"`    // PenaltyTypes:
}

// CancelPenalty is an object.
type CancelPenalty struct {
	Type         string           `json:"@type,omitempty"`        // Type:
	Deadline     *Deadline        `json:"Deadline,omitempty"`     // Deadline:
	HotelPenalty *HotelPenalty    `json:"HotelPenalty,omitempty"` // HotelPenalty:
	Refundable   YesNoUnknownEnum `json:"Refundable,omitempty"`   // Refundable: Yes , No , Unknown
}

// CancelPermitted is an object.
type CancelPermitted struct {
	Type                     string               `json:"@type"`                              // Type:
	Penalty                  []Penalty            `json:"Penalty,omitempty"`                  // Penalty:
	PenaltyAppliesTo         PenaltyAppliesToEnum `json:"PenaltyAppliesTo,omitempty"`         // PenaltyAppliesTo: Penalty applicable list
	ProductRefs              []string             `json:"ProductRefs,omitempty"`              // ProductRefs: The productRefs this change or cancel applies to
	SegmentSequence          []int32              `json:"SegmentSequence,omitempty"`          // SegmentSequence: The SegmentSequence of the product this change or cancel applies to
	HigherPenatltyAppliesInd bool                 `json:"higherPenatltyAppliesInd,omitempty"` // HigherPenatltyAppliesInd: If true, when an amount and a percent are specified in the Penalty then the higher of these apply
	PenaltyTypes             []PenaltyTypeEnum    `json:"penaltyTypes"`                       // PenaltyTypes:
}

// Cancellation is an object.
type Cancellation struct {
	Type string `json:"@type"` // Type:
}

// CancellationHold is an object.
type CancellationHold struct {
	Type    string   `json:"@type"`             // Type:
	Locator *Locator `json:"Locator,omitempty"` // Locator: Contains the locator (PNR or external locator) or cancellation number for the reservation, order, or offer
}

// cardNumberMaskPattern is the validation pattern for CardNumber.Mask
var cardNumberMaskPattern = regexp.MustCompile(`[0-9a-zA-Z]{1,32}`)

// cardNumberTokenPattern is the validation pattern for CardNumber.Token
var cardNumberTokenPattern = regexp.MustCompile(`[0-9a-zA-Z]{1,32}`)

// CardNumber is an object.
type CardNumber struct {
	Type                 string                      `json:"@type,omitempty"`                // Type:
	ErrorWarning         *ErrorWarning               `json:"ErrorWarning,omitempty"`         // ErrorWarning:
	PlainText            string                      `json:"PlainText,omitempty"`            // PlainText: Don't use this unless it is REALLY ok to not use encryption. Non-secure (plain text) value.
	AuthenticationMethod EncryptionTokenTypeAuthEnum `json:"authenticationMethod,omitempty"` // AuthenticationMethod: Type of Authentication
	EncryptedValue       string                      `json:"encryptedValue,omitempty"`       // EncryptedValue: Encrypted value
	EncryptionKey        string                      `json:"encryptionKey,omitempty"`        // EncryptionKey: Note: This contains a key required to retrieve the full payment instrument details compliant with PCI DSS standards.
	EncryptionKeyMethod  string                      `json:"encryptionKeyMethod,omitempty"`  // EncryptionKeyMethod: Developer: This contains a reference to the key generation method being used - this is NOT the key value.
	EncryptionMethod     string                      `json:"encryptionMethod,omitempty"`     // EncryptionMethod: OpenTravel Best Practice: Encryption Method: When using the OpenTravel Encryption element, it is RECOMMENDED that all trading partners be informed of all encryption methods being used in advance of implementation to ensure message processing compatibility.
	Mask                 string                      `json:"mask,omitempty"`                 // Mask: Masked Value
	Token                string                      `json:"token,omitempty"`                // Token: Token value
	TokenProviderID      string                      `json:"tokenProviderID,omitempty"`      // TokenProviderID: Developer: This contains a provider ID if multiple providers are used for secure information exchange.
}

// CarrierPreference is an object.
type CarrierPreference struct {
	Type           string                    `json:"@type,omitempty"`       // Type:
	Carriers       []string                  `json:"carriers"`              // Carriers: Carrier airline codes
	LegSequence    []int32                   `json:"legSequence,omitempty"` // LegSequence: Leg sequence
	PreferenceType CarrierPreferenceTypeEnum `json:"preferenceType"`        // PreferenceType:
}

// CatalogOffering is an object.
type CatalogOffering struct {
	Type               string      `json:"@type"`                        // Type:
	CatalogOfferingRef string      `json:"CatalogOfferingRef,omitempty"` // CatalogOfferingRef: Used to reference another instance of this object in the same message
	Identifier         *Identifier `json:"Identifier,omitempty"`         // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Price              struct {
		Type                string              `json:"@type,omitempty"`
		Base                float32             `json:"Base,omitempty"`
		CurrencyCode        CurrencyCode        `json:"CurrencyCode,omitempty"`
		PriceBreakdown      []PriceBreakdown    `json:"PriceBreakdown,omitempty"`
		TotalFees           float32             `json:"TotalFees,omitempty"`
		TotalPrice          float32             `json:"TotalPrice,omitempty"`
		TotalTaxes          float32             `json:"TotalTaxes,omitempty"`
		VendorCurrencyTotal VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"`
		Id                  string              `json:"id,omitempty"`
	} `json:"Price"` // Price:
	ProductOptions     []ProductOptions `json:"ProductOptions"` // ProductOptions:
	TermsAndConditions *struct {
		Type                  string            `json:"@type,omitempty"`
		CustomerLoyalty       []CustomerLoyalty `json:"CustomerLoyalty,omitempty"`
		ExpiryDate            time.Time         `json:"ExpiryDate,omitempty"`
		Identifier            Identifier        `json:"Identifier,omitempty"`
		Id                    string            `json:"id,omitempty"`
		TermsAndConditionsRef string            `json:"termsAndConditionsRef,omitempty"`
	} `json:"TermsAndConditions,omitempty"` // TermsAndConditions:
	Id string `json:"id,omitempty"` // Id: Local indentifier within a given message for this object.
}

// CatalogOfferingID is an object.
type CatalogOfferingID struct {
	Type               string      `json:"@type"`                        // Type:
	CatalogOfferingRef string      `json:"CatalogOfferingRef,omitempty"` // CatalogOfferingRef: Used to reference another instance of this object in the same message
	Identifier         *Identifier `json:"Identifier,omitempty"`         // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id                 string      `json:"id,omitempty"`                 // Id: Local indentifier within a given message for this object.
}

// CatalogOfferingIdentifier is an object.
type CatalogOfferingIdentifier struct {
	CatalogOfferingRef string      `json:"CatalogOfferingRef,omitempty"` // CatalogOfferingRef: Used to reference another instance of this object in the same message
	Identifier         *Identifier `json:"Identifier,omitempty"`         // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id                 string      `json:"id,omitempty"`                 // Id: Local indentifier within a given message for this object.
}

// CatalogOfferingModify is an object.
type CatalogOfferingModify struct {
	Type               string      `json:"@type"`                        // Type:
	CatalogOfferingRef string      `json:"CatalogOfferingRef,omitempty"` // CatalogOfferingRef: Used to reference another instance of this object in the same message
	Identifier         *Identifier `json:"Identifier,omitempty"`         // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ModifyPrice        ModifyPrice `json:"ModifyPrice"`                  // ModifyPrice:
	Price              struct {
		Type                string              `json:"@type,omitempty"`
		Base                float32             `json:"Base,omitempty"`
		CurrencyCode        CurrencyCode        `json:"CurrencyCode,omitempty"`
		PriceBreakdown      []PriceBreakdown    `json:"PriceBreakdown,omitempty"`
		TotalFees           float32             `json:"TotalFees,omitempty"`
		TotalPrice          float32             `json:"TotalPrice,omitempty"`
		TotalTaxes          float32             `json:"TotalTaxes,omitempty"`
		VendorCurrencyTotal VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"`
		Id                  string              `json:"id,omitempty"`
	} `json:"Price"` // Price:
	ProductOptions     []ProductOptions `json:"ProductOptions"` // ProductOptions:
	TermsAndConditions *struct {
		Type                  string            `json:"@type,omitempty"`
		CustomerLoyalty       []CustomerLoyalty `json:"CustomerLoyalty,omitempty"`
		ExpiryDate            time.Time         `json:"ExpiryDate,omitempty"`
		Identifier            Identifier        `json:"Identifier,omitempty"`
		Id                    string            `json:"id,omitempty"`
		TermsAndConditionsRef string            `json:"termsAndConditionsRef,omitempty"`
	} `json:"TermsAndConditions,omitempty"` // TermsAndConditions:
	Id string `json:"id,omitempty"` // Id: Local indentifier within a given message for this object.
}

// CatalogOfferings is an object.
type CatalogOfferings struct {
	Type              string                `json:"@type"`                       // Type:
	AncillaryOffering []AncillaryOfferingID `json:"AncillaryOffering,omitempty"` // AncillaryOffering:
	CatalogOffering   []CatalogOffering     `json:"CatalogOffering"`             // CatalogOffering:
	Identifier        *Identifier           `json:"Identifier,omitempty"`        // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id                string                `json:"id,omitempty"`                // Id: Local indentifier within a given message for this object.
}

// CatalogOfferingsAirChangeRequest is an object.
type CatalogOfferingsAirChangeRequest struct {
	Type                          string                         `json:"@type"`                                   // Type:
	PassengerCriteria             []PassengerCriteria            `json:"PassengerCriteria"`                       // PassengerCriteria:
	PricingModifiersAirChange     *PricingModifiersAirChange     `json:"PricingModifiersAirChange,omitempty"`     // PricingModifiersAirChange:
	SearchControlConsoleChannelID *SearchControlConsoleChannelID `json:"SearchControlConsoleChannelID,omitempty"` // SearchControlConsoleChannelID:
	SearchCriteriaFlight          []SearchCriteriaFlight         `json:"SearchCriteriaFlight"`                    // SearchCriteriaFlight:
	SearchModifiersAir            *SearchModifiersAir            `json:"SearchModifiersAir,omitempty"`            // SearchModifiersAir:
	CatalogOfferingsPerPage       int32                          `json:"catalogOfferingsPerPage,omitempty"`       // CatalogOfferingsPerPage: Catalog Offerings per page value
	DetailViewInd                 bool                           `json:"detailViewInd,omitempty"`                 // DetailViewInd:
	ReturnBrandedFaresInd         bool                           `json:"returnBrandedFaresInd,omitempty"`         // ReturnBrandedFaresInd:
	UpsellInd                     bool                           `json:"upsellInd,omitempty"`                     // UpsellInd:
}

// CatalogOfferingsAirChangeRequestReservation is an object.
type CatalogOfferingsAirChangeRequestReservation struct {
	Type                          string                         `json:"@type"`                                   // Type:
	BuildFromReservationWorkbench BuildFromReservationWorkbench  `json:"BuildFromReservationWorkbench"`           // BuildFromReservationWorkbench:
	PassengerCriteria             []PassengerCriteria            `json:"PassengerCriteria"`                       // PassengerCriteria:
	PricingModifiersAirChange     *PricingModifiersAirChange     `json:"PricingModifiersAirChange,omitempty"`     // PricingModifiersAirChange:
	SearchControlConsoleChannelID *SearchControlConsoleChannelID `json:"SearchControlConsoleChannelID,omitempty"` // SearchControlConsoleChannelID:
	SearchCriteriaFlight          []SearchCriteriaFlight         `json:"SearchCriteriaFlight"`                    // SearchCriteriaFlight:
	SearchModifiersAir            *SearchModifiersAir            `json:"SearchModifiersAir,omitempty"`            // SearchModifiersAir:
	CatalogOfferingsPerPage       int32                          `json:"catalogOfferingsPerPage,omitempty"`       // CatalogOfferingsPerPage: Catalog Offerings per page value
	DetailViewInd                 bool                           `json:"detailViewInd,omitempty"`                 // DetailViewInd:
	ReturnBrandedFaresInd         bool                           `json:"returnBrandedFaresInd,omitempty"`         // ReturnBrandedFaresInd:
	UpsellInd                     bool                           `json:"upsellInd,omitempty"`                     // UpsellInd:
}

// CatalogOfferingsAirChangeResponse is an object.
type CatalogOfferingsAirChangeResponse struct {
	CatalogOfferings       *CatalogOfferings        `json:"CatalogOfferings,omitempty"`       // CatalogOfferings:
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"` // CurrencyRateConversion:
	Identifier             *Identifier              `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`              // NextSteps:
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`          // ReferenceList:
	Result                 *Result                  `json:"Result,omitempty"`                 // Result:
	TraceId                string                   `json:"traceId,omitempty"`                // TraceId: Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TransactionId          string                   `json:"transactionId,omitempty"`          // TransactionId: Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
}

// CatalogOfferingsID is an object.
type CatalogOfferingsID struct {
	Type       string      `json:"@type"`                // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Local indentifier within a given message for this object.
}

// CatalogOfferingsIdentifier is an object.
type CatalogOfferingsIdentifier struct {
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Local indentifier within a given message for this object.
}

// CatalogOfferingsRequest is an object.
type CatalogOfferingsRequest struct {
	Type                          string                         `json:"@type,omitempty"`                         // Type:
	SearchControlConsoleChannelID *SearchControlConsoleChannelID `json:"SearchControlConsoleChannelID,omitempty"` // SearchControlConsoleChannelID:
}

// catalogOfferingsRequestHospitalityRequestedCurrencyPattern is the validation pattern for CatalogOfferingsRequestHospitality.RequestedCurrency
var catalogOfferingsRequestHospitalityRequestedCurrencyPattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// CatalogOfferingsRequestHospitality is an object.
type CatalogOfferingsRequestHospitality struct {
	Type                          string                         `json:"@type,omitempty"`                         // Type:
	HotelSearchCriterion          *HotelSearchCriterion          `json:"HotelSearchCriterion,omitempty"`          // HotelSearchCriterion:
	MaximumAmount                 *CurrencyAmount                `json:"MaximumAmount,omitempty"`                 // MaximumAmount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	MinimumAmount                 *CurrencyAmount                `json:"MinimumAmount,omitempty"`                 // MinimumAmount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	SearchControlConsoleChannelID *SearchControlConsoleChannelID `json:"SearchControlConsoleChannelID,omitempty"` // SearchControlConsoleChannelID:
	StayDates                     DateOrDateWindows              `json:"StayDates"`                               // StayDates: Indicates the expiry date.
	MaxResponseWaitTime           int32                          `json:"maxResponseWaitTime,omitempty"`           // MaxResponseWaitTime: Maximum time (in milliseconds) to wait for provider responses before returning a response to the consumer of this service
	RequestedCurrency             string                         `json:"requestedCurrency,omitempty"`             // RequestedCurrency: You can use requested currency to request conversion rate information. The response will return the currencyRateConversion object which will contain conversion rate of the requested currency.
	VerboseResponseInd            bool                           `json:"verboseResponseInd,omitempty"`            // VerboseResponseInd: Used to specify that a verbose response is to be returned.  Verbose responses repeat the Property information in each Product and do not return the reference list.
}

// CatalogProductOfferingIdentifier is an object.
type CatalogProductOfferingIdentifier struct {
	CatalogProductOfferingRef string      `json:"CatalogProductOfferingRef,omitempty"` // CatalogProductOfferingRef: Used to reference another instance of this object in the same message
	Identifier                *Identifier `json:"Identifier,omitempty"`                // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id                        string      `json:"id,omitempty"`                        // Id: Local indentifier within a given message for this object.
}

// CatalogProductOfferingSelection is an object.
type CatalogProductOfferingSelection struct {
	Type                             string                           `json:"@type,omitempty"`                  // Type:
	CatalogProductOfferingIdentifier CatalogProductOfferingIdentifier `json:"CatalogProductOfferingIdentifier"` // CatalogProductOfferingIdentifier:
	ProductBrandOfferingIdentifier   Identifier                       `json:"ProductBrandOfferingIdentifier"`   // ProductBrandOfferingIdentifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ProductIdentifier                []ProductIdentifier              `json:"ProductIdentifier,omitempty"`      // ProductIdentifier:
	SegmentSequence                  []int32                          `json:"SegmentSequence,omitempty"`        // SegmentSequence:
}

// CatalogProductOfferingsIdentifier is an object.
type CatalogProductOfferingsIdentifier struct {
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Local indentifier within a given message for this object.
}

// CatalogProductOfferingsRequest is an object.
type CatalogProductOfferingsRequest struct {
	Type                          string                         `json:"@type"`                                   // Type:
	SearchControlConsoleChannelID *SearchControlConsoleChannelID `json:"SearchControlConsoleChannelID,omitempty"` // SearchControlConsoleChannelID:
}

// CatalogProductOfferingsRequestAir is an object.
type CatalogProductOfferingsRequestAir struct {
	Type                          string                         `json:"@type"`                                   // Type:
	CustomResponseModifiersAir    *CustomResponseModifiersAir    `json:"CustomResponseModifiersAir,omitempty"`    // CustomResponseModifiersAir: Modifiers to customize the result set
	PassengerCriteria             []PassengerCriteria            `json:"PassengerCriteria,omitempty"`             // PassengerCriteria:
	PaymentCriteria               *PaymentCriteria               `json:"PaymentCriteria,omitempty"`               // PaymentCriteria:
	PricingModifiersAir           *PricingModifiersAir           `json:"PricingModifiersAir,omitempty"`           // PricingModifiersAir:
	PseudoCityInfo                *PseudoCityInfo                `json:"PseudoCityInfo,omitempty"`                // PseudoCityInfo: This field is obsolete. For pricing and ticketing PCC overrides use PricingModifiersAir
	SearchControlConsoleChannelID *SearchControlConsoleChannelID `json:"SearchControlConsoleChannelID,omitempty"` // SearchControlConsoleChannelID:
	SearchCriteriaFlight          []SearchCriteriaFlight         `json:"SearchCriteriaFlight,omitempty"`          // SearchCriteriaFlight:
	SearchModifiersAir            *SearchModifiersAir            `json:"SearchModifiersAir,omitempty"`            // SearchModifiersAir:
	SearchType                    SearchTypeEnum                 `json:"SearchType,omitempty"`                    // SearchType:
	ContentSourceList             []ContentSourceEnum            `json:"contentSourceList,omitempty"`             // ContentSourceList:
	DetailViewInd                 bool                           `json:"detailViewInd,omitempty"`                 // DetailViewInd: if true, detail view should be returned
	ExcludeMixedBrandsInd         bool                           `json:"excludeMixedBrandsInd,omitempty"`         // ExcludeMixedBrandsInd: If true, mixed brands will be inhibited from the response
	InhibitBrandContentInd        bool                           `json:"inhibitBrandContentInd,omitempty"`        // InhibitBrandContentInd: if true, brand infromation will be supressed.
	MaxNumberOfOffersToReturn     int32                          `json:"maxNumberOfOffersToReturn,omitempty"`     // MaxNumberOfOffersToReturn: This attribute is deprecated and not validated if sent
	MaxNumberOfUpsellsToReturn    int32                          `json:"maxNumberOfUpsellsToReturn,omitempty"`    // MaxNumberOfUpsellsToReturn: The maximum number of upsells to return
	NumberOfDownsellsToReturn     int32                          `json:"numberOfDownsellsToReturn,omitempty"`     // NumberOfDownsellsToReturn: The number of downsells to return
	OffersPerPage                 int32                          `json:"offersPerPage,omitempty"`                 // OffersPerPage: Number of offers per page
}

// Change is an object.
type Change struct {
	Type            string   `json:"@type"`                     // Type:
	ProductRefs     []string `json:"ProductRefs,omitempty"`     // ProductRefs: The productRefs this change or cancel applies to
	SegmentSequence []int32  `json:"SegmentSequence,omitempty"` // SegmentSequence: The SegmentSequence of the product this change or cancel applies to
}

// ChangeFeeCollectionMethod is an object.
type ChangeFeeCollectionMethod struct {
	ChangeFeeIssuedSeparatelyInd bool                 `json:"changeFeeIssuedSeparatelyInd,omitempty"` // ChangeFeeIssuedSeparatelyInd: if true, the change fee will be issued as a separate transaction to the residual amount
	Code                         string               `json:"code"`                                   // Code: The code value
	Description                  string               `json:"description,omitempty"`                  // Description: The description value
	SubCode                      string               `json:"subCode,omitempty"`                      // SubCode: The subcode value
	TaxIncludedInBaseAmountInd   bool                 `json:"taxIncludedInBaseAmountInd,omitempty"`   // TaxIncludedInBaseAmountInd: If true, the tax  on the fee will be included in the base fee amount and sent as a single value to the supplier for fulfilment
	Value                        *ChangeFeeMethodEnum `json:"value,omitempty"`                        // Value:
}

// ChangeFeeMethodEnum is an object.
type ChangeFeeMethodEnum struct {
	Value ChangeFeeMethodEnumBase `json:"value,omitempty"` // Value:
}

// ChangeIndeterminate is an object.
type ChangeIndeterminate struct {
	Type             string   `json:"@type"`                      // Type:
	IndeterminateInd bool     `json:"IndeterminateInd,omitempty"` // IndeterminateInd: structured fare rules could not be determined for this category
	ProductRefs      []string `json:"ProductRefs,omitempty"`      // ProductRefs: The productRefs this change or cancel applies to
	SegmentSequence  []int32  `json:"SegmentSequence,omitempty"`  // SegmentSequence: The SegmentSequence of the product this change or cancel applies to
}

// ChangeNotPermitted is an object.
type ChangeNotPermitted struct {
	Type            string            `json:"@type"`                     // Type:
	NotPermittedInd bool              `json:"NotPermittedInd,omitempty"` // NotPermittedInd: Changes are not permitted
	ProductRefs     []string          `json:"ProductRefs,omitempty"`     // ProductRefs: The productRefs this change or cancel applies to
	SegmentSequence []int32           `json:"SegmentSequence,omitempty"` // SegmentSequence: The SegmentSequence of the product this change or cancel applies to
	PenaltyTypes    []PenaltyTypeEnum `json:"penaltyTypes,omitempty"`    // PenaltyTypes:
}

// ChangeOptions is an object.
type ChangeOptions struct {
	ChangePenaltyRange *ChangePenaltyRange `json:"ChangePenaltyRange,omitempty"` // ChangePenaltyRange:
	ChangeTypes        []ChangeTypeENUM    `json:"changeTypes"`                  // ChangeTypes:
}

// ChangePenaltyRange is an object.
type ChangePenaltyRange struct {
	Type    string         `json:"@type,omitempty"`   // Type:
	Maximum *AmountPercent `json:"Maximum,omitempty"` // Maximum:
	Minimum *AmountPercent `json:"Minimum,omitempty"` // Minimum:
}

// ChangePermitted is an object.
type ChangePermitted struct {
	Type                     string               `json:"@type"`                              // Type:
	Penalty                  []Penalty            `json:"Penalty,omitempty"`                  // Penalty:
	PenaltyAppliesTo         PenaltyAppliesToEnum `json:"PenaltyAppliesTo,omitempty"`         // PenaltyAppliesTo: Penalty applicable list
	ProductRefs              []string             `json:"ProductRefs,omitempty"`              // ProductRefs: The productRefs this change or cancel applies to
	SegmentSequence          []int32              `json:"SegmentSequence,omitempty"`          // SegmentSequence: The SegmentSequence of the product this change or cancel applies to
	HigherPenatltyAppliesInd bool                 `json:"higherPenatltyAppliesInd,omitempty"` // HigherPenatltyAppliesInd: If true, when an amount and a percent are specified in the Penalty then the higher of these apply
	PenaltyTypes             []PenaltyTypeEnum    `json:"penaltyTypes"`                       // PenaltyTypes:
}

// checkInOutPolicyCheckInTimePattern is the validation pattern for CheckInOutPolicy.CheckInTime
var checkInOutPolicyCheckInTimePattern = regexp.MustCompile(`(([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)((:?)[0-5]\d)?([\.,]\d+(?!:))?`)

// checkInOutPolicyCheckOutTimePattern is the validation pattern for CheckInOutPolicy.CheckOutTime
var checkInOutPolicyCheckOutTimePattern = regexp.MustCompile(`(([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)((:?)[0-5]\d)?([\.,]\d+(?!:))?`)

// CheckInOutPolicy is an object.
type CheckInOutPolicy struct {
	Type         string                    `json:"@type,omitempty"`       // Type:
	Description  []TextTitleAndDescription `json:"Description,omitempty"` // Description:
	CheckInTime  string                    `json:"checkInTime"`           // CheckInTime: Check-in time
	CheckOutTime string                    `json:"checkOutTime"`          // CheckOutTime: Check-out time
	MinimumAge   int32                     `json:"minimumAge,omitempty"`  // MinimumAge: Minimum age of guest checking in or out
}

// classOfServiceAvailabilityValuePattern is the validation pattern for ClassOfServiceAvailability.Value
var classOfServiceAvailabilityValuePattern = regexp.MustCompile(`([a-zA-Z0-9]{1,2})`)

// ClassOfServiceAvailability is an object. The class of service
type ClassOfServiceAvailability struct {
	Number int32                   `json:"number,omitempty"` // Number: The class of service number value
	Status *AvailabilityStatusENUM `json:"status,omitempty"` // Status:
	Value  string                  `json:"value,omitempty"`  // Value:
}

// ClassOfServicePreference is an object.
type ClassOfServicePreference struct {
	Type             string                           `json:"@type,omitempty"`          // Type:
	ClassesOfService []string                         `json:"ClassesOfService"`         // ClassesOfService: Allows user to specify which class(s) of service they want returned in CatalogOfferings
	PreferenceType   ClassOfServicePreferenceTypeEnum `json:"PreferenceType,omitempty"` // PreferenceType:
	LegSequence      []int32                          `json:"legSequence,omitempty"`    // LegSequence: The legSequence value
}

// Code is an object. Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
type Code struct {
	CodeContext string `json:"codeContext,omitempty"` // CodeContext: Code Context
	Value       string `json:"value,omitempty"`       // Value:
}

// Comment is an object. Textual information.
type Comment struct {
	Id       string `json:"id,omitempty"`       // Id: Local indentifier within a given message for this object.
	Language string `json:"language,omitempty"` // Language: Language code using ISO-639 standard
	Name     string `json:"name,omitempty"`     // Name: Title
	Value    string `json:"value,omitempty"`    // Value:
}

// Comments is an object.
type Comments struct {
	Type    string    `json:"@type,omitempty"`   // Type:
	Comment []Comment `json:"Comment,omitempty"` // Comment:
}

// Commission is an object.
type Commission struct {
	Type        string         `json:"@type"`                 // Type:
	Application CommissionEnum `json:"application,omitempty"` // Application: Type of commission
}

// CommissionAmount is an object.
type CommissionAmount struct {
	Type        string          `json:"@type"`                 // Type:
	Amount      *CurrencyAmount `json:"Amount,omitempty"`      // Amount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	Application CommissionEnum  `json:"application,omitempty"` // Application: Type of commission
}

// CommissionPercent is an object.
type CommissionPercent struct {
	Type        string         `json:"@type"`                 // Type:
	Percent     float32        `json:"Percent,omitempty"`     // Percent: Percent amount of commission
	Application CommissionEnum `json:"application,omitempty"` // Application: Type of commission
}

// Commissions is an object.
type Commissions struct {
	Type                  string                  `json:"@type,omitempty"`                 // Type:
	ApplyTo               CommissionApplyEnum     `json:"ApplyTo,omitempty"`               // ApplyTo:
	Commission            Commission              `json:"Commission"`                      // Commission:
	TravelerIdentifierRef []TravelerIdentifierRef `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
}

// CompanyName is an object. Identifies a company by name.
type CompanyName struct {
	Code           string   `json:"code,omitempty"`           // Code: Identifies a company by the company code
	CodeContext    string   `json:"codeContext,omitempty"`    // CodeContext: Identifies the context of the identifying code,such as DUNS,IATA or internal code
	Department     string   `json:"department,omitempty"`     // Department: The department name or ID with which the contact is associated
	Division       string   `json:"division,omitempty"`       // Division: The division name or ID with which the contact is associated
	Id             string   `json:"id,omitempty"`             // Id: Use this id to internally identify this company in NextSteps
	ShortName      string   `json:"shortName,omitempty"`      // ShortName: Used to provide the company common name
	SystemOfRecord []string `json:"systemOfRecord,omitempty"` // SystemOfRecord: The system(s) that maintain the data
	Value          string   `json:"value,omitempty"`          // Value:
}

// Confirmation is an object.
type Confirmation struct {
	Type string `json:"@type"` // Type:
}

// ConfirmationHold is an object.
type ConfirmationHold struct {
	Type                      string       `json:"@type"`                 // Type:
	Locator                   Locator      `json:"Locator"`               // Locator: Contains the locator (PNR or external locator) or cancellation number for the reservation, order, or offer
	OfferStatus               *OfferStatus `json:"OfferStatus,omitempty"` // OfferStatus:
	ShoppingCartProductStatus *struct {
		Type      string      `json:"@type,omitempty"`
		StatusAir []StatusAir `json:"StatusAir,omitempty"`
	} `json:"ShoppingCartProductStatus,omitempty"` // ShoppingCartProductStatus:
}

// ConnectionPreferences is an object.
type ConnectionPreferences struct {
	Type                  string                            `json:"@type,omitempty"`                 // Type:
	ConnectionPoint       []string                          `json:"ConnectionPoint,omitempty"`       // ConnectionPoint:
	LegSequence           []int32                           `json:"legSequence,omitempty"`           // LegSequence: Leg sequence
	MaxConnectionDuration string                            `json:"maxConnectionDuration,omitempty"` // MaxConnectionDuration: The maximum acceptable duration of the connection ISO8601
	MaxOvernightDuration  string                            `json:"maxOvernightDuration,omitempty"`  // MaxOvernightDuration: The maximum acceptable overnight duration of the connection ISO8601
	PreferenceType        ConnectionPointPreferenceTypeENUM `json:"preferenceType,omitempty"`        // PreferenceType: Preference type - preferred, permitted or prohibited. Preferred is not permitted as a preference type and will be ignored
}

// ConnectionPreferencesAir is an object.
type ConnectionPreferencesAir struct {
	Type                  string                            `json:"@type,omitempty"`                 // Type:
	ConnectionPoint       []string                          `json:"ConnectionPoint,omitempty"`       // ConnectionPoint:
	FlightType            *FlightType                       `json:"FlightType,omitempty"`            // FlightType:
	LegSequence           []int32                           `json:"legSequence,omitempty"`           // LegSequence: Leg sequence
	MaxConnectionDuration string                            `json:"maxConnectionDuration,omitempty"` // MaxConnectionDuration: The maximum acceptable duration of the connection ISO8601
	MaxOvernightDuration  string                            `json:"maxOvernightDuration,omitempty"`  // MaxOvernightDuration: The maximum acceptable overnight duration of the connection ISO8601
	PreferenceType        ConnectionPointPreferenceTypeENUM `json:"preferenceType,omitempty"`        // PreferenceType: Preference type - preferred, permitted or prohibited. Preferred is not permitted as a preference type and will be ignored
}

// ConversionRate is an object. A conversion metric from standard to another with the contextual authority such as IATA, OAG, ISO, etc.
type ConversionRate struct {
	RateAsOf      time.Time `json:"rateAsOf,omitempty"`      // RateAsOf: Rate as of
	RateAuthority string    `json:"rateAuthority,omitempty"` // RateAuthority: Rate authority
	Value         float32   `json:"value,omitempty"`         // Value:
}

// countryValuePattern is the validation pattern for Country.Value
var countryValuePattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// Country is an object. ISO 3166 code for a country with optional name
type Country struct {
	CodeContext string `json:"codeContext,omitempty"` // CodeContext: The source of a code
	Id          string `json:"id,omitempty"`          // Id: Use this id to internally identify this country in NextSteps
	Name        string `json:"name,omitempty"`        // Name: The name or code of a country
	Value       string `json:"value,omitempty"`       // Value:
}

// CumulativeValue is an object.
type CumulativeValue struct {
	Type           string             `json:"@type,omitempty"`          // Type:
	Base           float32            `json:"Base,omitempty"`           // Base: The price prior to all applicable taxes of a product such as the rate for a room or fare for a flight.
	CurrencyCode   *CurrencyCode      `json:"CurrencyCode,omitempty"`   // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	Fees           *Fees              `json:"Fees,omitempty"`           // Fees:
	Taxes          *Taxes             `json:"Taxes,omitempty"`          // Taxes:
	Total          float32            `json:"Total,omitempty"`          // Total: Specifies the total price including base + taxes + fees
	ApproximateInd bool               `json:"approximateInd,omitempty"` // ApproximateInd: True if this amount has been converted from the original amount
	CurrencySource CurrencySourceEnum `json:"currencySource,omitempty"` // CurrencySource: The system requesting or returning the currency code specified in the attribute
}

// currencyAmountCodePattern is the validation pattern for CurrencyAmount.Code
var currencyAmountCodePattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// currencyAmountMinorUnitPattern is the validation pattern for CurrencyAmount.MinorUnit
var currencyAmountMinorUnitPattern = regexp.MustCompile(`([0-4])`)

// CurrencyAmount is an object. A monetary amount, up to 4 decimal places. Decimal place needs to be included.
type CurrencyAmount struct {
	ApproximateInd bool               `json:"approximateInd,omitempty"` // ApproximateInd: True if the currency amount has been converted from the original amount
	Code           string             `json:"code,omitempty"`           // Code: An ISO 4217 alpha character code that specifies a money unit
	CurrencySource CurrencySourceEnum `json:"currencySource,omitempty"` // CurrencySource: The system requesting or returning the currency code specified in the attribute
	MinorUnit      int32              `json:"minorUnit,omitempty"`      // MinorUnit: Minor units are a mechanism for expressing the relationship between a major currency unit and its corresponding minor currency unit.
	Value          float32            `json:"value,omitempty"`          // Value:
}

// currencyCodeDecimalPlacePattern is the validation pattern for CurrencyCode.DecimalPlace
var currencyCodeDecimalPlacePattern = regexp.MustCompile(`([0-4])`)

// currencyCodeValuePattern is the validation pattern for CurrencyCode.Value
var currencyCodeValuePattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// CurrencyCode is an object. Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
type CurrencyCode struct {
	CodeAuthority    string `json:"codeAuthority,omitempty"`    // CodeAuthority: Currency code authority
	DecimalAuthority string `json:"decimalAuthority,omitempty"` // DecimalAuthority: Currency code decimal authority
	DecimalPlace     int32  `json:"decimalPlace,omitempty"`     // DecimalPlace: Currency code decimal place
	Value            string `json:"value,omitempty"`            // Value:
}

// CurrencyRateConversion is an object.
type CurrencyRateConversion struct {
	Type           string         `json:"@type,omitempty"` // Type:
	ConversionRate ConversionRate `json:"ConversionRate"`  // ConversionRate: A conversion metric from standard to another with the contextual authority such as IATA, OAG, ISO, etc.
	SourceCurrency CurrencyCode   `json:"SourceCurrency"`  // SourceCurrency: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	TargetCurrency CurrencyCode   `json:"TargetCurrency"`  // TargetCurrency: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
}

// CustomResponseModifiersAir is an object. Modifiers to customize the result set
type CustomResponseModifiersAir struct {
	Type                      string                    `json:"@type,omitempty"`                     // Type:
	BrandAttributeInclusion   []BrandAttributeInclusion `json:"BrandAttributeInclusion,omitempty"`   // BrandAttributeInclusion:
	SearchRepresentation      SearchRepresentationENUM  `json:"SearchRepresentation,omitempty"`      // SearchRepresentation: Customize search result set as leg or journey based
	ExcludeBaggageFeesInd     bool                      `json:"excludeBaggageFeesInd,omitempty"`     // ExcludeBaggageFeesInd: If true, Baggage Fees will be inhibited from the response
	ExcludePenaltiesInd       bool                      `json:"excludePenaltiesInd,omitempty"`       // ExcludePenaltiesInd: If true, Penalties will be excluded from the response
	ExcludeSurchargesInd      bool                      `json:"excludeSurchargesInd,omitempty"`      // ExcludeSurchargesInd: If true, the surcharge breakdown will be excluded from Price_Detail
	ExcludeUnbundledFaresInd  bool                      `json:"excludeUnbundledFaresInd,omitempty"`  // ExcludeUnbundledFaresInd: If true, unbundled fares will not be returned in the response
	IncludeFareCalculationInd bool                      `json:"includeFareCalculationInd,omitempty"` // IncludeFareCalculationInd: if true, the fare calculation string will be returned in the response
}

// CustomerLoyalty is an object. Specifies the ID for the membership program.
type CustomerLoyalty struct {
	CardHolderName    string   `json:"cardHolderName,omitempty"`    // CardHolderName: The card holder name
	Id                string   `json:"id,omitempty"`                // Id: Customer Loyality Id
	Priority          int32    `json:"priority,omitempty"`          // Priority: Numeric Priority Code
	ProgramId         string   `json:"programId,omitempty"`         // ProgramId: Specifies an identifier to indicate the company owner of the loyalty program
	ProgramName       string   `json:"programName,omitempty"`       // ProgramName: Supplier's loyalty program name such as Frontier-EarlyReturns
	ShareWithSupplier []string `json:"shareWithSupplier,omitempty"` // ShareWithSupplier: The list of suppliers that the CustomerLoyalty number is shared.
	Supplier          string   `json:"supplier,omitempty"`          // Supplier: Supplier of a loyalty program
	SupplierType      string   `json:"supplierType,omitempty"`      // SupplierType: The kind of supplier of a loyalty program
	Tier              string   `json:"tier,omitempty"`              // Tier: Customer Loyalty tier level
	ValidatedInd      bool     `json:"validatedInd,omitempty"`      // ValidatedInd: Customer loyalty number has been validated by the supplier
	Value             string   `json:"value,omitempty"`             // Value:
}

// CustomerLoyaltyCredit is an object.
type CustomerLoyaltyCredit struct {
	Type            string          `json:"@type,omitempty"` // Type:
	CustomerLoyalty CustomerLoyalty `json:"CustomerLoyalty"` // CustomerLoyalty: Specifies the ID for the membership program.
	Earned          string          `json:"Earned"`          // Earned: Represents the amount of award credit awarded for this offer\/offering. Award credit can be used for purchasing goods and services through a customer loyalty program
	Status          string          `json:"Status"`          // Status: Represents the amount of status credit awarded for this offer\/offering. Status credit allow a customer to move up the ladder of a customer loyalty program
}

// DateCreateModify is an object. Time stamp of the creation.
type DateCreateModify struct {
	CreatorID      string    `json:"creatorID,omitempty"`      // CreatorID: ID of creator. Software system identifier or an employee id
	LastModifierID string    `json:"lastModifierID,omitempty"` // LastModifierID: Identifies the last software system or person to modify a record
	LastModify     time.Time `json:"lastModify,omitempty"`     // LastModify: Time stamp of last modification.
	Purge          string    `json:"purge,omitempty"`          // Purge: Date an item will be purged from a system of record
	Value          time.Time `json:"value,omitempty"`          // Value:
}

// DateEffectiveExpire is an object. Used to identify the effective date and\/or expiration date.
type DateEffectiveExpire struct {
	Effective              string `json:"effective,omitempty"`              // Effective: Indicates the starting date.
	Expire                 string `json:"expire,omitempty"`                 // Expire: Indicates the ending date.
	ExpireDateExclusiveInd bool   `json:"expireDateExclusiveInd,omitempty"` // ExpireDateExclusiveInd: When true, indicates that the ExpireDate is the first day after the applicable period (e.g. when expire date is Oct 15  the last date of the period is Oct 14).
}

// DateOrDateWindows is an object. Indicates the expiry date.
type DateOrDateWindows struct {
	Duration     string           `json:"duration,omitempty"`     // Duration: Duration from  start date.
	DurationUnit DurationUnitEnum `json:"durationUnit,omitempty"` // DurationUnit: Defines the Units that can be applied to Stay restrictions.
	End          string           `json:"end,omitempty"`          // End: The earliest and latest dates acceptable for the end date.
	Specific     string           `json:"specific,omitempty"`     // Specific: A specific date. When used with a windows must fall between start and end.
	Start        string           `json:"start,omitempty"`        // Start: The earliest and latest dates acceptable for the start date.
}

// DateRange is an object. Specifies the begin and end date of an event
type DateRange struct {
	End   string `json:"end"`   // End: Specifies the end date an event, such as a booking
	Start string `json:"start"` // Start: Specifies the start date for an event, such as a booking
}

// Deadline is an object.
type Deadline struct {
	Type         string            `json:"@type,omitempty"` // Type:
	SpecificDate DateOrDateWindows `json:"SpecificDate"`    // SpecificDate: Indicates the expiry date.
	Time         string            `json:"Time,omitempty"`  // Time: Local time of the property
}

// departureLocationPattern is the validation pattern for Departure.Location
var departureLocationPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// Departure is an object.
type Departure struct {
	Type     string `json:"@type"`    // Type:
	Date     string `json:"date"`     // Date: Local date of for arrival or departure
	Location string `json:"location"` // Location: Location of departure or arrival
	Time     string `json:"time"`     // Time: Local time Date of for arrival or departure
}

// departureDetailCountryPattern is the validation pattern for DepartureDetail.Country
var departureDetailCountryPattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// departureDetailLocationPattern is the validation pattern for DepartureDetail.Location
var departureDetailLocationPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// departureDetailTerminalPattern is the validation pattern for DepartureDetail.Terminal
var departureDetailTerminalPattern = regexp.MustCompile(`([0-9a-zA-Z]+)?`)

// DepartureDetail is an object.
type DepartureDetail struct {
	Type     string `json:"@type"`              // Type:
	Country  string `json:"country,omitempty"`  // Country: Country of to departure or arrival
	Date     string `json:"date"`               // Date: Local date of for arrival or departure
	Location string `json:"location"`           // Location: Location of departure or arrival
	Terminal string `json:"terminal,omitempty"` // Terminal: Departure/Arrival terminal
	Time     string `json:"time"`               // Time: Local time Date of for arrival or departure
}

// Deposit is an object.
type Deposit struct {
	Type         string    `json:"@type"`                  // Type:
	Date         time.Time `json:"Date,omitempty"`         // Date: The date and time the deposit is due
	RemainderInd bool      `json:"remainderInd,omitempty"` // RemainderInd: If present and true, the date is when the remainder of the deposit is due
}

// DepositAmount is an object.
type DepositAmount struct {
	Type           string         `json:"@type"`                  // Type:
	CurrencyAmount CurrencyAmount `json:"CurrencyAmount"`         // CurrencyAmount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	Date           time.Time      `json:"Date,omitempty"`         // Date: The date and time the deposit is due
	RemainderInd   bool           `json:"remainderInd,omitempty"` // RemainderInd: If present and true, the date is when the remainder of the deposit is due
}

// DepositNumberOfNights is an object.
type DepositNumberOfNights struct {
	Type           string    `json:"@type"`                  // Type:
	Date           time.Time `json:"Date,omitempty"`         // Date: The date and time the deposit is due
	NumberOfNights int32     `json:"NumberOfNights"`         // NumberOfNights: The number of nights that must be paid for by deposit
	RemainderInd   bool      `json:"remainderInd,omitempty"` // RemainderInd: If present and true, the date is when the remainder of the deposit is due
}

// DepositPercent is an object.
type DepositPercent struct {
	Type         string    `json:"@type"`                  // Type:
	Date         time.Time `json:"Date,omitempty"`         // Date: The date and time the deposit is due
	Percent      float32   `json:"Percent,omitempty"`      // Percent: The percentage of the price that must be paid for by deposit
	RemainderInd bool      `json:"remainderInd,omitempty"` // RemainderInd: If present and true, the date is when the remainder of the deposit is due
}

// DepositPolicy is an object.
type DepositPolicy struct {
	Type    string    `json:"@type,omitempty"` // Type:
	Deposit []Deposit `json:"Deposit"`         // Deposit:
}

// DestinationPurpose is an object.
type DestinationPurpose struct {
	Type        string          `json:"@type,omitempty"` // Type:
	Destination DestinationEnum `json:"destination"`     // Destination:
	Purpose     PurposeEnum     `json:"purpose"`         // Purpose:
}

// discountCurrencyCodePattern is the validation pattern for Discount.CurrencyCode
var discountCurrencyCodePattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// discountDecimalPlacePattern is the validation pattern for Discount.DecimalPlace
var discountDecimalPlacePattern = regexp.MustCompile(`([0-4])`)

// Discount is an object. Corporate or Other discount
type Discount struct {
	CodeAuthority    string  `json:"codeAuthority"`              // CodeAuthority: Code Authority
	CurrencyCode     string  `json:"currencyCode"`               // CurrencyCode: Currency code of the city.
	DecimalAuthority string  `json:"decimalAuthority,omitempty"` // DecimalAuthority: Decimal Authority
	DecimalPlace     int32   `json:"decimalPlace"`               // DecimalPlace: Number of decimal places
	Description      string  `json:"description,omitempty"`      // Description: The marketing description for the discount
	Value            float32 `json:"value,omitempty"`            // Value:
}

// DisplaySequence is an object.
type DisplaySequence struct {
	Type            string `json:"@type,omitempty"`      // Type:
	OfferRef        string `json:"OfferRef"`             // OfferRef: Offer reference
	ProductRef      string `json:"ProductRef,omitempty"` // ProductRef: Product reference. If blank, display sequence applies to all products within the Offer.
	Sequence        int32  `json:"Sequence,omitempty"`   // Sequence: Segment sequence, if blank, display sequence applies to all segments within the product
	DisplaySequence string `json:"displaySequence"`      // DisplaySequence: The sequence the products are to be displayed for sequential date ordering
}

// documentIssuingCityPattern is the validation pattern for Document.IssuingCity
var documentIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentIssuingIATAPattern is the validation pattern for Document.IssuingIATA
var documentIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentIssuingPCCPattern is the validation pattern for Document.IssuingPCC
var documentIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// Document is an object.
type Document struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
}

// documentAgencyServiceFeeIssuingCityPattern is the validation pattern for DocumentAgencyServiceFee.IssuingCity
var documentAgencyServiceFeeIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentAgencyServiceFeeIssuingIATAPattern is the validation pattern for DocumentAgencyServiceFee.IssuingIATA
var documentAgencyServiceFeeIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentAgencyServiceFeeIssuingPCCPattern is the validation pattern for DocumentAgencyServiceFee.IssuingPCC
var documentAgencyServiceFeeIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentAgencyServiceFee is an object.
type DocumentAgencyServiceFee struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	RelatedDocumentNumber *DocumentNumber  `json:"RelatedDocumentNumber,omitempty"` // RelatedDocumentNumber:
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
}

// documentAgencyServiceFeeVoidIssuingCityPattern is the validation pattern for DocumentAgencyServiceFeeVoid.IssuingCity
var documentAgencyServiceFeeVoidIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentAgencyServiceFeeVoidIssuingIATAPattern is the validation pattern for DocumentAgencyServiceFeeVoid.IssuingIATA
var documentAgencyServiceFeeVoidIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentAgencyServiceFeeVoidIssuingPCCPattern is the validation pattern for DocumentAgencyServiceFeeVoid.IssuingPCC
var documentAgencyServiceFeeVoidIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentAgencyServiceFeeVoid is an object.
type DocumentAgencyServiceFeeVoid struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
	VoidInd               bool             `json:"voidInd,omitempty"`               // VoidInd: If true, this agency service fee has been voided
}

// documentEMDIssuingCityPattern is the validation pattern for DocumentEMD.IssuingCity
var documentEMDIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentEMDIssuingIATAPattern is the validation pattern for DocumentEMD.IssuingIATA
var documentEMDIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentEMDIssuingPCCPattern is the validation pattern for DocumentEMD.IssuingPCC
var documentEMDIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentEMD is an object.
type DocumentEMD struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	EMDDescription        EMDDescription   `json:"EMDDescription"`                  // EMDDescription: A description of the ancillary with two description codes
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
}

// documentEMDExchangeIssuingCityPattern is the validation pattern for DocumentEMDExchange.IssuingCity
var documentEMDExchangeIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentEMDExchangeIssuingIATAPattern is the validation pattern for DocumentEMDExchange.IssuingIATA
var documentEMDExchangeIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentEMDExchangeIssuingPCCPattern is the validation pattern for DocumentEMDExchange.IssuingPCC
var documentEMDExchangeIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentEMDExchange is an object.
type DocumentEMDExchange struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
	ExchangeInd           bool             `json:"exchangeInd,omitempty"`           // ExchangeInd: If true this document has been exchanged for a new document
}

// documentEMDRefundIssuingCityPattern is the validation pattern for DocumentEMDRefund.IssuingCity
var documentEMDRefundIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentEMDRefundIssuingIATAPattern is the validation pattern for DocumentEMDRefund.IssuingIATA
var documentEMDRefundIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentEMDRefundIssuingPCCPattern is the validation pattern for DocumentEMDRefund.IssuingPCC
var documentEMDRefundIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentEMDRefund is an object.
type DocumentEMDRefund struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	EMDDescription        EMDDescription   `json:"EMDDescription"`                  // EMDDescription: A description of the ancillary with two description codes
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
}

// documentEMDVoidIssuingCityPattern is the validation pattern for DocumentEMDVoid.IssuingCity
var documentEMDVoidIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentEMDVoidIssuingIATAPattern is the validation pattern for DocumentEMDVoid.IssuingIATA
var documentEMDVoidIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentEMDVoidIssuingPCCPattern is the validation pattern for DocumentEMDVoid.IssuingPCC
var documentEMDVoidIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentEMDVoid is an object.
type DocumentEMDVoid struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	VoidInd               bool             `json:"VoidInd,omitempty"`               // VoidInd: If true the EMD has been voided
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
}

// documentForfeitIssuingCityPattern is the validation pattern for DocumentForfeit.IssuingCity
var documentForfeitIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentForfeitIssuingIATAPattern is the validation pattern for DocumentForfeit.IssuingIATA
var documentForfeitIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentForfeitIssuingPCCPattern is the validation pattern for DocumentForfeit.IssuingPCC
var documentForfeitIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentForfeit is an object.
type DocumentForfeit struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
	ForfeitInd            bool             `json:"forfeitInd,omitempty"`            // ForfeitInd: If true the value of the document has been forfeited
}

// documentMCOIssuingCityPattern is the validation pattern for DocumentMCO.IssuingCity
var documentMCOIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentMCOIssuingIATAPattern is the validation pattern for DocumentMCO.IssuingIATA
var documentMCOIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentMCOIssuingPCCPattern is the validation pattern for DocumentMCO.IssuingPCC
var documentMCOIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentMCO is an object.
type DocumentMCO struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
	McoInd                bool             `json:"mcoInd,omitempty"`                // McoInd: If true, the document issues is an MCO
}

// documentMCOExchangeIssuingCityPattern is the validation pattern for DocumentMCOExchange.IssuingCity
var documentMCOExchangeIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentMCOExchangeIssuingIATAPattern is the validation pattern for DocumentMCOExchange.IssuingIATA
var documentMCOExchangeIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentMCOExchangeIssuingPCCPattern is the validation pattern for DocumentMCOExchange.IssuingPCC
var documentMCOExchangeIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentMCOExchange is an object.
type DocumentMCOExchange struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
	ExchangeInd           bool             `json:"exchangeInd,omitempty"`           // ExchangeInd: If true, this MCO has been exchanged
}

// documentMCORefundIssuingCityPattern is the validation pattern for DocumentMCORefund.IssuingCity
var documentMCORefundIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentMCORefundIssuingIATAPattern is the validation pattern for DocumentMCORefund.IssuingIATA
var documentMCORefundIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentMCORefundIssuingPCCPattern is the validation pattern for DocumentMCORefund.IssuingPCC
var documentMCORefundIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentMCORefund is an object.
type DocumentMCORefund struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	ControlNumber         string           `json:"ControlNumber"`                   // ControlNumber: Reference for tracking refund
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
}

// documentMCOVoidIssuingCityPattern is the validation pattern for DocumentMCOVoid.IssuingCity
var documentMCOVoidIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentMCOVoidIssuingIATAPattern is the validation pattern for DocumentMCOVoid.IssuingIATA
var documentMCOVoidIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentMCOVoidIssuingPCCPattern is the validation pattern for DocumentMCOVoid.IssuingPCC
var documentMCOVoidIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentMCOVoid is an object.
type DocumentMCOVoid struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	VoidInd               bool             `json:"VoidInd,omitempty"`               // VoidInd: If true, the MCO has been voided
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
}

// documentNumberValuePattern is the validation pattern for DocumentNumber.Value
var documentNumberValuePattern = regexp.MustCompile(`([0-9]+)?`)

// DocumentNumber is an object.
type DocumentNumber struct {
	DocumentIssuer string           `json:"documentIssuer,omitempty"` // DocumentIssuer: Document issuer
	DocumentType   DocumentTypeEnum `json:"documentType,omitempty"`   // DocumentType: Document type like EMD, MCO
	Value          string           `json:"value,omitempty"`          // Value:
}

// DocumentOverridesID is an object.
type DocumentOverridesID struct {
	Type                 string      `json:"@type"`                          // Type:
	DocumentOverridesRef string      `json:"DocumentOverridesRef,omitempty"` // DocumentOverridesRef:
	Identifier           *Identifier `json:"Identifier,omitempty"`           // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id                   string      `json:"id,omitempty"`                   // Id: The reporting number.
}

// DocumentOverridesResponse is an object.
type DocumentOverridesResponse struct {
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"` // CurrencyRateConversion:
	DocumentOverrides      *DocumentOverridesID     `json:"DocumentOverrides,omitempty"`      // DocumentOverrides:
	Identifier             *Identifier              `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`              // NextSteps:
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`          // ReferenceList:
	Result                 *Result                  `json:"Result,omitempty"`                 // Result:
	TraceId                string                   `json:"traceId,omitempty"`                // TraceId: Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TransactionId          string                   `json:"transactionId,omitempty"`          // TransactionId: Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
}

// documentTicketIssuingCityPattern is the validation pattern for DocumentTicket.IssuingCity
var documentTicketIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentTicketIssuingIATAPattern is the validation pattern for DocumentTicket.IssuingIATA
var documentTicketIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentTicketIssuingPCCPattern is the validation pattern for DocumentTicket.IssuingPCC
var documentTicketIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentTicket is an object.
type DocumentTicket struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	ControlNumber         string           `json:"ControlNumber"`                   // ControlNumber: The control number assigned to the Ticket
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
	HistoricInd           bool             `json:"historicInd,omitempty"`           // HistoricInd: If true, this document has been superseded by a new Document facet and may have been exchanged, voided or refunded
}

// documentTicketExchangeIssuingCityPattern is the validation pattern for DocumentTicketExchange.IssuingCity
var documentTicketExchangeIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentTicketExchangeIssuingIATAPattern is the validation pattern for DocumentTicketExchange.IssuingIATA
var documentTicketExchangeIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentTicketExchangeIssuingPCCPattern is the validation pattern for DocumentTicketExchange.IssuingPCC
var documentTicketExchangeIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentTicketExchange is an object.
type DocumentTicketExchange struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
	ExchangeInd           bool             `json:"exchangeInd,omitempty"`           // ExchangeInd: If true this document has been exchanged for a new document
	HistoricInd           bool             `json:"historicInd,omitempty"`           // HistoricInd: if true this document exchange has been cancelled\/voided
}

// documentTicketRefundIssuingCityPattern is the validation pattern for DocumentTicketRefund.IssuingCity
var documentTicketRefundIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentTicketRefundIssuingIATAPattern is the validation pattern for DocumentTicketRefund.IssuingIATA
var documentTicketRefundIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentTicketRefundIssuingPCCPattern is the validation pattern for DocumentTicketRefund.IssuingPCC
var documentTicketRefundIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentTicketRefund is an object.
type DocumentTicketRefund struct {
	Type                           string               `json:"@type"`                                    // Type:
	Amount                         *Amount              `json:"Amount,omitempty"`                         // Amount:
	Commission                     *Commission          `json:"Commission,omitempty"`                     // Commission:
	ControlNumber                  string               `json:"ControlNumber"`                            // ControlNumber: The control number assigned to the TicketRefund
	CumulativeValue                *CumulativeValue     `json:"CumulativeValue,omitempty"`                // CumulativeValue:
	FareGuaranteePolicy            *FareGuaranteePolicy `json:"FareGuaranteePolicy,omitempty"`            // FareGuaranteePolicy:
	FiledAmount                    *FiledAmount         `json:"FiledAmount,omitempty"`                    // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity                    string               `json:"IssuingCity,omitempty"`                    // IssuingCity: Document issuing city
	IssuingIATA                    string               `json:"IssuingIATA,omitempty"`                    // IssuingIATA: Document issuing IATA
	IssuingPCC                     string               `json:"IssuingPCC,omitempty"`                     // IssuingPCC: Document issuing pcc
	Number                         string               `json:"Number,omitempty"`                         // Number: The identifying number of the document
	TravelerIdentifierRef          *IdentifierRef       `json:"TravelerIdentifierRef,omitempty"`          // TravelerIdentifierRef:
	WaiverCode                     *WaiverCode          `json:"WaiverCode,omitempty"`                     // WaiverCode:
	AgencySettlementNotReportedInd bool                 `json:"agencySettlementNotReportedInd,omitempty"` // AgencySettlementNotReportedInd: If true, this refund is settled by the agency directly with the traveler. Transaction is not reported to BSP or ARC. Ticket coupon is updated to RFND status
	HistoricInd                    bool                 `json:"historicInd,omitempty"`                    // HistoricInd: if true this document refund has been cancelled/voided
	PartialRefundInd               bool                 `json:"partialRefundInd,omitempty"`               // PartialRefundInd: if true, the ticket has been partially refunded
	RefundGuaranteedInd            bool                 `json:"refundGuaranteedInd,omitempty"`            // RefundGuaranteedInd: if true, this refund amount is guaranteed by Travelport JSON API automated refunds
}

// documentTicketRetainedIssuingCityPattern is the validation pattern for DocumentTicketRetained.IssuingCity
var documentTicketRetainedIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentTicketRetainedIssuingIATAPattern is the validation pattern for DocumentTicketRetained.IssuingIATA
var documentTicketRetainedIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentTicketRetainedIssuingPCCPattern is the validation pattern for DocumentTicketRetained.IssuingPCC
var documentTicketRetainedIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentTicketRetained is an object.
type DocumentTicketRetained struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	ValueRetainedInd      bool             `json:"ValueRetainedInd,omitempty"`      // ValueRetainedInd: If true the Document Ticket Value has been retained for future use
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
}

// documentTicketVoidIssuingCityPattern is the validation pattern for DocumentTicketVoid.IssuingCity
var documentTicketVoidIssuingCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// documentTicketVoidIssuingIATAPattern is the validation pattern for DocumentTicketVoid.IssuingIATA
var documentTicketVoidIssuingIATAPattern = regexp.MustCompile(`([0-9]{8})`)

// documentTicketVoidIssuingPCCPattern is the validation pattern for DocumentTicketVoid.IssuingPCC
var documentTicketVoidIssuingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// DocumentTicketVoid is an object.
type DocumentTicketVoid struct {
	Type                  string           `json:"@type"`                           // Type:
	Amount                *Amount          `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission      `json:"Commission,omitempty"`            // Commission:
	CumulativeValue       *CumulativeValue `json:"CumulativeValue,omitempty"`       // CumulativeValue:
	FiledAmount           *FiledAmount     `json:"FiledAmount,omitempty"`           // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	IssuingCity           string           `json:"IssuingCity,omitempty"`           // IssuingCity: Document issuing city
	IssuingIATA           string           `json:"IssuingIATA,omitempty"`           // IssuingIATA: Document issuing IATA
	IssuingPCC            string           `json:"IssuingPCC,omitempty"`            // IssuingPCC: Document issuing pcc
	Number                string           `json:"Number,omitempty"`                // Number: The identifying number of the document
	TravelerIdentifierRef *IdentifierRef   `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	VoidInd               bool             `json:"VoidInd,omitempty"`               // VoidInd: If true the document has been voided
	WaiverCode            *WaiverCode      `json:"WaiverCode,omitempty"`            // WaiverCode:
}

// DocumentValidDateRange is an object.
type DocumentValidDateRange struct {
	Type                string              `json:"@type,omitempty"`               // Type:
	ProductRef          []ProductIdentifier `json:"ProductRef,omitempty"`          // ProductRef:
	SegmentSequenceList []int32             `json:"SegmentSequenceList,omitempty"` // SegmentSequenceList: The segmentSequence within the product the action is being requested for. Used when multiple flights exist within a product. Only one product may be selected with this option.
	ValidDateRange      *DateRange          `json:"ValidDateRange,omitempty"`      // ValidDateRange: Specifies the begin and end date of an event
}

// driverInfoCountryOfDocIssuePattern is the validation pattern for DriverInfo.CountryOfDocIssue
var driverInfoCountryOfDocIssuePattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// DriverInfo is an object. Basic information (metadata) about the intended driver of the vehicle
type DriverInfo struct {
	Age               int32  `json:"age,omitempty"`               // Age: Assigned Type: c-1100:NumberDoubleDigit
	CountryOfDocIssue string `json:"countryOfDocIssue,omitempty"` // CountryOfDocIssue: Assigned Type: c-1100:CountryCodeISO
}

// EMD is an object.
type EMD struct {
	Type                   string            `json:"@type"`                            // Type:
	AgencyInfo             AgencyInfo        `json:"AgencyInfo"`                       // AgencyInfo: Detail of the travel agency that issues the ticket
	AssociatedTicketNumber *TicketNumber     `json:"AssociatedTicketNumber,omitempty"` // AssociatedTicketNumber: The ticketNumber that will be used as partial payment for this Offer\/Offering
	EMDRef                 string            `json:"EMDRef,omitempty"`                 // EMDRef:
	EMDSegment             []EMDSegment      `json:"EMDSegment"`                       // EMDSegment:
	ESAC                   string            `json:"ESAC,omitempty"`                   // ESAC: The BSP ESAC code assign for a void or refund transaction\nThe BSP E
	FormOfPayment          FormOfPayment     `json:"FormOfPayment"`                    // FormOfPayment:
	Identifier             *Identifier       `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	PersonName             PersonName        `json:"PersonName"`                       // PersonName:
	ReservationLocator     []SupplierLocator `json:"ReservationLocator,omitempty"`     // ReservationLocator:
	Restrictions           []string          `json:"Restrictions,omitempty"`           // Restrictions:
	TotalAmount            *TotalAmount      `json:"TotalAmount,omitempty"`            // TotalAmount:
	Id                     string            `json:"id,omitempty"`                     // Id:
}

// eMDDescriptionCodePattern is the validation pattern for EMDDescription.Code
var eMDDescriptionCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// eMDDescriptionSubCodePattern is the validation pattern for EMDDescription.SubCode
var eMDDescriptionSubCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// EMDDescription is an object. A description of the ancillary with two description codes
type EMDDescription struct {
	Code        string `json:"code,omitempty"`        // Code: A description of the ancillary with two description codes
	CodeContext string `json:"codeContext,omitempty"` // CodeContext: Code context
	SubCode     string `json:"subCode,omitempty"`     // SubCode: EMD number sub code
	Value       string `json:"value,omitempty"`       // Value:
}

// EMDID is an object.
type EMDID struct {
	Type       string      `json:"@type"`                // Type:
	EMDRef     string      `json:"EMDRef,omitempty"`     // EMDRef:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id:
}

// EMDListResponse is an object.
type EMDListResponse struct {
	EMDID []EMD `json:"EMDID,omitempty"` // EMDID:
}

// eMDQueryUpdateEMDAgencyCodePattern is the validation pattern for EMDQueryUpdateEMD.AgencyCode
var eMDQueryUpdateEMDAgencyCodePattern = regexp.MustCompile(`([0-9]{8})`)

// eMDSegmentDateOfServicePattern is the validation pattern for EMDSegment.DateOfService
var eMDSegmentDateOfServicePattern = regexp.MustCompile(`(\d{4}-\d{2}-\d{2})`)

// eMDSegmentPresentAtPattern is the validation pattern for EMDSegment.PresentAt
var eMDSegmentPresentAtPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// eMDSegmentPresentToPattern is the validation pattern for EMDSegment.PresentTo
var eMDSegmentPresentToPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,3})`)

// EMDSegment is an object.
type EMDSegment struct {
	Type           string          `json:"@type,omitempty"`          // Type:
	Amount         *Amount         `json:"Amount,omitempty"`         // Amount:
	DateOfService  string          `json:"DateOfService,omitempty"`  // DateOfService: The date of service the service is available for
	EMDDescription *EMDDescription `json:"EMDDescription,omitempty"` // EMDDescription: A description of the ancillary with two description codes
	PresentAt      string          `json:"PresentAt,omitempty"`      // PresentAt: The location the EMD should be presented to to supply the service
	PresentTo      string          `json:"PresentTo,omitempty"`      // PresentTo: The airline the EMD should be presented to to supply the service
	Routing        string          `json:"Routing,omitempty"`        // Routing: The routing the service is valid on
	Status         EMDStatusENUM   `json:"Status,omitempty"`         // Status:
	Quantity       int32           `json:"quantity,omitempty"`       // Quantity: The quantity of the ancillary available on this EMDSegment
	Sequence       int32           `json:"sequence"`                 // Sequence: Sequence of EMDSegment
}

// emailEmailTypePattern is the validation pattern for Email.EmailType
var emailEmailTypePattern = regexp.MustCompile(`[0-9A-Z]{1,3}(\.[A-Z]{3}(\.X){0,1}){0,1}`)

// Email is an object. Electronic email addresses, in IETF specified format.
type Email struct {
	Comment         string           `json:"comment,omitempty"`         // Comment: Assigned Type: c-1100:StringText
	EmailType       string           `json:"emailType,omitempty"`       // EmailType: Type of the e-mail
	Id              string           `json:"id,omitempty"`              // Id: Electronic email addresses, in IETF specified format.
	OptInDate       time.Time        `json:"optInDate,omitempty"`       // OptInDate: The datetime of receiving the opt in notice
	OptInStatus     OptInStatusEnum  `json:"optInStatus,omitempty"`     // OptInStatus: Used to indicate marketing preferences, OptIn, OptOut
	OptOutDate      time.Time        `json:"optOutDate,omitempty"`      // OptOutDate: The datetime the opt out notice was received
	OptOutInd       YesNoInheritEnum `json:"optOutInd,omitempty"`       // OptOutInd: Used to indicate marketing preferences, Yes, No, Inherit
	PreferredFormat string           `json:"preferredFormat,omitempty"` // PreferredFormat: Mime media type
	ProvisionedInd  bool             `json:"provisionedInd,omitempty"`  // ProvisionedInd: If true then the email address came from the provisioning process
	ShareMarketing  YesNoInheritEnum `json:"shareMarketing,omitempty"`  // ShareMarketing: Used to indicate marketing preferences, Yes, No, Inherit
	ShareSync       YesNoInheritEnum `json:"shareSync,omitempty"`       // ShareSync: Used to indicate marketing preferences, Yes, No, Inherit
	ValidInd        bool             `json:"validInd,omitempty"`        // ValidInd: If true, this is a valid email address that has been system verified via a successful email transmission.
	Value           string           `json:"value,omitempty"`           // Value:
}

// Error is an object.
type Error struct {
	Type          string          `json:"@type"`                   // Type:
	Message       string          `json:"Message,omitempty"`       // Message: The Travelport standardized error or warning message
	NameValuePair []NameValuePair `json:"NameValuePair,omitempty"` // NameValuePair:
	StatusCode    int32           `json:"StatusCode,omitempty"`    // StatusCode: Http standard response code
}

// ErrorDetail is an object.
type ErrorDetail struct {
	Type              string          `json:"@type"`                       // Type:
	Message           string          `json:"Message,omitempty"`           // Message: The Travelport standardized error or warning message
	NameValuePair     []NameValuePair `json:"NameValuePair,omitempty"`     // NameValuePair:
	SourceCode        string          `json:"SourceCode,omitempty"`        // SourceCode: The error or warning code returned by the source airline or host system
	SourceDescription string          `json:"SourceDescription,omitempty"` // SourceDescription: The error or warning message as it is returned by the source airline or host system
	SourceID          string          `json:"SourceID"`                    // SourceID: The identifier of the source system sending the error or warning
	StatusCode        int32           `json:"StatusCode,omitempty"`        // StatusCode: Http standard response code
}

// ErrorResponse is an object.
type ErrorResponse struct {
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"` // CurrencyRateConversion:
	Identifier             *Identifier              `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`              // NextSteps:
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`          // ReferenceList:
	Result                 *Result                  `json:"Result,omitempty"`                 // Result:
	TraceId                string                   `json:"traceId,omitempty"`                // TraceId: Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TransactionId          string                   `json:"transactionId,omitempty"`          // TransactionId: Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
}

// ErrorWarning is an object.
type ErrorWarning struct {
	Type          string          `json:"@type"`                   // Type:
	Message       string          `json:"Message,omitempty"`       // Message: The Travelport standardized error or warning message
	NameValuePair []NameValuePair `json:"NameValuePair,omitempty"` // NameValuePair:
	StatusCode    int32           `json:"StatusCode,omitempty"`    // StatusCode: Http standard response code
}

// ErrorWarningDetail is an object.
type ErrorWarningDetail struct {
	Type              string          `json:"@type"`                       // Type:
	Message           string          `json:"Message,omitempty"`           // Message: The Travelport standardized error or warning message
	NameValuePair     []NameValuePair `json:"NameValuePair,omitempty"`     // NameValuePair:
	SourceCode        string          `json:"SourceCode,omitempty"`        // SourceCode: The error or warning code returned by the source airline or host system
	SourceDescription string          `json:"SourceDescription,omitempty"` // SourceDescription: The error or warning message as it is returned by the source airline or host system
	SourceID          string          `json:"SourceID"`                    // SourceID: The identifier of the source system sending the error or warning
	StatusCode        int32           `json:"StatusCode,omitempty"`        // StatusCode: Http standard response code
}

// ExtendedPayment is an object. Note this field is deprecated in Payment schema and should be passed in FormOfPaymentPaymentCardExtendPayment schema
type ExtendedPayment struct {
	FirstInstallment     float32 `json:"FirstInstallment,omitempty"` // FirstInstallment: For Pagos Parceledos, specify the first installment amount. This will be the same currency as the payment
	NumberOfInstallments int32   `json:"NumberOfInstallments"`       // NumberOfInstallments: The number of installment payments to be charged by the payment card provider
	OTATOCode            string  `json:"OTATOCode,omitempty"`        // OTATOCode: For Pagos Parceledos the OTATOCode
	RemainingAmount      float32 `json:"RemainingAmount,omitempty"`  // RemainingAmount: For Pagos Parceledos, specify the remaining amount to be charged that will be spread across the number of installments. This will be the same currency as the payment
}

// FareGuaranteePolicy is an object.
type FareGuaranteePolicy struct {
	Type                 string           `json:"@type,omitempty"`              // Type:
	Code                 *Code            `json:"Code,omitempty"`               // Code: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	EligibleforADMReview YesNoUnknownEnum `json:"EligibleforADMReview"`         // EligibleforADMReview: Yes , No , Unknown
	PassengerTypeCodes   []string         `json:"passengerTypeCodes,omitempty"` // PassengerTypeCodes: The list of passenger type codes
}

// FareQualifierENUM is an object.
type FareQualifierENUM struct {
	Value FareQualifierENUMBase `json:"value,omitempty"` // Value:
}

// FareSelection is an object.
type FareSelection struct {
	Type          string             `json:"@type"`                   // Type:
	ChangeOptions *ChangeOptions     `json:"ChangeOptions,omitempty"` // ChangeOptions:
	FareQualifier *FareQualifierENUM `json:"FareQualifier,omitempty"` // FareQualifier:
	RefundOptions *RefundOptions     `json:"RefundOptions,omitempty"` // RefundOptions:
	FareType      FaresFilterEnum    `json:"fareType,omitempty"`      // FareType: Defines the type of fares to return (Only public fares, Only private fares, Only agency private fares, Only
}

// fareSelectionDetailValidatingCarrierPattern is the validation pattern for FareSelectionDetail.ValidatingCarrier
var fareSelectionDetailValidatingCarrierPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,3})`)

// FareSelectionDetail is an object.
type FareSelectionDetail struct {
	Type                            string             `json:"@type"`                                     // Type:
	ChangeOptions                   *ChangeOptions     `json:"ChangeOptions,omitempty"`                   // ChangeOptions:
	FareQualifier                   *FareQualifierENUM `json:"FareQualifier,omitempty"`                   // FareQualifier:
	RefundOptions                   *RefundOptions     `json:"RefundOptions,omitempty"`                   // RefundOptions:
	FareType                        FaresFilterEnum    `json:"fareType,omitempty"`                        // FareType: Defines the type of fares to return (Only public fares, Only private fares, Only agency private fares, Only
	ProhibitAdvancePurchaseFaresInd bool               `json:"prohibitAdvancePurchaseFaresInd,omitempty"` // ProhibitAdvancePurchaseFaresInd: If present and true, fares with advance purchase requirements will not be returned
	ProhibitMaxStayFaresInd         bool               `json:"prohibitMaxStayFaresInd,omitempty"`         // ProhibitMaxStayFaresInd: If present and true, fares with maximum stays will not be returned
	ProhibitMinStayFaresInd         bool               `json:"prohibitMinStayFaresInd,omitempty"`         // ProhibitMinStayFaresInd: If present and true, fares with minimum stays will not be returned
	ProhibitRefundableFaresInd      bool               `json:"prohibitRefundableFaresInd,omitempty"`      // ProhibitRefundableFaresInd: This field is deprecated. Use RefundOptions for refunadability options
	ProhibitUnbundledFaresInd       bool               `json:"prohibitUnbundledFaresInd,omitempty"`       // ProhibitUnbundledFaresInd:
	RefundableOnlyInd               bool               `json:"refundableOnlyInd,omitempty"`               // RefundableOnlyInd: This field is deprecated. Use RefundOptions for refunadability options
	ValidatingCarrier               string             `json:"validatingCarrier,omitempty"`               // ValidatingCarrier: Airline code
}

// Fee is an object.
type Fee struct {
	Type               string             `json:"@type,omitempty"`              // Type:
	FeeAmountOrPercent FeeAmountOrPercent `json:"FeeAmountOrPercent"`           // FeeAmountOrPercent:
	Tax                []Tax              `json:"Tax,omitempty"`                // Tax:
	Description        string             `json:"description,omitempty"`        // Description: Fee description
	FeeApplication     ApplicationEnum    `json:"feeApplication,omitempty"`     // FeeApplication: Application values like perperson , peroom
	FeeCode            string             `json:"feeCode,omitempty"`            // FeeCode: Fee code
	FeeFrequency       FrequencyEnum      `json:"feeFrequency,omitempty"`       // FeeFrequency: Stay frequency like PerNight, PerDay
	IncludedinBaseInd  bool               `json:"includedinBaseInd,omitempty"`  // IncludedinBaseInd: If the fee is included in the Base Price
	Purpose            string             `json:"purpose,omitempty"`            // Purpose: Fee purpose
	ReportingAuthority string             `json:"reportingAuthority,omitempty"` // ReportingAuthority: Identifies the reporting authority.
}

// FeeAmountOrPercent is an object.
type FeeAmountOrPercent struct {
	Type        string         `json:"@type"`                 // Type:
	Application CommissionEnum `json:"application,omitempty"` // Application: Type of commission
}

// FeeAmountOrPercentAmount is an object.
type FeeAmountOrPercentAmount struct {
	Type        string          `json:"@type"`                 // Type:
	Amount      *CurrencyAmount `json:"Amount,omitempty"`      // Amount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	Application CommissionEnum  `json:"application,omitempty"` // Application: Type of commission
}

// FeeAmountOrPercentPercent is an object.
type FeeAmountOrPercentPercent struct {
	Type        string         `json:"@type"`                 // Type:
	Percent     float32        `json:"Percent,omitempty"`     // Percent: Percent amount of commission
	Application CommissionEnum `json:"application,omitempty"` // Application: Type of commission
}

// Fees is an object.
type Fees struct {
	Type      string  `json:"@type"`               // Type:
	TotalFees float32 `json:"TotalFees,omitempty"` // TotalFees: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
}

// FeesDetail is an object.
type FeesDetail struct {
	Type      string  `json:"@type"`               // Type:
	Fee       []Fee   `json:"Fee,omitempty"`       // Fee:
	TotalFees float32 `json:"TotalFees,omitempty"` // TotalFees: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
}

// filedAmountCurrencyCodePattern is the validation pattern for FiledAmount.CurrencyCode
var filedAmountCurrencyCodePattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// filedAmountDecimalPlacePattern is the validation pattern for FiledAmount.DecimalPlace
var filedAmountDecimalPlacePattern = regexp.MustCompile(`([0-4])`)

// FiledAmount is an object. The base amount of a ticket price or net price that is filed in local currency
type FiledAmount struct {
	CodeAuthority    string  `json:"codeAuthority"`              // CodeAuthority: Filed amount currency code authority
	CurrencyCode     string  `json:"currencyCode,omitempty"`     // CurrencyCode: Filed amount currency code
	DecimalAuthority string  `json:"decimalAuthority,omitempty"` // DecimalAuthority: ISO 4217 standard decimal authority
	DecimalPlace     int32   `json:"decimalPlace"`               // DecimalPlace: ISO 4217 standard has a different number of decimals
	Value            float32 `json:"value,omitempty"`            // Value: Filed amount value
}

// flightCarrierPattern is the validation pattern for Flight.Carrier
var flightCarrierPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,3})`)

// flightEquipmentPattern is the validation pattern for Flight.Equipment
var flightEquipmentPattern = regexp.MustCompile(`([A-Z0-9]{3})?`)

// flightNumberPattern is the validation pattern for Flight.Number
var flightNumberPattern = regexp.MustCompile(`[0-9]{1,4}[A-Z]?|OPEN|ARNK`)

// flightOperatingCarrierPattern is the validation pattern for Flight.OperatingCarrier
var flightOperatingCarrierPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,3})`)

// Flight is an object.
type Flight struct {
	Type                           string             `json:"@type"`                                    // Type:
	Arrival                        Arrival            `json:"Arrival"`                                  // Arrival:
	Departure                      Departure          `json:"Departure"`                                // Departure:
	FlightRef                      string             `json:"FlightRef,omitempty"`                      // FlightRef: Reference to a Flight object eslewhere in the message
	Identifier                     *Identifier        `json:"Identifier,omitempty"`                     // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	IntermediateStop               []IntermediateStop `json:"IntermediateStop,omitempty"`               // IntermediateStop:
	Carrier                        string             `json:"carrier"`                                  // Carrier: The airline code
	Distance                       int32              `json:"distance,omitempty"`                       // Distance: The flight travelled distance
	Duration                       string             `json:"duration,omitempty"`                       // Duration: Elapsed flight time represented in ISO 8601 format
	Equipment                      string             `json:"equipment,omitempty"`                      // Equipment: Air Equip Code IATA
	Id                             string             `json:"id,omitempty"`                             // Id: Internal ID
	Number                         string             `json:"number"`                                   // Number: Flight number
	OperatingCarrier               string             `json:"operatingCarrier,omitempty"`               // OperatingCarrier: The airline code
	OperatingCarrierName           string             `json:"operatingCarrierName,omitempty"`           // OperatingCarrierName: The operating carrier name
	Stops                          int32              `json:"stops,omitempty"`                          // Stops: Number of stops
	SubjectToGovernmentApprovalInd bool               `json:"subjectToGovernmentApprovalInd,omitempty"` // SubjectToGovernmentApprovalInd: If true, this flight schedule is yet to receive government approval
}

// flightDetailCarrierPattern is the validation pattern for FlightDetail.Carrier
var flightDetailCarrierPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,3})`)

// flightDetailEquipmentPattern is the validation pattern for FlightDetail.Equipment
var flightDetailEquipmentPattern = regexp.MustCompile(`([A-Z0-9]{3})?`)

// flightDetailNumberPattern is the validation pattern for FlightDetail.Number
var flightDetailNumberPattern = regexp.MustCompile(`[0-9]{1,4}[A-Z]?|OPEN|ARNK`)

// flightDetailOperatingCarrierPattern is the validation pattern for FlightDetail.OperatingCarrier
var flightDetailOperatingCarrierPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,3})`)

// FlightDetail is an object.
type FlightDetail struct {
	Type                           string                     `json:"@type"`                                    // Type:
	Arrival                        Arrival                    `json:"Arrival"`                                  // Arrival:
	AvailabilitySourceCode         AvailabilitySourceCodeENUM `json:"AvailabilitySourceCode,omitempty"`         // AvailabilitySourceCode: A code representing the source of the flight availability
	Departure                      Departure                  `json:"Departure"`                                // Departure:
	FlightRef                      string                     `json:"FlightRef,omitempty"`                      // FlightRef: Reference to a Flight object eslewhere in the message
	Identifier                     *Identifier                `json:"Identifier,omitempty"`                     // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	IntermediateStop               []IntermediateStop         `json:"IntermediateStop,omitempty"`               // IntermediateStop:
	Carrier                        string                     `json:"carrier"`                                  // Carrier: The airline code
	Distance                       int32                      `json:"distance,omitempty"`                       // Distance: The flight travelled distance
	Duration                       string                     `json:"duration,omitempty"`                       // Duration: Elapsed flight time represented in ISO 8601 format
	Equipment                      string                     `json:"equipment,omitempty"`                      // Equipment: Air Equip Code IATA
	Id                             string                     `json:"id,omitempty"`                             // Id: Internal ID
	Number                         string                     `json:"number"`                                   // Number: Flight number
	OperatingCarrier               string                     `json:"operatingCarrier,omitempty"`               // OperatingCarrier: The airline code
	OperatingCarrierName           string                     `json:"operatingCarrierName,omitempty"`           // OperatingCarrierName: The operating carrier name
	Stops                          int32                      `json:"stops,omitempty"`                          // Stops: Number of stops
	SubjectToGovernmentApprovalInd bool                       `json:"subjectToGovernmentApprovalInd,omitempty"` // SubjectToGovernmentApprovalInd: If true, this flight schedule is yet to receive government approval
}

// FlightID is an object.
type FlightID struct {
	Type       string      `json:"@type"`                // Type:
	FlightRef  string      `json:"FlightRef,omitempty"`  // FlightRef: Reference to a Flight object eslewhere in the message
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Internal ID
}

// flightProductClassOfServicePattern is the validation pattern for FlightProduct.ClassOfService
var flightProductClassOfServicePattern = regexp.MustCompile(`([a-zA-Z0-9]{1,2})`)

// FlightProduct is an object.
type FlightProduct struct {
	Type                       string                       `json:"@type,omitempty"`                      // Type:
	Brand                      *BrandID                     `json:"Brand,omitempty"`                      // Brand:
	ClassOfServiceAvailability []ClassOfServiceAvailability `json:"ClassOfServiceAvailability,omitempty"` // ClassOfServiceAvailability:
	CustomerLoyaltyCredit      []CustomerLoyaltyCredit      `json:"CustomerLoyaltyCredit,omitempty"`      // CustomerLoyaltyCredit:
	FareQualifier              *FareQualifierENUM           `json:"FareQualifier,omitempty"`              // FareQualifier:
	Cabin                      CabinAirEnum                 `json:"cabin,omitempty"`                      // Cabin: Specifies the cabin type (e.g. first, business, economy).
	CarCode                    string                       `json:"carCode,omitempty"`                    // CarCode: The car code
	ClassOfService             string                       `json:"classOfService,omitempty"`             // ClassOfService: The class of service
	FareBasisCode              string                       `json:"fareBasisCode,omitempty"`              // FareBasisCode: Fare basis code
	FareType                   FareTypeEnum                 `json:"fareType,omitempty"`                   // FareType: Defines the type of fares to return (Only public fares, Only private fares, Only agency private fares, Only
	FareTypeCode               string                       `json:"fareTypeCode,omitempty"`               // FareTypeCode: The ATPCO fare type code
	SegmentSequence            []int32                      `json:"segmentSequence"`                      // SegmentSequence: The Segment sequence
	StopoverPriced             YesNoUnknownEnum             `json:"stopoverPriced,omitempty"`             // StopoverPriced: Yes , No , Unknown
	TicketDesignator           string                       `json:"ticketDesignator,omitempty"`           // TicketDesignator: The ticket designator
	ValueCode                  string                       `json:"valueCode,omitempty"`                  // ValueCode: The value code
}

// FlightSegment is an object.
type FlightSegment struct {
	Type               string                 `json:"@type,omitempty"`              // Type:
	CO2Actual          *Measurement           `json:"CO2Actual,omitempty"`          // CO2Actual: Used for dimensional units (width, height, depth) or weight
	Flight             FlightID               `json:"Flight"`                       // Flight:
	OperationalStatus  *OperationalStatusENUM `json:"OperationalStatus,omitempty"`  // OperationalStatus:
	BoundFlightsInd    bool                   `json:"boundFlightsInd,omitempty"`    // BoundFlightsInd: If present and true, the Segments in this Connection must be sold and cancelled together.
	ConnectionDuration string                 `json:"connectionDuration,omitempty"` // ConnectionDuration: The actual duration (in minutes) between
	Id                 string                 `json:"id,omitempty"`                 // Id: Local indentifier within a given message for this object.
	Sequence           int32                  `json:"sequence"`                     // Sequence: Segment sequence
}

// FlightType is an object.
type FlightType struct {
	Type                           string             `json:"@type,omitempty"`                          // Type:
	ConnectionType                 ConnectionTypeEnum `json:"connectionType,omitempty"`                 // ConnectionType:
	ExcludeInterlineConnectionsInd bool               `json:"excludeInterlineConnectionsInd,omitempty"` // ExcludeInterlineConnectionsInd: If present and true, exclude interline connections
}

// formOfPaymentDocumentNumberPattern is the validation pattern for FormOfPayment.DocumentNumber
var formOfPaymentDocumentNumberPattern = regexp.MustCompile(`([0-9]+)?`)

// FormOfPayment is an object.
type FormOfPayment struct {
	DocumentIssuer string                `json:"documentIssuer,omitempty"` // DocumentIssuer: Document issuer
	DocumentNumber string                `json:"documentNumber,omitempty"` // DocumentNumber: Payment document number
	DocumentType   DocumentTypeEnum      `json:"documentType,omitempty"`   // DocumentType: Document type like EMD, MCO
	EncryptedValue string                `json:"encryptedValue,omitempty"` // EncryptedValue: Encrypted value
	Value          FormOfPaymentTypeEnum `json:"value,omitempty"`          // Value: The list of valid forms of payment.
}

// FormOfPaymentAgencyAccount is an object.
type FormOfPaymentAgencyAccount struct {
	Type             string      `json:"@type"`                      // Type:
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"` // FormOfPaymentRef:
	Identifier       *Identifier `json:"Identifier,omitempty"`       // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	AgencyId         string      `json:"agencyId,omitempty"`         // AgencyId: The agency Id
	Id               string      `json:"id,omitempty"`               // Id:
}

// FormOfPaymentBSP is an object.
type FormOfPaymentBSP struct {
	Type             string      `json:"@type"`                      // Type:
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"` // FormOfPaymentRef:
	Identifier       *Identifier `json:"Identifier,omitempty"`       // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	AccountNumber    string      `json:"accountNumber,omitempty"`    // AccountNumber: The account number for the Form of payment BSP
	Id               string      `json:"id,omitempty"`               // Id:
}

// FormOfPaymentCash is an object.
type FormOfPaymentCash struct {
	Type                  string      `json:"@type"`                           // Type:
	FormOfPaymentRef      string      `json:"FormOfPaymentRef,omitempty"`      // FormOfPaymentRef:
	Identifier            *Identifier `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	AgentNonRefundableInd bool        `json:"agentNonRefundableInd,omitempty"` // AgentNonRefundableInd: This indicates that the Cash payment should not be refunded
	Id                    string      `json:"id,omitempty"`                    // Id:
}

// FormOfPaymentDocument is an object.
type FormOfPaymentDocument struct {
	Type             string          `json:"@type"`                      // Type:
	DocumentNumber   *DocumentNumber `json:"DocumentNumber,omitempty"`   // DocumentNumber:
	FormOfPaymentRef string          `json:"FormOfPaymentRef,omitempty"` // FormOfPaymentRef:
	Identifier       *Identifier     `json:"Identifier,omitempty"`       // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id               string          `json:"id,omitempty"`               // Id:
}

// FormOfPaymentForfeit is an object.
type FormOfPaymentForfeit struct {
	Type             string      `json:"@type"`                      // Type:
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"` // FormOfPaymentRef:
	Identifier       *Identifier `json:"Identifier,omitempty"`       // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ForfeitInd       bool        `json:"forfeitInd,omitempty"`       // ForfeitInd: If true, this form of payment instruction is to forfeit residual amounts specified in an Offer. Used in conjunction with Payment to specify which amounts to be forfeited
	Id               string      `json:"id,omitempty"`               // Id:
}

// FormOfPaymentIdentifier is an object.
type FormOfPaymentIdentifier struct {
	Type             string      `json:"@type,omitempty"`            // Type:
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"` // FormOfPaymentRef:
	Identifier       *Identifier `json:"Identifier,omitempty"`       // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id               string      `json:"id,omitempty"`               // Id:
}

// FormOfPaymentInvoice is an object.
type FormOfPaymentInvoice struct {
	Type             string      `json:"@type"`                      // Type:
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"` // FormOfPaymentRef:
	Identifier       *Identifier `json:"Identifier,omitempty"`       // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	InvoiceNumber    string      `json:"InvoiceNumber,omitempty"`    // InvoiceNumber: The invoice number applicable to this form of payment. Send null or empty string if no invoice number specified.
	Id               string      `json:"id,omitempty"`               // Id:
}

// FormOfPaymentListResponse is an object.
type FormOfPaymentListResponse struct {
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"` // CurrencyRateConversion:
	FormOfPaymentID        []FormOfPaymentID        `json:"FormOfPaymentID,omitempty"`        // FormOfPaymentID:
	Identifier             *Identifier              `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`              // NextSteps:
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`          // ReferenceList:
	Result                 *Result                  `json:"Result,omitempty"`                 // Result:
	TraceId                string                   `json:"traceId,omitempty"`                // TraceId: Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TransactionId          string                   `json:"transactionId,omitempty"`          // TransactionId: Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
}

// FormOfPaymentPaymentCard is an object.
type FormOfPaymentPaymentCard struct {
	Type                               string           `json:"@type"`                                        // Type:
	ExtendedPayment                    *ExtendedPayment `json:"ExtendedPayment,omitempty"`                    // ExtendedPayment: Note this field is deprecated in Payment schema and should be passed in FormOfPaymentPaymentCardExtendPayment schema
	FormOfPaymentRef                   string           `json:"FormOfPaymentRef,omitempty"`                   // FormOfPaymentRef:
	Identifier                         *Identifier      `json:"Identifier,omitempty"`                         // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	PaymentCard                        *PaymentCard     `json:"PaymentCard,omitempty"`                        // PaymentCard:
	Id                                 string           `json:"id,omitempty"`                                 // Id:
	InhibitPaymentCardAuthorizationInd bool             `json:"inhibitPaymentCardAuthorizationInd,omitempty"` // InhibitPaymentCardAuthorizationInd: If true, the payment card will not go through card authorization process
}

// FormOfPaymentResponse is an object.
type FormOfPaymentResponse struct {
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"` // CurrencyRateConversion:
	FormOfPayment          *FormOfPaymentID         `json:"FormOfPayment,omitempty"`          // FormOfPayment:
	Identifier             *Identifier              `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`              // NextSteps:
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`          // ReferenceList:
	Result                 *Result                  `json:"Result,omitempty"`                 // Result:
	TraceId                string                   `json:"traceId,omitempty"`                // TraceId: Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TransactionId          string                   `json:"transactionId,omitempty"`          // TransactionId: Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
}

// FormOfPaymentVirtualPaymentAccount is an object.
type FormOfPaymentVirtualPaymentAccount struct {
	Type                    string           `json:"@type"`                             // Type:
	AccountID               string           `json:"AccountID,omitempty"`               // AccountID:
	AlternateEmailAddress   []Email          `json:"AlternateEmailAddress,omitempty"`   // AlternateEmailAddress: The alternate agency email to be used for correspondence with this virtual payment
	AlternateHotelFax       []Telephone      `json:"AlternateHotelFax,omitempty"`       // AlternateHotelFax: Hotel fax number to be used if the hotel fax is unknown or not provided in Property details
	FormOfPaymentRef        string           `json:"FormOfPaymentRef,omitempty"`        // FormOfPaymentRef:
	Identifier              *Identifier      `json:"Identifier,omitempty"`              // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	IncidentalCharges       []string         `json:"IncidentalCharges,omitempty"`       // IncidentalCharges: List of incidentals that are permitted to be charged to the virtual payment card.
	MaximumChargeableAmount []CurrencyAmount `json:"MaximumChargeableAmount,omitempty"` // MaximumChargeableAmount: The maximum amount that the supplier may charge to the payment card including room rate and any incidentals specified
	PaymentComment          string           `json:"PaymentComment,omitempty"`          // PaymentComment: Optional text to be sent to the supplier
	Supplier                string           `json:"Supplier,omitempty"`                // Supplier:
	Id                      string           `json:"id,omitempty"`                      // Id:
}

// FormOfPaymentWaiverCode is an object.
type FormOfPaymentWaiverCode struct {
	Type             string      `json:"@type"`                      // Type:
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"` // FormOfPaymentRef:
	Identifier       *Identifier `json:"Identifier,omitempty"`       // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	WaiverCode       *WaiverCode `json:"WaiverCode,omitempty"`       // WaiverCode:
	Id               string      `json:"id,omitempty"`               // Id:
}

// fromToValuePattern is the validation pattern for FromTo.Value
var fromToValuePattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// FromTo is an object. Location code
type FromTo struct {
	CityOrAirport CityOrAirportEnum `json:"cityOrAirport,omitempty"` // CityOrAirport: Clarification of how the airport or city code is used
	Value         string            `json:"value,omitempty"`         // Value:
}

// FulfillmentMethod is an object.
type FulfillmentMethod struct {
	Type                      string                     `json:"@type,omitempty"`                     // Type:
	ChangeFeeCollectionMethod *ChangeFeeCollectionMethod `json:"ChangeFeeCollectionMethod,omitempty"` // ChangeFeeCollectionMethod:
	RefundMethod              *RefundMethodEnum          `json:"RefundMethod,omitempty"`              // RefundMethod:
	ProductRefs               []string                   `json:"productRefs,omitempty"`               // ProductRefs: The product(s) the Fulfillment Method applies to. If blank applies to all products in the Offer
	SegmentSequenceList       []int32                    `json:"segmentSequenceList,omitempty"`       // SegmentSequenceList: List of segment sequence
}

// gSTRegistrationNumberCountryPattern is the validation pattern for GSTRegistrationNumber.Country
var gSTRegistrationNumberCountryPattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// GSTRegistrationNumber is an object. The GST Registration Number for this Organization
type GSTRegistrationNumber struct {
	Address     string `json:"address,omitempty"`     // Address: Address of the GST customer
	CompanyName string `json:"companyName,omitempty"` // CompanyName: Name of the Company
	Country     string `json:"country"`               // Country: Country
	Email       string `json:"email,omitempty"`       // Email: E-Mail
	Telephone   string `json:"telephone,omitempty"`   // Telephone: Telephone Number
	Value       string `json:"value,omitempty"`       // Value:
}

// GeoPoliticalArea is an object. The location code of the geographical location. Codes from Ref Pub
type GeoPoliticalArea struct {
	Id    string                    `json:"id,omitempty"`    // Id: Optional internally referenced id
	Level GeoPoliticalAreaLevelEnum `json:"level,omitempty"` // Level: Represents the type of geopolitical area (country, Continent, State etc)
	Value string                    `json:"value,omitempty"` // Value:
}

// Guarantee is an object.
type Guarantee struct {
	Type                   string            `json:"@type,omitempty"`                  // Type:
	Code                   string            `json:"code,omitempty"`                   // Code: Guarantee code
	CredentialsRequiredInd bool              `json:"credentialsRequiredInd,omitempty"` // CredentialsRequiredInd:
	GuaranteeType          GuaranteeTypeEnum `json:"guaranteeType,omitempty"`          // GuaranteeType: An enumerated type defining the guarantee to be applied to this reservation.
}

// GuestCount is an object.
type GuestCount struct {
	Type              string `json:"@type,omitempty"`             // Type:
	Age               int32  `json:"age,omitempty"`               // Age: The age of the guest
	AgeQualifyingCode string `json:"ageQualifyingCode,omitempty"` // AgeQualifyingCode: Enter 10 for an adult or 08 for a child
	Count             int32  `json:"count,omitempty"`             // Count: The number of guests in one AgeQualifyingCode or Count.
}

// GuestCounts is an object.
type GuestCounts struct {
	Type       string       `json:"@type,omitempty"` // Type:
	GuestCount []GuestCount `json:"GuestCount"`      // GuestCount:
}

// HotelPenalty is an object.
type HotelPenalty struct {
	Type         string           `json:"@type"`                  // Type:
	SubjectToTax YesNoUnknownEnum `json:"subjectToTax,omitempty"` // SubjectToTax: Yes , No , Unknown
}

// HotelPenaltyAmount is an object.
type HotelPenaltyAmount struct {
	Type         string           `json:"@type"`                  // Type:
	Amount       []CurrencyAmount `json:"Amount"`                 // Amount:
	IncludesTax  YesNoUnknownEnum `json:"includesTax,omitempty"`  // IncludesTax: Yes , No , Unknown
	SubjectToTax YesNoUnknownEnum `json:"subjectToTax,omitempty"` // SubjectToTax: Yes , No , Unknown
}

// HotelPenaltyNights is an object.
type HotelPenaltyNights struct {
	Type         string           `json:"@type"`                  // Type:
	Nights       int32            `json:"Nights"`                 // Nights: The number of nights that will be charged as a penalty
	SubjectToTax YesNoUnknownEnum `json:"subjectToTax,omitempty"` // SubjectToTax: Yes , No , Unknown
}

// HotelPenaltyPercent is an object.
type HotelPenaltyPercent struct {
	Type         string           `json:"@type"`                  // Type:
	Amount       []CurrencyAmount `json:"Amount,omitempty"`       // Amount:
	Nights       int32            `json:"Nights,omitempty"`       // Nights: The number of nights the percentage needs to be applied to determine cancel penalty amount
	Percent      float32          `json:"Percent"`                // Percent: A percentage charged as a Penalty
	AppliesTo    PercentAppliesTo `json:"appliesTo"`              // AppliesTo: The increment the percent applies to. Default value is Amount
	SubjectToTax YesNoUnknownEnum `json:"subjectToTax,omitempty"` // SubjectToTax: Yes , No , Unknown
}

// HotelSearchCriterion is an object.
type HotelSearchCriterion struct {
	Type               string              `json:"@type,omitempty"`              // Type:
	PropertyRequest    []PropertyRequest   `json:"PropertyRequest"`              // PropertyRequest:
	RateCandidates     *RateCandidates     `json:"RateCandidates,omitempty"`     // RateCandidates:
	RoomStayCandidates *RoomStayCandidates `json:"RoomStayCandidates,omitempty"` // RoomStayCandidates:
	NumberOfRooms      int32               `json:"numberOfRooms,omitempty"`      // NumberOfRooms: Number of rooms requested
}

// Identifier is an object. Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
type Identifier struct {
	Authority string `json:"authority,omitempty"` // Authority: Name of the authoritative system that created this Guid
	Value     string `json:"value,omitempty"`     // Value:
}

// IdentifierRef is an object.
type IdentifierRef struct {
	Description string   `json:"description,omitempty"` // Description: Descriptive text used to identify the contents of a target object
	Id          string   `json:"id,omitempty"`          // Id: A locally referenced ID
	Uris        []string `json:"uris,omitempty"`        // Uris: Uniform Resource Identifier
	Value       string   `json:"value,omitempty"`       // Value:
}

// Image is an object. URL of the image
type Image struct {
	Caption           string        `json:"caption,omitempty"`           // Caption: Image title
	DimensionCategory string        `json:"dimensionCategory,omitempty"` // DimensionCategory: Deprecated and replaced by Image Size
	Height            int32         `json:"height,omitempty"`            // Height: Height
	ImageSize         ImageSizeEnum `json:"imageSize,omitempty"`         // ImageSize: Indicates the size of the image. Hospitality APIs no longer support thumbnail
	PictureCategory   int32         `json:"pictureCategory,omitempty"`   // PictureCategory: deprecated and replaced by pictureOf
	PictureOf         PictureofEnum `json:"pictureOf,omitempty"`         // PictureOf:
	Value             string        `json:"value,omitempty"`             // Value:
	Width             int32         `json:"width,omitempty"`             // Width: Width of image
}

// intermediateStopArrivalDatePattern is the validation pattern for IntermediateStop.ArrivalDate
var intermediateStopArrivalDatePattern = regexp.MustCompile(`(\d{4}-\d{2}-\d{2})`)

// intermediateStopArrivalTerminalPattern is the validation pattern for IntermediateStop.ArrivalTerminal
var intermediateStopArrivalTerminalPattern = regexp.MustCompile(`([0-9a-zA-Z]+)?`)

// intermediateStopArrivalTimePattern is the validation pattern for IntermediateStop.ArrivalTime
var intermediateStopArrivalTimePattern = regexp.MustCompile(`(([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)((:?)[0-5]\d)?([\.,]\d+(?!:))?`)

// intermediateStopDepartureDatePattern is the validation pattern for IntermediateStop.DepartureDate
var intermediateStopDepartureDatePattern = regexp.MustCompile(`(\d{4}-\d{2}-\d{2})`)

// intermediateStopDepartureTerminalPattern is the validation pattern for IntermediateStop.DepartureTerminal
var intermediateStopDepartureTerminalPattern = regexp.MustCompile(`([0-9a-zA-Z]+)?`)

// intermediateStopDeparturelTimePattern is the validation pattern for IntermediateStop.DeparturelTime
var intermediateStopDeparturelTimePattern = regexp.MustCompile(`(([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)((:?)[0-5]\d)?([\.,]\d+(?!:))?`)

// intermediateStopEquipmentPattern is the validation pattern for IntermediateStop.Equipment
var intermediateStopEquipmentPattern = regexp.MustCompile(`([A-Z0-9]{3})?`)

// intermediateStopValuePattern is the validation pattern for IntermediateStop.Value
var intermediateStopValuePattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// IntermediateStop is an object. An intermediate stop location and duration
type IntermediateStop struct {
	ArrivalDate             string `json:"arrivalDate,omitempty"`             // ArrivalDate: The local date the aircraft arrives at the intermediate stop
	ArrivalFlightDuration   string `json:"arrivalFlightDuration,omitempty"`   // ArrivalFlightDuration: ArrivalFlight Duration stopped at this location
	ArrivalTerminal         string `json:"arrivalTerminal,omitempty"`         // ArrivalTerminal: Arrival Terminal of the Airport
	ArrivalTime             string `json:"arrivalTime,omitempty"`             // ArrivalTime: The local time the aircraft arrives at the intermediate stop
	DepartureDate           string `json:"departureDate,omitempty"`           // DepartureDate: The local date the aircraft departs from the intermediate stop
	DepartureFlightDuration string `json:"departureFlightDuration,omitempty"` // DepartureFlightDuration: DepartureFlight Duration stopped at this location
	DepartureTerminal       string `json:"departureTerminal,omitempty"`       // DepartureTerminal: Departure Terminal of the Airport
	DeparturelTime          string `json:"departurelTime,omitempty"`          // DeparturelTime: The local time the aircraft departs from the intermediate stop
	Duration                string `json:"duration,omitempty"`                // Duration: Duration stopped at this location
	Equipment               string `json:"equipment,omitempty"`               // Equipment: Aircraft equipment.
	Value                   string `json:"value,omitempty"`                   // Value:
}

// ListPaymentCardIssuerEnum is an object.
type ListPaymentCardIssuerEnum struct {
	Value ListPaymentCardIssuerEnumBase `json:"value,omitempty"` // Value: Source: OpenTravel
}

// locatorOtaTypePattern is the validation pattern for Locator.OtaType
var locatorOtaTypePattern = regexp.MustCompile(`[0-9A-Z]{1,3}(\.[A-Z]{3}(\.X){0,1}){0,1}`)

// locatorValuePattern is the validation pattern for Locator.Value
var locatorValuePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// Locator is an object. Contains the locator (PNR or external locator) or cancellation number for the reservation, order, or offer
type Locator struct {
	CreationDate  string `json:"creationDate,omitempty"`  // CreationDate: Reservation Creation date
	LocatorType   string `json:"locatorType,omitempty"`   // LocatorType: Specifies the type of reservation ID
	OtaType       string `json:"otaType,omitempty"`       // OtaType: Used for codes
	Source        string `json:"source,omitempty"`        // Source: Specifies a unique identifier to indicate the source system which generated the resid
	SourceContext string `json:"sourceContext,omitempty"` // SourceContext: Specifies the context of the source
	Value         string `json:"value,omitempty"`         // Value:
}

// magneticStripeMaskPattern is the validation pattern for MagneticStripe.Mask
var magneticStripeMaskPattern = regexp.MustCompile(`[0-9a-zA-Z]{1,32}`)

// magneticStripeTokenPattern is the validation pattern for MagneticStripe.Token
var magneticStripeTokenPattern = regexp.MustCompile(`[0-9a-zA-Z]{1,32}`)

// MagneticStripe is an object.
type MagneticStripe struct {
	Type                 string                      `json:"@type,omitempty"`                // Type:
	ErrorWarning         *ErrorWarning               `json:"ErrorWarning,omitempty"`         // ErrorWarning:
	PlainText            string                      `json:"PlainText,omitempty"`            // PlainText: Don't use this unless it is REALLY ok to not use encryption. Non-secure (plain text) value.
	AuthenticationMethod EncryptionTokenTypeAuthEnum `json:"authenticationMethod,omitempty"` // AuthenticationMethod: Type of Authentication
	EncryptedValue       string                      `json:"encryptedValue,omitempty"`       // EncryptedValue: Encrypted value
	EncryptionKey        string                      `json:"encryptionKey,omitempty"`        // EncryptionKey: Note: This contains a key required to retrieve the full payment instrument details compliant with PCI DSS standards.
	EncryptionKeyMethod  string                      `json:"encryptionKeyMethod,omitempty"`  // EncryptionKeyMethod: Developer: This contains a reference to the key generation method being used - this is NOT the key value.
	EncryptionMethod     string                      `json:"encryptionMethod,omitempty"`     // EncryptionMethod: OpenTravel Best Practice: Encryption Method: When using the OpenTravel Encryption element, it is RECOMMENDED that all trading partners be informed of all encryption methods being used in advance of implementation to ensure message processing compatibility.
	Mask                 string                      `json:"mask,omitempty"`                 // Mask: Masked Value
	Token                string                      `json:"token,omitempty"`                // Token: Token value
	TokenProviderID      string                      `json:"tokenProviderID,omitempty"`      // TokenProviderID: Developer: This contains a provider ID if multiple providers are used for secure information exchange.
}

// MealsIncluded is an object. Indicates if a meal is included or not
type MealsIncluded struct {
	BreakfastInd bool `json:"breakfastInd,omitempty"` // BreakfastInd:
	DinnerInd    bool `json:"dinnerInd,omitempty"`    // DinnerInd:
	LunchInd     bool `json:"lunchInd,omitempty"`     // LunchInd:
}

// Measurement is an object. Used for dimensional units (width, height, depth) or weight
type Measurement struct {
	MeasurementType MeasurementTypeEnum `json:"measurementType,omitempty"` // MeasurementType: The type of measurement such as width, height, weight
	Unit            UnitOfMeasureEnum   `json:"unit,omitempty"`            // Unit: The unit of measure in a code format. Refer to OpenTravel Code List Unit of Measure Code (UOM).
	Value           float32             `json:"value,omitempty"`           // Value:
}

// ModifyPrice is an object.
type ModifyPrice struct {
	Type                string               `json:"@type"`                         // Type:
	Base                float32              `json:"Base,omitempty"`                // Base: The total amount not including taxes and\/or fees
	CurrencyCode        *CurrencyCode        `json:"CurrencyCode,omitempty"`        // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	TotalFees           float32              `json:"TotalFees,omitempty"`           // TotalFees: The total of the fees included in the total price
	TotalPrice          float32              `json:"TotalPrice,omitempty"`          // TotalPrice: The total price of the product in the currency indicated
	TotalTaxes          float32              `json:"TotalTaxes,omitempty"`          // TotalTaxes: The total of the taxes included in the total price
	VendorCurrencyTotal *VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"` // VendorCurrencyTotal:
	Id                  string               `json:"id,omitempty"`                  // Id: Internally referenced id
}

// ModifyPriceDetail is an object.
type ModifyPriceDetail struct {
	Type                string               `json:"@type"`                         // Type:
	Base                float32              `json:"Base,omitempty"`                // Base: The total amount not including taxes and\/or fees
	CurrencyCode        *CurrencyCode        `json:"CurrencyCode,omitempty"`        // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	PriceBreakdown      []PriceBreakdown     `json:"PriceBreakdown,omitempty"`      // PriceBreakdown:
	TotalFees           float32              `json:"TotalFees,omitempty"`           // TotalFees: The total of the fees included in the total price
	TotalPrice          float32              `json:"TotalPrice,omitempty"`          // TotalPrice: The total price of the product in the currency indicated
	TotalTaxes          float32              `json:"TotalTaxes,omitempty"`          // TotalTaxes: The total of the taxes included in the total price
	VendorCurrencyTotal *VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"` // VendorCurrencyTotal:
	Id                  string               `json:"id,omitempty"`                  // Id: Internally referenced id
}

// NameValuePair is an object. Used for data stored in Name Value pairs
type NameValuePair struct {
	Id    string `json:"id,omitempty"`    // Id: Optional internally referenced id
	Name  string `json:"name"`            // Name: Key
	Value string `json:"value,omitempty"` // Value:
}

// NetFareBreakdownConstruction is an object.
type NetFareBreakdownConstruction struct {
	FareType       string  `json:"fareType,omitempty"`       // FareType: Assigned Type: c-1100:StringTiny
	RateOfExchange float32 `json:"rateOfExchange,omitempty"` // RateOfExchange: The rate of exchange applied to the fare breakdown
	Value          string  `json:"value,omitempty"`          // Value:
}

// netFareInfoPassengerTypeCodePattern is the validation pattern for NetFareInfo.PassengerTypeCode
var netFareInfoPassengerTypeCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{3,5})`)

// NetFareInfo is an object.
type NetFareInfo struct {
	NetFareBreakdownConstruction []NetFareBreakdownConstruction `json:"NetFareBreakdownConstruction,omitempty"` // NetFareBreakdownConstruction:
	TicketBaseAudit              *FiledAmount                   `json:"TicketBaseAudit,omitempty"`              // TicketBaseAudit: The base amount of a ticket price or net price that is filed in local currency
	TicketBasePassenger          *TicketBasePassenger           `json:"TicketBasePassenger,omitempty"`          // TicketBasePassenger: The monetary value
	PassengerTypeCode            string                         `json:"passengerTypeCode,omitempty"`            // PassengerTypeCode: PassengerTypeCode
}

// NetRemitInfo is an object.
type NetRemitInfo struct {
	Type              string       `json:"@type,omitempty"`             // Type:
	ActualSellingFare float32      `json:"ActualSellingFare,omitempty"` // ActualSellingFare: The actual selling fare which will override the Offer base fare on the document
	CarCode           string       `json:"CarCode,omitempty"`           // CarCode: The CAR code applied to this product for use in net remit
	NetBaseAmount     *FiledAmount `json:"NetBaseAmount,omitempty"`     // NetBaseAmount: The base amount of a ticket price or net price that is filed in local currency
	ValueCode         string       `json:"ValueCode,omitempty"`         // ValueCode: The Value code applied to this product for use in net remit
}

// nextResultReferenceProviderCodePattern is the validation pattern for NextResultReference.ProviderCode
var nextResultReferenceProviderCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,5})`)

// NextResultReference is an object. Container to return\/send additional search results
type NextResultReference struct {
	ProviderCode string `json:"providerCode,omitempty"` // ProviderCode: Assigned Type: c-1100:SupplierCode
	Value        string `json:"value,omitempty"`        // Value:
}

// NextStep is an object. A URL that describes a step that can be applied to the resource containing the next step structure.
type NextStep struct {
	Action      string             `json:"action"`                // Action: The action this next step is intended to achieve
	Description string             `json:"description,omitempty"` // Description: Additional clarification for the next step
	Id          string             `json:"id,omitempty"`          // Id: Identifier for the Next Step
	Method      NextStepMethodEnum `json:"method"`                // Method: Describes the set of potential methods that can be taken after an operation.
	Value       string             `json:"value,omitempty"`       // Value:
}

// NextSteps is an object.
type NextSteps struct {
	NextStep []NextStep `json:"NextStep"`     // NextStep:
	BaseURI  string     `json:"baseURI"`      // BaseURI: The base portion of the uri in order to shorten the uri's in the individual steps
	Id       string     `json:"id,omitempty"` // Id: Optional internally referenced id
}

// NightlyRate is an object.
type NightlyRate struct {
	Type      string `json:"@type,omitempty"`  // Type:
	Amount    Amount `json:"Amount"`           // Amount:
	Nights    int32  `json:"nights,omitempty"` // Nights: Number of nights this rate applies
	StartDate string `json:"startDate"`        // StartDate: Start date
}

// Notification is an object.
type Notification struct {
	Date        string        `json:"Date,omitempty"`        // Date: The notification date is equivalent to ticket time limit and will place the Reservation on the defined queue for the date specified. Sending a new notificiation date at commit step will update the existing notificationDate. Sending 000/00/00 will delete an existing notificationDate.
	QueueNumber []QueueNumber `json:"QueueNumber,omitempty"` // QueueNumber:
}

// OfferID is an object.
type OfferID struct {
	Type       string      `json:"@type"`                // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Local indentifier within a given message for this object.
	OfferRef   string      `json:"offerRef,omitempty"`   // OfferRef: Used to reference another instance of this object in the same message
}

// OfferIdentifier is an object.
type OfferIdentifier struct {
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Local indentifier within a given message for this object.
	OfferRef   string      `json:"offerRef,omitempty"`   // OfferRef: Used to reference another instance of this object in the same message
}

// OfferLink is an object.
type OfferLink struct {
	Type        string       `json:"@type,omitempty"`       // Type:
	OfferRef    string       `json:"OfferRef,omitempty"`    // OfferRef:
	ParentOffer *ParentOffer `json:"ParentOffer,omitempty"` // ParentOffer:
}

// OfferListResponse is an object.
type OfferListResponse struct {
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"` // CurrencyRateConversion:
	Identifier             *Identifier              `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`              // NextSteps:
	OfferID                []Offer                  `json:"OfferID,omitempty"`                // OfferID:
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`          // ReferenceList:
	Result                 *Result                  `json:"Result,omitempty"`                 // Result:
	TraceId                string                   `json:"traceId,omitempty"`                // TraceId: Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TransactionId          string                   `json:"transactionId,omitempty"`          // TransactionId: Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
}

// OfferModify is an object.
type OfferModify struct {
	Type          string       `json:"@type"`                   // Type:
	ConsumedPrice *PriceDetail `json:"ConsumedPrice,omitempty"` // ConsumedPrice:
	Identifier    *Identifier  `json:"Identifier,omitempty"`    // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ModifyPrice   struct {
		Type                string              `json:"@type,omitempty"`
		Base                float32             `json:"Base,omitempty"`
		CurrencyCode        CurrencyCode        `json:"CurrencyCode,omitempty"`
		PriceBreakdown      []PriceBreakdown    `json:"PriceBreakdown,omitempty"`
		TotalFees           float32             `json:"TotalFees,omitempty"`
		TotalPrice          float32             `json:"TotalPrice,omitempty"`
		TotalTaxes          float32             `json:"TotalTaxes,omitempty"`
		VendorCurrencyTotal VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"`
		Id                  string              `json:"id,omitempty"`
	} `json:"ModifyPrice"` // ModifyPrice:
	OfferIdentifier       *OfferIdentifier `json:"OfferIdentifier,omitempty"` // OfferIdentifier:
	Price                 Price            `json:"Price"`                     // Price:
	Product               []ProductID      `json:"Product"`                   // Product:
	SupplierRetainedPrice *struct {
		Type                string              `json:"@type,omitempty"`
		Base                float32             `json:"Base,omitempty"`
		CurrencyCode        CurrencyCode        `json:"CurrencyCode,omitempty"`
		PriceBreakdown      []PriceBreakdown    `json:"PriceBreakdown,omitempty"`
		TotalFees           float32             `json:"TotalFees,omitempty"`
		TotalPrice          float32             `json:"TotalPrice,omitempty"`
		TotalTaxes          float32             `json:"TotalTaxes,omitempty"`
		VendorCurrencyTotal VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"`
		Id                  string              `json:"id,omitempty"`
	} `json:"SupplierRetainedPrice,omitempty"` // SupplierRetainedPrice:
	TermsAndConditionsFull []TermsAndConditionsFull `json:"TermsAndConditionsFull"`       // TermsAndConditionsFull:
	Id                     string                   `json:"id,omitempty"`                 // Id: Local indentifier within a given message for this object.
	OfferRef               string                   `json:"offerRef,omitempty"`           // OfferRef: Used to reference another instance of this object in the same message
	ParentOfferRef         string                   `json:"parentOfferRef,omitempty"`     // ParentOfferRef: A reference to the Offer this offer is sold in conjunction with
	PassiveOfferInd        bool                     `json:"passiveOfferInd,omitempty"`    // PassiveOfferInd: If true, the Offer is passive for booking purposes.
	PriceUpdatedInd        bool                     `json:"priceUpdatedInd,omitempty"`    // PriceUpdatedInd: If present and true, the price in the host copy of the reservation has already been updated
	ProductsUpdatedInd     bool                     `json:"productsUpdatedInd,omitempty"` // ProductsUpdatedInd: If present and true, the products in the host copy of the reservation have already been updated
	RetainedValueInd       bool                     `json:"retainedValueInd,omitempty"`   // RetainedValueInd: If true, the value will be retained on a document for future use
	ScheduleChangeInd      bool                     `json:"scheduleChangeInd,omitempty"`  // ScheduleChangeInd: If true the Offer_Modify is as a result of an airline initiated schedule change
}

// OfferQueryBuildAncillaryOffersFromCatalogOfferings is an object.
type OfferQueryBuildAncillaryOffersFromCatalogOfferings struct {
	Type                                     string                                     `json:"@type"`                                    // Type:
	BuildAncillaryOffersFromCatalogOfferings []BuildAncillaryOffersFromCatalogOfferings `json:"BuildAncillaryOffersFromCatalogOfferings"` // BuildAncillaryOffersFromCatalogOfferings:
}

// OfferStatus is an object.
type OfferStatus struct {
	Type string `json:"@type"` // Type:
}

// OfferStatusAir is an object.
type OfferStatusAir struct {
	Type      string      `json:"@type"`               // Type:
	StatusAir []StatusAir `json:"StatusAir,omitempty"` // StatusAir:
}

// OfferStatusAncillary is an object.
type OfferStatusAncillary struct {
	Type            string            `json:"@type"`           // Type:
	StatusAncillary []StatusAncillary `json:"StatusAncillary"` // StatusAncillary:
}

// OfferStatusHospitality is an object.
type OfferStatusHospitality struct {
	Type   string          `json:"@type"`            // Type:
	Status OfferStatusEnum `json:"Status,omitempty"` // Status: Offer Status like confirmed ,Pending etc
}

// OfferStatusVehicle is an object.
type OfferStatusVehicle struct {
	Type   string          `json:"@type"`            // Type:
	Status OfferStatusEnum `json:"Status,omitempty"` // Status: Offer Status like confirmed ,Pending etc
}

// OfferUpsell is an object.
type OfferUpsell struct {
	Type                   string                   `json:"@type"`                     // Type:
	Identifier             *Identifier              `json:"Identifier,omitempty"`      // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OfferUpsellInd         bool                     `json:"OfferUpsellInd,omitempty"`  // OfferUpsellInd: If true, the OfferUpsell contains ancillary offerings offered in connection with the ParentOffer
	Price                  Price                    `json:"Price"`                     // Price:
	Product                []ProductID              `json:"Product"`                   // Product:
	TermsAndConditionsFull []TermsAndConditionsFull `json:"TermsAndConditionsFull"`    // TermsAndConditionsFull:
	Id                     string                   `json:"id,omitempty"`              // Id: Local indentifier within a given message for this object.
	OfferRef               string                   `json:"offerRef,omitempty"`        // OfferRef: Used to reference another instance of this object in the same message
	ParentOfferRef         string                   `json:"parentOfferRef,omitempty"`  // ParentOfferRef: A reference to the Offer this offer is sold in conjunction with
	PassiveOfferInd        bool                     `json:"passiveOfferInd,omitempty"` // PassiveOfferInd: If true, the Offer is passive for booking purposes.
}

// OperationTimes is an object.
type OperationTimes struct {
	Type       string          `json:"@type,omitempty"`      // Type:
	CloseTime  string          `json:"closeTime,omitempty"`  // CloseTime:
	DaysOfWeek []DayOfWeekEnum `json:"daysOfWeek,omitempty"` // DaysOfWeek:
	OpenTime   string          `json:"openTime,omitempty"`   // OpenTime:
}

// OperationalStatusENUM is an object.
type OperationalStatusENUM struct {
	Value OperationalStatusENUMBase `json:"value,omitempty"` // Value:
}

// organizationIdentifierSupplierPattern is the validation pattern for OrganizationIdentifier.Supplier
var organizationIdentifierSupplierPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,3})`)

// OrganizationIdentifier is an object. The organization identifier
type OrganizationIdentifier struct {
	AccountCodeFaresOnlyInd bool                     `json:"accountCodeFaresOnlyInd,omitempty"` // AccountCodeFaresOnlyInd: If true, account code only fares will be returned
	OrganizationCodeType    OrganizationCodeTypeEnum `json:"organizationCodeType,omitempty"`    // OrganizationCodeType: Defines the type of code given to the Organization to obtain discounts or additional benefits
	ProductRef              []string                 `json:"productRef,omitempty"`              // ProductRef: The productRef which the OrganizationIdentifier applies to
	SegmentSequenceList     []int32                  `json:"segmentSequenceList,omitempty"`     // SegmentSequenceList: SegmentSequenceList
	Supplier                string                   `json:"supplier,omitempty"`                // Supplier: Supplier code
	Value                   string                   `json:"value,omitempty"`                   // Value:
}

// OrganizationInformation is an object.
type OrganizationInformation struct {
	Type                   string                   `json:"@type,omitempty"`                  // Type:
	GSTRegistrationNumber  []GSTRegistrationNumber  `json:"GSTRegistrationNumber,omitempty"`  // GSTRegistrationNumber:
	OrganizationIdentifier []OrganizationIdentifier `json:"OrganizationIdentifier,omitempty"` // OrganizationIdentifier:
}

// OrganizationLoyaltyProgram is an object.
type OrganizationLoyaltyProgram struct {
	Type              string      `json:"@type"`                // Type:
	Identifier        *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	LoyaltyIdentifier string      `json:"LoyaltyIdentifier"`    // LoyaltyIdentifier: Loyalty Identifier
	Supplier          string      `json:"Supplier"`             // Supplier: The supplier of the loyalty program
	Id                string      `json:"id,omitempty"`         // Id:
}

// OrganizationLoyaltyProgramID is an object.
type OrganizationLoyaltyProgramID struct {
	Type       string      `json:"@type"`                // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id:
}

// PaidTaxes is an object.
type PaidTaxes struct {
	Type       string  `json:"@type"`                // Type:
	TotalTaxes float32 `json:"TotalTaxes,omitempty"` // TotalTaxes: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
}

// PaidTaxesDetail is an object.
type PaidTaxesDetail struct {
	Type       string      `json:"@type"`                // Type:
	Tax        []Tax       `json:"Tax,omitempty"`        // Tax:
	TaxPercent *TaxPercent `json:"TaxPercent,omitempty"` // TaxPercent:
	TotalTaxes float32     `json:"TotalTaxes,omitempty"` // TotalTaxes: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
}

// ParentOffer is an object.
type ParentOffer struct {
	Type       string `json:"@type"`                // Type:
	OfferRef   string `json:"OfferRef,omitempty"`   // OfferRef:
	ProductRef string `json:"ProductRef,omitempty"` // ProductRef:
}

// ParentOfferAir is an object.
type ParentOfferAir struct {
	Type             string `json:"@type"`                      // Type:
	FlightSegmentRef string `json:"FlightSegmentRef,omitempty"` // FlightSegmentRef:
	OfferRef         string `json:"OfferRef,omitempty"`         // OfferRef:
	ProductRef       string `json:"ProductRef,omitempty"`       // ProductRef:
}

// passengerCriteriaPassengerTypeCodePattern is the validation pattern for PassengerCriteria.PassengerTypeCode
var passengerCriteriaPassengerTypeCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{3,5})`)

// PassengerCriteria is an object.
type PassengerCriteria struct {
	Type                              string                      `json:"@type,omitempty"`                             // Type:
	CustomerLoyalty                   []CustomerLoyalty           `json:"CustomerLoyalty,omitempty"`                   // CustomerLoyalty:
	TravelerGeographicLocation        *TravelerGeographicLocation `json:"TravelerGeographicLocation,omitempty"`        // TravelerGeographicLocation: Specifies which location the Traveler resides in. Used for resident fares
	Age                               int32                       `json:"age,omitempty"`                               // Age: Age
	Number                            int32                       `json:"number,omitempty"`                            // Number: Number
	PassengerTypeCode                 string                      `json:"passengerTypeCode,omitempty"`                 // PassengerTypeCode: Passenger Type Code
	SpecifiedPassengerTypeCodeOnlyInd bool                        `json:"specifiedPassengerTypeCodeOnlyInd,omitempty"` // SpecifiedPassengerTypeCodeOnlyInd: If true then the Offering\/Offer will only be returned for the specific passengerTypeCode
}

// passengerCriteriaAirChangePassengerTypeCodePattern is the validation pattern for PassengerCriteriaAirChange.PassengerTypeCode
var passengerCriteriaAirChangePassengerTypeCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{3,5})`)

// PassengerCriteriaAirChange is an object.
type PassengerCriteriaAirChange struct {
	Type                              string                      `json:"@type,omitempty"`                             // Type:
	CustomerLoyalty                   []CustomerLoyalty           `json:"CustomerLoyalty,omitempty"`                   // CustomerLoyalty:
	TravelerGeographicLocation        *TravelerGeographicLocation `json:"TravelerGeographicLocation,omitempty"`        // TravelerGeographicLocation: Specifies which location the Traveler resides in. Used for resident fares
	TravelerIdentifierRef             IdentifierRef               `json:"TravelerIdentifierRef"`                       // TravelerIdentifierRef:
	Age                               int32                       `json:"age,omitempty"`                               // Age: Age
	BirthDate                         string                      `json:"birthDate,omitempty"`                         // BirthDate: The date of birth of the passenger. May be used in age validation for fares with age restrictions
	Number                            int32                       `json:"number,omitempty"`                            // Number: Number
	PassengerTypeCode                 string                      `json:"passengerTypeCode,omitempty"`                 // PassengerTypeCode: Passenger Type Code
	SpecifiedPassengerTypeCodeOnlyInd bool                        `json:"specifiedPassengerTypeCodeOnlyInd,omitempty"` // SpecifiedPassengerTypeCodeOnlyInd: If true then the Offering\/Offer will only be returned for the specific passengerTypeCode
}

// passengerFlightPassengerTypeCodePattern is the validation pattern for PassengerFlight.PassengerTypeCode
var passengerFlightPassengerTypeCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{3,5})`)

// PassengerFlight is an object.
type PassengerFlight struct {
	Type              string          `json:"@type,omitempty"`             // Type:
	FlightProduct     []FlightProduct `json:"FlightProduct,omitempty"`     // FlightProduct:
	PassengerQuantity int32           `json:"passengerQuantity,omitempty"` // PassengerQuantity: Number of passengers of the specified passenger type code
	PassengerTypeCode string          `json:"passengerTypeCode,omitempty"` // PassengerTypeCode: Passenger type code
}

// passwordMaskPattern is the validation pattern for Password.Mask
var passwordMaskPattern = regexp.MustCompile(`[0-9a-zA-Z]{1,32}`)

// passwordTokenPattern is the validation pattern for Password.Token
var passwordTokenPattern = regexp.MustCompile(`[0-9a-zA-Z]{1,32}`)

// Password is an object.
type Password struct {
	Type                 string                      `json:"@type,omitempty"`                // Type:
	ErrorWarning         *ErrorWarning               `json:"ErrorWarning,omitempty"`         // ErrorWarning:
	PlainText            string                      `json:"PlainText,omitempty"`            // PlainText: Don't use this unless it is REALLY ok to not use encryption. Non-secure (plain text) value.
	AuthenticationMethod EncryptionTokenTypeAuthEnum `json:"authenticationMethod,omitempty"` // AuthenticationMethod: Type of Authentication
	EncryptedValue       string                      `json:"encryptedValue,omitempty"`       // EncryptedValue: Encrypted value
	EncryptionKey        string                      `json:"encryptionKey,omitempty"`        // EncryptionKey: Note: This contains a key required to retrieve the full payment instrument details compliant with PCI DSS standards.
	EncryptionKeyMethod  string                      `json:"encryptionKeyMethod,omitempty"`  // EncryptionKeyMethod: Developer: This contains a reference to the key generation method being used - this is NOT the key value.
	EncryptionMethod     string                      `json:"encryptionMethod,omitempty"`     // EncryptionMethod: OpenTravel Best Practice: Encryption Method: When using the OpenTravel Encryption element, it is RECOMMENDED that all trading partners be informed of all encryption methods being used in advance of implementation to ensure message processing compatibility.
	Mask                 string                      `json:"mask,omitempty"`                 // Mask: Masked Value
	Token                string                      `json:"token,omitempty"`                // Token: Token value
	TokenProviderID      string                      `json:"tokenProviderID,omitempty"`      // TokenProviderID: Developer: This contains a provider ID if multiple providers are used for secure information exchange.
}

// paymentCardCardCodePattern is the validation pattern for PaymentCard.CardCode
var paymentCardCardCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// paymentCardApprovalCodePattern is the validation pattern for PaymentCard.ApprovalCode
var paymentCardApprovalCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// paymentCardExpireDatePattern is the validation pattern for PaymentCard.ExpireDate
var paymentCardExpireDatePattern = regexp.MustCompile(`(0[1-9]|1[0-2])[0-9][0-9]`)

// PaymentCard is an object.
type PaymentCard struct {
	Type           string                 `json:"@type"`                    // Type:
	CardBrand      *PaymentCardTypeIssuer `json:"CardBrand,omitempty"`      // CardBrand: This object contains Cards details for Payment
	CardCode       string                 `json:"CardCode,omitempty"`       // CardCode: Specifies the two character code (MC, VI, AX, etc) for the payment card (open enumeration)
	CardHolderName string                 `json:"CardHolderName,omitempty"` // CardHolderName: Name as displayed on Payment Card
	CardIssuer     *PaymentCardTypeIssuer `json:"CardIssuer,omitempty"`     // CardIssuer: This object contains Cards details for Payment
	CardNumber     *CardNumber            `json:"CardNumber,omitempty"`     // CardNumber:
	CardType       PaymentCardTypeEnum    `json:"CardType,omitempty"`       // CardType: Credit, Debit, etc.
	MagneticStripe []MagneticStripe       `json:"MagneticStripe,omitempty"` // MagneticStripe:
	PrivacyGroup   *Privacy               `json:"PrivacyGroup,omitempty"`   // PrivacyGroup: Confidential details for marketing purpose
	SeriesCode     *SeriesCode            `json:"SeriesCode,omitempty"`     // SeriesCode:
	ApprovalCode   string                 `json:"approvalCode,omitempty"`   // ApprovalCode: The approval code value
	EffectiveDate  string                 `json:"effectiveDate,omitempty"`  // EffectiveDate: Indicated starting date.
	ExpireDate     string                 `json:"expireDate,omitempty"`     // ExpireDate: The expiration date value
	Id             string                 `json:"id,omitempty"`             // Id: Payment card reference ID.
	SecureInd      bool                   `json:"secureInd,omitempty"`      // SecureInd: Implementer: If true, all or a portion of this data is secure, via tokenization, encryption and\/or masking.
}

// paymentCardDetailCardCodePattern is the validation pattern for PaymentCardDetail.CardCode
var paymentCardDetailCardCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// paymentCardDetailApprovalCodePattern is the validation pattern for PaymentCardDetail.ApprovalCode
var paymentCardDetailApprovalCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// paymentCardDetailBankCountryCodePattern is the validation pattern for PaymentCardDetail.BankCountryCode
var paymentCardDetailBankCountryCodePattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// paymentCardDetailBankStateCodePattern is the validation pattern for PaymentCardDetail.BankStateCode
var paymentCardDetailBankStateCodePattern = regexp.MustCompile(`([a-zA-Z]{2})`)

// paymentCardDetailCompanyCardReferencePattern is the validation pattern for PaymentCardDetail.CompanyCardReference
var paymentCardDetailCompanyCardReferencePattern = regexp.MustCompile(`([0-9a-zA-Z]+)?`)

// paymentCardDetailCountryOfIssuePattern is the validation pattern for PaymentCardDetail.CountryOfIssue
var paymentCardDetailCountryOfIssuePattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// paymentCardDetailExpireDatePattern is the validation pattern for PaymentCardDetail.ExpireDate
var paymentCardDetailExpireDatePattern = regexp.MustCompile(`(0[1-9]|1[0-2])[0-9][0-9]`)

// PaymentCardDetail is an object.
type PaymentCardDetail struct {
	Type                  string                 `json:"@type"`                           // Type:
	Address               *Address               `json:"Address,omitempty"`               // Address:
	CardBrand             *PaymentCardTypeIssuer `json:"CardBrand,omitempty"`             // CardBrand: This object contains Cards details for Payment
	CardCode              string                 `json:"CardCode,omitempty"`              // CardCode: Specifies the two character code (MC, VI, AX, etc) for the payment card (open enumeration)
	CardHolderId          *Identifier            `json:"CardHolderId,omitempty"`          // CardHolderId: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	CardHolderName        string                 `json:"CardHolderName,omitempty"`        // CardHolderName: Name as displayed on Payment Card
	CardIssuer            *PaymentCardTypeIssuer `json:"CardIssuer,omitempty"`            // CardIssuer: This object contains Cards details for Payment
	CardNumber            *CardNumber            `json:"CardNumber,omitempty"`            // CardNumber:
	CardType              PaymentCardTypeEnum    `json:"CardType,omitempty"`              // CardType: Credit, Debit, etc.
	CustomerLoyalty       []CustomerLoyalty      `json:"CustomerLoyalty,omitempty"`       // CustomerLoyalty:
	Email                 []Email                `json:"Email,omitempty"`                 // Email:
	MagneticStripe        []MagneticStripe       `json:"MagneticStripe,omitempty"`        // MagneticStripe:
	PersonName            *PersonName            `json:"PersonName,omitempty"`            // PersonName:
	PrivacyGroup          *Privacy               `json:"PrivacyGroup,omitempty"`          // PrivacyGroup: Confidential details for marketing purpose
	SeriesCode            *SeriesCode            `json:"SeriesCode,omitempty"`            // SeriesCode:
	SignatureOnFile       *SignatureOnFile       `json:"SignatureOnFile,omitempty"`       // SignatureOnFile:
	Telephone             []Telephone            `json:"Telephone,omitempty"`             // Telephone:
	ThreeDomainSecurity   *ThreeDomainSecurity   `json:"ThreeDomainSecurity,omitempty"`   // ThreeDomainSecurity:
	AcceptanceOverrideInd bool                   `json:"acceptanceOverrideInd,omitempty"` // AcceptanceOverrideInd: If true, override airline restriction on the payment card
	ApprovalCode          string                 `json:"approvalCode,omitempty"`          // ApprovalCode: The approval code value
	BankCountryCode       string                 `json:"bankCountryCode,omitempty"`       // BankCountryCode: The bank country code ISO
	BankName              string                 `json:"bankName,omitempty"`              // BankName: The bank name value
	BankStateCode         string                 `json:"bankStateCode,omitempty"`         // BankStateCode: The bank state code ISO
	CompanyCardReference  string                 `json:"companyCardReference,omitempty"`  // CompanyCardReference: The company card reference
	CountryOfIssue        string                 `json:"countryOfIssue,omitempty"`        // CountryOfIssue: The country code ISO
	EffectiveDate         string                 `json:"effectiveDate,omitempty"`         // EffectiveDate: Indicated starting date.
	EnettInd              bool                   `json:"enettInd,omitempty"`              // EnettInd: True if this payment card has been issued through Enett
	ExpireDate            string                 `json:"expireDate,omitempty"`            // ExpireDate: The expiration date value
	ExtendedPaymentInd    bool                   `json:"extendedPaymentInd,omitempty"`    // ExtendedPaymentInd: Implementer: If true, the credit card company is requested to delay the date on which the amount of this transaction is applied to the customer's account.
	Id                    string                 `json:"id,omitempty"`                    // Id: Payment card reference ID.
	SecureInd             bool                   `json:"secureInd,omitempty"`             // SecureInd: Implementer: If true, all or a portion of this data is secure, via tokenization, encryption and\/or masking.
	ThirdPartyInd         bool                   `json:"thirdPartyInd,omitempty"`         // ThirdPartyInd: If true, then the payment card holder is not one of the travelers in the reservation
}

// PaymentCardTypeIssuer is an object. This object contains Cards details for Payment
type PaymentCardTypeIssuer struct {
	IssueNumber                 int32                      `json:"issueNumber,omitempty"`                 // IssueNumber: Assigned Type: c-1100:NumberDoubleDigit
	PaymentCardIssuers          *ListPaymentCardIssuerEnum `json:"paymentCardIssuers,omitempty"`          // PaymentCardIssuers:
	PaymentCardIssuersExtension string                     `json:"paymentCardIssuersExtension,omitempty"` // PaymentCardIssuersExtension:
}

// paymentCriteriaIssuerIdentificationNumberPattern is the validation pattern for PaymentCriteria.IssuerIdentificationNumber
var paymentCriteriaIssuerIdentificationNumberPattern = regexp.MustCompile(`[0-9]{6,11}`)

// paymentCriteriaPaymentCardCodePattern is the validation pattern for PaymentCriteria.PaymentCardCode
var paymentCriteriaPaymentCardCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// PaymentCriteria is an object.
type PaymentCriteria struct {
	Type                       string           `json:"@type,omitempty"`                      // Type:
	DocumentNumber             []DocumentNumber `json:"DocumentNumber,omitempty"`             // DocumentNumber:
	IssuerIdentificationNumber string           `json:"IssuerIdentificationNumber,omitempty"` // IssuerIdentificationNumber: This the BIN/IIN
	PaymentCardCode            string           `json:"PaymentCardCode,omitempty"`            // PaymentCardCode: A two character code for a credit card, like MC, AX
	AgencyAccountInd           bool             `json:"agencyAccountInd,omitempty"`           // AgencyAccountInd: If true, payment will be made by agency account
	BspInd                     bool             `json:"bspInd,omitempty"`                     // BspInd: If true, payment will be made by BSP
	CashInd                    bool             `json:"cashInd,omitempty"`                    // CashInd: If true, payment will be made by cash
	InvoiceInd                 bool             `json:"invoiceInd,omitempty"`                 // InvoiceInd: If true, payment will be made by invoice
}

// PaymentID is an object.
type PaymentID struct {
	Type       string      `json:"@type,omitempty"`      // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	PaymentRef string      `json:"PaymentRef,omitempty"` // PaymentRef:
	Id         string      `json:"id,omitempty"`         // Id:
}

// PaymentIdentifier is an object.
type PaymentIdentifier struct {
	Type       string      `json:"@type,omitempty"`      // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	PaymentRef string      `json:"PaymentRef,omitempty"` // PaymentRef:
	Id         string      `json:"id,omitempty"`         // Id:
}

// PaymentListResponse is an object.
type PaymentListResponse struct {
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"` // CurrencyRateConversion:
	Identifier             *Identifier              `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`              // NextSteps:
	PaymentID              []PaymentID              `json:"PaymentID,omitempty"`              // PaymentID:
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`          // ReferenceList:
	Result                 *Result                  `json:"Result,omitempty"`                 // Result:
	TraceId                string                   `json:"traceId,omitempty"`                // TraceId: Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TransactionId          string                   `json:"transactionId,omitempty"`          // TransactionId: Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
}

// PaymentResponse is an object.
type PaymentResponse struct {
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"` // CurrencyRateConversion:
	Identifier             *Identifier              `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`              // NextSteps:
	Payment                *PaymentID               `json:"Payment,omitempty"`                // Payment:
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`          // ReferenceList:
	Result                 *Result                  `json:"Result,omitempty"`                 // Result:
	TraceId                string                   `json:"traceId,omitempty"`                // TraceId: Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TransactionId          string                   `json:"transactionId,omitempty"`          // TransactionId: Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
}

// Penalties is an object.
type Penalties struct {
	Type               string             `json:"@type,omitempty"`              // Type:
	Cancel             []Cancel           `json:"Cancel,omitempty"`             // Cancel:
	Change             []Change           `json:"Change,omitempty"`             // Change:
	PassengerTypeCodes []string           `json:"PassengerTypeCodes,omitempty"` // PassengerTypeCodes: The passenger type codes that this penalty applies to
	PenaltySourceCode  *PenaltySourceCode `json:"PenaltySourceCode,omitempty"`  // PenaltySourceCode:
	Waiver             []WaiverEnum       `json:"Waiver,omitempty"`             // Waiver:
	InvoluntaryInd     bool               `json:"involuntaryInd,omitempty"`     // InvoluntaryInd: Penalties apply for involuntary changes initiated by the airline
}

// Penalty is an object.
type Penalty struct {
	Type        string         `json:"@type"`                 // Type:
	Application CommissionEnum `json:"application,omitempty"` // Application: Type of commission
}

// PenaltyAmount is an object.
type PenaltyAmount struct {
	Type        string          `json:"@type"`                 // Type:
	Amount      *CurrencyAmount `json:"Amount,omitempty"`      // Amount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	Application CommissionEnum  `json:"application,omitempty"` // Application: Type of commission
}

// PenaltyPercent is an object.
type PenaltyPercent struct {
	Type        string         `json:"@type"`                 // Type:
	Percent     float32        `json:"Percent,omitempty"`     // Percent: Percent amount of commission
	Application CommissionEnum `json:"application,omitempty"` // Application: Type of commission
}

// PenaltySourceCode is an object.
type PenaltySourceCode struct {
	CodeContext string `json:"codeContext,omitempty"` // CodeContext: Penalty source code context
	Value       string `json:"value,omitempty"`       // Value:
}

// PermittedVendors is an object.
type PermittedVendors struct {
	Vendor []CompanyName `json:"Vendor"` // Vendor:
}

// PersonName is an object.
type PersonName struct {
	Type    string `json:"@type"`            // Type:
	Given   string `json:"Given,omitempty"`  // Given: Given name, first name or names.
	Middle  string `json:"Middle,omitempty"` // Middle: The middle name of the person name.
	Prefix  string `json:"Prefix,omitempty"` // Prefix: Salutation of honorific (e.g. Mr., Mrs., Ms., Miss, Dr.)
	Surname string `json:"Surname"`          // Surname: Family name, last name.
}

// PersonNameDetail is an object.
type PersonNameDetail struct {
	Type           string       `json:"@type"`                    // Type:
	Given          string       `json:"Given,omitempty"`          // Given: Given name, first name or names.
	Middle         string       `json:"Middle,omitempty"`         // Middle: The middle name of the person name.
	Prefix         string       `json:"Prefix,omitempty"`         // Prefix: Salutation of honorific (e.g. Mr., Mrs., Ms., Miss, Dr.)
	Privacy        *Privacy     `json:"Privacy,omitempty"`        // Privacy: Confidential details for marketing purpose
	Suffix         string       `json:"Suffix,omitempty"`         // Suffix: Hold various name suffixes and letters
	Surname        string       `json:"Surname"`                  // Surname: Family name, last name.
	SurnamePrefix  string       `json:"SurnamePrefix,omitempty"`  // SurnamePrefix: The surname prefix
	Title          string       `json:"Title,omitempty"`          // Title: Degree or honors
	DefaultInd     bool         `json:"defaultInd,omitempty"`     // DefaultInd: If true, this is the default or primary name within a collection of names.
	Language       string       `json:"language,omitempty"`       // Language: 'ISO639 code of the language the name is represented
	PersonNameType NameTypeEnum `json:"personNameType,omitempty"` // PersonNameType: OTA Code
}

// PersonNameUpdatable is an object.
type PersonNameUpdatable struct {
	Type    string `json:"@type,omitempty"` // Type:
	Given   string `json:"Given"`           // Given: The first given name of the person
	Middle  string `json:"Middle"`          // Middle: The middle name of the person
	Prefix  string `json:"Prefix"`          // Prefix: Salutation of honorific
	Suffix  string `json:"Suffix"`          // Suffix: Name suffix
	Surname string `json:"Surname"`         // Surname: Family name, last name
}

// Policy is an object.
type Policy struct {
	Type          string          `json:"@type"`                   // Type:
	SubPolicy     []SubPolicy     `json:"SubPolicy,omitempty"`     // SubPolicy:
	TextFormatted []TextFormatted `json:"TextFormatted,omitempty"` // TextFormatted:
	Title         string          `json:"title,omitempty"`         // Title: Assigned Type: c-1100:String
}

// Preference is an object.
type Preference struct {
	Type               string              `json:"@type"`                        // Type:
	AppliesTo          *AppliesTo          `json:"AppliesTo,omitempty"`          // AppliesTo:
	TravelerIdentifier *TravelerIdentifier `json:"TravelerIdentifier,omitempty"` // TravelerIdentifier:
	Id                 string              `json:"id,omitempty"`                 // Id:
}

// PreferenceAirSeat is an object.
type PreferenceAirSeat struct {
	Type               string              `json:"@type"`                        // Type:
	AppliesTo          *AppliesTo          `json:"AppliesTo,omitempty"`          // AppliesTo:
	SeatLocation       SeatLocationEnum    `json:"SeatLocation"`                 // SeatLocation: Window, aisle, middle, etc.
	TravelerIdentifier *TravelerIdentifier `json:"TravelerIdentifier,omitempty"` // TravelerIdentifier:
	Id                 string              `json:"id,omitempty"`                 // Id:
}

// PreferenceID is an object.
type PreferenceID struct {
	Type string `json:"@type"`        // Type:
	Id   string `json:"id,omitempty"` // Id:
}

// Price is an object.
type Price struct {
	Type                string               `json:"@type"`                         // Type:
	Base                float32              `json:"Base,omitempty"`                // Base: The total amount not including taxes and\/or fees
	CurrencyCode        *CurrencyCode        `json:"CurrencyCode,omitempty"`        // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	TotalFees           float32              `json:"TotalFees,omitempty"`           // TotalFees: The total of the fees included in the total price
	TotalPrice          float32              `json:"TotalPrice,omitempty"`          // TotalPrice: The total price of the product in the currency indicated
	TotalTaxes          float32              `json:"TotalTaxes,omitempty"`          // TotalTaxes: The total of the taxes included in the total price
	VendorCurrencyTotal *VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"` // VendorCurrencyTotal:
	Id                  string               `json:"id,omitempty"`                  // Id: Internally referenced id
}

// PriceBreakdown is an object.
type PriceBreakdown struct {
	Type       string      `json:"@type"`                // Type:
	Amount     *Amount     `json:"Amount,omitempty"`     // Amount:
	Commission *Commission `json:"Commission,omitempty"` // Commission:
}

// priceBreakdownAirRequestedPassengerTypePattern is the validation pattern for PriceBreakdownAir.RequestedPassengerType
var priceBreakdownAirRequestedPassengerTypePattern = regexp.MustCompile(`([a-zA-Z0-9]{3,5})`)

// PriceBreakdownAir is an object.
type PriceBreakdownAir struct {
	Type                   string         `json:"@type"`                            // Type:
	Amount                 *Amount        `json:"Amount,omitempty"`                 // Amount:
	Commission             *Commission    `json:"Commission,omitempty"`             // Commission:
	Discount               *Discount      `json:"Discount,omitempty"`               // Discount: Corporate or Other discount
	FareCalculation        string         `json:"FareCalculation,omitempty"`        // FareCalculation:
	FiledAmount            *FiledAmount   `json:"FiledAmount,omitempty"`            // FiledAmount: The base amount of a ticket price or net price that is filed in local currency
	NetBaseAmount          *FiledAmount   `json:"NetBaseAmount,omitempty"`          // NetBaseAmount: The base amount of a ticket price or net price that is filed in local currency
	NetFareInfo            *NetFareInfo   `json:"NetFareInfo,omitempty"`            // NetFareInfo:
	Surcharges             *Surcharges    `json:"Surcharges,omitempty"`             // Surcharges:
	TravelerIdentifierRef  *IdentifierRef `json:"TravelerIdentifierRef,omitempty"`  // TravelerIdentifierRef:
	Quantity               int32          `json:"quantity,omitempty"`               // Quantity: The quantity value
	RequestedPassengerType string         `json:"requestedPassengerType,omitempty"` // RequestedPassengerType: The requested passenger type code
}

// PriceBreakdownAncillary is an object.
type PriceBreakdownAncillary struct {
	Type                  string                `json:"@type"`                           // Type:
	Amount                *Amount               `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission           `json:"Commission,omitempty"`            // Commission:
	Description           *AncillaryDescription `json:"Description,omitempty"`           // Description: A description of the ancillary with two description codes
	Discount              *Discount             `json:"Discount,omitempty"`              // Discount: Corporate or Other discount
	ProductRef            string                `json:"ProductRef,omitempty"`            // ProductRef: The product ref this PriceBreakdown applies to. If no productRef exists then the PriceBreakdown applies to all Products within the Offer.
	TravelerIdentifierRef *IdentifierRef        `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	Quantity              int32                 `json:"quantity,omitempty"`              // Quantity: The quantity of ancillary items included in this PriceBreakdown
}

// priceBreakdownAncillaryAirPassengerTypeCodePattern is the validation pattern for PriceBreakdownAncillaryAir.PassengerTypeCode
var priceBreakdownAncillaryAirPassengerTypeCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{3,5})`)

// PriceBreakdownAncillaryAir is an object.
type PriceBreakdownAncillaryAir struct {
	Type                  string                `json:"@type"`                           // Type:
	Amount                *Amount               `json:"Amount,omitempty"`                // Amount:
	Commission            *Commission           `json:"Commission,omitempty"`            // Commission:
	Description           *AncillaryDescription `json:"Description,omitempty"`           // Description: A description of the ancillary with two description codes
	Discount              *Discount             `json:"Discount,omitempty"`              // Discount: Corporate or Other discount
	PassengerTypeCode     string                `json:"PassengerTypeCode,omitempty"`     // PassengerTypeCode: The passenger type code the ancillary is valid for
	ProductRef            string                `json:"ProductRef,omitempty"`            // ProductRef: The product ref this PriceBreakdown applies to. If no productRef exists then the PriceBreakdown applies to all Products within the Offer.
	TravelerIdentifierRef *IdentifierRef        `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
	ApproximateInd        bool                  `json:"approximateInd,omitempty"`        // ApproximateInd: Used to indicate that the Price is approximate. Often used to allow for currency fluctuations when supplier currency is different to agency currency.
	Quantity              int32                 `json:"quantity,omitempty"`              // Quantity: The quantity of ancillary items included in this PriceBreakdown
}

// PriceBreakdownAncillaryVehicle is an object.
type PriceBreakdownAncillaryVehicle struct {
	Type                        string                `json:"@type"`                                 // Type:
	Amount                      *Amount               `json:"Amount,omitempty"`                      // Amount:
	Commission                  *Commission           `json:"Commission,omitempty"`                  // Commission:
	Description                 *AncillaryDescription `json:"Description,omitempty"`                 // Description: A description of the ancillary with two description codes
	Discount                    *Discount             `json:"Discount,omitempty"`                    // Discount: Corporate or Other discount
	ProductRef                  string                `json:"ProductRef,omitempty"`                  // ProductRef: The product ref this PriceBreakdown applies to. If no productRef exists then the PriceBreakdown applies to all Products within the Offer.
	RatePeriod                  RatePeriodEnum        `json:"RatePeriod,omitempty"`                  // RatePeriod: The time period for a rate such as daily, weekly, monthly
	TravelerIdentifierRef       *IdentifierRef        `json:"TravelerIdentifierRef,omitempty"`       // TravelerIdentifierRef:
	IncludedInEstimatedTotalInd bool                  `json:"includedInEstimatedTotalInd,omitempty"` // IncludedInEstimatedTotalInd: If true the AncillaryVehicle is included in the estimated totalPrice.
	PayNowInd                   bool                  `json:"payNowInd,omitempty"`                   // PayNowInd: If true the vehicle ancillary must be paid now and is included in the totalPrice calculation
	Quantity                    int32                 `json:"quantity,omitempty"`                    // Quantity: The quantity of ancillary items included in this PriceBreakdown
}

// PriceBreakdownHospitality is an object.
type PriceBreakdownHospitality struct {
	Type              string `json:"@type"` // Type:
	AmenitySurcharges *struct {
		Type            string      `json:"@type,omitempty"`
		Surcharge       []Surcharge `json:"Surcharge,omitempty"`
		TotalSurcharges float32     `json:"TotalSurcharges,omitempty"`
		ApproximateInd  bool        `json:"approximateInd,omitempty"`
	} `json:"AmenitySurcharges,omitempty"` // AmenitySurcharges:
	Amount                    *Amount          `json:"Amount,omitempty"`                    // Amount:
	AverageNightlyRate        []CurrencyAmount `json:"AverageNightlyRate,omitempty"`        // AverageNightlyRate:
	Commission                *Commission      `json:"Commission,omitempty"`                // Commission:
	Description               string           `json:"Description,omitempty"`               // Description:
	NightlyRate               []NightlyRate    `json:"NightlyRate,omitempty"`               // NightlyRate:
	PriceChangesDuringStayInd bool             `json:"priceChangesDuringStayInd,omitempty"` // PriceChangesDuringStayInd: If present and true, indicates the nightly price changes one or more times during the stay
	RoomPricingType           PricingEnum      `json:"roomPricingType,omitempty"`           // RoomPricingType: An enumerated type that defines how a service is priced.
}

// PriceBreakdownVehicleCharges is an object.
type PriceBreakdownVehicleCharges struct {
	Type           string          `json:"@type"`                    // Type:
	Amount         *Amount         `json:"Amount,omitempty"`         // Amount:
	Commission     *Commission     `json:"Commission,omitempty"`     // Commission:
	VehicleCharges *VehicleCharges `json:"VehicleCharges,omitempty"` // VehicleCharges:
}

// PriceBreakdownVehicleDeposit is an object.
type PriceBreakdownVehicleDeposit struct {
	Type                 string      `json:"@type"`                          // Type:
	Amount               *Amount     `json:"Amount,omitempty"`               // Amount:
	Commission           *Commission `json:"Commission,omitempty"`           // Commission:
	TotalPayableLaterInd bool        `json:"totalPayableLaterInd,omitempty"` // TotalPayableLaterInd: If True the Amount is the total amount payable later
	TotalPayableNowInd   bool        `json:"totalPayableNowInd,omitempty"`   // TotalPayableNowInd: If True the Amount is the total amount payable now
}

// PriceBreakdownVehiclePrice is an object.
type PriceBreakdownVehiclePrice struct {
	Type              string       `json:"@type"`                       // Type:
	Amount            *Amount      `json:"Amount,omitempty"`            // Amount:
	Commission        *Commission  `json:"Commission,omitempty"`        // Commission:
	VehiclePrice      VehiclePrice `json:"VehiclePrice"`                // VehiclePrice:
	RateGuaranteedInd bool         `json:"rateGuaranteedInd,omitempty"` // RateGuaranteedInd:
}

// PriceDetail is an object.
type PriceDetail struct {
	Type                string               `json:"@type"`                         // Type:
	Base                float32              `json:"Base,omitempty"`                // Base: The total amount not including taxes and\/or fees
	CurrencyCode        *CurrencyCode        `json:"CurrencyCode,omitempty"`        // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	PriceBreakdown      []PriceBreakdown     `json:"PriceBreakdown,omitempty"`      // PriceBreakdown:
	TotalFees           float32              `json:"TotalFees,omitempty"`           // TotalFees: The total of the fees included in the total price
	TotalPrice          float32              `json:"TotalPrice,omitempty"`          // TotalPrice: The total price of the product in the currency indicated
	TotalTaxes          float32              `json:"TotalTaxes,omitempty"`          // TotalTaxes: The total of the taxes included in the total price
	VendorCurrencyTotal *VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"` // VendorCurrencyTotal:
	Id                  string               `json:"id,omitempty"`                  // Id: Internally referenced id
}

// pricingAgencyCodePattern is the validation pattern for PricingAgency.Code
var pricingAgencyCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// PricingAgency is an object.
type PricingAgency struct {
	Type                string   `json:"@type,omitempty"`               // Type:
	Code                string   `json:"Code"`                          // Code: The Pricing Agency PCC
	ProductRef          []string `json:"ProductRef,omitempty"`          // ProductRef:
	SegmentSequenceList []int32  `json:"SegmentSequenceList,omitempty"` // SegmentSequenceList:
}

// pricingModifiersAirPricingPCCPattern is the validation pattern for PricingModifiersAir.PricingPCC
var pricingModifiersAirPricingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// pricingModifiersAirSellCityPattern is the validation pattern for PricingModifiersAir.SellCity
var pricingModifiersAirSellCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// pricingModifiersAirTicketCityPattern is the validation pattern for PricingModifiersAir.TicketCity
var pricingModifiersAirTicketCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// pricingModifiersAirTicketingPCCPattern is the validation pattern for PricingModifiersAir.TicketingPCC
var pricingModifiersAirTicketingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// pricingModifiersAirCurrencyCodePattern is the validation pattern for PricingModifiersAir.CurrencyCode
var pricingModifiersAirCurrencyCodePattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// PricingModifiersAir is an object.
type PricingModifiersAir struct {
	Type                          string                   `json:"@type"`                                   // Type:
	FareSelection                 *FareSelectionDetail     `json:"FareSelection,omitempty"`                 // FareSelection:
	OrganizationInformation       *OrganizationInformation `json:"OrganizationInformation,omitempty"`       // OrganizationInformation:
	PricingPCC                    string                   `json:"PricingPCC,omitempty"`                    // PricingPCC:
	PromotionalCode               []PromotionalCode        `json:"PromotionalCode,omitempty"`               // PromotionalCode:
	SellCity                      string                   `json:"SellCity,omitempty"`                      // SellCity: Overrides the sell city of the requestor.
	TaxExemption                  *TaxExemption            `json:"TaxExemption,omitempty"`                  // TaxExemption:
	TicketCity                    string                   `json:"TicketCity,omitempty"`                    // TicketCity: Overrides the ticket city of the requestor.
	TicketingPCC                  string                   `json:"TicketingPCC,omitempty"`                  // TicketingPCC:
	CurrencyCode                  string                   `json:"currencyCode,omitempty"`                  // CurrencyCode: Currency Code ISO
	IncludeSplitPaymentInd        bool                     `json:"includeSplitPaymentInd,omitempty"`        // IncludeSplitPaymentInd: If true, split payment (split ticket) offerings\/offers will be returned
	ReturnMostRestrictiveBrandInd bool                     `json:"returnMostRestrictiveBrandInd,omitempty"` // ReturnMostRestrictiveBrandInd: if true, the most restrictive brand will be returned in the response when there are different brands present in the Offering
	SplitPaymentOfferings         float32                  `json:"splitPaymentOfferings,omitempty"`         // SplitPaymentOfferings: The percentage, between 0 and 99, of round trip offerings the user would like returned in the result set.
}

// pricingModifiersAirChangePricingPCCPattern is the validation pattern for PricingModifiersAirChange.PricingPCC
var pricingModifiersAirChangePricingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// pricingModifiersAirChangeSellCityPattern is the validation pattern for PricingModifiersAirChange.SellCity
var pricingModifiersAirChangeSellCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// pricingModifiersAirChangeTicketCityPattern is the validation pattern for PricingModifiersAirChange.TicketCity
var pricingModifiersAirChangeTicketCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// pricingModifiersAirChangeTicketingPCCPattern is the validation pattern for PricingModifiersAirChange.TicketingPCC
var pricingModifiersAirChangeTicketingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// pricingModifiersAirChangeCurrencyCodePattern is the validation pattern for PricingModifiersAirChange.CurrencyCode
var pricingModifiersAirChangeCurrencyCodePattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// PricingModifiersAirChange is an object.
type PricingModifiersAirChange struct {
	Type                          string                    `json:"@type,omitempty"`                         // Type:
	CalculatedFareAdjustment      *CalculatedFareAdjustment `json:"CalculatedFareAdjustment,omitempty"`      // CalculatedFareAdjustment:
	FareSelection                 *FareSelection            `json:"FareSelection,omitempty"`                 // FareSelection:
	OrganizationInformation       *OrganizationInformation  `json:"OrganizationInformation,omitempty"`       // OrganizationInformation:
	PricingPCC                    string                    `json:"PricingPCC,omitempty"`                    // PricingPCC:
	PromotionalCode               *PromotionalCode          `json:"PromotionalCode,omitempty"`               // PromotionalCode:
	SellCity                      string                    `json:"SellCity,omitempty"`                      // SellCity: Overrides the sell city of the requestor.
	TaxExemption                  *TaxExemption             `json:"TaxExemption,omitempty"`                  // TaxExemption:
	TicketCity                    string                    `json:"TicketCity,omitempty"`                    // TicketCity: Overrides the ticket city of the requestor.
	TicketingPCC                  string                    `json:"TicketingPCC,omitempty"`                  // TicketingPCC:
	WaiverCode                    *WaiverCode               `json:"WaiverCode,omitempty"`                    // WaiverCode:
	CurrencyCode                  string                    `json:"currencyCode,omitempty"`                  // CurrencyCode: ISO 4217 currency code
	KeepToBrandInd                bool                      `json:"keepToBrandInd,omitempty"`                // KeepToBrandInd: If true, the offerings returned will be of the same brand as the original Offer
	ReturnMostRestrictiveBrandInd bool                      `json:"returnMostRestrictiveBrandInd,omitempty"` // ReturnMostRestrictiveBrandInd: if true, the most restrictive brand will be returned in the response when there are different brands present in the Offering
}

// pricingModifiersAirDetailPricingPCCPattern is the validation pattern for PricingModifiersAirDetail.PricingPCC
var pricingModifiersAirDetailPricingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// pricingModifiersAirDetailSellCityPattern is the validation pattern for PricingModifiersAirDetail.SellCity
var pricingModifiersAirDetailSellCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// pricingModifiersAirDetailTicketCityPattern is the validation pattern for PricingModifiersAirDetail.TicketCity
var pricingModifiersAirDetailTicketCityPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// pricingModifiersAirDetailTicketingPCCPattern is the validation pattern for PricingModifiersAirDetail.TicketingPCC
var pricingModifiersAirDetailTicketingPCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// pricingModifiersAirDetailCurrencyCodePattern is the validation pattern for PricingModifiersAirDetail.CurrencyCode
var pricingModifiersAirDetailCurrencyCodePattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// PricingModifiersAirDetail is an object.
type PricingModifiersAirDetail struct {
	Type  string `json:"@type"` // Type:
	Brand *struct {
		Type                     string                     `json:"@type,omitempty"`
		AdditionalBrandAttribute []AdditionalBrandAttribute `json:"AdditionalBrandAttribute,omitempty"`
		BrandAttribute           []BrandAttribute           `json:"BrandAttribute,omitempty"`
		BrandRef                 string                     `json:"BrandRef,omitempty"`
		Identifier               Identifier                 `json:"Identifier,omitempty"`
		ImageURL                 []string                   `json:"ImageURL,omitempty"`
		Id                       string                     `json:"id,omitempty"`
		Name                     string                     `json:"name,omitempty"`
		ShelfNumbers             []int32                    `json:"shelfNumbers,omitempty"`
		Tier                     int32                      `json:"tier,omitempty"`
	} `json:"Brand,omitempty"` // Brand:
	FareSelection *struct {
		Type                            string            `json:"@type,omitempty"`
		ChangeOptions                   ChangeOptions     `json:"ChangeOptions,omitempty"`
		FareQualifier                   FareQualifierENUM `json:"FareQualifier,omitempty"`
		RefundOptions                   RefundOptions     `json:"RefundOptions,omitempty"`
		FareType                        FaresFilterEnum   `json:"fareType,omitempty"`
		ProhibitAdvancePurchaseFaresInd bool              `json:"prohibitAdvancePurchaseFaresInd,omitempty"`
		ProhibitMaxStayFaresInd         bool              `json:"prohibitMaxStayFaresInd,omitempty"`
		ProhibitMinStayFaresInd         bool              `json:"prohibitMinStayFaresInd,omitempty"`
		ProhibitRefundableFaresInd      bool              `json:"prohibitRefundableFaresInd,omitempty"`
		ProhibitUnbundledFaresInd       bool              `json:"prohibitUnbundledFaresInd,omitempty"`
		RefundableOnlyInd               bool              `json:"refundableOnlyInd,omitempty"`
		ValidatingCarrier               string            `json:"validatingCarrier,omitempty"`
	} `json:"FareSelection,omitempty"` // FareSelection:
	OrganizationInformation       *OrganizationInformation `json:"OrganizationInformation,omitempty"`       // OrganizationInformation:
	PricingPCC                    string                   `json:"PricingPCC,omitempty"`                    // PricingPCC:
	PromotionalCode               []PromotionalCode        `json:"PromotionalCode,omitempty"`               // PromotionalCode:
	SellCity                      string                   `json:"SellCity,omitempty"`                      // SellCity: Overrides the sell city of the requestor.
	TaxExemption                  *TaxExemption            `json:"TaxExemption,omitempty"`                  // TaxExemption:
	TicketCity                    string                   `json:"TicketCity,omitempty"`                    // TicketCity: Overrides the ticket city of the requestor.
	TicketingPCC                  string                   `json:"TicketingPCC,omitempty"`                  // TicketingPCC:
	CurrencyCode                  string                   `json:"currencyCode,omitempty"`                  // CurrencyCode: Currency Code ISO
	IncludeSplitPaymentInd        bool                     `json:"includeSplitPaymentInd,omitempty"`        // IncludeSplitPaymentInd: If true, split payment (split ticket) offerings\/offers will be returned
	ReturnMostRestrictiveBrandInd bool                     `json:"returnMostRestrictiveBrandInd,omitempty"` // ReturnMostRestrictiveBrandInd: if true, the most restrictive brand will be returned in the response when there are different brands present in the Offering
	SplitPaymentOfferings         float32                  `json:"splitPaymentOfferings,omitempty"`         // SplitPaymentOfferings: The percentage, between 0 and 99, of round trip offerings the user would like returned in the result set.
}

// PrimaryContact is an object.
type PrimaryContact struct {
	Type                         string              `json:"@type"`                                  // Type:
	Email                        *Email              `json:"Email,omitempty"`                        // Email: Electronic email addresses, in IETF specified format.
	Telephone                    *Telephone          `json:"Telephone,omitempty"`                    // Telephone:
	TravelerIdentifier           *TravelerIdentifier `json:"TravelerIdentifier,omitempty"`           // TravelerIdentifier:
	ContactInformationRefusedInd bool                `json:"contactInformationRefusedInd,omitempty"` // ContactInformationRefusedInd: If true, the passenger has refused to provide emergency contact details
	Id                           string              `json:"id,omitempty"`                           // Id:
	ShareWith                    ShareWithEnum       `json:"shareWith,omitempty"`                    // ShareWith: Share with like Supplier,agency etc
	ShareWithSupplier            []string            `json:"shareWithSupplier,omitempty"`            // ShareWithSupplier: Primary contact shared with supplier
}

// PrimaryContactID is an object.
type PrimaryContactID struct {
	Type string `json:"@type"`        // Type:
	Id   string `json:"id,omitempty"` // Id:
}

// Privacy is an object. Confidential details for marketing purpose
type Privacy struct {
	Id             string           `json:"id,omitempty"`             // Id: Optional internally referenced id
	OptInDate      time.Time        `json:"optInDate,omitempty"`      // OptInDate: The datetime of receiving the opt in notice
	OptInStatus    OptInStatusEnum  `json:"optInStatus,omitempty"`    // OptInStatus: Used to indicate marketing preferences, OptIn, OptOut
	OptOutDate     time.Time        `json:"optOutDate,omitempty"`     // OptOutDate: The datetime the opt out notice was received
	OptOutInd      YesNoInheritEnum `json:"optOutInd,omitempty"`      // OptOutInd: Used to indicate marketing preferences, Yes, No, Inherit
	ShareMarketing YesNoInheritEnum `json:"shareMarketing,omitempty"` // ShareMarketing: Used to indicate marketing preferences, Yes, No, Inherit
	ShareSync      YesNoInheritEnum `json:"shareSync,omitempty"`      // ShareSync: Used to indicate marketing preferences, Yes, No, Inherit
}

// Product is an object.
type Product struct {
	Type       string      `json:"@type"`                // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Quantity   int32       `json:"Quantity,omitempty"`   // Quantity: The quantity of the product
	Id         string      `json:"id,omitempty"`         // Id: Local indentifier within a given message for this object.
	ProductRef string      `json:"productRef,omitempty"` // ProductRef: Used to reference another instance of this object in the same message
}

// ProductAir is an object.
type ProductAir struct {
	Type            string            `json:"@type"`                   // Type:
	FlightSegment   []FlightSegment   `json:"FlightSegment"`           // FlightSegment:
	Identifier      *Identifier       `json:"Identifier,omitempty"`    // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	PassengerFlight []PassengerFlight `json:"PassengerFlight"`         // PassengerFlight:
	Quantity        int32             `json:"Quantity,omitempty"`      // Quantity: The quantity of the product
	Id              string            `json:"id,omitempty"`            // Id: Local indentifier within a given message for this object.
	ProductRef      string            `json:"productRef,omitempty"`    // ProductRef: Used to reference another instance of this object in the same message
	TotalDuration   string            `json:"totalDuration,omitempty"` // TotalDuration: Total duration of all Segments that are part of this ProductAir
}

// ProductAncillary is an object.
type ProductAncillary struct {
	Type                 string      `json:"@type"`                          // Type:
	Ancillary            *Ancillary  `json:"Ancillary,omitempty"`            // Ancillary:
	Identifier           *Identifier `json:"Identifier,omitempty"`           // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Quantity             int32       `json:"Quantity,omitempty"`             // Quantity: The quantity of the product
	Id                   string      `json:"id,omitempty"`                   // Id: Local indentifier within a given message for this object.
	ProductRef           string      `json:"productRef,omitempty"`           // ProductRef: Used to reference another instance of this object in the same message
	SelectedByDefaultInd bool        `json:"selectedByDefaultInd,omitempty"` // SelectedByDefaultInd:
}

// ProductAncillaryVehicle is an object.
type ProductAncillaryVehicle struct {
	Type                        string      `json:"@type"`                                 // Type:
	Ancillary                   *Ancillary  `json:"Ancillary,omitempty"`                   // Ancillary:
	EquipmentTypeCode           *Code       `json:"EquipmentTypeCode,omitempty"`           // EquipmentTypeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	Identifier                  *Identifier `json:"Identifier,omitempty"`                  // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Quantity                    int32       `json:"Quantity,omitempty"`                    // Quantity: The quantity of the product
	FreeQuantityIncludedInPrice int32       `json:"freeQuantityIncludedInPrice,omitempty"` // FreeQuantityIncludedInPrice: The mount of this ancillary that is included with the vehicle rental
	Id                          string      `json:"id,omitempty"`                          // Id: Local indentifier within a given message for this object.
	MaxBookableQuantity         int32       `json:"maxBookableQuantity,omitempty"`         // MaxBookableQuantity: The maximum amount of this ancillary that may be booked with the vehicle rental
	ProductRef                  string      `json:"productRef,omitempty"`                  // ProductRef: Used to reference another instance of this object in the same message
	SelectedByDefaultInd        bool        `json:"selectedByDefaultInd,omitempty"`        // SelectedByDefaultInd:
}

// ProductCriteriaAir is an object.
type ProductCriteriaAir struct {
	Type                   string                   `json:"@type,omitempty"`        // Type:
	SpecificFlightCriteria []SpecificFlightCriteria `json:"SpecificFlightCriteria"` // SpecificFlightCriteria:
	Sequence               int32                    `json:"sequence"`               // Sequence: Product criteria air sequence
}

// ProductHospitality is an object.
type ProductHospitality struct {
	Type           string           `json:"@type"`                    // Type:
	DateRange      *DateRange       `json:"DateRange,omitempty"`      // DateRange: Specifies the begin and end date of an event
	Identifier     *Identifier      `json:"Identifier,omitempty"`     // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	PropertyKey    PropertyKey      `json:"PropertyKey"`              // PropertyKey:
	Quantity       int32            `json:"Quantity,omitempty"`       // Quantity: The quantity of the product
	RoomType       *RoomType        `json:"RoomType,omitempty"`       // RoomType:
	AdaCompliant   YesNoUnknownEnum `json:"adaCompliant,omitempty"`   // AdaCompliant: Yes , No , Unknown
	BookingCode    string           `json:"bookingCode,omitempty"`    // BookingCode: Booking code retrieved from the Availability response.
	Guests         int32            `json:"guests,omitempty"`         // Guests: Total number of guests
	Id             string           `json:"id,omitempty"`             // Id: Local indentifier within a given message for this object.
	MoreRatesToken string           `json:"moreRatesToken,omitempty"` // MoreRatesToken: More rates token
	ProductRef     string           `json:"productRef,omitempty"`     // ProductRef: Used to reference another instance of this object in the same message
	PropertyName   string           `json:"propertyName,omitempty"`   // PropertyName: The name of the hotel property
}

// ProductID is an object.
type ProductID struct {
	Type       string      `json:"@type"`                // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Local indentifier within a given message for this object.
	ProductRef string      `json:"productRef,omitempty"` // ProductRef: Used to reference another instance of this object in the same message
}

// ProductIdentifier is an object.
type ProductIdentifier struct {
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Local indentifier within a given message for this object.
	ProductRef string      `json:"productRef,omitempty"` // ProductRef: Used to reference another instance of this object in the same message
}

// ProductInclusionPreference is an object.
type ProductInclusionPreference struct {
	Type                     string                    `json:"@type,omitempty"`                    // Type:
	AdditionalClassification []string                  `json:"AdditionalClassification,omitempty"` // AdditionalClassification:
	Classification           []BrandClassificationEnum `json:"Classification"`                     // Classification:
	BestMatchInd             bool                      `json:"bestMatchInd,omitempty"`             // BestMatchInd: If true, the bestMatch will be returned according to the select product inclusions
	ExactMatchInd            bool                      `json:"exactMatchInd,omitempty"`            // ExactMatchInd: This indicator is deprecated. The default behavior will be to provide an exact match to the product inclusion preferences
	LegSequence              []int32                   `json:"legSequence,omitempty"`              // LegSequence: The legSequence value
}

// ProductOptions is an object.
type ProductOptions struct {
	Type              string      `json:"@type"`                       // Type:
	Identifier        *Identifier `json:"Identifier,omitempty"`        // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Product           []ProductID `json:"Product"`                     // Product:
	ProductOptionsRef string      `json:"ProductOptionsRef,omitempty"` // ProductOptionsRef: Used to reference another instance of this object in the same message
	Id                string      `json:"id,omitempty"`                // Id: Local indentifier within a given message for this object.
	Sequence          int32       `json:"sequence,omitempty"`          // Sequence: NonnegativeInteger
}

// ProductOptionsID is an object.
type ProductOptionsID struct {
	Type              string      `json:"@type"`                       // Type:
	Identifier        *Identifier `json:"Identifier,omitempty"`        // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ProductOptionsRef string      `json:"ProductOptionsRef,omitempty"` // ProductOptionsRef: Used to reference another instance of this object in the same message
	Id                string      `json:"id,omitempty"`                // Id: Local indentifier within a given message for this object.
}

// ProductRateCodeInfo is an object.
type ProductRateCodeInfo struct {
	Type         string       `json:"@type,omitempty"`      // Type:
	ProductRef   string       `json:"ProductRef,omitempty"` // ProductRef: Product reference
	RateCodeInfo RateCodeInfo `json:"RateCodeInfo"`         // RateCodeInfo: Rate Code
}

// ProductVehicle is an object.
type ProductVehicle struct {
	Type       string      `json:"@type"`                // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Quantity   int32       `json:"Quantity,omitempty"`   // Quantity: The quantity of the product
	Vehicle    *Vehicle    `json:"Vehicle,omitempty"`    // Vehicle:
	Id         string      `json:"id,omitempty"`         // Id: Local indentifier within a given message for this object.
	ProductRef string      `json:"productRef,omitempty"` // ProductRef: Used to reference another instance of this object in the same message
}

// ProhibitedVendors is an object.
type ProhibitedVendors struct {
	Vendor []CompanyName `json:"Vendor"` // Vendor:
}

// promotionalCodeSupplierCodePattern is the validation pattern for PromotionalCode.SupplierCode
var promotionalCodeSupplierCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,5})`)

// PromotionalCode is an object.
type PromotionalCode struct {
	SupplierCode string `json:"supplierCode"`    // SupplierCode: Promotional code supplier code
	Value        string `json:"value,omitempty"` // Value:
}

// propertyKeyChainCodePattern is the validation pattern for PropertyKey.ChainCode
var propertyKeyChainCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,5})`)

// PropertyKey is an object.
type PropertyKey struct {
	Type         string `json:"@type,omitempty"` // Type:
	ChainCode    string `json:"chainCode"`       // ChainCode: Chain code for the property.
	PropertyCode string `json:"propertyCode"`    // PropertyCode: Code for the property within the hotel chain.
}

// PropertyRequest is an object.
type PropertyRequest struct {
	Type           string      `json:"@type,omitempty"`          // Type:
	PropertyKey    PropertyKey `json:"PropertyKey"`              // PropertyKey:
	MoreRatesToken string      `json:"moreRatesToken,omitempty"` // MoreRatesToken: More rates token
}

// pseudoCityInfoProviderCodePattern is the validation pattern for PseudoCityInfo.ProviderCode
var pseudoCityInfoProviderCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,5})`)

// pseudoCityInfoValuePattern is the validation pattern for PseudoCityInfo.Value
var pseudoCityInfoValuePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// PseudoCityInfo is an object. a pseudo city information contains the details about the corporate user of a computer reservation system (CRS) or global distribution system (GDS), typically a travel agency.
type PseudoCityInfo struct {
	ProviderCode string `json:"providerCode,omitempty"` // ProviderCode: Assigned Type: c-1100:SupplierCode
	Value        string `json:"value,omitempty"`        // Value:
}

// queueNumberCategoryPattern is the validation pattern for QueueNumber.Category
var queueNumberCategoryPattern = regexp.MustCompile(`([0-9a-zA-Z]+)?`)

// queueNumberOverridePCCPattern is the validation pattern for QueueNumber.OverridePCC
var queueNumberOverridePCCPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// QueueNumber is an object. The queue number
type QueueNumber struct {
	Category    string `json:"category,omitempty"`    // Category: The Queue Category
	OverridePCC string `json:"overridePCC,omitempty"` // OverridePCC: Use PCC to override to queue the Reservation to another PCC
	SubCategory string `json:"subCategory,omitempty"` // SubCategory: Date range subCategory
	Value       int32  `json:"value,omitempty"`       // Value:
}

// railDiscountCardSupplierCodePattern is the validation pattern for RailDiscountCard.SupplierCode
var railDiscountCardSupplierCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,5})`)

// RailDiscountCard is an object. The name of the Rail Discount
type RailDiscountCard struct {
	ReferenceNumber string `json:"referenceNumber,omitempty"` // ReferenceNumber: ReferenceNumber
	SupplierCode    string `json:"supplierCode"`              // SupplierCode: Code of the Supplier
	Value           string `json:"value,omitempty"`           // Value:
}

// rateCandidateChainCodePattern is the validation pattern for RateCandidate.ChainCode
var rateCandidateChainCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,5})`)

// RateCandidate is an object.
type RateCandidate struct {
	Type         string           `json:"@type"`                  // Type:
	ChainCode    string           `json:"chainCode,omitempty"`    // ChainCode: The hotel chain code
	Priority     int32            `json:"priority,omitempty"`     // Priority: rate candidate priority
	PropertyCode string           `json:"propertyCode,omitempty"` // PropertyCode: The hotel chain code
	RateCategory RateCategoryEnum `json:"rateCategory,omitempty"` // RateCategory: Rate Category
	RateCode     string           `json:"rateCode,omitempty"`     // RateCode: The rateCode to be applied to the request
}

// rateCandidateDetailChainCodePattern is the validation pattern for RateCandidateDetail.ChainCode
var rateCandidateDetailChainCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,5})`)

// RateCandidateDetail is an object.
type RateCandidateDetail struct {
	Type            string           `json:"@type"`                     // Type:
	CustomerLoyalty *CustomerLoyalty `json:"CustomerLoyalty,omitempty"` // CustomerLoyalty: Specifies the ID for the membership program.
	ChainCode       string           `json:"chainCode,omitempty"`       // ChainCode: The hotel chain code
	Priority        int32            `json:"priority,omitempty"`        // Priority: rate candidate priority
	PropertyCode    string           `json:"propertyCode,omitempty"`    // PropertyCode: The hotel chain code
	RateCategory    RateCategoryEnum `json:"rateCategory,omitempty"`    // RateCategory: Rate Category
	RateCode        string           `json:"rateCode,omitempty"`        // RateCode: The rateCode to be applied to the request
	RateID          string           `json:"rateID,omitempty"`          // RateID: ID of the rate plan associated with the negotiated rate.
}

// RateCandidates is an object.
type RateCandidates struct {
	Type                string          `json:"@type"`                         // Type:
	RateCandidate       []RateCandidate `json:"RateCandidate"`                 // RateCandidate:
	PostPayRatesOnlyInd bool            `json:"postPayRatesOnlyInd,omitempty"` // PostPayRatesOnlyInd: If true, only postpay rates will be returned
	PrePayRatesOnlyInd  bool            `json:"prePayRatesOnlyInd,omitempty"`  // PrePayRatesOnlyInd: If true, only prepay rates will be returned
}

// RateCandidatesDetail is an object.
type RateCandidatesDetail struct {
	Type                string          `json:"@type"`                         // Type:
	RateCandidate       []RateCandidate `json:"RateCandidate"`                 // RateCandidate:
	NumberOfRatePlans   int32           `json:"numberOfRatePlans,omitempty"`   // NumberOfRatePlans: Minimum number rate plans requested in response
	PostPayRatesOnlyInd bool            `json:"postPayRatesOnlyInd,omitempty"` // PostPayRatesOnlyInd: If true, only postpay rates will be returned
	PrePayRatesOnlyInd  bool            `json:"prePayRatesOnlyInd,omitempty"`  // PrePayRatesOnlyInd: If true, only prepay rates will be returned
}

// RateCodeInfo is an object. Rate Code
type RateCodeInfo struct {
	RateCategory RateCategoryEnum `json:"rateCategory,omitempty"` // RateCategory: Rate Category
	RateID       string           `json:"rateID,omitempty"`       // RateID: Identifier for the rate code
	RateName     string           `json:"rateName,omitempty"`     // RateName: Rate code name
	Value        string           `json:"value,omitempty"`        // Value:
}

// RateDistance is an object. Rate for the period defined by the attributes
type RateDistance struct {
	Allowance               int32              `json:"allowance,omitempty"`               // Allowance: Assigned Type: c-1100:NumberTripleDigit
	DistanceUnits           UnitOfDistanceEnum `json:"distanceUnits,omitempty"`           // DistanceUnits: Miles, Kilometers, etc.
	RatePeriod              RatePeriodEnum     `json:"ratePeriod,omitempty"`              // RatePeriod: The time period for a rate such as daily, weekly, monthly
	RequestedCodeAppliedInd bool               `json:"requestedCodeAppliedInd,omitempty"` // RequestedCodeAppliedInd: Assigned Type: c-1100:OptionalIndicator
	UnlimitedDistanceInd    bool               `json:"unlimitedDistanceInd,omitempty"`    // UnlimitedDistanceInd: Assigned Type: c-1100:OptionalIndicator
	Value                   float32            `json:"value,omitempty"`                   // Value:
}

// Receipt is an object.
type Receipt struct {
	Type       string      `json:"@type"`                // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OfferRef   []string    `json:"OfferRef,omitempty"`   // OfferRef: List of offers which links with the receipt
	ProductRef string      `json:"ProductRef,omitempty"` // ProductRef: Reference of product
	ReceiptRef string      `json:"ReceiptRef,omitempty"` // ReceiptRef:
	DateTime   time.Time   `json:"dateTime,omitempty"`   // DateTime: Receipt date time
	Id         string      `json:"id,omitempty"`         // Id: The verification number.
}

// ReceiptCancellation is an object.
type ReceiptCancellation struct {
	Type         string        `json:"@type"`                  // Type:
	Cancellation *Cancellation `json:"Cancellation,omitempty"` // Cancellation:
	Identifier   *Identifier   `json:"Identifier,omitempty"`   // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OfferRef     []string      `json:"OfferRef,omitempty"`     // OfferRef: List of offers which links with the receipt
	ProductRef   string        `json:"ProductRef,omitempty"`   // ProductRef: Reference of product
	ReceiptRef   string        `json:"ReceiptRef,omitempty"`   // ReceiptRef:
	DateTime     time.Time     `json:"dateTime,omitempty"`     // DateTime: Receipt date time
	Id           string        `json:"id,omitempty"`           // Id: The verification number.
}

// ReceiptConfirmation is an object.
type ReceiptConfirmation struct {
	Type                string        `json:"@type"`                         // Type:
	Confirmation        *Confirmation `json:"Confirmation,omitempty"`        // Confirmation:
	Identifier          *Identifier   `json:"Identifier,omitempty"`          // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OfferRef            []string      `json:"OfferRef,omitempty"`            // OfferRef: List of offers which links with the receipt
	ProductRef          string        `json:"ProductRef,omitempty"`          // ProductRef: Reference of product
	ReceiptRef          string        `json:"ReceiptRef,omitempty"`          // ReceiptRef:
	SegmentSequenceList []int32       `json:"SegmentSequenceList,omitempty"` // SegmentSequenceList: The segmentSequenceList the ReceiptConfirmation applies to
	DateTime            time.Time     `json:"dateTime,omitempty"`            // DateTime: Receipt date time
	Id                  string        `json:"id,omitempty"`                  // Id: The verification number.
}

// ReceiptConfirmationDivide is an object.
type ReceiptConfirmationDivide struct {
	Type                string        `json:"@type"`                         // Type:
	Confirmation        *Confirmation `json:"Confirmation,omitempty"`        // Confirmation:
	Identifier          *Identifier   `json:"Identifier,omitempty"`          // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OfferRef            []string      `json:"OfferRef,omitempty"`            // OfferRef: List of offers which links with the receipt
	ParentLocator       Locator       `json:"ParentLocator"`                 // ParentLocator: Contains the locator (PNR or external locator) or cancellation number for the reservation, order, or offer
	ProductRef          string        `json:"ProductRef,omitempty"`          // ProductRef: Reference of product
	ReceiptRef          string        `json:"ReceiptRef,omitempty"`          // ReceiptRef:
	SegmentSequenceList []int32       `json:"SegmentSequenceList,omitempty"` // SegmentSequenceList: The segmentSequenceList the ReceiptConfirmation applies to
	DateTime            time.Time     `json:"dateTime,omitempty"`            // DateTime: Receipt date time
	Id                  string        `json:"id,omitempty"`                  // Id: The verification number.
}

// ReceiptID is an object.
type ReceiptID struct {
	Type       string      `json:"@type"`                // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ReceiptRef string      `json:"ReceiptRef,omitempty"` // ReceiptRef:
	Id         string      `json:"id,omitempty"`         // Id: The verification number.
}

// ReceiptListResponse is an object.
type ReceiptListResponse struct {
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"` // CurrencyRateConversion:
	Identifier             *Identifier              `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`              // NextSteps:
	ReceiptID              []Receipt                `json:"ReceiptID,omitempty"`              // ReceiptID:
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`          // ReferenceList:
	Result                 *Result                  `json:"Result,omitempty"`                 // Result:
	TraceId                string                   `json:"traceId,omitempty"`                // TraceId: Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TransactionId          string                   `json:"transactionId,omitempty"`          // TransactionId: Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
}

// ReceiptPayment is an object.
type ReceiptPayment struct {
	Type              string             `json:"@type"`                       // Type:
	Document          []Document         `json:"Document,omitempty"`          // Document:
	Identifier        *Identifier        `json:"Identifier,omitempty"`        // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OfferRef          []string           `json:"OfferRef,omitempty"`          // OfferRef: List of offers which links with the receipt
	PaymentIdentifier *PaymentIdentifier `json:"PaymentIdentifier,omitempty"` // PaymentIdentifier:
	ProductRef        string             `json:"ProductRef,omitempty"`        // ProductRef: Reference of product
	ReceiptRef        string             `json:"ReceiptRef,omitempty"`        // ReceiptRef:
	DateTime          time.Time          `json:"dateTime,omitempty"`          // DateTime: Receipt date time
	Id                string             `json:"id,omitempty"`                // Id: The verification number.
}

// ReferenceList is an object.
type ReferenceList struct {
	Type string `json:"@type"`        // Type:
	Id   string `json:"id,omitempty"` // Id: Uniquely identifies for the Reference List
}

// ReferenceListBrand is an object.
type ReferenceListBrand struct {
	Type  string  `json:"@type"`           // Type:
	Brand []Brand `json:"Brand,omitempty"` // Brand:
	Id    string  `json:"id,omitempty"`    // Id: Uniquely identifies for the Reference List
}

// ReferenceListExchangedPrice is an object.
type ReferenceListExchangedPrice struct {
	Type  string  `json:"@type"`        // Type:
	Price []Price `json:"Price"`        // Price:
	Id    string  `json:"id,omitempty"` // Id: Uniquely identifies for the Reference List
}

// ReferenceListFlight is an object.
type ReferenceListFlight struct {
	Type   string   `json:"@type"`        // Type:
	Flight []Flight `json:"Flight"`       // Flight:
	Id     string   `json:"id,omitempty"` // Id: Uniquely identifies for the Reference List
}

// RefundMethodEnum is an object.
type RefundMethodEnum struct {
	Value RefundMethodEnumBase `json:"value,omitempty"` // Value:
}

// RefundOptions is an object.
type RefundOptions struct {
	Type               string              `json:"@type,omitempty"`              // Type:
	RefundPenaltyRange *RefundPenaltyRange `json:"RefundPenaltyRange,omitempty"` // RefundPenaltyRange:
	RefundTypes        []RefundTypeEnum    `json:"refundTypes"`                  // RefundTypes:
}

// RefundPenaltyRange is an object.
type RefundPenaltyRange struct {
	Type    string         `json:"@type,omitempty"`   // Type:
	Maximum *AmountPercent `json:"Maximum,omitempty"` // Maximum:
	Minimum *AmountPercent `json:"Minimum,omitempty"` // Minimum:
}

// RefundabilityEnum is an object.
type RefundabilityEnum struct {
	Value RefundabilityEnumBase `json:"value,omitempty"` // Value:
}

// RentalPickupReturn is an object.
type RentalPickupReturn struct {
	Type           string          `json:"@type,omitempty"`          // Type:
	VendorLocation *VendorLocation `json:"VendorLocation,omitempty"` // VendorLocation: The vendor's location number for pickup or return
	Date           string          `json:"date"`                     // Date: The local date of the pickup or return
	Time           string          `json:"time,omitempty"`           // Time: The local time of the pickup or return
}

// Reservation is an object.
type Reservation struct {
	Type                       string                       `json:"@type"`                                // Type:
	AgencyServiceFee           []AgencyServiceFee           `json:"AgencyServiceFee,omitempty"`           // AgencyServiceFee:
	FormOfPayment              []FormOfPaymentID            `json:"FormOfPayment,omitempty"`              // FormOfPayment:
	GroupName                  string                       `json:"GroupName,omitempty"`                  // GroupName: A name assigned to a Reservation containing an offer with Passengerflight/Flight Quantity equal to or greater than 10
	Identifier                 *Identifier                  `json:"Identifier,omitempty"`                 // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Offer                      []Offer                      `json:"Offer,omitempty"`                      // Offer:
	OfferLink                  []OfferLink                  `json:"OfferLink,omitempty"`                  // OfferLink:
	OrganizationLoyaltyProgram []OrganizationLoyaltyProgram `json:"OrganizationLoyaltyProgram,omitempty"` // OrganizationLoyaltyProgram:
	Payment                    []Payment                    `json:"Payment,omitempty"`                    // Payment:
	Preference                 *struct {
		Type               string             `json:"@type,omitempty"`
		AppliesTo          AppliesTo          `json:"AppliesTo,omitempty"`
		TravelerIdentifier TravelerIdentifier `json:"TravelerIdentifier,omitempty"`
		Id                 string             `json:"id,omitempty"`
	} `json:"Preference,omitempty"` // Preference:
	PrimaryContact             []PrimaryContact            `json:"PrimaryContact,omitempty"`             // PrimaryContact:
	Receipt                    []Receipt                   `json:"Receipt,omitempty"`                    // Receipt:
	ReservationComment         []ReservationComment        `json:"ReservationComment,omitempty"`         // ReservationComment:
	ReservationDisplaySequence *ReservationDisplaySequence `json:"ReservationDisplaySequence,omitempty"` // ReservationDisplaySequence:
	ShoppingCart               *ShoppingCart               `json:"ShoppingCart,omitempty"`               // ShoppingCart:
	SpecialService             []SpecialService            `json:"SpecialService,omitempty"`             // SpecialService:
	TravelAgency               *struct {
		Type                  string               `json:"@type,omitempty"`
		CorporateCode         string               `json:"CorporateCode,omitempty"`
		Identifier            Identifier           `json:"Identifier,omitempty"`
		OrganizationName      CompanyName          `json:"OrganizationName,omitempty"`
		ProfileName           []string             `json:"ProfileName,omitempty"`
		TravelOrganizationRef string               `json:"TravelOrganizationRef,omitempty"`
		Id                    string               `json:"id,omitempty"`
		OrganizationType      OrganizationTypeEnum `json:"organizationType,omitempty"`
	} `json:"TravelAgency,omitempty"` // TravelAgency:
	Traveler         []Traveler        `json:"Traveler,omitempty"`         // Traveler:
	TravelerProduct  []TravelerProduct `json:"TravelerProduct,omitempty"`  // TravelerProduct:
	AutoDeleteDate   string            `json:"autoDeleteDate,omitempty"`   // AutoDeleteDate: The auto delete date represents the date that the Reservation will be kept active. Also known as retention segment or retention date.
	Id               string            `json:"id,omitempty"`               // Id: Internal ID
	NotificationDate string            `json:"notificationDate,omitempty"` // NotificationDate: The notification date represents the date that the Reservation should be reviewed. Also known as ticket time limit date.
}

// ReservationBuild is an object.
type ReservationBuild struct {
	Type                string                 `json:"@type"`                         // Type:
	Accounting          *AccountingID          `json:"Accounting,omitempty"`          // Accounting:
	DocumentOverrides   []DocumentOverridesID  `json:"DocumentOverrides,omitempty"`   // DocumentOverrides:
	FormOfPayment       []FormOfPaymentID      `json:"FormOfPayment,omitempty"`       // FormOfPayment:
	Payment             []PaymentID            `json:"Payment,omitempty"`             // Payment:
	Preference          []PreferenceID         `json:"Preference,omitempty"`          // Preference:
	PrimaryContact      []PrimaryContactID     `json:"PrimaryContact,omitempty"`      // PrimaryContact:
	ReceiptConfirmation *ReceiptConfirmation   `json:"ReceiptConfirmation,omitempty"` // ReceiptConfirmation:
	ReservationComment  []ReservationCommentID `json:"ReservationComment,omitempty"`  // ReservationComment:
	SpecialService      []SpecialServiceID     `json:"SpecialService,omitempty"`      // SpecialService:
	TravelAgency        *TravelAgency          `json:"TravelAgency,omitempty"`        // TravelAgency:
	Traveler            []TravelerID           `json:"Traveler"`                      // Traveler:
}

// ReservationBuildFromCatalogOffering is an object.
type ReservationBuildFromCatalogOffering struct {
	Type                                string                               `json:"@type"`                                         // Type:
	Accounting                          *AccountingID                        `json:"Accounting,omitempty"`                          // Accounting:
	BuildFromCatalogOfferingHospitality *BuildFromCatalogOfferingHospitality `json:"BuildFromCatalogOfferingHospitality,omitempty"` // BuildFromCatalogOfferingHospitality:
	DocumentOverrides                   []DocumentOverridesID                `json:"DocumentOverrides,omitempty"`                   // DocumentOverrides:
	FormOfPayment                       []FormOfPaymentID                    `json:"FormOfPayment,omitempty"`                       // FormOfPayment:
	Payment                             []PaymentID                          `json:"Payment,omitempty"`                             // Payment:
	Preference                          []PreferenceID                       `json:"Preference,omitempty"`                          // Preference:
	PrimaryContact                      []PrimaryContactID                   `json:"PrimaryContact,omitempty"`                      // PrimaryContact:
	ReceiptConfirmation                 *struct {
		Type                string       `json:"@type,omitempty"`
		Confirmation        Confirmation `json:"Confirmation,omitempty"`
		Identifier          Identifier   `json:"Identifier,omitempty"`
		OfferRef            []string     `json:"OfferRef,omitempty"`
		ProductRef          string       `json:"ProductRef,omitempty"`
		ReceiptRef          string       `json:"ReceiptRef,omitempty"`
		SegmentSequenceList []int32      `json:"SegmentSequenceList,omitempty"`
		DateTime            time.Time    `json:"dateTime,omitempty"`
		Id                  string       `json:"id,omitempty"`
	} `json:"ReceiptConfirmation,omitempty"` // ReceiptConfirmation:
	ReservationComment []ReservationCommentID `json:"ReservationComment,omitempty"` // ReservationComment:
	SpecialService     []SpecialServiceID     `json:"SpecialService,omitempty"`     // SpecialService:
	TravelAgency       *struct {
		Type                  string               `json:"@type,omitempty"`
		CorporateCode         string               `json:"CorporateCode,omitempty"`
		Identifier            Identifier           `json:"Identifier,omitempty"`
		OrganizationName      CompanyName          `json:"OrganizationName,omitempty"`
		ProfileName           []string             `json:"ProfileName,omitempty"`
		TravelOrganizationRef string               `json:"TravelOrganizationRef,omitempty"`
		Id                    string               `json:"id,omitempty"`
		OrganizationType      OrganizationTypeEnum `json:"organizationType,omitempty"`
	} `json:"TravelAgency,omitempty"` // TravelAgency:
	Traveler []TravelerID `json:"Traveler"` // Traveler:
}

// ReservationBuildFromCatalogOfferings is an object.
type ReservationBuildFromCatalogOfferings struct {
	Type                             string                           `json:"@type"`                            // Type:
	Accounting                       *AccountingID                    `json:"Accounting,omitempty"`             // Accounting:
	BuildFromCatalogOfferingsRequest BuildFromCatalogOfferingsRequest `json:"BuildFromCatalogOfferingsRequest"` // BuildFromCatalogOfferingsRequest:
	DocumentOverrides                []DocumentOverridesID            `json:"DocumentOverrides,omitempty"`      // DocumentOverrides:
	FormOfPayment                    []FormOfPaymentID                `json:"FormOfPayment,omitempty"`          // FormOfPayment:
	Payment                          []PaymentID                      `json:"Payment,omitempty"`                // Payment:
	Preference                       []PreferenceID                   `json:"Preference,omitempty"`             // Preference:
	PrimaryContact                   []PrimaryContactID               `json:"PrimaryContact,omitempty"`         // PrimaryContact:
	ReceiptConfirmation              *ReceiptID                       `json:"ReceiptConfirmation,omitempty"`    // ReceiptConfirmation:
	ReservationComment               []ReservationCommentID           `json:"ReservationComment,omitempty"`     // ReservationComment:
	SpecialService                   []SpecialServiceID               `json:"SpecialService,omitempty"`         // SpecialService:
	TravelAgency                     *struct {
		Type                  string               `json:"@type,omitempty"`
		CorporateCode         string               `json:"CorporateCode,omitempty"`
		Identifier            Identifier           `json:"Identifier,omitempty"`
		OrganizationName      CompanyName          `json:"OrganizationName,omitempty"`
		ProfileName           []string             `json:"ProfileName,omitempty"`
		TravelOrganizationRef string               `json:"TravelOrganizationRef,omitempty"`
		Id                    string               `json:"id,omitempty"`
		OrganizationType      OrganizationTypeEnum `json:"organizationType,omitempty"`
	} `json:"TravelAgency,omitempty"` // TravelAgency:
	Traveler []TravelerID `json:"Traveler"` // Traveler:
}

// ReservationBuildFromCatalogOfferingsAir is an object.
type ReservationBuildFromCatalogOfferingsAir struct {
	Type                             string                           `json:"@type"`                            // Type:
	Accounting                       *AccountingID                    `json:"Accounting,omitempty"`             // Accounting:
	BuildFromCatalogOfferingsRequest BuildFromCatalogOfferingsRequest `json:"BuildFromCatalogOfferingsRequest"` // BuildFromCatalogOfferingsRequest:
	DocumentOverrides                []DocumentOverridesID            `json:"DocumentOverrides,omitempty"`      // DocumentOverrides:
	FormOfPayment                    []FormOfPaymentID                `json:"FormOfPayment,omitempty"`          // FormOfPayment:
	Payment                          []PaymentID                      `json:"Payment,omitempty"`                // Payment:
	Preference                       []PreferenceID                   `json:"Preference,omitempty"`             // Preference:
	PricingModifiersAir              PricingModifiersAir              `json:"PricingModifiersAir"`              // PricingModifiersAir:
	PrimaryContact                   []PrimaryContactID               `json:"PrimaryContact,omitempty"`         // PrimaryContact:
	ReceiptConfirmation              *ReceiptID                       `json:"ReceiptConfirmation,omitempty"`    // ReceiptConfirmation:
	ReservationComment               []ReservationCommentID           `json:"ReservationComment,omitempty"`     // ReservationComment:
	SpecialService                   []SpecialServiceID               `json:"SpecialService,omitempty"`         // SpecialService:
	TravelAgency                     *struct {
		Type                  string               `json:"@type,omitempty"`
		CorporateCode         string               `json:"CorporateCode,omitempty"`
		Identifier            Identifier           `json:"Identifier,omitempty"`
		OrganizationName      CompanyName          `json:"OrganizationName,omitempty"`
		ProfileName           []string             `json:"ProfileName,omitempty"`
		TravelOrganizationRef string               `json:"TravelOrganizationRef,omitempty"`
		Id                    string               `json:"id,omitempty"`
		OrganizationType      OrganizationTypeEnum `json:"organizationType,omitempty"`
	} `json:"TravelAgency,omitempty"` // TravelAgency:
	Traveler []TravelerID `json:"Traveler"` // Traveler:
}

// ReservationBuildFromCatalogProductOfferings is an object.
type ReservationBuildFromCatalogProductOfferings struct {
	Type                                    string                                   `json:"@type"`                                             // Type:
	Accounting                              *AccountingID                            `json:"Accounting,omitempty"`                              // Accounting:
	BuildFromCatalogProductOfferingsRequest *BuildFromCatalogProductOfferingsRequest `json:"BuildFromCatalogProductOfferingsRequest,omitempty"` // BuildFromCatalogProductOfferingsRequest:
	DocumentOverrides                       []DocumentOverridesID                    `json:"DocumentOverrides,omitempty"`                       // DocumentOverrides:
	FormOfPayment                           []FormOfPaymentID                        `json:"FormOfPayment,omitempty"`                           // FormOfPayment:
	Payment                                 []PaymentID                              `json:"Payment,omitempty"`                                 // Payment:
	Preference                              []PreferenceID                           `json:"Preference,omitempty"`                              // Preference:
	PrimaryContact                          []PrimaryContactID                       `json:"PrimaryContact,omitempty"`                          // PrimaryContact:
	ReceiptConfirmation                     *ReceiptID                               `json:"ReceiptConfirmation,omitempty"`                     // ReceiptConfirmation:
	ReservationComment                      []ReservationCommentID                   `json:"ReservationComment,omitempty"`                      // ReservationComment:
	SpecialService                          []SpecialServiceID                       `json:"SpecialService,omitempty"`                          // SpecialService:
	TravelAgency                            *struct {
		Type                  string               `json:"@type,omitempty"`
		CorporateCode         string               `json:"CorporateCode,omitempty"`
		Identifier            Identifier           `json:"Identifier,omitempty"`
		OrganizationName      CompanyName          `json:"OrganizationName,omitempty"`
		ProfileName           []string             `json:"ProfileName,omitempty"`
		TravelOrganizationRef string               `json:"TravelOrganizationRef,omitempty"`
		Id                    string               `json:"id,omitempty"`
		OrganizationType      OrganizationTypeEnum `json:"organizationType,omitempty"`
	} `json:"TravelAgency,omitempty"` // TravelAgency:
	Traveler []TravelerID `json:"Traveler"` // Traveler:
}

// ReservationBuildFromProducts is an object.
type ReservationBuildFromProducts struct {
	Type                string                 `json:"@type"`                         // Type:
	Accounting          *AccountingID          `json:"Accounting,omitempty"`          // Accounting:
	DocumentOverrides   []DocumentOverridesID  `json:"DocumentOverrides,omitempty"`   // DocumentOverrides:
	FormOfPayment       []FormOfPaymentID      `json:"FormOfPayment,omitempty"`       // FormOfPayment:
	Payment             []PaymentID            `json:"Payment,omitempty"`             // Payment:
	Preference          []PreferenceID         `json:"Preference,omitempty"`          // Preference:
	PricingModifiersAir *PricingModifiersAir   `json:"PricingModifiersAir,omitempty"` // PricingModifiersAir:
	PrimaryContact      []PrimaryContactID     `json:"PrimaryContact,omitempty"`      // PrimaryContact:
	ProductCriteriaAir  []ProductCriteriaAir   `json:"ProductCriteriaAir"`            // ProductCriteriaAir:
	ReceiptConfirmation *ReceiptID             `json:"ReceiptConfirmation,omitempty"` // ReceiptConfirmation:
	ReservationComment  []ReservationCommentID `json:"ReservationComment,omitempty"`  // ReservationComment:
	SpecialService      []SpecialServiceID     `json:"SpecialService,omitempty"`      // SpecialService:
	TravelAgency        *struct {
		Type                  string               `json:"@type,omitempty"`
		CorporateCode         string               `json:"CorporateCode,omitempty"`
		Identifier            Identifier           `json:"Identifier,omitempty"`
		OrganizationName      CompanyName          `json:"OrganizationName,omitempty"`
		ProfileName           []string             `json:"ProfileName,omitempty"`
		TravelOrganizationRef string               `json:"TravelOrganizationRef,omitempty"`
		Id                    string               `json:"id,omitempty"`
		OrganizationType      OrganizationTypeEnum `json:"organizationType,omitempty"`
	} `json:"TravelAgency,omitempty"` // TravelAgency:
	Traveler []TravelerID `json:"Traveler"` // Traveler:
}

// ReservationBuildVehicle is an object.
type ReservationBuildVehicle struct {
	Type                string                 `json:"@type"`                         // Type:
	Accounting          *AccountingID          `json:"Accounting,omitempty"`          // Accounting:
	ArrivalDetails      *ArrivalDetails        `json:"ArrivalDetails,omitempty"`      // ArrivalDetails:
	DocumentOverrides   []DocumentOverridesID  `json:"DocumentOverrides,omitempty"`   // DocumentOverrides:
	FormOfPayment       []FormOfPaymentID      `json:"FormOfPayment,omitempty"`       // FormOfPayment:
	Payment             []PaymentID            `json:"Payment,omitempty"`             // Payment:
	Preference          []PreferenceID         `json:"Preference,omitempty"`          // Preference:
	PrimaryContact      []PrimaryContactID     `json:"PrimaryContact,omitempty"`      // PrimaryContact:
	ReceiptConfirmation *ReceiptID             `json:"ReceiptConfirmation,omitempty"` // ReceiptConfirmation:
	ReservationComment  []ReservationCommentID `json:"ReservationComment,omitempty"`  // ReservationComment:
	SpecialService      []SpecialServiceID     `json:"SpecialService,omitempty"`      // SpecialService:
	TravelAgency        *struct {
		Type                  string               `json:"@type,omitempty"`
		CorporateCode         string               `json:"CorporateCode,omitempty"`
		Identifier            Identifier           `json:"Identifier,omitempty"`
		OrganizationName      CompanyName          `json:"OrganizationName,omitempty"`
		ProfileName           []string             `json:"ProfileName,omitempty"`
		TravelOrganizationRef string               `json:"TravelOrganizationRef,omitempty"`
		Id                    string               `json:"id,omitempty"`
		OrganizationType      OrganizationTypeEnum `json:"organizationType,omitempty"`
	} `json:"TravelAgency,omitempty"` // TravelAgency:
	Traveler              []TravelerID           `json:"Traveler"`                        // Traveler:
	VehicleTravelCriteria *VehicleTravelCriteria `json:"VehicleTravelCriteria,omitempty"` // VehicleTravelCriteria:
}

// ReservationComment is an object.
type ReservationComment struct {
	Type              string            `json:"@type"`                       // Type:
	AppliesTo         []AppliesTo       `json:"AppliesTo,omitempty"`         // AppliesTo:
	Comment           []Comment         `json:"Comment,omitempty"`           // Comment:
	CommentSource     CommentSourceEnum `json:"commentSource,omitempty"`     // CommentSource:
	Id                string            `json:"id,omitempty"`                // Id: Local indentifier within a given message for this object.
	ShareWith         ShareWithEnum     `json:"shareWith,omitempty"`         // ShareWith: Share with like Supplier,agency etc
	ShareWithSupplier []string          `json:"shareWithSupplier,omitempty"` // ShareWithSupplier: Reservation comment shared with supplier
}

// ReservationCommentID is an object.
type ReservationCommentID struct {
	Type string `json:"@type"`        // Type:
	Id   string `json:"id,omitempty"` // Id: Local indentifier within a given message for this object.
}

// ReservationDetail is an object.
type ReservationDetail struct {
	Type       string `json:"@type"` // Type:
	Accounting *struct {
		Type          string          `json:"@type,omitempty"`
		AccountingRef string          `json:"AccountingRef,omitempty"`
		Identifier    Identifier      `json:"Identifier,omitempty"`
		NameValuePair []NameValuePair `json:"NameValuePair,omitempty"`
		DataType      string          `json:"dataType,omitempty"`
		Id            string          `json:"id,omitempty"`
		Template      string          `json:"template,omitempty"`
	} `json:"Accounting,omitempty"` // Accounting:
	AgencyServiceFee           []AgencyServiceFee           `json:"AgencyServiceFee,omitempty"`           // AgencyServiceFee:
	DocumentOverrides          []DocumentOverrides          `json:"DocumentOverrides,omitempty"`          // DocumentOverrides:
	FormOfPayment              []FormOfPaymentID            `json:"FormOfPayment,omitempty"`              // FormOfPayment:
	GroupName                  string                       `json:"GroupName,omitempty"`                  // GroupName: A name assigned to a Reservation containing an offer with Passengerflight/Flight Quantity equal to or greater than 10
	Identifier                 *Identifier                  `json:"Identifier,omitempty"`                 // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Offer                      []Offer                      `json:"Offer,omitempty"`                      // Offer:
	OfferLink                  []OfferLink                  `json:"OfferLink,omitempty"`                  // OfferLink:
	OrganizationLoyaltyProgram []OrganizationLoyaltyProgram `json:"OrganizationLoyaltyProgram,omitempty"` // OrganizationLoyaltyProgram:
	Payment                    []Payment                    `json:"Payment,omitempty"`                    // Payment:
	Preference                 *struct {
		Type               string             `json:"@type,omitempty"`
		AppliesTo          AppliesTo          `json:"AppliesTo,omitempty"`
		TravelerIdentifier TravelerIdentifier `json:"TravelerIdentifier,omitempty"`
		Id                 string             `json:"id,omitempty"`
	} `json:"Preference,omitempty"` // Preference:
	PrimaryContact             []PrimaryContact            `json:"PrimaryContact,omitempty"`             // PrimaryContact:
	Receipt                    []Receipt                   `json:"Receipt,omitempty"`                    // Receipt:
	ReservationComment         []ReservationComment        `json:"ReservationComment,omitempty"`         // ReservationComment:
	ReservationDisplaySequence *ReservationDisplaySequence `json:"ReservationDisplaySequence,omitempty"` // ReservationDisplaySequence:
	ShoppingCart               *ShoppingCart               `json:"ShoppingCart,omitempty"`               // ShoppingCart:
	SpecialService             []SpecialService            `json:"SpecialService,omitempty"`             // SpecialService:
	TravelAgency               *struct {
		Type                  string               `json:"@type,omitempty"`
		CorporateCode         string               `json:"CorporateCode,omitempty"`
		Identifier            Identifier           `json:"Identifier,omitempty"`
		OrganizationName      CompanyName          `json:"OrganizationName,omitempty"`
		ProfileName           []string             `json:"ProfileName,omitempty"`
		TravelOrganizationRef string               `json:"TravelOrganizationRef,omitempty"`
		Id                    string               `json:"id,omitempty"`
		OrganizationType      OrganizationTypeEnum `json:"organizationType,omitempty"`
	} `json:"TravelAgency,omitempty"` // TravelAgency:
	Traveler         []Traveler        `json:"Traveler,omitempty"`         // Traveler:
	TravelerProduct  []TravelerProduct `json:"TravelerProduct,omitempty"`  // TravelerProduct:
	AutoDeleteDate   string            `json:"autoDeleteDate,omitempty"`   // AutoDeleteDate: The auto delete date represents the date that the Reservation will be kept active. Also known as retention segment or retention date.
	Id               string            `json:"id,omitempty"`               // Id: Internal ID
	NotificationDate string            `json:"notificationDate,omitempty"` // NotificationDate: The notification date represents the date that the Reservation should be reviewed. Also known as ticket time limit date.
}

// ReservationDisplaySequence is an object.
type ReservationDisplaySequence struct {
	Type                   string            `json:"@type"`                            // Type:
	DisplaySequence        []DisplaySequence `json:"DisplaySequence,omitempty"`        // DisplaySequence:
	AutoDeleteDateSequence int32             `json:"autoDeleteDateSequence,omitempty"` // AutoDeleteDateSequence: The sequence of the autoDeleteDate (retention segment) within the Reservation
}

// ReservationIdentifier is an object.
type ReservationIdentifier struct {
	Type       string      `json:"@type,omitempty"`      // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Internal ID
}

// ReservationQueryBuild is an object.
type ReservationQueryBuild struct {
	Type             string           `json:"@type"`            // Type:
	ReservationBuild ReservationBuild `json:"ReservationBuild"` // ReservationBuild:
}

// reservationQuerySearchCriteriaReservationArrivalPattern is the validation pattern for ReservationQuerySearchCriteriaReservation.Arrival
var reservationQuerySearchCriteriaReservationArrivalPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// reservationQuerySearchCriteriaReservationDeparturePattern is the validation pattern for ReservationQuerySearchCriteriaReservation.Departure
var reservationQuerySearchCriteriaReservationDeparturePattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// ReservationResponse is an object.
type ReservationResponse struct {
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"` // CurrencyRateConversion:
	Identifier             *Identifier              `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`              // NextSteps:
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`          // ReferenceList:
	Reservation            *Reservation             `json:"Reservation,omitempty"`            // Reservation:
	Result                 *Result                  `json:"Result,omitempty"`                 // Result:
	TraceId                string                   `json:"traceId,omitempty"`                // TraceId: Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TransactionId          string                   `json:"transactionId,omitempty"`          // TransactionId: Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
}

// Restriction is an object. Restriction
type Restriction struct {
	ProductRef          []string `json:"productRef,omitempty"`          // ProductRef: The productRef which the restriction applies to
	SegmentSequenceList []int32  `json:"segmentSequenceList,omitempty"` // SegmentSequenceList: Segment sequence list
	Value               string   `json:"value,omitempty"`               // Value:
}

// Restrictions is an object.
type Restrictions struct {
	Type                  string                  `json:"@type,omitempty"`                 // Type:
	DocumentType          DocumentTypeEnum        `json:"DocumentType,omitempty"`          // DocumentType: Document type like EMD, MCO
	Restriction           []string                `json:"Restriction"`                     // Restriction:
	TravelerIdentifierRef []TravelerIdentifierRef `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
}

// Result is an object.
type Result struct {
	Type    string           `json:"@type,omitempty"`   // Type:
	Error   []Error          `json:"Error,omitempty"`   // Error:
	Warning []Warning        `json:"Warning,omitempty"` // Warning:
	Status  ResultStatusEnum `json:"status,omitempty"`  // Status: The status of an error or warning
}

// roomAmenityCodePattern is the validation pattern for RoomAmenity.Code
var roomAmenityCodePattern = regexp.MustCompile(`[0-9A-Z]{1,3}(\\.[A-Z]{3}(\\.X){0,1}){0,1}`)

// RoomAmenity is an object.
type RoomAmenity struct {
	Type         string   `json:"@type,omitempty"`        // Type:
	Inclusion    []string `json:"Inclusion,omitempty"`    // Inclusion:
	Name         string   `json:"Name,omitempty"`         // Name: Room Amenity Name
	Code         string   `json:"code,omitempty"`         // Code: OTA code used to describe the room amenity. This is optional in the Properties Search request but mandatory in the response
	Description  string   `json:"description,omitempty"`  // Description: description of the room amenity
	IncludedInd  bool     `json:"includedInd,omitempty"`  // IncludedInd: To represent if the Amenity is included in the rate
	Quantity     int32    `json:"quantity,omitempty"`     // Quantity: quantity of amenity
	SurchargeInd bool     `json:"surchargeInd,omitempty"` // SurchargeInd: To represent if the Amenity attracts a surcharge
}

// roomCharacteristicsCategoryPattern is the validation pattern for RoomCharacteristics.Category
var roomCharacteristicsCategoryPattern = regexp.MustCompile(`[0-9A-Z]{1,3}(\.[A-Z]{3}(\.X){0,1}){0,1}`)

// roomCharacteristicsViewCodePattern is the validation pattern for RoomCharacteristics.ViewCode
var roomCharacteristicsViewCodePattern = regexp.MustCompile(`[0-9A-Z]{1,3}(\.[A-Z]{3}(\.X){0,1}){0,1}`)

// RoomCharacteristics is an object.
type RoomCharacteristics struct {
	Type             string             `json:"@type,omitempty"`            // Type:
	BedConfiguration []BedConfiguration `json:"BedConfiguration,omitempty"` // BedConfiguration:
	Category         string             `json:"category,omitempty"`         // Category: Category of the room.
	NonSmokingInd    bool               `json:"nonSmokingInd,omitempty"`    // NonSmokingInd:
	SmokingAllowed   YesNoUnknownEnum   `json:"smokingAllowed,omitempty"`   // SmokingAllowed: Yes , No , Unknown
	TypeCode         string             `json:"typeCode,omitempty"`         // TypeCode: Type code
	ViewCode         string             `json:"viewCode,omitempty"`         // ViewCode: Free text describing the view.
	WifiIncluded     YesNoUnknownEnum   `json:"wifiIncluded,omitempty"`     // WifiIncluded: Yes , No , Unknown
}

// RoomOccupancy is an object.
type RoomOccupancy struct {
	Type          string          `json:"@type,omitempty"`         // Type:
	AgeQualifying []AgeQualifying `json:"AgeQualifying,omitempty"` // AgeQualifying:
	MaxOccupancy  int32           `json:"maxOccupancy,omitempty"`  // MaxOccupancy: The maximum number of room occupants.
	MinOccupancy  int32           `json:"minOccupancy,omitempty"`  // MinOccupancy: The minimum occupancy
}

// RoomStayCandidate is an object.
type RoomStayCandidate struct {
	GuestCounts GuestCounts   `json:"GuestCounts"`           // GuestCounts:
	RoomAmenity []RoomAmenity `json:"RoomAmenity,omitempty"` // RoomAmenity:
}

// RoomStayCandidates is an object.
type RoomStayCandidates struct {
	RoomStayCandidate []RoomStayCandidate `json:"RoomStayCandidate"` // RoomStayCandidate:
}

// RoomType is an object.
type RoomType struct {
	Type                string                   `json:"@type"`                         // Type:
	Description         *TextTitleAndDescription `json:"Description,omitempty"`         // Description: Descriptive text
	RoomAmenity         []RoomAmenity            `json:"RoomAmenity,omitempty"`         // RoomAmenity:
	RoomCharacteristics *RoomCharacteristics     `json:"RoomCharacteristics,omitempty"` // RoomCharacteristics:
}

// RoomTypeDetail is an object.
type RoomTypeDetail struct {
	Type                string                   `json:"@type"`                         // Type:
	AdditionalDetails   *AdditionalDetails       `json:"AdditionalDetails,omitempty"`   // AdditionalDetails:
	Description         *TextTitleAndDescription `json:"Description,omitempty"`         // Description: Descriptive text
	RoomAmenity         []RoomAmenity            `json:"RoomAmenity,omitempty"`         // RoomAmenity:
	RoomCharacteristics *RoomCharacteristics     `json:"RoomCharacteristics,omitempty"` // RoomCharacteristics:
	RoomOccupancy       []RoomOccupancy          `json:"RoomOccupancy,omitempty"`       // RoomOccupancy:
	AlternateInd        bool                     `json:"alternateInd,omitempty"`        // AlternateInd: Indicates the room is an alternate room type to the requested room type when true.
	ConvertedInd        bool                     `json:"convertedInd,omitempty"`        // ConvertedInd: Indicates the room is converted when true.
	NumberOfUnits       int32                    `json:"numberOfUnits,omitempty"`       // NumberOfUnits: The number of rooms that have been combined to create this room type.
	ReqdGuaranteeType   string                   `json:"reqdGuaranteeType,omitempty"`   // ReqdGuaranteeType: TODO-(Should this be Guarantee?)Denotes the form of guarantee for this room.
	RoomInd             bool                     `json:"roomInd,omitempty"`             // RoomInd: Indicates the room is a sleeping room when true.
}

// SearchControlConsoleChannelID is an object.
type SearchControlConsoleChannelID struct {
	SccType string `json:"sccType,omitempty"` // SccType: Assigned Type: c-1100:StringTiny
	Value   string `json:"value,omitempty"`   // Value:
}

// SearchCriteriaFlight is an object.
type SearchCriteriaFlight struct {
	DepartureTimeRange *TimeRange `json:"DepartureTimeRange,omitempty"` // DepartureTimeRange: Specify time.
	From               FromTo     `json:"From"`                         // From: Location code
	To                 FromTo     `json:"To"`                           // To: Location code
	ArrivalDate        string     `json:"arrivalDate,omitempty"`        // ArrivalDate: Preferred local arrival date. Cannot be used in conjunction with departure date.
	ArrivalTime        time.Time  `json:"arrivalTime,omitempty"`        // ArrivalTime: Preferred local arrival time. Cannot be used in conjunction with departure time.
	DepartureDate      string     `json:"departureDate"`                // DepartureDate: Preferred local departure date. Cannot be used in conjunction with arrival date
	DepartureTime      string     `json:"departureTime,omitempty"`      // DepartureTime: Preferred local departure time. Cannot be used in conjunction with arrival time
	LegSequence        int32      `json:"legSequence,omitempty"`        // LegSequence: Leg sequence
}

// SearchModifiersAir is an object.
type SearchModifiersAir struct {
	Type                       string                       `json:"@type,omitempty"`                      // Type:
	CabinPreference            []CabinPreference            `json:"CabinPreference,omitempty"`            // CabinPreference:
	CarrierPreference          []CarrierPreference          `json:"CarrierPreference,omitempty"`          // CarrierPreference:
	ClassOfServicePreference   []ClassOfServicePreference   `json:"ClassOfServicePreference,omitempty"`   // ClassOfServicePreference:
	ConnectionPreferences      []ConnectionPreferencesAir   `json:"ConnectionPreferences,omitempty"`      // ConnectionPreferences:
	ProductInclusionPreference []ProductInclusionPreference `json:"ProductInclusionPreference,omitempty"` // ProductInclusionPreference:
	ExcludeGround              ExcludeGroundTypeEnum        `json:"excludeGround,omitempty"`              // ExcludeGround:
	ProhibitChangeOfAirportInd bool                         `json:"prohibitChangeOfAirportInd,omitempty"` // ProhibitChangeOfAirportInd: If present and true, connections that require a change of airports are not returned
}

// SearchVehicleAttributes is an object. The physical vehical attrirbutes that a person can search on.
type SearchVehicleAttributes struct {
	AirConditioningInd   bool   `json:"AirConditioningInd,omitempty"`   // AirConditioningInd: True if requesting air conditioning
	DoorCount            string `json:"DoorCount,omitempty"`            // DoorCount: Number of doors
	FuelTypeCode         *Code  `json:"FuelTypeCode,omitempty"`         // FuelTypeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	TransmissionTypeCode *Code  `json:"TransmissionTypeCode,omitempty"` // TransmissionTypeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	VehicleCategoryCode  *Code  `json:"VehicleCategoryCode,omitempty"`  // VehicleCategoryCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	VehicleSizeCode      *Code  `json:"VehicleSizeCode,omitempty"`      // VehicleSizeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
}

// Seat is an object.
type Seat struct {
	Type           string       `json:"@type,omitempty"`          // Type:
	Characteristic []string     `json:"Characteristic,omitempty"` // Characteristic:
	SeatFeature    SpaceFeature `json:"SeatFeature"`              // SeatFeature: Discriptive information about the seat.
	Location       string       `json:"location,omitempty"`       // Location: The seat location
	SeatType       string       `json:"seatType,omitempty"`       // SeatType: The type of seat
}

// SeatAssignment is an object.
type SeatAssignment struct {
	Type           string         `json:"@type,omitempty"`       // Type:
	Characteristic []string       `json:"Characteristic"`        // Characteristic:
	Seat           string         `json:"Seat"`                  // Seat: Seat
	SeatFeature    []SpaceFeature `json:"SeatFeature,omitempty"` // SeatFeature:
}

// seriesCodeMaskPattern is the validation pattern for SeriesCode.Mask
var seriesCodeMaskPattern = regexp.MustCompile(`[0-9a-zA-Z]{1,32}`)

// seriesCodeTokenPattern is the validation pattern for SeriesCode.Token
var seriesCodeTokenPattern = regexp.MustCompile(`[0-9a-zA-Z]{1,32}`)

// SeriesCode is an object.
type SeriesCode struct {
	Type                 string                      `json:"@type,omitempty"`                // Type:
	ErrorWarning         *ErrorWarning               `json:"ErrorWarning,omitempty"`         // ErrorWarning:
	PlainText            string                      `json:"PlainText,omitempty"`            // PlainText: Don't use this unless it is REALLY ok to not use encryption. Non-secure (plain text) value.
	AuthenticationMethod EncryptionTokenTypeAuthEnum `json:"authenticationMethod,omitempty"` // AuthenticationMethod: Type of Authentication
	EncryptedValue       string                      `json:"encryptedValue,omitempty"`       // EncryptedValue: Encrypted value
	EncryptionKey        string                      `json:"encryptionKey,omitempty"`        // EncryptionKey: Note: This contains a key required to retrieve the full payment instrument details compliant with PCI DSS standards.
	EncryptionKeyMethod  string                      `json:"encryptionKeyMethod,omitempty"`  // EncryptionKeyMethod: Developer: This contains a reference to the key generation method being used - this is NOT the key value.
	EncryptionMethod     string                      `json:"encryptionMethod,omitempty"`     // EncryptionMethod: OpenTravel Best Practice: Encryption Method: When using the OpenTravel Encryption element, it is RECOMMENDED that all trading partners be informed of all encryption methods being used in advance of implementation to ensure message processing compatibility.
	Mask                 string                      `json:"mask,omitempty"`                 // Mask: Masked Value
	Token                string                      `json:"token,omitempty"`                // Token: Token value
	TokenProviderID      string                      `json:"tokenProviderID,omitempty"`      // TokenProviderID: Developer: This contains a provider ID if multiple providers are used for secure information exchange.
}

// ShoppingCart is an object.
type ShoppingCart struct {
	Type    string       `json:"@type,omitempty"`   // Type:
	Product []ProductAir `json:"Product,omitempty"` // Product:
}

// ShoppingCartProductStatus is an object.
type ShoppingCartProductStatus struct {
	Type string `json:"@type"` // Type:
}

// ShoppingCartProductStatusAir is an object.
type ShoppingCartProductStatusAir struct {
	Type      string      `json:"@type"`               // Type:
	StatusAir []StatusAir `json:"StatusAir,omitempty"` // StatusAir:
}

// SignatureOnFile is an object.
type SignatureOnFile struct {
	Type                string               `json:"@type,omitempty"`                // Type:
	DateEffectiveExpire *DateEffectiveExpire `json:"Date_EffectiveExpire,omitempty"` // DateEffectiveExpire: Used to identify the effective date and\/or expiration date.
	SignatureOnFileInd  bool                 `json:"signatureOnFileInd,omitempty"`   // SignatureOnFileInd: When true, indicates a signature has been obtained.
}

// SpaceFeature is an object. Discriptive information about the seat.
type SpaceFeature struct {
	Context     string `json:"context,omitempty"`     // Context: The source of the code
	Description string `json:"description,omitempty"` // Description: The description of the space feature
	Power       string `json:"power,omitempty"`       // Power: The type of power provided, if any
	Rating      string `json:"rating,omitempty"`      // Rating: The seat guru rating of the seat
	SeatType    string `json:"seatType,omitempty"`    // SeatType: The type of object that occupies the space
	Value       string `json:"value,omitempty"`       // Value:
	Video       string `json:"video,omitempty"`       // Video: The type of video provided, if any
}

// SpecialEquipment is an object. A list of special equipment codes from OTA Equipment Type Code
type SpecialEquipment struct {
	CodeContext       string `json:"codeContext,omitempty"`       // CodeContext:
	EquipmentTypeCode string `json:"equipmentTypeCode,omitempty"` // EquipmentTypeCode:
	Quantity          int32  `json:"quantity,omitempty"`          // Quantity:
	Value             string `json:"value,omitempty"`             // Value:
}

// SpecialService is an object.
type SpecialService struct {
	Type               string              `json:"@type"`                        // Type:
	AppliesTo          *AppliesTo          `json:"AppliesTo,omitempty"`          // AppliesTo:
	Identifier         *Identifier         `json:"Identifier,omitempty"`         // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ServiceAnimalType  string              `json:"ServiceAnimalType,omitempty"`  // ServiceAnimalType: The type of service animal accompanying the Traveler. If no service animal leave blank.
	Status             *Status             `json:"Status,omitempty"`             // Status:
	TravelerIdentifier *TravelerIdentifier `json:"TravelerIdentifier,omitempty"` // TravelerIdentifier:
	Id                 string              `json:"id,omitempty"`                 // Id: Internal Id
}

// SpecialServiceBassinet is an object.
type SpecialServiceBassinet struct {
	Type                string              `json:"@type"`                         // Type:
	AppliesTo           *AppliesTo          `json:"AppliesTo,omitempty"`           // AppliesTo:
	Identifier          *Identifier         `json:"Identifier,omitempty"`          // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ServiceAnimalType   string              `json:"ServiceAnimalType,omitempty"`   // ServiceAnimalType: The type of service animal accompanying the Traveler. If no service animal leave blank.
	Status              *Status             `json:"Status,omitempty"`              // Status:
	TravelerIdentifier  *TravelerIdentifier `json:"TravelerIdentifier,omitempty"`  // TravelerIdentifier:
	BassinetRequiredInd bool                `json:"bassinetRequiredInd,omitempty"` // BassinetRequiredInd: If true, a bassinet is required
	Id                  string              `json:"id,omitempty"`                  // Id: Internal Id
}

// SpecialServiceBlind is an object.
type SpecialServiceBlind struct {
	Type                        string              `json:"@type"`                                 // Type:
	AppliesTo                   *AppliesTo          `json:"AppliesTo,omitempty"`                   // AppliesTo:
	Identifier                  *Identifier         `json:"Identifier,omitempty"`                  // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ServiceAnimalType           string              `json:"ServiceAnimalType,omitempty"`           // ServiceAnimalType: The type of service animal accompanying the Traveler. If no service animal leave blank.
	Status                      *Status             `json:"Status,omitempty"`                      // Status:
	TravelerIdentifier          *TravelerIdentifier `json:"TravelerIdentifier,omitempty"`          // TravelerIdentifier:
	BlindAssistanceRequestedInd bool                `json:"blindAssistanceRequestedInd,omitempty"` // BlindAssistanceRequestedInd: If true the Traveler is Blind and requests assistance
	Id                          string              `json:"id,omitempty"`                          // Id: Internal Id
}

// SpecialServiceDPNA is an object.
type SpecialServiceDPNA struct {
	Type               string              `json:"@type"`                        // Type:
	AppliesTo          *AppliesTo          `json:"AppliesTo,omitempty"`          // AppliesTo:
	Description        string              `json:"Description"`                  // Description: Describes the Traveler development or intellectual disability
	Identifier         *Identifier         `json:"Identifier,omitempty"`         // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ServiceAnimalType  string              `json:"ServiceAnimalType,omitempty"`  // ServiceAnimalType: The type of service animal accompanying the Traveler. If no service animal leave blank.
	Status             *Status             `json:"Status,omitempty"`             // Status:
	TravelerIdentifier *TravelerIdentifier `json:"TravelerIdentifier,omitempty"` // TravelerIdentifier:
	Id                 string              `json:"id,omitempty"`                 // Id: Internal Id
}

// SpecialServiceDeaf is an object.
type SpecialServiceDeaf struct {
	Type                       string              `json:"@type"`                                // Type:
	AppliesTo                  *AppliesTo          `json:"AppliesTo,omitempty"`                  // AppliesTo:
	Identifier                 *Identifier         `json:"Identifier,omitempty"`                 // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ServiceAnimalType          string              `json:"ServiceAnimalType,omitempty"`          // ServiceAnimalType: The type of service animal accompanying the Traveler. If no service animal leave blank.
	Status                     *Status             `json:"Status,omitempty"`                     // Status:
	TravelerIdentifier         *TravelerIdentifier `json:"TravelerIdentifier,omitempty"`         // TravelerIdentifier:
	DeafAssistanceRequestedInd bool                `json:"deafAssistanceRequestedInd,omitempty"` // DeafAssistanceRequestedInd: If True the Traveler is Deaf and requests assistance
	Id                         string              `json:"id,omitempty"`                         // Id: Internal Id
}

// SpecialServiceID is an object.
type SpecialServiceID struct {
	Type       string      `json:"@type"`                // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Internal Id
}

// SpecialServiceMeal is an object.
type SpecialServiceMeal struct {
	Type                string              `json:"@type"`                         // Type:
	AppliesTo           *AppliesTo          `json:"AppliesTo,omitempty"`           // AppliesTo:
	Identifier          *Identifier         `json:"Identifier,omitempty"`          // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ServiceAnimalType   string              `json:"ServiceAnimalType,omitempty"`   // ServiceAnimalType: The type of service animal accompanying the Traveler. If no service animal leave blank.
	SpecialMealTypeEnum SpecialMealTypeEnum `json:"SpecialMealTypeEnum,omitempty"` // SpecialMealTypeEnum: Special Meal Type
	Status              *Status             `json:"Status,omitempty"`              // Status:
	TravelerIdentifier  *TravelerIdentifier `json:"TravelerIdentifier,omitempty"`  // TravelerIdentifier:
	Id                  string              `json:"id,omitempty"`                  // Id: Internal Id
}

// SpecialServiceUnaccompaniedMinor is an object.
type SpecialServiceUnaccompaniedMinor struct {
	Type                  string              `json:"@type"`                           // Type:
	AppliesTo             *AppliesTo          `json:"AppliesTo,omitempty"`             // AppliesTo:
	Identifier            *Identifier         `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ServiceAnimalType     string              `json:"ServiceAnimalType,omitempty"`     // ServiceAnimalType: The type of service animal accompanying the Traveler. If no service animal leave blank.
	Status                *Status             `json:"Status,omitempty"`                // Status:
	TravelerIdentifier    *TravelerIdentifier `json:"TravelerIdentifier,omitempty"`    // TravelerIdentifier:
	Id                    string              `json:"id,omitempty"`                    // Id: Internal Id
	UnaccompaniedMinorInd bool                `json:"unaccompaniedMinorInd,omitempty"` // UnaccompaniedMinorInd: Indicates that the Traveler is an Unaccompanied minor
}

// SpecialServiceWheelchairAirlineSupplied is an object.
type SpecialServiceWheelchairAirlineSupplied struct {
	Type                  string              `json:"@type"`                           // Type:
	AppliesTo             *AppliesTo          `json:"AppliesTo,omitempty"`             // AppliesTo:
	Identifier            *Identifier         `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ServiceAnimalType     string              `json:"ServiceAnimalType,omitempty"`     // ServiceAnimalType: The type of service animal accompanying the Traveler. If no service animal leave blank.
	Status                *Status             `json:"Status,omitempty"`                // Status:
	TravelerIdentifier    *TravelerIdentifier `json:"TravelerIdentifier,omitempty"`    // TravelerIdentifier:
	CannotAscendStairsInd bool                `json:"cannotAscendStairsInd,omitempty"` // CannotAscendStairsInd: if true, traveler needs assistance ascending and descending stairs
	Id                    string              `json:"id,omitempty"`                    // Id: Internal Id
	TravelerImmobileInd   bool                `json:"travelerImmobileInd,omitempty"`   // TravelerImmobileInd: if true, traveler is completely immobile and requires assistance to\/from aircraft\/mobile lounge and ascending\/descending stairs
}

// SpecialServiceWheelchairTravelerSupplied is an object.
type SpecialServiceWheelchairTravelerSupplied struct {
	Type                  string              `json:"@type"`                           // Type:
	AppliesTo             *AppliesTo          `json:"AppliesTo,omitempty"`             // AppliesTo:
	BatteryTypeEnum       BatteryTypeEnum     `json:"BatteryTypeEnum,omitempty"`       // BatteryTypeEnum: The type of battery that is used in the device
	Identifier            *Identifier         `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Measurement           []Measurement       `json:"Measurement,omitempty"`           // Measurement:
	ServiceAnimalType     string              `json:"ServiceAnimalType,omitempty"`     // ServiceAnimalType: The type of service animal accompanying the Traveler. If no service animal leave blank.
	Status                *Status             `json:"Status,omitempty"`                // Status:
	TravelerIdentifier    *TravelerIdentifier `json:"TravelerIdentifier,omitempty"`    // TravelerIdentifier:
	CannotAscendStairsInd bool                `json:"cannotAscendStairsInd,omitempty"` // CannotAscendStairsInd: if true, traveler needs assistance ascending and descending stairs
	Id                    string              `json:"id,omitempty"`                    // Id: Internal Id
	TravelerImmobileInd   bool                `json:"travelerImmobileInd,omitempty"`   // TravelerImmobileInd: if true, traveler is completely immobile and requires assistance to\/from aircraft\/mobile lounge and ascending\/descending stairs
}

// specificFlightCriteriaArrivalTimePattern is the validation pattern for SpecificFlightCriteria.ArrivalTime
var specificFlightCriteriaArrivalTimePattern = regexp.MustCompile(`(([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)((:?)[0-5]\d)?([\.,]\d+(?!:))?`)

// specificFlightCriteriaCarrierPattern is the validation pattern for SpecificFlightCriteria.Carrier
var specificFlightCriteriaCarrierPattern = regexp.MustCompile(`([a-zA-Z0-9]{2,3})`)

// specificFlightCriteriaClassOfServicePattern is the validation pattern for SpecificFlightCriteria.ClassOfService
var specificFlightCriteriaClassOfServicePattern = regexp.MustCompile(`([a-zA-Z0-9]{1,2})`)

// specificFlightCriteriaDepartureTimePattern is the validation pattern for SpecificFlightCriteria.DepartureTime
var specificFlightCriteriaDepartureTimePattern = regexp.MustCompile(`(([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)((:?)[0-5]\d)?([\.,]\d+(?!:))?`)

// specificFlightCriteriaFlightNumberPattern is the validation pattern for SpecificFlightCriteria.FlightNumber
var specificFlightCriteriaFlightNumberPattern = regexp.MustCompile(`[0-9]{1,4}[A-Z]?|OPEN|ARNK`)

// specificFlightCriteriaFromPattern is the validation pattern for SpecificFlightCriteria.From
var specificFlightCriteriaFromPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// specificFlightCriteriaToPattern is the validation pattern for SpecificFlightCriteria.To
var specificFlightCriteriaToPattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// SpecificFlightCriteria is an object.
type SpecificFlightCriteria struct {
	Type                   string                     `json:"@type,omitempty"`                  // Type:
	AvailabilitySourceCode AvailabilitySourceCodeENUM `json:"AvailabilitySourceCode,omitempty"` // AvailabilitySourceCode: A code representing the source of the flight availability
	ContentSource          ContentSourceEnum          `json:"ContentSource,omitempty"`          // ContentSource: The source of the content to be returned in CatalogOfferings
	ArrivalDate            string                     `json:"arrivalDate,omitempty"`            // ArrivalDate: Arrival date
	ArrivalTime            string                     `json:"arrivalTime,omitempty"`            // ArrivalTime: The arrival time in local timezone
	BoundFlightsInd        bool                       `json:"boundFlightsInd,omitempty"`        // BoundFlightsInd: if true indicates that the flight availability was polled (availability check) using connecting segments. If false, flights were polled as point to point segments.
	BrandTier              int32                      `json:"brandTier,omitempty"`              // BrandTier: Brand tier
	Cabin                  CabinAirEnum               `json:"cabin,omitempty"`                  // Cabin: Specifies the cabin type (e.g. first, business, economy).
	Carrier                string                     `json:"carrier"`                          // Carrier: Carrier
	ClassOfService         string                     `json:"classOfService,omitempty"`         // ClassOfService: The class of service
	DepartureDate          string                     `json:"departureDate"`                    // DepartureDate: date of departure
	DepartureTime          string                     `json:"departureTime,omitempty"`          // DepartureTime: The departure time in local timezone
	FlightNumber           string                     `json:"flightNumber"`                     // FlightNumber: Flight Number
	From                   string                     `json:"from"`                             // From: From Airport Code IATA
	SegmentSequence        int32                      `json:"segmentSequence"`                  // SegmentSequence: Segment sequence
	To                     string                     `json:"to"`                               // To: to Airpor Code IATA
}

// StateProv is an object. The standard code or abbreviation for the state, province, or region with optional name
type StateProv struct {
	Name  string `json:"name,omitempty"`  // Name: State,province or region name or code needed to identify location
	Value string `json:"value,omitempty"` // Value:
}

// Status is an object.
type Status struct {
	SupplierText string                 `json:"supplierText,omitempty"` // SupplierText: Supplier status text
	Value        ConfirmationStatusEnum `json:"value,omitempty"`        // Value: Status returned in a response for a two or more phase commitment process
}

// StatusAir is an object.
type StatusAir struct {
	Code        string          `json:"code,omitempty"`        // Code: Status code
	FlightRefs  []string        `json:"flightRefs,omitempty"`  // FlightRefs: The flightRefs the status is applicable to within the Offer
	PastDateInd bool            `json:"pastDateInd,omitempty"` // PastDateInd: If true, the flight is considered to be past date
	Value       OfferStatusEnum `json:"value,omitempty"`       // Value: Offer Status like confirmed ,Pending etc
}

// StatusAncillary is an object.
type StatusAncillary struct {
	Code  string          `json:"code,omitempty"`  // Code: Assigned Type: c-1100:StringTiny
	Value OfferStatusEnum `json:"value,omitempty"` // Value: Offer Status like confirmed ,Pending etc
}

// SubPolicy is an object.
type SubPolicy struct {
	Type          string          `json:"@type,omitempty"`         // Type:
	TextFormatted []TextFormatted `json:"TextFormatted,omitempty"` // TextFormatted:
	Id            string          `json:"id,omitempty"`            // Id:
	Title         string          `json:"title,omitempty"`         // Title:
}

// supplierLocatorSupplierCodePattern is the validation pattern for SupplierLocator.SupplierCode
var supplierLocatorSupplierCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,5})`)

// supplierLocatorValuePattern is the validation pattern for SupplierLocator.Value
var supplierLocatorValuePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// SupplierLocator is an object. The supplier and the supplier's locator code for a product
type SupplierLocator struct {
	SupplierCode string `json:"supplierCode,omitempty"` // SupplierCode: Supplier Code
	SupplierName string `json:"supplierName,omitempty"` // SupplierName: Name of the supplier
	Value        string `json:"value,omitempty"`        // Value:
}

// SupplierRate is an object.
type SupplierRate struct {
	Type                 string          `json:"@type,omitempty"`                // Type:
	BaseRate             *CurrencyAmount `json:"BaseRate,omitempty"`             // BaseRate: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	DropOffCharge        *CurrencyAmount `json:"DropOffCharge,omitempty"`        // DropOffCharge: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	EstimatedTotalAmount *CurrencyAmount `json:"EstimatedTotalAmount,omitempty"` // EstimatedTotalAmount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	ExtraMileageCharge   *CurrencyAmount `json:"ExtraMileageCharge,omitempty"`   // ExtraMileageCharge: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	RateForPeriod        *CurrencyAmount `json:"RateForPeriod,omitempty"`        // RateForPeriod: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
}

// surchargeCurrencyCodePattern is the validation pattern for Surcharge.CurrencyCode
var surchargeCurrencyCodePattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// surchargeDecimalPlacePattern is the validation pattern for Surcharge.DecimalPlace
var surchargeDecimalPlacePattern = regexp.MustCompile(`([0-4])`)

// Surcharge is an object. The fee amount with feecode and reporting informtion
type Surcharge struct {
	CodeAuthority        string          `json:"codeAuthority,omitempty"`        // CodeAuthority: Surcharge code authority
	CurrencyCode         string          `json:"currencyCode,omitempty"`         // CurrencyCode: Sur charge currency code
	DecimalAuthority     string          `json:"decimalAuthority,omitempty"`     // DecimalAuthority: Currency code decimal authority
	DecimalPlace         int32           `json:"decimalPlace,omitempty"`         // DecimalPlace: Decimal place for the currency unit
	Description          string          `json:"description,omitempty"`          // Description: Description
	Purpose              string          `json:"purpose,omitempty"`              // Purpose: Sur charge purpose
	ReportingAuthority   string          `json:"reportingAuthority,omitempty"`   // ReportingAuthority: Sur charge reporting authority
	SurchargeApplication ApplicationEnum `json:"surchargeApplication,omitempty"` // SurchargeApplication: Application values like perperson , peroom
	SurchargeCode        string          `json:"surchargeCode,omitempty"`        // SurchargeCode: Sur charge code
	SurchargeFrequency   FrequencyEnum   `json:"surchargeFrequency,omitempty"`   // SurchargeFrequency: Stay frequency like PerNight, PerDay
	Value                float32         `json:"value,omitempty"`                // Value:
}

// Surcharges is an object.
type Surcharges struct {
	Type            string  `json:"@type"`                     // Type:
	TotalSurcharges float32 `json:"TotalSurcharges,omitempty"` // TotalSurcharges: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	ApproximateInd  bool    `json:"approximateInd,omitempty"`  // ApproximateInd: if true, the surcharge amounts are approximate
}

// SurchargesDetail is an object.
type SurchargesDetail struct {
	Type            string      `json:"@type"`                     // Type:
	Surcharge       []Surcharge `json:"Surcharge"`                 // Surcharge:
	TotalSurcharges float32     `json:"TotalSurcharges,omitempty"` // TotalSurcharges: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	ApproximateInd  bool        `json:"approximateInd,omitempty"`  // ApproximateInd: if true, the surcharge amounts are approximate
}

// taxCurrencyCodePattern is the validation pattern for Tax.CurrencyCode
var taxCurrencyCodePattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// taxDecimalPlacePattern is the validation pattern for Tax.DecimalPlace
var taxDecimalPlacePattern = regexp.MustCompile(`([0-4])`)

// Tax is an object.
type Tax struct {
	CodeAuthority      string           `json:"codeAuthority,omitempty"`      // CodeAuthority: Code Authority
	CurrencyCode       string           `json:"currencyCode,omitempty"`       // CurrencyCode: Currency code of the city.
	DecimalAuthority   string           `json:"decimalAuthority,omitempty"`   // DecimalAuthority: Decimal Authority
	DecimalPlace       int32            `json:"decimalPlace,omitempty"`       // DecimalPlace: Allowed number of decimals.
	Description        string           `json:"description,omitempty"`        // Description: additional information
	ExemptInd          bool             `json:"exemptInd,omitempty"`          // ExemptInd: If true, this tax is exempt
	IncludedInBase     YesNoUnknownEnum `json:"includedInBase,omitempty"`     // IncludedInBase: Yes , No , Unknown
	Purpose            string           `json:"purpose,omitempty"`            // Purpose: purpose
	ReportingAuthority string           `json:"reportingAuthority,omitempty"` // ReportingAuthority: Identifies the reporting authority such as airport code for XF taxes.
	TaxCode            string           `json:"taxCode,omitempty"`            // TaxCode: Tax code of the city
	Value              float32          `json:"value,omitempty"`              // Value:
}

// taxBreakdownAirportCodePattern is the validation pattern for TaxBreakdown.AirportCode
var taxBreakdownAirportCodePattern = regexp.MustCompile(`([a-zA-Z]{3})`)

// TaxBreakdown is an object.
type TaxBreakdown struct {
	AirportCode  string        `json:"AirportCode"`            // AirportCode: The airport location the tax applies to
	Amount       float32       `json:"Amount,omitempty"`       // Amount: The amount of the tax applied
	CurrencyCode *CurrencyCode `json:"CurrencyCode,omitempty"` // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
}

// TaxExemption is an object.
type TaxExemption struct {
	Type              string   `json:"@type,omitempty"`             // Type:
	AllTaxesExemptInd bool     `json:"allTaxesExemptInd,omitempty"` // AllTaxesExemptInd: If true, the Offer/Offering is exempt from all taxes
	Countries         []string `json:"countries,omitempty"`         // Countries: ISO country code
	TaxCodes          []string `json:"taxCodes,omitempty"`          // TaxCodes: Tax codes
}

// TaxInfo is an object.
type TaxInfo struct {
	Amount       float32        `json:"Amount"`                 // Amount: The amount of the tax applied
	CurrencyCode *CurrencyCode  `json:"CurrencyCode,omitempty"` // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	TaxBreakdown []TaxBreakdown `json:"TaxBreakdown"`           // TaxBreakdown: The breakdown of the tax for this tax code
	TaxCode      string         `json:"TaxCode"`                // TaxCode: The tax code
}

// TaxPercent is an object.
type TaxPercent struct {
	Description        string           `json:"description,omitempty"`        // Description: Description
	IncludedInBase     YesNoUnknownEnum `json:"includedInBase,omitempty"`     // IncludedInBase: Yes , No , Unknown
	Purpose            string           `json:"purpose,omitempty"`            // Purpose: Purpose of tax
	ReportingAuthority string           `json:"reportingAuthority,omitempty"` // ReportingAuthority: Tax reporting authority
	TaxCode            string           `json:"taxCode,omitempty"`            // TaxCode: Tax code
	Value              float32          `json:"value,omitempty"`              // Value:
}

// Taxes is an object.
type Taxes struct {
	Type       string    `json:"@type"`                // Type:
	TaxInfo    []TaxInfo `json:"TaxInfo,omitempty"`    // TaxInfo:
	TotalTaxes float32   `json:"TotalTaxes,omitempty"` // TotalTaxes: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
}

// TaxesDetail is an object.
type TaxesDetail struct {
	Type       string      `json:"@type"`                // Type:
	Tax        []Tax       `json:"Tax,omitempty"`        // Tax:
	TaxInfo    []TaxInfo   `json:"TaxInfo,omitempty"`    // TaxInfo:
	TaxPercent *TaxPercent `json:"TaxPercent,omitempty"` // TaxPercent:
	TotalTaxes float32     `json:"TotalTaxes,omitempty"` // TotalTaxes: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
}

// telephoneAreaCityCodePattern is the validation pattern for Telephone.AreaCityCode
var telephoneAreaCityCodePattern = regexp.MustCompile(`[0-9]{1,8}`)

// telephoneCityCodePattern is the validation pattern for Telephone.CityCode
var telephoneCityCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// telephoneCountryAccessCodePattern is the validation pattern for Telephone.CountryAccessCode
var telephoneCountryAccessCodePattern = regexp.MustCompile(`[0-9]{1,3}`)

// telephoneExtensionPattern is the validation pattern for Telephone.Extension
var telephoneExtensionPattern = regexp.MustCompile(`[0-9]{0,5}`)

// Telephone is an object.
type Telephone struct {
	Type              string            `json:"@type"`                       // Type:
	AreaCityCode      string            `json:"areaCityCode,omitempty"`      // AreaCityCode: Telephone Area CityCode
	CityCode          string            `json:"cityCode,omitempty"`          // CityCode: City Code
	CountryAccessCode string            `json:"countryAccessCode,omitempty"` // CountryAccessCode: TelephoneCountry AccessCode
	Extension         string            `json:"extension,omitempty"`         // Extension: Telephone extension number
	Id                string            `json:"id,omitempty"`                // Id: UOptional internally referenced id
	PhoneNumber       string            `json:"phoneNumber"`                 // PhoneNumber: Mobile/Telephone Number
	Role              EnumTelephoneRole `json:"role,omitempty"`              // Role:
}

// telephoneDetailAreaCityCodePattern is the validation pattern for TelephoneDetail.AreaCityCode
var telephoneDetailAreaCityCodePattern = regexp.MustCompile(`[0-9]{1,8}`)

// telephoneDetailCityCodePattern is the validation pattern for TelephoneDetail.CityCode
var telephoneDetailCityCodePattern = regexp.MustCompile(`([A-Z0-9]+)?`)

// telephoneDetailCountryAccessCodePattern is the validation pattern for TelephoneDetail.CountryAccessCode
var telephoneDetailCountryAccessCodePattern = regexp.MustCompile(`[0-9]{1,3}`)

// telephoneDetailExtensionPattern is the validation pattern for TelephoneDetail.Extension
var telephoneDetailExtensionPattern = regexp.MustCompile(`[0-9]{0,5}`)

// telephoneDetailPhoneLocationTypePattern is the validation pattern for TelephoneDetail.PhoneLocationType
var telephoneDetailPhoneLocationTypePattern = regexp.MustCompile(`[0-9A-Z]{1,3}(\.[A-Z]{3}(\.X){0,1}){0,1}`)

// telephoneDetailPhoneTechTypePattern is the validation pattern for TelephoneDetail.PhoneTechType
var telephoneDetailPhoneTechTypePattern = regexp.MustCompile(`[0-9A-Z]{1,3}(\.[A-Z]{3}(\.X){0,1}){0,1}`)

// telephoneDetailPhoneUseTypePattern is the validation pattern for TelephoneDetail.PhoneUseType
var telephoneDetailPhoneUseTypePattern = regexp.MustCompile(`[0-9A-Z]{1,3}(\.[A-Z]{3}(\.X){0,1}){0,1}`)

// TelephoneDetail is an object.
type TelephoneDetail struct {
	Type              string            `json:"@type"`                        // Type:
	Comment           *Comment          `json:"Comment,omitempty"`            // Comment: Textual information.
	EnumTelephoneRole EnumTelephoneRole `json:"Enum_TelephoneRole,omitempty"` // EnumTelephoneRole:
	Privacy           *Privacy          `json:"Privacy,omitempty"`            // Privacy: Confidential details for marketing purpose
	AreaCityCode      string            `json:"areaCityCode,omitempty"`       // AreaCityCode: Telephone Area CityCode
	CityCode          string            `json:"cityCode,omitempty"`           // CityCode: City Code
	CountryAccessCode string            `json:"countryAccessCode,omitempty"`  // CountryAccessCode: TelephoneCountry AccessCode
	DefaultInd        bool              `json:"defaultInd,omitempty"`         // DefaultInd: When true, indicates a default value should be used.
	Extension         string            `json:"extension,omitempty"`          // Extension: Telephone extension number
	Id                string            `json:"id,omitempty"`                 // Id: UOptional internally referenced id
	PhoneLocationType string            `json:"phoneLocationType,omitempty"`  // PhoneLocationType: Location of the phone
	PhoneNumber       string            `json:"phoneNumber"`                  // PhoneNumber: Mobile/Telephone Number
	PhoneTechType     string            `json:"phoneTechType,omitempty"`      // PhoneTechType: Indicates the type of technology associated with the telephone number
	PhoneUseType      string            `json:"phoneUseType,omitempty"`       // PhoneUseType: Use of the phone
	Pin               string            `json:"pin,omitempty"`                // Pin: Additional codes used for telephone
	Priority          int32             `json:"priority,omitempty"`           // Priority: Priority
	ProvisionedInd    bool              `json:"provisionedInd,omitempty"`     // ProvisionedInd: true indicates this phone number was created through provisioned
	Role              EnumTelephoneRole `json:"role,omitempty"`               // Role:
}

// TermsAndConditions is an object.
type TermsAndConditions struct {
	Type                  string            `json:"@type"`                           // Type:
	CustomerLoyalty       []CustomerLoyalty `json:"CustomerLoyalty,omitempty"`       // CustomerLoyalty:
	ExpiryDate            time.Time         `json:"ExpiryDate,omitempty"`            // ExpiryDate: The data and time the offer will expire
	Identifier            *Identifier       `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id                    string            `json:"id,omitempty"`                    // Id: Local indentifier within a given message for this object.
	TermsAndConditionsRef string            `json:"termsAndConditionsRef,omitempty"` // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
}

// TermsAndConditionsAir is an object.
type TermsAndConditionsAir struct {
	Type                                 string                   `json:"@type"`                                          // Type:
	BaggageAllowance                     []BaggageAllowance       `json:"BaggageAllowance,omitempty"`                     // BaggageAllowance:
	BaggageRecheck                       []BaggageRecheck         `json:"BaggageRecheck,omitempty"`                       // BaggageRecheck:
	CustomerLoyalty                      []CustomerLoyalty        `json:"CustomerLoyalty,omitempty"`                      // CustomerLoyalty:
	DocumentValidDateRange               *DocumentValidDateRange  `json:"DocumentValidDateRange,omitempty"`               // DocumentValidDateRange:
	ExpiryDate                           time.Time                `json:"ExpiryDate,omitempty"`                           // ExpiryDate: The data and time the offer will expire
	FareGuaranteePolicy                  []FareGuaranteePolicy    `json:"FareGuaranteePolicy,omitempty"`                  // FareGuaranteePolicy:
	FareRuleIdentifierRef                *IdentifierRef           `json:"FareRuleIdentifierRef,omitempty"`                // FareRuleIdentifierRef:
	Identifier                           *Identifier              `json:"Identifier,omitempty"`                           // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OrganizationInformation              *OrganizationInformation `json:"OrganizationInformation,omitempty"`              // OrganizationInformation:
	PaymentTimeLimit                     time.Time                `json:"PaymentTimeLimit,omitempty"`                     // PaymentTimeLimit: The date and time by which the Offer must be paid for once the Reservation is completed
	Penalties                            []Penalties              `json:"Penalties,omitempty"`                            // Penalties:
	PricingAgency                        []PricingAgency          `json:"PricingAgency,omitempty"`                        // PricingAgency:
	PromotionalCode                      *PromotionalCode         `json:"PromotionalCode,omitempty"`                      // PromotionalCode:
	Restriction                          []Restriction            `json:"Restriction,omitempty"`                          // Restriction:
	TicketingAgency                      []TicketingAgency        `json:"TicketingAgency,omitempty"`                      // TicketingAgency:
	TourCodes                            []TourCodes              `json:"TourCodes,omitempty"`                            // TourCodes:
	ValidatingAirline                    []ValidatingAirline      `json:"ValidatingAirline,omitempty"`                    // ValidatingAirline:
	Id                                   string                   `json:"id,omitempty"`                                   // Id: Local indentifier within a given message for this object.
	InstantPurchaseInd                   bool                     `json:"instantPurchaseInd,omitempty"`                   // InstantPurchaseInd: If true the Offer\/Offering must be paid for at the same time as creating the Reservation
	SecureFlightPassengerDataRequiredInd bool                     `json:"secureFlightPassengerDataRequiredInd,omitempty"` // SecureFlightPassengerDataRequiredInd: If true, Secure Flight Passenger Data must be input for all Travelers to complete the Reservation
	TermsAndConditionsRef                string                   `json:"termsAndConditionsRef,omitempty"`                // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
}

// TermsAndConditionsAirChange is an object.
type TermsAndConditionsAirChange struct {
	Type                                 string                   `json:"@type"`                                          // Type:
	BaggageAllowance                     []BaggageAllowance       `json:"BaggageAllowance,omitempty"`                     // BaggageAllowance:
	BaggageRecheck                       []BaggageRecheck         `json:"BaggageRecheck,omitempty"`                       // BaggageRecheck:
	CustomerLoyalty                      []CustomerLoyalty        `json:"CustomerLoyalty,omitempty"`                      // CustomerLoyalty:
	DocumentValidDateRange               *DocumentValidDateRange  `json:"DocumentValidDateRange,omitempty"`               // DocumentValidDateRange:
	ExpiryDate                           time.Time                `json:"ExpiryDate,omitempty"`                           // ExpiryDate: The data and time the offer will expire
	FareGuaranteePolicy                  []FareGuaranteePolicy    `json:"FareGuaranteePolicy,omitempty"`                  // FareGuaranteePolicy:
	FareRuleIdentifierRef                *IdentifierRef           `json:"FareRuleIdentifierRef,omitempty"`                // FareRuleIdentifierRef:
	FulfillmentMethod                    []FulfillmentMethod      `json:"FulfillmentMethod,omitempty"`                    // FulfillmentMethod:
	Identifier                           *Identifier              `json:"Identifier,omitempty"`                           // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OrganizationInformation              *OrganizationInformation `json:"OrganizationInformation,omitempty"`              // OrganizationInformation:
	PaymentTimeLimit                     time.Time                `json:"PaymentTimeLimit,omitempty"`                     // PaymentTimeLimit: The date and time by which the Offer must be paid for once the Reservation is completed
	Penalties                            []Penalties              `json:"Penalties,omitempty"`                            // Penalties:
	PricingAgency                        []PricingAgency          `json:"PricingAgency,omitempty"`                        // PricingAgency:
	PromotionalCode                      *PromotionalCode         `json:"PromotionalCode,omitempty"`                      // PromotionalCode:
	Restriction                          []Restriction            `json:"Restriction,omitempty"`                          // Restriction:
	TicketingAgency                      []TicketingAgency        `json:"TicketingAgency,omitempty"`                      // TicketingAgency:
	TourCodes                            []TourCodes              `json:"TourCodes,omitempty"`                            // TourCodes:
	ValidatingAirline                    []ValidatingAirline      `json:"ValidatingAirline,omitempty"`                    // ValidatingAirline:
	Id                                   string                   `json:"id,omitempty"`                                   // Id: Local indentifier within a given message for this object.
	InstantPurchaseInd                   bool                     `json:"instantPurchaseInd,omitempty"`                   // InstantPurchaseInd: If true the Offer\/Offering must be paid for at the same time as creating the Reservation
	SecureFlightPassengerDataRequiredInd bool                     `json:"secureFlightPassengerDataRequiredInd,omitempty"` // SecureFlightPassengerDataRequiredInd: If true, Secure Flight Passenger Data must be input for all Travelers to complete the Reservation
	TermsAndConditionsRef                string                   `json:"termsAndConditionsRef,omitempty"`                // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
}

// TermsAndConditionsAncillary is an object.
type TermsAndConditionsAncillary struct {
	Type                  string             `json:"@type"`                           // Type:
	ApplicationLimit      ApplicationLimit   `json:"ApplicationLimit"`                // ApplicationLimit:
	CustomerLoyalty       []CustomerLoyalty  `json:"CustomerLoyalty,omitempty"`       // CustomerLoyalty:
	ExpiryDate            time.Time          `json:"ExpiryDate,omitempty"`            // ExpiryDate: The data and time the offer will expire
	Identifier            *Identifier        `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Refundability         *RefundabilityEnum `json:"Refundability,omitempty"`         // Refundability:
	Id                    string             `json:"id,omitempty"`                    // Id: Local indentifier within a given message for this object.
	TermsAndConditionsRef string             `json:"termsAndConditionsRef,omitempty"` // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
	UnsellableInd         bool               `json:"unsellableInd,omitempty"`         // UnsellableInd: If true, this ancillary product can not be sold through Travelport systems
}

// TermsAndConditionsFull is an object.
type TermsAndConditionsFull struct {
	Type                  string            `json:"@type"`                           // Type:
	CustomerLoyalty       []CustomerLoyalty `json:"CustomerLoyalty,omitempty"`       // CustomerLoyalty:
	ExpiryDate            time.Time         `json:"ExpiryDate,omitempty"`            // ExpiryDate: The data and time the offer will expire
	Identifier            *Identifier       `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	TextBlock             []TextBlock       `json:"TextBlock,omitempty"`             // TextBlock:
	Id                    string            `json:"id,omitempty"`                    // Id: Local indentifier within a given message for this object.
	TermsAndConditionsRef string            `json:"termsAndConditionsRef,omitempty"` // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
}

// TermsAndConditionsFullAir is an object.
type TermsAndConditionsFullAir struct {
	Type                                 string                   `json:"@type"`                                          // Type:
	BaggageAllowance                     []BaggageAllowance       `json:"BaggageAllowance,omitempty"`                     // BaggageAllowance:
	BaggageRecheck                       []BaggageRecheck         `json:"BaggageRecheck,omitempty"`                       // BaggageRecheck:
	CustomerLoyalty                      []CustomerLoyalty        `json:"CustomerLoyalty,omitempty"`                      // CustomerLoyalty:
	DocumentValidDateRange               *DocumentValidDateRange  `json:"DocumentValidDateRange,omitempty"`               // DocumentValidDateRange:
	ExpiryDate                           time.Time                `json:"ExpiryDate,omitempty"`                           // ExpiryDate: The data and time the offer will expire
	FareGuaranteePolicy                  []FareGuaranteePolicy    `json:"FareGuaranteePolicy,omitempty"`                  // FareGuaranteePolicy:
	FareRuleIdentifierRef                *IdentifierRef           `json:"FareRuleIdentifierRef,omitempty"`                // FareRuleIdentifierRef:
	Identifier                           *Identifier              `json:"Identifier,omitempty"`                           // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OrganizationInformation              *OrganizationInformation `json:"OrganizationInformation,omitempty"`              // OrganizationInformation:
	PaymentTimeLimit                     time.Time                `json:"PaymentTimeLimit,omitempty"`                     // PaymentTimeLimit: The date and time by which the Offer must be paid for once the Reservation is completed
	Penalties                            []Penalties              `json:"Penalties,omitempty"`                            // Penalties:
	PricingAgency                        []PricingAgency          `json:"PricingAgency,omitempty"`                        // PricingAgency:
	PromotionalCode                      *PromotionalCode         `json:"PromotionalCode,omitempty"`                      // PromotionalCode:
	Restriction                          []Restriction            `json:"Restriction,omitempty"`                          // Restriction:
	TextBlock                            []TextBlock              `json:"TextBlock,omitempty"`                            // TextBlock:
	TicketingAgency                      []TicketingAgency        `json:"TicketingAgency,omitempty"`                      // TicketingAgency:
	TourCodes                            []TourCodes              `json:"TourCodes,omitempty"`                            // TourCodes:
	ValidatingAirline                    []ValidatingAirline      `json:"ValidatingAirline,omitempty"`                    // ValidatingAirline:
	Id                                   string                   `json:"id,omitempty"`                                   // Id: Local indentifier within a given message for this object.
	InstantPurchaseInd                   bool                     `json:"instantPurchaseInd,omitempty"`                   // InstantPurchaseInd: If true the Offer\/Offering must be paid for at the same time as creating the Reservation
	SecureFlightPassengerDataRequiredInd bool                     `json:"secureFlightPassengerDataRequiredInd,omitempty"` // SecureFlightPassengerDataRequiredInd: If true, Secure Flight Passenger Data must be input for all Travelers to complete the Reservation
	TermsAndConditionsRef                string                   `json:"termsAndConditionsRef,omitempty"`                // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
}

// TermsAndConditionsFullAirChange is an object.
type TermsAndConditionsFullAirChange struct {
	Type                                 string                   `json:"@type"`                                          // Type:
	BaggageAllowance                     []BaggageAllowance       `json:"BaggageAllowance,omitempty"`                     // BaggageAllowance:
	BaggageRecheck                       []BaggageRecheck         `json:"BaggageRecheck,omitempty"`                       // BaggageRecheck:
	CustomerLoyalty                      []CustomerLoyalty        `json:"CustomerLoyalty,omitempty"`                      // CustomerLoyalty:
	DocumentValidDateRange               *DocumentValidDateRange  `json:"DocumentValidDateRange,omitempty"`               // DocumentValidDateRange:
	ExpiryDate                           time.Time                `json:"ExpiryDate,omitempty"`                           // ExpiryDate: The data and time the offer will expire
	FareGuaranteePolicy                  []FareGuaranteePolicy    `json:"FareGuaranteePolicy,omitempty"`                  // FareGuaranteePolicy:
	FareRuleIdentifierRef                *IdentifierRef           `json:"FareRuleIdentifierRef,omitempty"`                // FareRuleIdentifierRef:
	FulfillmentMethod                    []FulfillmentMethod      `json:"FulfillmentMethod,omitempty"`                    // FulfillmentMethod:
	Identifier                           *Identifier              `json:"Identifier,omitempty"`                           // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OrganizationInformation              *OrganizationInformation `json:"OrganizationInformation,omitempty"`              // OrganizationInformation:
	PaymentTimeLimit                     time.Time                `json:"PaymentTimeLimit,omitempty"`                     // PaymentTimeLimit: The date and time by which the Offer must be paid for once the Reservation is completed
	Penalties                            []Penalties              `json:"Penalties,omitempty"`                            // Penalties:
	PricingAgency                        []PricingAgency          `json:"PricingAgency,omitempty"`                        // PricingAgency:
	PromotionalCode                      *PromotionalCode         `json:"PromotionalCode,omitempty"`                      // PromotionalCode:
	Restriction                          []Restriction            `json:"Restriction,omitempty"`                          // Restriction:
	TextBlock                            []TextBlock              `json:"TextBlock,omitempty"`                            // TextBlock:
	TicketingAgency                      []TicketingAgency        `json:"TicketingAgency,omitempty"`                      // TicketingAgency:
	TourCodes                            []TourCodes              `json:"TourCodes,omitempty"`                            // TourCodes:
	ValidatingAirline                    []ValidatingAirline      `json:"ValidatingAirline,omitempty"`                    // ValidatingAirline:
	Id                                   string                   `json:"id,omitempty"`                                   // Id: Local indentifier within a given message for this object.
	InstantPurchaseInd                   bool                     `json:"instantPurchaseInd,omitempty"`                   // InstantPurchaseInd: If true the Offer\/Offering must be paid for at the same time as creating the Reservation
	SecureFlightPassengerDataRequiredInd bool                     `json:"secureFlightPassengerDataRequiredInd,omitempty"` // SecureFlightPassengerDataRequiredInd: If true, Secure Flight Passenger Data must be input for all Travelers to complete the Reservation
	TermsAndConditionsRef                string                   `json:"termsAndConditionsRef,omitempty"`                // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
}

// TermsAndConditionsFullAncillary is an object.
type TermsAndConditionsFullAncillary struct {
	Type                  string             `json:"@type"`                           // Type:
	ApplicationLimit      ApplicationLimit   `json:"ApplicationLimit"`                // ApplicationLimit:
	CustomerLoyalty       []CustomerLoyalty  `json:"CustomerLoyalty,omitempty"`       // CustomerLoyalty:
	ExpiryDate            time.Time          `json:"ExpiryDate,omitempty"`            // ExpiryDate: The data and time the offer will expire
	Identifier            *Identifier        `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Refundability         *RefundabilityEnum `json:"Refundability,omitempty"`         // Refundability:
	TextBlock             []TextBlock        `json:"TextBlock,omitempty"`             // TextBlock:
	Id                    string             `json:"id,omitempty"`                    // Id: Local indentifier within a given message for this object.
	TermsAndConditionsRef string             `json:"termsAndConditionsRef,omitempty"` // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
	UnsellableInd         bool               `json:"unsellableInd,omitempty"`         // UnsellableInd: If true, this ancillary product can not be sold through Travelport systems
}

// TermsAndConditionsFullHospitality is an object.
type TermsAndConditionsFullHospitality struct {
	Type                  string                `json:"@type"`                           // Type:
	AcceptedCreditCard    []AcceptedCreditCard  `json:"AcceptedCreditCard,omitempty"`    // AcceptedCreditCard:
	CancelPenalty         []CancelPenalty       `json:"CancelPenalty,omitempty"`         // CancelPenalty:
	CheckInOutPolicy      *CheckInOutPolicy     `json:"CheckInOutPolicy,omitempty"`      // CheckInOutPolicy:
	CustomerLoyalty       []CustomerLoyalty     `json:"CustomerLoyalty,omitempty"`       // CustomerLoyalty:
	DepositPolicy         *DepositPolicy        `json:"DepositPolicy,omitempty"`         // DepositPolicy:
	Description           []string              `json:"Description,omitempty"`           // Description:
	ExpiryDate            time.Time             `json:"ExpiryDate,omitempty"`            // ExpiryDate: The data and time the offer will expire
	Guarantee             []Guarantee           `json:"Guarantee,omitempty"`             // Guarantee:
	Identifier            *Identifier           `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	MealsIncluded         *MealsIncluded        `json:"MealsIncluded,omitempty"`         // MealsIncluded: Indicates if a meal is included or not
	ProductRateCodeInfo   []ProductRateCodeInfo `json:"ProductRateCodeInfo,omitempty"`   // ProductRateCodeInfo:
	RatePaymentInfo       RatePaymentEnum       `json:"RatePaymentInfo,omitempty"`       // RatePaymentInfo: Payment Rate
	TextBlock             []TextBlock           `json:"TextBlock,omitempty"`             // TextBlock:
	Id                    string                `json:"id,omitempty"`                    // Id: Local indentifier within a given message for this object.
	TermsAndConditionsRef string                `json:"termsAndConditionsRef,omitempty"` // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
}

// TermsAndConditionsFullID is an object.
type TermsAndConditionsFullID struct {
	Type                  string      `json:"@type"`                           // Type:
	Identifier            *Identifier `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id                    string      `json:"id,omitempty"`                    // Id: Local indentifier within a given message for this object.
	TermsAndConditionsRef string      `json:"termsAndConditionsRef,omitempty"` // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
}

// TermsAndConditionsFullVehicle is an object.
type TermsAndConditionsFullVehicle struct {
	Type                          string                `json:"@type"`                                   // Type:
	CustomerLoyalty               []CustomerLoyalty     `json:"CustomerLoyalty,omitempty"`               // CustomerLoyalty:
	DriverInfo                    *DriverInfo           `json:"DriverInfo,omitempty"`                    // DriverInfo: Basic information (metadata) about the intended driver of the vehicle
	ExpiryDate                    time.Time             `json:"ExpiryDate,omitempty"`                    // ExpiryDate: The data and time the offer will expire
	Identifier                    *Identifier           `json:"Identifier,omitempty"`                    // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Policy                        *Policy               `json:"Policy,omitempty"`                        // Policy:
	ProductRateCodeInfo           []ProductRateCodeInfo `json:"ProductRateCodeInfo,omitempty"`           // ProductRateCodeInfo:
	TermsAndConditionsSubCategory *TextBlock            `json:"TermsAndConditionsSubCategory,omitempty"` // TermsAndConditionsSubCategory:
	TextBlock                     []TextBlock           `json:"TextBlock,omitempty"`                     // TextBlock:
	FlightRestrictionInd          bool                  `json:"flightRestrictionInd,omitempty"`          // FlightRestrictionInd: if true, the traveler must have a valid flight to qualify for this Offer
	Id                            string                `json:"id,omitempty"`                            // Id: Local indentifier within a given message for this object.
	TermsAndConditionsRef         string                `json:"termsAndConditionsRef,omitempty"`         // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
}

// TermsAndConditionsID is an object.
type TermsAndConditionsID struct {
	Type                  string      `json:"@type"`                           // Type:
	Identifier            *Identifier `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id                    string      `json:"id,omitempty"`                    // Id: Local indentifier within a given message for this object.
	TermsAndConditionsRef string      `json:"termsAndConditionsRef,omitempty"` // TermsAndConditionsRef: Used to reference another instance of this object in the same message.
}

// TextBlock is an object.
type TextBlock struct {
	Type          string          `json:"@type"`           // Type:
	TextFormatted []TextFormatted `json:"TextFormatted"`   // TextFormatted:
	Id            string          `json:"id,omitempty"`    // Id: Internally referenced id
	Title         string          `json:"title,omitempty"` // Title: Title
}

// TextBlockDetail is an object.
type TextBlockDetail struct {
	Type             string           `json:"@type"`                 // Type:
	DateCreateModify DateCreateModify `json:"DateCreateModify"`      // DateCreateModify: Time stamp of the creation.
	Image            *Image           `json:"Image,omitempty"`       // Image: URL of the image
	TextFormatted    []TextFormatted  `json:"TextFormatted"`         // TextFormatted:
	URL              string           `json:"URL,omitempty"`         // URL: A URL for this block
	Description      string           `json:"description,omitempty"` // Description: Text block detail description
	Id               string           `json:"id,omitempty"`          // Id: Internally referenced id
	Sequence         int32            `json:"sequence,omitempty"`    // Sequence: The order of the text block, if there are more than one block.
	Title            string           `json:"title,omitempty"`       // Title: Title
}

// TextFormatted is an object. Provides text and indicates whether it is formatted or not.
type TextFormatted struct {
	Language   string         `json:"language,omitempty"`   // Language: The language in which the text is provided.
	TextFormat TextFormatEnum `json:"textFormat,omitempty"` // TextFormat: Describes the format of text such as plain text or html
	Value      string         `json:"value,omitempty"`      // Value:
}

// TextTitleAndDescription is an object. Descriptive text
type TextTitleAndDescription struct {
	Languages []string `json:"languages,omitempty"` // Languages: Language of the text
	Title     string   `json:"title,omitempty"`     // Title: Title of the Text
	Value     string   `json:"value,omitempty"`     // Value:
}

// ThreeDomainSecurity is an object.
type ThreeDomainSecurity struct {
	Type                       string                     `json:"@type,omitempty"`            // Type:
	ThreeDomainSecurityGateway ThreeDomainSecurityGateway `json:"ThreeDomainSecurityGateway"` // ThreeDomainSecurityGateway:
	ThreeDomainSecurityResults ThreeDomainSecurityResults `json:"ThreeDomainSecurityResults"` // ThreeDomainSecurityResults:
}

// ThreeDomainSecurityGateway is an object.
type ThreeDomainSecurityGateway struct {
	Type                       string                      `json:"@type,omitempty"`                      // Type:
	AuthenticationVerification *AuthenticationVerification `json:"AuthenticationVerification,omitempty"` // AuthenticationVerification:
	Password                   *Password                   `json:"Password,omitempty"`                   // Password:
	ECI                        string                      `json:"eCI,omitempty"`                        // ECI: The eCI value
	MerchantID                 string                      `json:"merchantID,omitempty"`                 // MerchantID: The merchant ID value
	ProcessorID                string                      `json:"processorID,omitempty"`                // ProcessorID: The processor ID value
	URL                        string                      `json:"uRL,omitempty"`                        // URL: Transaction URL.
}

// ThreeDomainSecurityResults is an object.
type ThreeDomainSecurityResults struct {
	Type                 string `json:"@type,omitempty"`                // Type:
	CAVV                 string `json:"cAVV,omitempty"`                 // CAVV: The cAVV value
	ECI                  string `json:"eCI,omitempty"`                  // ECI: Electronic Commerce Indicator - 3-D secure data, contact your authenticator for rules and downline processing.
	PAResStatus          string `json:"pAResStatus,omitempty"`          // PAResStatus: The pAResStatus value
	SignatureVerfication string `json:"signatureVerfication,omitempty"` // SignatureVerfication: The signature Verification value
	TransactionID        string `json:"transactionID,omitempty"`        // TransactionID: The transaction ID
	UCAFIndicator        string `json:"uCAFIndicator,omitempty"`        // UCAFIndicator: Universal Card Authentication Field™ MasterCard only UCAF is the mechanism that is used to transmit the AAV from the merchant to issuer for authentication purposes during the authorization process
	XID                  string `json:"xID,omitempty"`                  // XID: Merchants must ensure that each Payer Authentication Request (PAReq) contains a unique combination of account ID and XID
}

// ticketBasePassengerCurrencyCodePattern is the validation pattern for TicketBasePassenger.CurrencyCode
var ticketBasePassengerCurrencyCodePattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// ticketBasePassengerDecimalPlacePattern is the validation pattern for TicketBasePassenger.DecimalPlace
var ticketBasePassengerDecimalPlacePattern = regexp.MustCompile(`([0-4])`)

// TicketBasePassenger is an object. The monetary value
type TicketBasePassenger struct {
	BTInd            bool    `json:"bTInd,omitempty"`            // BTInd: If true, this is a bulk ticket fare
	CodeAuthority    string  `json:"codeAuthority"`              // CodeAuthority: Assigned Type: c-1100:CodeContext
	CurrencyCode     string  `json:"currencyCode,omitempty"`     // CurrencyCode: Assigned Type: c-1100:CurrencyCode
	DecimalAuthority string  `json:"decimalAuthority,omitempty"` // DecimalAuthority: Assigned Type: c-1100:CodeContext
	DecimalPlace     int32   `json:"decimalPlace"`               // DecimalPlace: Assigned Type: c-1100:CurrencyMinorUnit
	ITInd            bool    `json:"iTInd,omitempty"`            // ITInd: If true, this is an inclusive tour fare
	Value            float32 `json:"value,omitempty"`            // Value:
}

// ticketChangeEligibilityPassengerTypeCodePattern is the validation pattern for TicketChangeEligibility.PassengerTypeCode
var ticketChangeEligibilityPassengerTypeCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{3,5})`)

// TicketChangeEligibility is an object.
type TicketChangeEligibility struct {
	Type                      string          `json:"@type,omitempty"`                     // Type:
	Identifier                *Identifier     `json:"Identifier,omitempty"`                // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	PassengerTypeCode         string          `json:"PassengerTypeCode,omitempty"`         // PassengerTypeCode: Pessanger Type Code
	Penalties                 *Penalties      `json:"Penalties,omitempty"`                 // Penalties:
	AutomationNotSupportedInd bool            `json:"automationNotSupportedInd,omitempty"` // AutomationNotSupportedInd: automation process not supported for this eligibility
	Exchangeable              AllSomeNoneEnum `json:"exchangeable,omitempty"`              // Exchangeable: Used to indicate if all, some, or none of the ticket can be exchanged or         refunded
	ObjID                     string          `json:"objID,omitempty"`                     // ObjID:
	Refundable                AllSomeNoneEnum `json:"refundable,omitempty"`                // Refundable: Used to indicate if all, some, or none of the ticket can be exchanged or         refunded
}

// TicketChangeEligibilityID is an object.
type TicketChangeEligibilityID struct {
	Type       string      `json:"@type,omitempty"`      // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	ObjID      string      `json:"objID,omitempty"`      // ObjID:
}

// TicketChangeEligibilityListResponse is an object.
type TicketChangeEligibilityListResponse struct {
	TicketChangeEligibility []TicketChangeEligibility `json:"TicketChangeEligibility,omitempty"` // TicketChangeEligibility:
}

// ticketNumberValuePattern is the validation pattern for TicketNumber.Value
var ticketNumberValuePattern = regexp.MustCompile(`([0-9]+)?`)

// TicketNumber is an object. The ticketNumber that will be used as partial payment for this Offer\/Offering
type TicketNumber struct {
	TicketIssuer string `json:"ticketIssuer,omitempty"` // TicketIssuer: Ticket issuer
	Value        string `json:"value,omitempty"`        // Value: Ticket number
}

// ticketingAgencyCodePattern is the validation pattern for TicketingAgency.Code
var ticketingAgencyCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,10})`)

// TicketingAgency is an object.
type TicketingAgency struct {
	Type                string   `json:"@type,omitempty"`               // Type:
	Code                string   `json:"Code"`                          // Code: The code of the ticketing agency
	ProductRef          []string `json:"ProductRef,omitempty"`          // ProductRef: The Product Ref the TicketingAgency applies to
	SegmentSequenceList []int32  `json:"SegmentSequenceList,omitempty"` // SegmentSequenceList: The segmentSequenceList the TicketingAgency applies to
}

// TimeRange is an object. Specify time.
type TimeRange struct {
	End   string `json:"end,omitempty"`   // End: endTime
	Start string `json:"start,omitempty"` // Start: start time
}

// TotalAmount is an object.
type TotalAmount struct {
	Type           string             `json:"@type,omitempty"`          // Type:
	Base           float32            `json:"Base,omitempty"`           // Base: The price prior to all applicable taxes of a product such as the rate for a room or fare for a flight.
	CurrencyCode   *CurrencyCode      `json:"CurrencyCode,omitempty"`   // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	Fees           *Fees              `json:"Fees,omitempty"`           // Fees:
	Taxes          *Taxes             `json:"Taxes,omitempty"`          // Taxes:
	Total          float32            `json:"Total,omitempty"`          // Total: Specifies the total price including base + taxes + fees
	ApproximateInd bool               `json:"approximateInd,omitempty"` // ApproximateInd: True if this amount has been converted from the original amount
	CurrencySource CurrencySourceEnum `json:"currencySource,omitempty"` // CurrencySource: The system requesting or returning the currency code specified in the attribute
}

// TourCode is an object. Tour code
type TourCode struct {
	TourCodeType TourCodeTypeEnum `json:"tourCodeType,omitempty"` // TourCodeType:
	Value        string           `json:"value,omitempty"`        // Value:
}

// TourCodes is an object.
type TourCodes struct {
	Type                  string                  `json:"@type,omitempty"`                 // Type:
	TourCode              TourCode                `json:"TourCode"`                        // TourCode: Tour code
	TravelerIdentifierRef []TravelerIdentifierRef `json:"TravelerIdentifierRef,omitempty"` // TravelerIdentifierRef:
}

// TravelAgency is an object.
type TravelAgency struct {
	Type                  string               `json:"@type"`                           // Type:
	CorporateCode         string               `json:"CorporateCode,omitempty"`         // CorporateCode: A reference assigned by the Travel Agency to identify the corporate organization
	Identifier            *Identifier          `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OrganizationName      CompanyName          `json:"OrganizationName"`                // OrganizationName: Identifies a company by name.
	ProfileName           []string             `json:"ProfileName,omitempty"`           // ProfileName:
	TravelOrganizationRef string               `json:"TravelOrganizationRef,omitempty"` // TravelOrganizationRef: An organization that has a name and a structure and members and directly works in the travel industry
	Id                    string               `json:"id,omitempty"`                    // Id: Simple xsd id, not for external use
	OrganizationType      OrganizationTypeEnum `json:"organizationType,omitempty"`      // OrganizationType: The type of organization such as an Agency, Branch, Company, Supplier, Provider
}

// TravelAgencyID is an object.
type TravelAgencyID struct {
	Type                  string      `json:"@type"`                           // Type:
	Identifier            *Identifier `json:"Identifier,omitempty"`            // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	TravelOrganizationRef string      `json:"TravelOrganizationRef,omitempty"` // TravelOrganizationRef: An organization that has a name and a structure and members and directly works in the travel industry
	Id                    string      `json:"id,omitempty"`                    // Id: Simple xsd id, not for external use
}

// travelDocumentNationalityPattern is the validation pattern for TravelDocument.Nationality
var travelDocumentNationalityPattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// travelDocumentBirthCountryPattern is the validation pattern for TravelDocument.BirthCountry
var travelDocumentBirthCountryPattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// travelDocumentIssueCountryPattern is the validation pattern for TravelDocument.IssueCountry
var travelDocumentIssueCountryPattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// TravelDocument is an object.
type TravelDocument struct {
	Type          string          `json:"@type"`                   // Type:
	Gender        GenderEnum      `json:"Gender"`                  // Gender: Gender Type Male, Female etc. This field is not used by Hotel APIs and will be ignored
	Nationality   string          `json:"Nationality,omitempty"`   // Nationality: Specifies a 2 character country code as defined in ISO3166.
	PersonName    PersonName      `json:"PersonName"`              // PersonName:
	BirthCountry  string          `json:"birthCountry,omitempty"`  // BirthCountry: Birth country on Country Code ISO value
	BirthDate     string          `json:"birthDate,omitempty"`     // BirthDate: The date of birth of the document holder
	BirthPlace    string          `json:"birthPlace,omitempty"`    // BirthPlace: Birth place value
	DocNumber     string          `json:"docNumber"`               // DocNumber: Document number value
	DocType       DocTypeCodeEnum `json:"docType,omitempty"`       // DocType: Codes from OTA DOC - Document Type
	ExpireDate    string          `json:"expireDate,omitempty"`    // ExpireDate: Date of expiration
	Id            string          `json:"id,omitempty"`            // Id: Locally referenced id
	IssueCountry  string          `json:"issueCountry,omitempty"`  // IssueCountry: Issue country on Country Code ISO
	IssueDate     string          `json:"issueDate,omitempty"`     // IssueDate: Date of Issue
	PlaceOfIssue  string          `json:"placeOfIssue,omitempty"`  // PlaceOfIssue: Place of issue value
	Residence     string          `json:"residence,omitempty"`     // Residence: Residence value
	StateProvCode string          `json:"stateProvCode,omitempty"` // StateProvCode: State Province Code value
}

// travelDocumentDetailNationalityPattern is the validation pattern for TravelDocumentDetail.Nationality
var travelDocumentDetailNationalityPattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// travelDocumentDetailBirthCountryPattern is the validation pattern for TravelDocumentDetail.BirthCountry
var travelDocumentDetailBirthCountryPattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// travelDocumentDetailIssueCountryPattern is the validation pattern for TravelDocumentDetail.IssueCountry
var travelDocumentDetailIssueCountryPattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// TravelDocumentDetail is an object.
type TravelDocumentDetail struct {
	Type                      string            `json:"@type"`                               // Type:
	Address                   *Address          `json:"Address,omitempty"`                   // Address:
	Gender                    GenderEnum        `json:"Gender"`                              // Gender: Gender Type Male, Female etc. This field is not used by Hotel APIs and will be ignored
	IssuedForGeoPoliticalArea *GeoPoliticalArea `json:"IssuedForGeoPoliticalArea,omitempty"` // IssuedForGeoPoliticalArea: The location code of the geographical location. Codes from Ref Pub
	Nationality               string            `json:"Nationality,omitempty"`               // Nationality: Specifies a 2 character country code as defined in ISO3166.
	PersonName                PersonName        `json:"PersonName"`                          // PersonName:
	BirthCountry              string            `json:"birthCountry,omitempty"`              // BirthCountry: Birth country on Country Code ISO value
	BirthDate                 string            `json:"birthDate,omitempty"`                 // BirthDate: The date of birth of the document holder
	BirthPlace                string            `json:"birthPlace,omitempty"`                // BirthPlace: Birth place value
	DocNumber                 string            `json:"docNumber"`                           // DocNumber: Document number value
	DocType                   DocTypeCodeEnum   `json:"docType,omitempty"`                   // DocType: Codes from OTA DOC - Document Type
	ExpireDate                string            `json:"expireDate,omitempty"`                // ExpireDate: Date of expiration
	Id                        string            `json:"id,omitempty"`                        // Id: Locally referenced id
	IssueCountry              string            `json:"issueCountry,omitempty"`              // IssueCountry: Issue country on Country Code ISO
	IssueDate                 string            `json:"issueDate,omitempty"`                 // IssueDate: Date of Issue
	PlaceOfIssue              string            `json:"placeOfIssue,omitempty"`              // PlaceOfIssue: Place of issue value
	Residence                 string            `json:"residence,omitempty"`                 // Residence: Residence value
	StateProvCode             string            `json:"stateProvCode,omitempty"`             // StateProvCode: State Province Code value
}

// travelerNationalityPattern is the validation pattern for Traveler.Nationality
var travelerNationalityPattern = regexp.MustCompile(`[a-zA-Z]{2}`)

// travelerPassengerTypeCodePattern is the validation pattern for Traveler.PassengerTypeCode
var travelerPassengerTypeCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{3,5})`)

// Traveler is an object.
type Traveler struct {
	Type                   string             `json:"@type"`                            // Type:
	Address                []Address          `json:"Address,omitempty"`                // Address:
	AlternateContact       []AlternateContact `json:"AlternateContact,omitempty"`       // AlternateContact:
	Comments               *Comments          `json:"Comments,omitempty"`               // Comments:
	CustomerLoyalty        []CustomerLoyalty  `json:"CustomerLoyalty,omitempty"`        // CustomerLoyalty:
	Email                  []Email            `json:"Email,omitempty"`                  // Email:
	Identifier             *Identifier        `json:"Identifier,omitempty"`             // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	PersonName             PersonName         `json:"PersonName"`                       // PersonName:
	RailDiscountCard       []RailDiscountCard `json:"RailDiscountCard,omitempty"`       // RailDiscountCard:
	Telephone              []Telephone        `json:"Telephone,omitempty"`              // Telephone:
	TravelDocument         []TravelDocument   `json:"TravelDocument,omitempty"`         // TravelDocument:
	TravelerRef            string             `json:"TravelerRef,omitempty"`            // TravelerRef:
	AccompaniedByInfantInd bool               `json:"accompaniedByInfantInd,omitempty"` // AccompaniedByInfantInd:
	BirthDate              string             `json:"birthDate,omitempty"`              // BirthDate: Date of Birth YYYY-MM-DD
	Gender                 GenderEnum         `json:"gender,omitempty"`                 // Gender: Gender Type Male, Female etc. This field is not used by Hotel APIs and will be ignored
	Id                     string             `json:"id,omitempty"`                     // Id:
	Nationality            string             `json:"nationality,omitempty"`            // Nationality: Nationality on country code ISO
	PassengerTypeCode      string             `json:"passengerTypeCode,omitempty"`      // PassengerTypeCode: Passenger type code
}

// TravelerGeographicLocation is an object. Specifies which location the Traveler resides in. Used for resident fares
type TravelerGeographicLocation struct {
	GeneralLargeFamilyResidentDiscountInd bool                       `json:"generalLargeFamilyResidentDiscountInd,omitempty"` // GeneralLargeFamilyResidentDiscountInd: if true, this request qualifies for general large family resident discount. General large families (up to 3 children) from Spain, from the EU/EEA or of any other nationality, whose residency in Spain is recognised and who are in possession of a large-family certificate issued by the autonomous community in which they live.
	ResidentGeographicCode                string                     `json:"residentGeographicCode,omitempty"`                // ResidentGeographicCode: Resident code, currently used to handle Spanish residency fares for NDC channel where this code is required in addition to the city of residence
	SpecialLargeFamilyResidentDiscountInd bool                       `json:"specialLargeFamilyResidentDiscountInd,omitempty"` // SpecialLargeFamilyResidentDiscountInd: if true, this request qualifies for special large family resident discount. Special large families (4 or more children) from Spain, from the EU/EEA or of any other nationality, whose residency in Spain is recognised and who are in possession of a large-family certificate issued by the autonomous community in which they live.
	TravelerGeographicLocationType        TravelerGeographicTypeEnum `json:"travelerGeographicLocationType,omitempty"`        // TravelerGeographicLocationType:
	Value                                 string                     `json:"value,omitempty"`                                 // Value:
}

// TravelerID is an object.
type TravelerID struct {
	Type        string      `json:"@type"`                 // Type:
	Identifier  *Identifier `json:"Identifier,omitempty"`  // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	TravelerRef string      `json:"TravelerRef,omitempty"` // TravelerRef:
	Id          string      `json:"id,omitempty"`          // Id:
}

// TravelerIdentifier is an object.
type TravelerIdentifier struct {
	Identifier  *Identifier `json:"Identifier,omitempty"`  // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	TravelerRef string      `json:"TravelerRef,omitempty"` // TravelerRef:
	Id          string      `json:"id,omitempty"`          // Id:
}

// travelerIdentifierRefPassengerTypeCodePattern is the validation pattern for TravelerIdentifierRef.PassengerTypeCode
var travelerIdentifierRefPassengerTypeCodePattern = regexp.MustCompile(`([a-zA-Z0-9]{3,5})`)

// TravelerIdentifierRef is an object.
type TravelerIdentifierRef struct {
	Description       string   `json:"description,omitempty"`       // Description: Descriptive text used to identify the contents of a target object
	Id                string   `json:"id,omitempty"`                // Id: A locally referenced ID
	Name              string   `json:"name,omitempty"`              // Name: Traveler identifier
	PassengerTypeCode string   `json:"passengerTypeCode,omitempty"` // PassengerTypeCode: Passenger Type code
	Uris              []string `json:"uris,omitempty"`              // Uris: Uniform Resource Identifier
	Value             string   `json:"value,omitempty"`             // Value:
}

// TravelerProduct is an object.
type TravelerProduct struct {
	Type                   string                 `json:"@type"`                            // Type:
	ConfirmationStatusEnum ConfirmationStatusEnum `json:"ConfirmationStatusEnum,omitempty"` // ConfirmationStatusEnum: Status returned in a response for a two or more phase commitment process
	OfferRef               string                 `json:"OfferRef,omitempty"`               // OfferRef: A pointer to the Offer id
	ProductRef             string                 `json:"ProductRef,omitempty"`             // ProductRef: A pointer to the product id
	TravelerRef            string                 `json:"TravelerRef,omitempty"`            // TravelerRef: A pointer to the traveler id
	Id                     string                 `json:"id,omitempty"`                     // Id:
}

// TravelerProductID is an object.
type TravelerProductID struct {
	Type string `json:"@type"`        // Type:
	Id   string `json:"id,omitempty"` // Id:
}

// TravelerUpdatedItem is an object.
type TravelerUpdatedItem struct {
	Type                    string `json:"@type"`                             // Type:
	TravelerUpdatableItemID string `json:"TravelerUpdatableItemID,omitempty"` // TravelerUpdatableItemID: A unique GUID to identify the TravelerUpdatedItem
	AddInd                  bool   `json:"addInd,omitempty"`                  // AddInd: If true the TravelerUpdatedItem is being added to the Traveler
	DeleteInd               bool   `json:"deleteInd,omitempty"`               // DeleteInd: If true the TravelerUpdatedItem is being deleted from the Traveler
	ModifyInd               bool   `json:"modifyInd,omitempty"`               // ModifyInd: If true the TravelerUpdatedItem is being modified in the Traveler
}

// TravelerUpdatedItemAddress is an object.
type TravelerUpdatedItemAddress struct {
	Type                    string  `json:"@type"`                             // Type:
	Address                 Address `json:"Address"`                           // Address:
	TravelerUpdatableItemID string  `json:"TravelerUpdatableItemID,omitempty"` // TravelerUpdatableItemID: A unique GUID to identify the TravelerUpdatedItem
	AddInd                  bool    `json:"addInd,omitempty"`                  // AddInd: If true the TravelerUpdatedItem is being added to the Traveler
	DeleteInd               bool    `json:"deleteInd,omitempty"`               // DeleteInd: If true the TravelerUpdatedItem is being deleted from the Traveler
	ModifyInd               bool    `json:"modifyInd,omitempty"`               // ModifyInd: If true the TravelerUpdatedItem is being modified in the Traveler
}

// TravelerUpdatedItemBirthDate is an object.
type TravelerUpdatedItemBirthDate struct {
	Type                    string `json:"@type"`                             // Type:
	BirthDateUpdatable      string `json:"BirthDateUpdatable,omitempty"`      // BirthDateUpdatable: The updated birth date
	TravelerUpdatableItemID string `json:"TravelerUpdatableItemID,omitempty"` // TravelerUpdatableItemID: A unique GUID to identify the TravelerUpdatedItem
	AddInd                  bool   `json:"addInd,omitempty"`                  // AddInd: If true the TravelerUpdatedItem is being added to the Traveler
	DeleteInd               bool   `json:"deleteInd,omitempty"`               // DeleteInd: If true the TravelerUpdatedItem is being deleted from the Traveler
	ModifyInd               bool   `json:"modifyInd,omitempty"`               // ModifyInd: If true the TravelerUpdatedItem is being modified in the Traveler
}

// TravelerUpdatedItemCustomerLoyalty is an object.
type TravelerUpdatedItemCustomerLoyalty struct {
	Type                    string          `json:"@type"`                             // Type:
	CustomerLoyalty         CustomerLoyalty `json:"CustomerLoyalty"`                   // CustomerLoyalty: Specifies the ID for the membership program.
	TravelerUpdatableItemID string          `json:"TravelerUpdatableItemID,omitempty"` // TravelerUpdatableItemID: A unique GUID to identify the TravelerUpdatedItem
	AddInd                  bool            `json:"addInd,omitempty"`                  // AddInd: If true the TravelerUpdatedItem is being added to the Traveler
	DeleteInd               bool            `json:"deleteInd,omitempty"`               // DeleteInd: If true the TravelerUpdatedItem is being deleted from the Traveler
	ModifyInd               bool            `json:"modifyInd,omitempty"`               // ModifyInd: If true the TravelerUpdatedItem is being modified in the Traveler
}

// TravelerUpdatedItemEmail is an object.
type TravelerUpdatedItemEmail struct {
	Type                    string `json:"@type"`                             // Type:
	Email                   Email  `json:"Email"`                             // Email: Electronic email addresses, in IETF specified format.
	TravelerUpdatableItemID string `json:"TravelerUpdatableItemID,omitempty"` // TravelerUpdatableItemID: A unique GUID to identify the TravelerUpdatedItem
	AddInd                  bool   `json:"addInd,omitempty"`                  // AddInd: If true the TravelerUpdatedItem is being added to the Traveler
	DeleteInd               bool   `json:"deleteInd,omitempty"`               // DeleteInd: If true the TravelerUpdatedItem is being deleted from the Traveler
	ModifyInd               bool   `json:"modifyInd,omitempty"`               // ModifyInd: If true the TravelerUpdatedItem is being modified in the Traveler
}

// TravelerUpdatedItemGender is an object.
type TravelerUpdatedItemGender struct {
	Type                    string     `json:"@type"`                             // Type:
	GenderUpdatable         GenderEnum `json:"GenderUpdatable,omitempty"`         // GenderUpdatable: Gender Type Male, Female etc. This field is not used by Hotel APIs and will be ignored
	TravelerUpdatableItemID string     `json:"TravelerUpdatableItemID,omitempty"` // TravelerUpdatableItemID: A unique GUID to identify the TravelerUpdatedItem
	AddInd                  bool       `json:"addInd,omitempty"`                  // AddInd: If true the TravelerUpdatedItem is being added to the Traveler
	DeleteInd               bool       `json:"deleteInd,omitempty"`               // DeleteInd: If true the TravelerUpdatedItem is being deleted from the Traveler
	ModifyInd               bool       `json:"modifyInd,omitempty"`               // ModifyInd: If true the TravelerUpdatedItem is being modified in the Traveler
}

// TravelerUpdatedItemPersonName is an object.
type TravelerUpdatedItemPersonName struct {
	Type                    string              `json:"@type"`                             // Type:
	PersonNameUpdatable     PersonNameUpdatable `json:"PersonNameUpdatable"`               // PersonNameUpdatable:
	TravelerUpdatableItemID string              `json:"TravelerUpdatableItemID,omitempty"` // TravelerUpdatableItemID: A unique GUID to identify the TravelerUpdatedItem
	AddInd                  bool                `json:"addInd,omitempty"`                  // AddInd: If true the TravelerUpdatedItem is being added to the Traveler
	DeleteInd               bool                `json:"deleteInd,omitempty"`               // DeleteInd: If true the TravelerUpdatedItem is being deleted from the Traveler
	ModifyInd               bool                `json:"modifyInd,omitempty"`               // ModifyInd: If true the TravelerUpdatedItem is being modified in the Traveler
}

// TravelerUpdatedItemTelephone is an object.
type TravelerUpdatedItemTelephone struct {
	Type                    string    `json:"@type"`                             // Type:
	Telephone               Telephone `json:"Telephone"`                         // Telephone:
	TravelerUpdatableItemID string    `json:"TravelerUpdatableItemID,omitempty"` // TravelerUpdatableItemID: A unique GUID to identify the TravelerUpdatedItem
	AddInd                  bool      `json:"addInd,omitempty"`                  // AddInd: If true the TravelerUpdatedItem is being added to the Traveler
	DeleteInd               bool      `json:"deleteInd,omitempty"`               // DeleteInd: If true the TravelerUpdatedItem is being deleted from the Traveler
	ModifyInd               bool      `json:"modifyInd,omitempty"`               // ModifyInd: If true the TravelerUpdatedItem is being modified in the Traveler
}

// TravelerUpdatedItemTravelDocument is an object.
type TravelerUpdatedItemTravelDocument struct {
	Type                    string         `json:"@type"`                             // Type:
	TravelDocument          TravelDocument `json:"TravelDocument"`                    // TravelDocument:
	TravelerUpdatableItemID string         `json:"TravelerUpdatableItemID,omitempty"` // TravelerUpdatableItemID: A unique GUID to identify the TravelerUpdatedItem
	AddInd                  bool           `json:"addInd,omitempty"`                  // AddInd: If true the TravelerUpdatedItem is being added to the Traveler
	DeleteInd               bool           `json:"deleteInd,omitempty"`               // DeleteInd: If true the TravelerUpdatedItem is being deleted from the Traveler
	ModifyInd               bool           `json:"modifyInd,omitempty"`               // ModifyInd: If true the TravelerUpdatedItem is being modified in the Traveler
}

// UpsellOfferingIdentifier is an object.
type UpsellOfferingIdentifier struct {
	CatalogProductOfferingRef string      `json:"CatalogProductOfferingRef,omitempty"` // CatalogProductOfferingRef: Used to reference another instance of this object in the same message
	Identifier                *Identifier `json:"Identifier,omitempty"`                // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id                        string      `json:"id,omitempty"`                        // Id: Local indentifier within a given message for this object.
}

// validatingAirlineValidatingAirlinePattern is the validation pattern for ValidatingAirline.ValidatingAirline
var validatingAirlineValidatingAirlinePattern = regexp.MustCompile(`([a-zA-Z0-9]{2,3})`)

// ValidatingAirline is an object.
type ValidatingAirline struct {
	Type                string   `json:"@type,omitempty"`               // Type:
	ProductRef          []string `json:"ProductRef,omitempty"`          // ProductRef: The productRef the validatingCarrier applies to
	SegmentSequenceList []int32  `json:"SegmentSequenceList,omitempty"` // SegmentSequenceList: The segmentSequenceList the validatingCarrier applies to
	ValidatingAirline   string   `json:"ValidatingAirline"`             // ValidatingAirline: Represents the airline responsible for ticketing\/fulfillment of this Offer
}

// Vehicle is an object.
type Vehicle struct {
	Type                 string            `json:"@type"`                          // Type:
	FuelTypeCode         *Code             `json:"FuelTypeCode,omitempty"`         // FuelTypeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	Image                *Image            `json:"Image,omitempty"`                // Image: URL of the image
	TransmissionTypeCode *Code             `json:"TransmissionTypeCode,omitempty"` // TransmissionTypeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	VehicleCategoryCode  *Code             `json:"VehicleCategoryCode,omitempty"`  // VehicleCategoryCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	VehicleClassCode     *Code             `json:"VehicleClassCode,omitempty"`     // VehicleClassCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	VehicleMakeModel     *VehicleMakeModel `json:"VehicleMakeModel,omitempty"`     // VehicleMakeModel: The make and model of the vehicle along with a description
	VehicleSizeCode      *Code             `json:"VehicleSizeCode,omitempty"`      // VehicleSizeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	AirConditioningInd   bool              `json:"airConditioningInd,omitempty"`   // AirConditioningInd: True if vehicle has air conditioning
	BaggageQuantity      int32             `json:"baggageQuantity,omitempty"`      // BaggageQuantity: Baggage Quantity that is able to fit into the car with passengers
	DoorCount            string            `json:"doorCount,omitempty"`            // DoorCount: The number of doors for the vehicle
	PassengerQuantity    string            `json:"passengerQuantity,omitempty"`    // PassengerQuantity:
}

// VehicleCharges is an object.
type VehicleCharges struct {
	Type                        string           `json:"@type,omitempty"`                       // Type:
	Calculation                 []Calculation    `json:"Calculation,omitempty"`                 // Calculation:
	RateCategory                RateCategoryEnum `json:"RateCategory,omitempty"`                // RateCategory: Rate Category
	VehicleChargePurposeCode    *Code            `json:"VehicleChargePurposeCode,omitempty"`    // VehicleChargePurposeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	VehicleCoverageTypeCode     *Code            `json:"VehicleCoverageTypeCode,omitempty"`     // VehicleCoverageTypeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	Description                 string           `json:"description,omitempty"`                 // Description:
	GuaranteedInd               bool             `json:"guaranteedInd,omitempty"`               // GuaranteedInd:
	IncludedInBaseRateInd       bool             `json:"includedInBaseRateInd,omitempty"`       // IncludedInBaseRateInd: If true the Vehicle Charge has been included in the base rate of the Vehicle price
	IncludedInEstimatedTotalInd bool             `json:"includedInEstimatedTotalInd,omitempty"` // IncludedInEstimatedTotalInd:
	PayNowInd                   bool             `json:"payNowInd,omitempty"`                   // PayNowInd: If true the vehicle charge must be paid now and is included in the totalPrice calculation
	RateChargeInfo              string           `json:"rateChargeInfo,omitempty"`              // RateChargeInfo:
	RatePeriod                  RatePeriodEnum   `json:"ratePeriod,omitempty"`                  // RatePeriod: The time period for a rate such as daily, weekly, monthly
	TaxInclusiveInd             bool             `json:"taxInclusiveInd,omitempty"`             // TaxInclusiveInd:
}

// VehicleDateLocation is an object.
type VehicleDateLocation struct {
	Type         string             `json:"@type,omitempty"` // Type:
	RentalPickup RentalPickupReturn `json:"RentalPickup"`    // RentalPickup:
	RentalReturn RentalPickupReturn `json:"RentalReturn"`    // RentalReturn:
}

// VehicleDetail is an object.
type VehicleDetail struct {
	Type                 string               `json:"@type"`                          // Type:
	FuelTypeCode         *Code                `json:"FuelTypeCode,omitempty"`         // FuelTypeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	Image                *Image               `json:"Image,omitempty"`                // Image: URL of the image
	TransmissionTypeCode *Code                `json:"TransmissionTypeCode,omitempty"` // TransmissionTypeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	VehicleCategoryCode  *Code                `json:"VehicleCategoryCode,omitempty"`  // VehicleCategoryCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	VehicleClassCode     *Code                `json:"VehicleClassCode,omitempty"`     // VehicleClassCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	VehicleDateLocation  *VehicleDateLocation `json:"VehicleDateLocation,omitempty"`  // VehicleDateLocation:
	VehicleMakeModel     *VehicleMakeModel    `json:"VehicleMakeModel,omitempty"`     // VehicleMakeModel: The make and model of the vehicle along with a description
	VehicleSizeCode      *Code                `json:"VehicleSizeCode,omitempty"`      // VehicleSizeCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	AirConditioningInd   bool                 `json:"airConditioningInd,omitempty"`   // AirConditioningInd: True if vehicle has air conditioning
	BaggageQuantity      int32                `json:"baggageQuantity,omitempty"`      // BaggageQuantity: Baggage Quantity that is able to fit into the car with passengers
	DoorCount            string               `json:"doorCount,omitempty"`            // DoorCount: The number of doors for the vehicle
	PassengerQuantity    string               `json:"passengerQuantity,omitempty"`    // PassengerQuantity:
}

// VehicleMakeModel is an object. The make and model of the vehicle along with a description
type VehicleMakeModel struct {
	Code                  string `json:"code,omitempty"`                  // Code: Assigned Type: c-1100:StringTiny
	OperatingSupplierCode string `json:"operatingSupplierCode,omitempty"` // OperatingSupplierCode: Assigned Type: c-1100:StringTiny
	OperatingSupplierName string `json:"operatingSupplierName,omitempty"` // OperatingSupplierName: Assigned Type: c-1100:StringTiny
	SupplierReference     string `json:"supplierReference,omitempty"`     // SupplierReference: Assigned Type: ctvh-1100:SupplierReference
	Value                 string `json:"value,omitempty"`                 // Value:
	VendorCode            string `json:"vendorCode,omitempty"`            // VendorCode: Assigned Type: c-1100:StringTiny
}

// VehiclePrice is an object.
type VehiclePrice struct {
	Type              string               `json:"@type,omitempty"`             // Type:
	ApproximateRate   *ApproximateRate     `json:"ApproximateRate,omitempty"`   // ApproximateRate:
	CustomerLoyalty   *CustomerLoyalty     `json:"CustomerLoyalty,omitempty"`   // CustomerLoyalty: Specifies the ID for the membership program.
	RateAvailability  RateAvailabilityEnum `json:"RateAvailability,omitempty"`  // RateAvailability: Options are available to sell, need to call, or closed
	RateDescription   []TextBlock          `json:"RateDescription,omitempty"`   // RateDescription:
	RateDistance      *RateDistance        `json:"RateDistance,omitempty"`      // RateDistance: Rate for the period defined by the attributes
	RateQualifier     RateQualifierEnum    `json:"RateQualifier,omitempty"`     // RateQualifier: A closed enumeration of possible rate qualifiers for vehicle rental
	SupplierRate      SupplierRate         `json:"SupplierRate"`                // SupplierRate:
	Id                string               `json:"id,omitempty"`                // Id: Internal ID
	RateGuaranteedInd bool                 `json:"rateGuaranteedInd,omitempty"` // RateGuaranteedInd: Assigned Type: c-1100:OptionalIndicator
	RatePeriod        RatePeriodEnum       `json:"ratePeriod,omitempty"`        // RatePeriod: The time period for a rate such as daily, weekly, monthly
	RateSource        string               `json:"rateSource,omitempty"`        // RateSource: Assigned Type: c-1100:StringTiny
}

// VehicleSearchModifiers is an object.
type VehicleSearchModifiers struct {
	Type                  string `json:"@type,omitempty"`                 // Type:
	AvailableRatesOnlyInd bool   `json:"availableRatesOnlyInd,omitempty"` // AvailableRatesOnlyInd: Assigned Type: c-1100:OptionalIndicator
	SellableRatesOnlyInd  bool   `json:"sellableRatesOnlyInd,omitempty"`  // SellableRatesOnlyInd: Assigned Type: c-1100:OptionalIndicator
}

// vehicleTravelCriteriaCurrencyCodePattern is the validation pattern for VehicleTravelCriteria.CurrencyCode
var vehicleTravelCriteriaCurrencyCodePattern = regexp.MustCompile(`[a-zA-Z]{3}`)

// VehicleTravelCriteria is an object.
type VehicleTravelCriteria struct {
	Type                    string                   `json:"@type,omitempty"`                   // Type:
	BookingSource           *Code                    `json:"BookingSource,omitempty"`           // BookingSource: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	CustomerLoyalty         []CustomerLoyalty        `json:"CustomerLoyalty,omitempty"`         // CustomerLoyalty:
	DriverInfo              DriverInfo               `json:"DriverInfo"`                        // DriverInfo: Basic information (metadata) about the intended driver of the vehicle
	NextResultReference     *NextResultReference     `json:"NextResultReference,omitempty"`     // NextResultReference: Container to return\/send additional search results
	PermittedVendors        *PermittedVendors        `json:"PermittedVendors,omitempty"`        // PermittedVendors:
	ProhibitedVendors       *ProhibitedVendors       `json:"ProhibitedVendors,omitempty"`       // ProhibitedVendors:
	RateCodeInfo            []RateCodeInfo           `json:"RateCodeInfo,omitempty"`            // RateCodeInfo:
	SearchVehicleAttributes *SearchVehicleAttributes `json:"SearchVehicleAttributes,omitempty"` // SearchVehicleAttributes: The physical vehical attrirbutes that a person can search on.
	SpecialEquipment        []SpecialEquipment       `json:"SpecialEquipment,omitempty"`        // SpecialEquipment:
	VehicleDateLocation     *VehicleDateLocation     `json:"VehicleDateLocation,omitempty"`     // VehicleDateLocation:
	VehicleMakeModel        []VehicleMakeModel       `json:"VehicleMakeModel,omitempty"`        // VehicleMakeModel:
	VehicleSearchModifiers  *VehicleSearchModifiers  `json:"VehicleSearchModifiers,omitempty"`  // VehicleSearchModifiers:
	CurrencyCode            string                   `json:"currencyCode,omitempty"`            // CurrencyCode: Requested currency code
	RateCategory            RateCategoryEnum         `json:"rateCategory,omitempty"`            // RateCategory: Rate Category
	RateGuaranteedInd       bool                     `json:"rateGuaranteedInd,omitempty"`       // RateGuaranteedInd: If true, this rate is guaranteed
	RatePeriod              RatePeriodEnum           `json:"ratePeriod,omitempty"`              // RatePeriod: The time period for a rate such as daily, weekly, monthly
	UnlimitedMileageInd     bool                     `json:"unlimitedMileageInd,omitempty"`     // UnlimitedMileageInd: If true, this rate includes unlimited mileage
}

// VendorCurrencyTotal is an object.
type VendorCurrencyTotal struct {
	Type           string             `json:"@type,omitempty"`          // Type:
	Base           float32            `json:"Base,omitempty"`           // Base: The price prior to all applicable taxes of a product such as the rate for a room or fare for a flight.
	CurrencyCode   *CurrencyCode      `json:"CurrencyCode,omitempty"`   // CurrencyCode: Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
	Fees           *Fees              `json:"Fees,omitempty"`           // Fees:
	Taxes          *Taxes             `json:"Taxes,omitempty"`          // Taxes:
	Total          float32            `json:"Total,omitempty"`          // Total: Specifies the total price including base + taxes + fees
	ApproximateInd bool               `json:"approximateInd,omitempty"` // ApproximateInd: True if this amount has been converted from the original amount
	CurrencySource CurrencySourceEnum `json:"currencySource,omitempty"` // CurrencySource: The system requesting or returning the currency code specified in the attribute
}

// vendorLocationRentalLocationNumberPattern is the validation pattern for VendorLocation.RentalLocationNumber
var vendorLocationRentalLocationNumberPattern = regexp.MustCompile(`([0-9a-zA-Z]+)?`)

// VendorLocation is an object. The vendor's location number for pickup or return
type VendorLocation struct {
	Type                   string           `json:"@type"`                            // Type:
	AdditionalInstructions string           `json:"AdditionalInstructions,omitempty"` // AdditionalInstructions: Additional instructions regarding the vendor location
	Address                *Address         `json:"Address,omitempty"`                // Address:
	CounterLocationCode    *Code            `json:"CounterLocationCode,omitempty"`    // CounterLocationCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	Description            string           `json:"Description,omitempty"`            // Description: Detailed location information on where to pick up and return a vehicle
	Directions             string           `json:"Directions,omitempty"`             // Directions: Directions for collecting the vehicle
	OperationTimes         []OperationTimes `json:"OperationTimes,omitempty"`         // OperationTimes:
	RentalLocationCode     *Code            `json:"RentalLocationCode,omitempty"`     // RentalLocationCode: Any code used to specify an item, for example a type of traveler, service code, room amenity, etc.
	ShuttleService         string           `json:"ShuttleService,omitempty"`         // ShuttleService: Information on shuttle service
	Telephone              []Telephone      `json:"Telephone,omitempty"`              // Telephone:
	Code                   string           `json:"code,omitempty"`                   // Code:
	RentalLocationName     string           `json:"rentalLocationName,omitempty"`     // RentalLocationName:
	RentalLocationNumber   string           `json:"rentalLocationNumber,omitempty"`   // RentalLocationNumber:
	VendorCode             string           `json:"vendorCode,omitempty"`             // VendorCode:
}

// WaiverCode is an object.
type WaiverCode struct {
	ReasonCode int32  `json:"reasonCode,omitempty"` // ReasonCode: A code assigned to identify the reason for disruption
	Value      string `json:"value,omitempty"`      // Value:
}

// Warning is an object.
type Warning struct {
	Type          string          `json:"@type"`                   // Type:
	Message       string          `json:"Message,omitempty"`       // Message: The Travelport standardized error or warning message
	NameValuePair []NameValuePair `json:"NameValuePair,omitempty"` // NameValuePair:
	StatusCode    int32           `json:"StatusCode,omitempty"`    // StatusCode: Http standard response code
}

// WarningDetail is an object.
type WarningDetail struct {
	Type              string          `json:"@type"`                       // Type:
	Message           string          `json:"Message,omitempty"`           // Message: The Travelport standardized error or warning message
	NameValuePair     []NameValuePair `json:"NameValuePair,omitempty"`     // NameValuePair:
	SourceCode        string          `json:"SourceCode,omitempty"`        // SourceCode: The error or warning code returned by the source airline or host system
	SourceDescription string          `json:"SourceDescription,omitempty"` // SourceDescription: The error or warning message as it is returned by the source airline or host system
	SourceID          string          `json:"SourceID"`                    // SourceID: The identifier of the source system sending the error or warning
	StatusCode        int32           `json:"StatusCode,omitempty"`        // StatusCode: Http standard response code
}
