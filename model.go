package travelport

import "time"

// Credit card code
type AcceptedCreditCard struct {
	Value string `json:"value,omitempty"`
	// Country Code ISO
	ProcessingCountry string `json:"processingCountry,omitempty"`
}

// Detail of the travel agency that issues the ticket
type AgencyInfo struct {
	// Ticketed date
	TicketedDate string `json:"ticketedDate,omitempty"`
	// Name of the Agency
	Name string `json:"name"`
	// Place of the agency
	Place string `json:"place"`
	// Ticketing PCC
	TicketingPCC string `json:"ticketingPCC,omitempty"`
	// Agency code
	Code string `json:"code,omitempty"`
	// Sales type
	SalesType string `json:"salesType,omitempty"`
	// Ticketing country
	TicketingCountry string `json:"ticketingCountry"`
	// Ticketing city
	TicketingCity string `json:"ticketingCity"`
}

type Accounting struct {
	// Accounting data type
	DataType string `json:"dataType,omitempty"`
	// Accounting template
	Template      string          `json:"template,omitempty"`
	NameValuePair []NameValuePair `json:"NameValuePair,omitempty"`
	Type_         string          `json:"@type"`
	Id            string          `json:"id,omitempty"`
	// Accounting reference
	AccountingRef string      `json:"AccountingRef,omitempty"`
	Identifier    *Identifier `json:"Identifier,omitempty"`
}

type AccountingId struct {
	Type_ string `json:"@type"`
	Id    string `json:"id,omitempty"`
	// Accounting reference
	AccountingRef string      `json:"AccountingRef,omitempty"`
	Identifier    *Identifier `json:"Identifier,omitempty"`
}

type AdditionalDetail struct {
	Type_ string `json:"@type,omitempty"`
	// OTA Code
	DetailType string `json:"detailType,omitempty"`
	// Partner code
	Code        string                   `json:"code,omitempty"`
	Amount      *CurrencyAmount          `json:"Amount,omitempty"`
	Description *TextTitleAndDescription `json:"Description,omitempty"`
}

type AdditionalDetails struct {
	Type_            string             `json:"@type,omitempty"`
	AdditionalDetail []AdditionalDetail `json:"AdditionalDetail"`
}

type Address struct {
	Type_ string `json:"@type"`
	// Internally referenced id
	Id       string               `json:"id,omitempty"`
	BldgRoom *AddressBldgRoom     `json:"BldgRoom,omitempty"`
	Number   *AddressStreetNumber `json:"Number,omitempty"`
	// May contain the street number when the Street number element is missing.
	Street string `json:"Street,omitempty"`
	// Additional address line details
	AddressLine []string `json:"AddressLine,omitempty"`
	// City (e.g., Dublin), town, or postal station (i.e., a postal service territory, often used in a military address).
	City string `json:"City"`
	// County or Region Name (e.g., Fairfax).
	County    string     `json:"County,omitempty"`
	StateProv *StateProv `json:"StateProv,omitempty"`
	Country   *Country   `json:"Country,omitempty"`
	// Post Office Code number.
	PostalCode string `json:"PostalCode,omitempty"`
	// The name of the company or person to be addressed
	Addressee string           `json:"Addressee,omitempty"`
	Role      *EnumAddressRole `json:"role,omitempty"`
}

// Address with building and room number
type AddressBldgRoom struct {
	Value string `json:"value,omitempty"`
	// When true, the information is a building name. When false, it is an apartment or room #
	BuldingInd bool `json:"buldingInd,omitempty"`
}

// The street number alone is the numerical number that precedes the street name in the address
type AddressStreetNumber struct {
	Value string `json:"value,omitempty"`
	// Street Number Suffix
	StreetNmbrSuffix string `json:"streetNmbrSuffix,omitempty"`
	// Dircetion of the Street
	StreetDirection string `json:"streetDirection,omitempty"`
	// RuralRoute Number
	RuralRouteNmbr string `json:"ruralRouteNmbr,omitempty"`
	// PO Box Number
	PoBox string `json:"po_Box,omitempty"`
}

type AgeQualifying struct {
	Type_ string `json:"@type,omitempty"`
	// MinAge: The minimum age to qualify for AgeQualifyingCode.
	MinAge int32 `json:"minAge,omitempty"`
	// Max Age: The maximum age to qualify for AgeQualifyingCode.
	MaxAge int32 `json:"maxAge,omitempty"`
	// The age bucket
	AgeBucket string `json:"ageBucket,omitempty"`
	// The number of age qualifying
	Count int32 `json:"count,omitempty"`
	// Enter 10 for an adult or 08 for a child
	AgeQualifyingCode string `json:"ageQualifyingCode,omitempty"`
}

type AgencyServiceFee struct {
	// The service fee expiry date. Once expiry date has been reached, the service fee information will only be stored in the ReservationReceipt
	ExpiryDate time.Time `json:"ExpiryDate,omitempty"`
	// The description of the service fee
	Description           string          `json:"Description,omitempty"`
	Amount                *CurrencyAmount `json:"Amount"`
	Tax                   []Tax           `json:"Tax,omitempty"`
	RelatedDocumentNumber *DocumentNumber `json:"RelatedDocumentNumber,omitempty"`
	// Reference to a Traveler within the Reservation that this service fee applies to
	TravelerRef string `json:"TravelerRef,omitempty"`
	// Reference to an Offer within the Reservation that this service fee applies to
	OfferRef string `json:"OfferRef,omitempty"`
	Type_    string `json:"@type"`
	// Unique id for this object within a message
	Id string `json:"id,omitempty"`
}

type AgencyServiceFeeId struct {
	Type_ string `json:"@type"`
	// Unique id for this object within a message
	Id string `json:"id,omitempty"`
}

type AgencyServiceFeeIdentifier struct {
	// Unique id for this object within a message
	Id string `json:"id,omitempty"`
}

type AlternateContact struct {
	Type_ string `json:"@type,omitempty"`
	Id    string `json:"id,omitempty"`
	// Contact type value
	ContactType string `json:"contactType,omitempty"`
	// Relation value
	Relation   string      `json:"relation,omitempty"`
	PersonName *PersonName `json:"PersonName"`
	Address    []Address   `json:"Address,omitempty"`
	Telephone  []Telephone `json:"Telephone,omitempty"`
	Email      []Email     `json:"Email,omitempty"`
	// This is the contact in case of an emergency
	EmergencyInd bool `json:"emergencyInd,omitempty"`
	// This is the default contact
	DefaultInd bool `json:"defaultInd,omitempty"`
}

type AmenitySurcharges struct {
	Type_ string `json:"@type"`
	// A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	TotalSurcharges float32 `json:"TotalSurcharges,omitempty"`
	// if true, the surcharge amounts are approximate
	ApproximateInd bool `json:"approximateInd,omitempty"`
}

type AmenitySurchargesDetail struct {
	Surcharge []Surcharge `json:"Surcharge"`
	Type_     string      `json:"@type"`
	// A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	TotalSurcharges float32 `json:"TotalSurcharges,omitempty"`
	// if true, the surcharge amounts are approximate
	ApproximateInd bool `json:"approximateInd,omitempty"`
}

type Amount struct {
	Type_          string              `json:"@type,omitempty"`
	CurrencySource *CurrencySourceEnum `json:"currencySource,omitempty"`
	CurrencyCode   *CurrencyCode       `json:"CurrencyCode,omitempty"`
	// The price prior to all applicable taxes of a product such as the rate for a room or fare for a flight.
	Base  float32 `json:"Base,omitempty"`
	Taxes *Taxes  `json:"Taxes,omitempty"`
	Fees  *Fees   `json:"Fees,omitempty"`
	// Specifies the total price including base + taxes + fees
	Total float32 `json:"Total,omitempty"`
	// True if this amount has been converted from the original amount
	ApproximateInd bool `json:"approximateInd,omitempty"`
}

type AmountPercent struct {
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type AmountPercentAmount struct {
	Amount      *CurrencyAmount `json:"Amount,omitempty"`
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type AmountPercentPercent struct {
	// Percent amount of commission
	Percent     float32         `json:"Percent,omitempty"`
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type AncillaryOfferingId struct {
	Type_ string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message
	CatalogOfferingRef   string      `json:"CatalogOfferingRef,omitempty"`
	AncillaryOfferingRef string      `json:"AncillaryOfferingRef,omitempty"`
	Identifier           *Identifier `json:"Identifier,omitempty"`
}

type AncillaryOfferingIdentifier struct {
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message
	CatalogOfferingRef   string      `json:"CatalogOfferingRef,omitempty"`
	AncillaryOfferingRef string      `json:"AncillaryOfferingRef,omitempty"`
	Identifier           *Identifier `json:"Identifier,omitempty"`
}

type AppliesTo struct {
	Type_ string `json:"@type"`
}

type AppliesToOffer struct {
	OfferIdentifier []OfferIdentifier `json:"OfferIdentifier,omitempty"`
	Type_           string            `json:"@type"`
}

type AppliesToOfferProduct struct {
	OfferIdentifier   *OfferIdentifier    `json:"OfferIdentifier,omitempty"`
	ProductIdentifier []ProductIdentifier `json:"ProductIdentifier,omitempty"`
	Type_             string              `json:"@type"`
}

type Attraction struct {
	Type_ string `json:"@type,omitempty"`
	// The name of the attraction
	Name     string    `json:"name"`
	Distance *Distance `json:"Distance,omitempty"`
}

type AuthenticationVerification struct {
	Type_ string `json:"@type,omitempty"`
	// Note: This contains a key required to retrieve the full payment instrument details compliant with PCI DSS standards.
	EncryptionKey string `json:"encryptionKey,omitempty"`
	// Developer: This contains a reference to the key generation method being used - this is NOT the key value.
	EncryptionKeyMethod string `json:"encryptionKeyMethod,omitempty"`
	// OpenTravel Best Practice: Encryption Method: When using the OpenTravel Encryption element, it is RECOMMENDED that all trading partners be informed of all encryption methods being used in advance of implementation to ensure message processing compatibility.
	EncryptionMethod string `json:"encryptionMethod,omitempty"`
	// Encrypted value
	EncryptedValue string `json:"encryptedValue,omitempty"`
	// Masked Value
	Mask string `json:"mask,omitempty"`
	// Token value
	Token string `json:"token,omitempty"`
	// Developer: This contains a provider ID if multiple providers are used for secure information exchange.
	TokenProviderID      string                       `json:"tokenProviderID,omitempty"`
	AuthenticationMethod *EncryptionTokenTypeAuthEnum `json:"authenticationMethod,omitempty"`
	// Don't use this unless it is REALLY ok to not use encryption. Non-secure (plain text) value.
	PlainText    string        `json:"PlainText,omitempty"`
	ErrorWarning *ErrorWarning `json:"ErrorWarning,omitempty"`
}

type BedConfiguration struct {
	// The number of bed of this type and size in the room
	Quantity int32 `json:"quantity,omitempty"`
	// Configuration of bed(s) in room.
	BedType string `json:"bedType,omitempty"`
	// Size of bed(s) in the room.
	Size string `json:"size,omitempty"`
}

type BrandId struct {
	Type_ string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message
	BrandRef   string      `json:"BrandRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type BuildFromCatalogOfferingHospitality struct {
	Type_                     string      `json:"@type,omitempty"`
	CatalogOfferingIdentifier *Identifier `json:"CatalogOfferingIdentifier,omitempty"`
	// Number of rooms required.
	NumberOfRooms int32 `json:"NumberOfRooms,omitempty"`
}

type BuildFromProperties struct {
	Type_                string      `json:"@type,omitempty"`
	PropertiesIdentifier *Identifier `json:"PropertiesIdentifier,omitempty"`
	PropertyInfoIds      []string    `json:"PropertyInfoIds,omitempty"`
}

type CancelPenalty struct {
	Type_        string            `json:"@type,omitempty"`
	Deadline     *Deadline         `json:"Deadline,omitempty"`
	HotelPenalty *HotelPenalty     `json:"HotelPenalty,omitempty"`
	Refundable   *YesNoUnknownEnum `json:"Refundable,omitempty"`
}

type Cancellation struct {
	Type_ string `json:"@type"`
}

type CancellationHold struct {
	Locator *Locator `json:"Locator,omitempty"`
	Type_   string   `json:"@type"`
}

type CardNumber struct {
	Type_ string `json:"@type,omitempty"`
	// Note: This contains a key required to retrieve the full payment instrument details compliant with PCI DSS standards.
	EncryptionKey string `json:"encryptionKey,omitempty"`
	// Developer: This contains a reference to the key generation method being used - this is NOT the key value.
	EncryptionKeyMethod string `json:"encryptionKeyMethod,omitempty"`
	// OpenTravel Best Practice: Encryption Method: When using the OpenTravel Encryption element, it is RECOMMENDED that all trading partners be informed of all encryption methods being used in advance of implementation to ensure message processing compatibility.
	EncryptionMethod string `json:"encryptionMethod,omitempty"`
	// Encrypted value
	EncryptedValue string `json:"encryptedValue,omitempty"`
	// Masked Value
	Mask string `json:"mask,omitempty"`
	// Token value
	Token string `json:"token,omitempty"`
	// Developer: This contains a provider ID if multiple providers are used for secure information exchange.
	TokenProviderID      string                       `json:"tokenProviderID,omitempty"`
	AuthenticationMethod *EncryptionTokenTypeAuthEnum `json:"authenticationMethod,omitempty"`
	// Don't use this unless it is REALLY ok to not use encryption. Non-secure (plain text) value.
	PlainText    string        `json:"PlainText,omitempty"`
	ErrorWarning *ErrorWarning `json:"ErrorWarning,omitempty"`
}

type CatalogOffering struct {
	ProductOptions     []ProductOptions    `json:"ProductOptions"`
	Price              *PriceDetail        `json:"Price"`
	TermsAndConditions *TermsAndConditions `json:"TermsAndConditions,omitempty"`
	Type_              string              `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message
	CatalogOfferingRef string      `json:"CatalogOfferingRef,omitempty"`
	Identifier         *Identifier `json:"Identifier,omitempty"`
}

type CatalogOfferingHospitality struct {
	// Sequence determines the order in which the offer shall be displayed to the customer
	Sequence  int32      `json:"Sequence,omitempty"`
	NextSteps *NextSteps `json:"NextSteps,omitempty"`
	// If true the offering/Offer has limited availability
	LimitedAvailabilityInd bool                `json:"limitedAvailabilityInd,omitempty"`
	ProductOptions         []ProductOptions    `json:"ProductOptions"`
	Price                  *PriceDetail        `json:"Price"`
	TermsAndConditions     *TermsAndConditions `json:"TermsAndConditions,omitempty"`
}

type CatalogOfferingId struct {
	Type_ string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message
	CatalogOfferingRef string      `json:"CatalogOfferingRef,omitempty"`
	Identifier         *Identifier `json:"Identifier,omitempty"`
}

type CatalogOfferingIdentifier struct {
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message
	CatalogOfferingRef string      `json:"CatalogOfferingRef,omitempty"`
	Identifier         *Identifier `json:"Identifier,omitempty"`
}

type CatalogOfferings struct {
	CatalogOffering   []CatalogOffering     `json:"CatalogOffering"`
	AncillaryOffering []AncillaryOfferingId `json:"AncillaryOffering,omitempty"`
	Type_             string                `json:"@type"`
	// Local indentifier within a given message for this object.
	Id         string      `json:"id,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type CatalogOfferingsId struct {
	Type_ string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id         string      `json:"id,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type CatalogOfferingsIdentifier struct {
	// Local indentifier within a given message for this object.
	Id         string      `json:"id,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type CatalogOfferingsQueryBuildFromProperties struct {
	Type_                               string               `json:"@type"`
	BuildFromCatalogOfferingHospitality *BuildFromProperties `json:"BuildFromCatalogOfferingHospitality,omitempty"`
}

type CatalogOfferingsQueryBuildFromPropertiesWrapper struct {
	CatalogOfferingsQueryBuildFromProperties *CatalogOfferingsQueryBuildFromProperties `json:"CatalogOfferingsQueryBuildFromProperties,omitempty"`
}

type ChangeFeeCollectionMethod struct {
	Value *ChangeFeeMethodEnum `json:"value,omitempty"`
	// The code value
	Code string `json:"code"`
	// The subcode value
	SubCode string `json:"subCode,omitempty"`
	// The description value
	Description string `json:"description,omitempty"`
	// if true, the change fee will be issued as a separate transaction to the residual amount
	ChangeFeeIssuedSeparatelyInd bool `json:"changeFeeIssuedSeparatelyInd,omitempty"`
	// If true, the tax  on the fee will be included in the base fee amount and sent as a single value to the supplier for fulfilment
	TaxIncludedInBaseAmountInd bool `json:"taxIncludedInBaseAmountInd,omitempty"`
}

type ChangeFeeMethodEnum struct {
	Value *ChangeFeeMethodEnumBase `json:"value,omitempty"`
}

type CheckInOutPolicy struct {
	Type_ string `json:"@type,omitempty"`
	// Check-in time
	CheckInTime string `json:"checkInTime"`
	// Check-out time
	CheckOutTime string `json:"checkOutTime"`
	// Minimum age of guest checking in or out
	MinimumAge  int32                     `json:"minimumAge,omitempty"`
	Description []TextTitleAndDescription `json:"Description,omitempty"`
}

// The class of service
type ClassOfServiceAvailability struct {
	Value string `json:"value,omitempty"`
	// The class of service number value
	Number int32                   `json:"number,omitempty"`
	Status *AvailabilityStatusEnum `json:"status,omitempty"`
}

type Commission struct {
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type CommissionAmount struct {
	Amount      *CurrencyAmount `json:"Amount,omitempty"`
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type CommissionPercent struct {
	// Percent amount of commission
	Percent     float32         `json:"Percent,omitempty"`
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type Commissions struct {
	Type_                 string                  `json:"@type,omitempty"`
	TravelerIdentifierRef []TravelerIdentifierRef `json:"TravelerIdentifierRef,omitempty"`
	Commission            *Commission             `json:"Commission"`
	ApplyTo               *CommissionApplyEnum    `json:"ApplyTo,omitempty"`
}

// Textual information.
type Comment struct {
	Value string `json:"value,omitempty"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Title
	Name string `json:"name,omitempty"`
	// Language code using ISO-639 standard
	Language string `json:"language,omitempty"`
}

type Comments struct {
	Type_   string    `json:"@type,omitempty"`
	Comment []Comment `json:"Comment,omitempty"`
}

// Identifies a company by name.
type CompanyName struct {
	Value string `json:"value,omitempty"`
	// Use this id to internally identify this company in NextSteps
	Id string `json:"id,omitempty"`
	// The division name or ID with which the contact is associated
	Division string `json:"division,omitempty"`
	// The department name or ID with which the contact is associated
	Department string `json:"department,omitempty"`
	// Used to provide the company common name
	ShortName string `json:"shortName,omitempty"`
	// Identifies a company by the company code
	Code string `json:"code,omitempty"`
	// Identifies the context of the identifying code,such as DUNS,IATA or internal code
	CodeContext string `json:"codeContext,omitempty"`
	// The system(s) that maintain the data
	SystemOfRecord []string `json:"systemOfRecord,omitempty"`
}

type Confirmation struct {
	Type_ string `json:"@type"`
}

type ConfirmationHold struct {
	Locator                   *Locator                      `json:"Locator"`
	OfferStatus               *OfferStatus                  `json:"OfferStatus,omitempty"`
	ShoppingCartProductStatus *ShoppingCartProductStatusAir `json:"ShoppingCartProductStatus,omitempty"`
	Type_                     string                        `json:"@type"`
}

// A conversion metric from standard to another with the contextual authority such as IATA, OAG, ISO, etc.
type ConversionRate struct {
	Value float32 `json:"value,omitempty"`
	// Rate authority
	RateAuthority string `json:"rateAuthority,omitempty"`
	// Rate as of
	RateAsOf time.Time `json:"rateAsOf,omitempty"`
}

// ISO 3166 code for a country with optional name
type Country struct {
	Value string `json:"value,omitempty"`
	// Use this id to internally identify this country in NextSteps
	Id string `json:"id,omitempty"`
	// The name or code of a country
	Name string `json:"name,omitempty"`
	// The source of a code
	CodeContext string `json:"codeContext,omitempty"`
}

type CumulativeValue struct {
	Type_          string              `json:"@type,omitempty"`
	CurrencySource *CurrencySourceEnum `json:"currencySource,omitempty"`
	CurrencyCode   *CurrencyCode       `json:"CurrencyCode,omitempty"`
	// The price prior to all applicable taxes of a product such as the rate for a room or fare for a flight.
	Base  float32 `json:"Base,omitempty"`
	Taxes *Taxes  `json:"Taxes,omitempty"`
	Fees  *Fees   `json:"Fees,omitempty"`
	// Specifies the total price including base + taxes + fees
	Total float32 `json:"Total,omitempty"`
	// True if this amount has been converted from the original amount
	ApproximateInd bool `json:"approximateInd,omitempty"`
}

// A monetary amount, up to 4 decimal places. Decimal place needs to be included.
type CurrencyAmount struct {
	Value float32 `json:"value,omitempty"`
	// An ISO 4217 alpha character code that specifies a money unit
	Code string `json:"code,omitempty"`
	// Minor units are a mechanism for expressing the relationship between a major currency unit and its corresponding minor currency unit.
	MinorUnit      int32               `json:"minorUnit,omitempty"`
	CurrencySource *CurrencySourceEnum `json:"currencySource,omitempty"`
	// True if the currency amount has been converted from the original amount
	ApproximateInd bool `json:"approximateInd,omitempty"`
}

// Currency codes are the three-letter alphabetic codes that represent the various currencies used throughout the world.
type CurrencyCode struct {
	Value string `json:"value,omitempty"`
	// Currency code authority
	CodeAuthority string `json:"codeAuthority,omitempty"`
	// Currency code decimal place
	DecimalPlace int32 `json:"decimalPlace,omitempty"`
	// Currency code decimal authority
	DecimalAuthority string `json:"decimalAuthority,omitempty"`
}

type CurrencyRateConversion struct {
	Type_          string          `json:"@type,omitempty"`
	SourceCurrency *CurrencyCode   `json:"SourceCurrency"`
	TargetCurrency *CurrencyCode   `json:"TargetCurrency"`
	ConversionRate *ConversionRate `json:"ConversionRate"`
}

// Specifies the ID for the membership program.
type CustomerLoyalty struct {
	Value string `json:"value,omitempty"`
	// Customer Loyality Id
	Id string `json:"id,omitempty"`
	// Numeric Priority Code
	Priority int32 `json:"priority,omitempty"`
	// Specifies an identifier to indicate the company owner of the loyalty program
	ProgramId string `json:"programId,omitempty"`
	// Supplier's loyalty program name such as Frontier-EarlyReturns
	ProgramName string `json:"programName,omitempty"`
	// The kind of supplier of a loyalty program
	SupplierType string `json:"supplierType,omitempty"`
	// Supplier of a loyalty program
	Supplier string `json:"supplier,omitempty"`
	// Customer Loyalty tier level
	Tier string `json:"tier,omitempty"`
	// The list of suppliers that the CustomerLoyalty number is shared.
	ShareWithSupplier []string `json:"shareWithSupplier,omitempty"`
	// The card holder name
	CardHolderName string `json:"cardHolderName,omitempty"`
	// Customer loyalty number has been validated by the supplier
	ValidatedInd bool `json:"validatedInd,omitempty"`
}

type CustomerLoyaltyCredit struct {
	Type_           string           `json:"@type,omitempty"`
	CustomerLoyalty *CustomerLoyalty `json:"CustomerLoyalty"`
	// Represents the amount of award credit awarded for this offer/offering. Award credit can be used for purchasing goods and services through a customer loyalty program
	Earned string `json:"Earned"`
	// Represents the amount of status credit awarded for this offer/offering. Status credit allow a customer to move up the ladder of a customer loyalty program
	Status string `json:"Status"`
}

// Used to identify the effective date and/or expiration date.
type DateEffectiveExpire struct {
	// Indicates the starting date.
	Effective string `json:"effective,omitempty"`
	// Indicates the ending date.
	Expire string `json:"expire,omitempty"`
	// When true, indicates that the ExpireDate is the first day after the applicable period (e.g. when expire date is Oct 15  the last date of the period is Oct 14).
	ExpireDateExclusiveInd bool `json:"expireDateExclusiveInd,omitempty"`
}

// Indicates the expiry date.
type DateOrDateWindows struct {
	// A specific date. When used with a windows must fall between start and end.
	Specific string `json:"specific,omitempty"`
	// The earliest and latest dates acceptable for the start date.
	Start string `json:"start,omitempty"`
	// The earliest and latest dates acceptable for the end date.
	End string `json:"end,omitempty"`
	// Duration from  start date.
	Duration     string            `json:"duration,omitempty"`
	DurationUnit *DurationUnitEnum `json:"durationUnit,omitempty"`
}

// Specifies the begin and end date of an event
type DateRange struct {
	// Specifies the start date for an event, such as a booking
	Start string `json:"start"`
	// Specifies the end date an event, such as a booking
	End string `json:"end"`
}

type Deadline struct {
	Type_        string             `json:"@type,omitempty"`
	SpecificDate *DateOrDateWindows `json:"SpecificDate"`
	// Local time of the property
	Time string `json:"Time,omitempty"`
}

type Deposit struct {
	Type_ string `json:"@type"`
	// The date and time the deposit is due
	Date time.Time `json:"Date,omitempty"`
	// If present and true, the date is when the remainder of the deposit is due
	RemainderInd bool `json:"remainderInd,omitempty"`
}

type DepositAmount struct {
	CurrencyAmount *CurrencyAmount `json:"CurrencyAmount"`
	Type_          string          `json:"@type"`
	// The date and time the deposit is due
	Date time.Time `json:"Date,omitempty"`
	// If present and true, the date is when the remainder of the deposit is due
	RemainderInd bool `json:"remainderInd,omitempty"`
}

type DepositNumberOfNights struct {
	// The number of nights that must be paid for by deposit
	NumberOfNights int32  `json:"NumberOfNights"`
	Type_          string `json:"@type"`
	// The date and time the deposit is due
	Date time.Time `json:"Date,omitempty"`
	// If present and true, the date is when the remainder of the deposit is due
	RemainderInd bool `json:"remainderInd,omitempty"`
}

type DepositPercent struct {
	// The percentage of the price that must be paid for by deposit
	Percent float32 `json:"Percent,omitempty"`
	Type_   string  `json:"@type"`
	// The date and time the deposit is due
	Date time.Time `json:"Date,omitempty"`
	// If present and true, the date is when the remainder of the deposit is due
	RemainderInd bool `json:"remainderInd,omitempty"`
}

type DepositPolicy struct {
	Type_   string    `json:"@type,omitempty"`
	Deposit []Deposit `json:"Deposit"`
}

type DestinationPurpose struct {
	Type_       string           `json:"@type,omitempty"`
	Destination *DestinationEnum `json:"destination"`
	Purpose     *PurposeEnum     `json:"purpose"`
}

type DisplaySequence struct {
	Type_ string `json:"@type,omitempty"`
	// The sequence the products are to be displayed for sequential date ordering
	DisplaySequence string `json:"displaySequence"`
	// Offer reference
	OfferRef string `json:"OfferRef"`
	// Product reference. If blank, display sequence applies to all products within the Offer.
	ProductRef string `json:"ProductRef,omitempty"`
	// Segment sequence, if blank, display sequence applies to all segments within the product
	Sequence int32 `json:"Sequence,omitempty"`
}

// A search radius
type Distance struct {
	// When using distance as a property search parameter, the maximum distance is 25 regardless of unit of distance
	Value          float64             `json:"value,omitempty"`
	UnitOfDistance *UnitOfDistanceEnum `json:"unitOfDistance,omitempty"`
}

type Document struct {
	Type_ string `json:"@type"`
	// The identifying number of the document
	Number                string                 `json:"Number,omitempty"`
	TravelerIdentifierRef *TravelerIdentifierRef `json:"TravelerIdentifierRef,omitempty"`
	Amount                *Amount                `json:"Amount,omitempty"`
	WaiverCode            *WaiverCode            `json:"WaiverCode,omitempty"`
	Commission            *Commission            `json:"Commission,omitempty"`
	CumulativeValue       *CumulativeValue       `json:"CumulativeValue,omitempty"`
	// Document issuing pcc
	IssuingPCC string `json:"IssuingPCC,omitempty"`
	// Document issuing IATA
	IssuingIATA string `json:"IssuingIATA,omitempty"`
	// Document issuing city
	IssuingCity string       `json:"IssuingCity,omitempty"`
	FiledAmount *FiledAmount `json:"FiledAmount,omitempty"`
}

type DocumentNumber struct {
	Value string `json:"value,omitempty"`
	// Document issuer
	DocumentIssuer string            `json:"documentIssuer,omitempty"`
	DocumentType   *DocumentTypeEnum `json:"documentType,omitempty"`
}

type DocumentOverrides struct {
	OfferIdentifier           *OfferIdentifier           `json:"OfferIdentifier,omitempty"`
	ProductIdentifier         *ProductIdentifier         `json:"ProductIdentifier,omitempty"`
	Commissions               []Commissions              `json:"Commissions,omitempty"`
	DestinationPurpose        []DestinationPurpose       `json:"DestinationPurpose,omitempty"`
	Restrictions              []Restrictions             `json:"Restrictions,omitempty"`
	TourCodes                 []TourCodes                `json:"TourCodes,omitempty"`
	ChangeFeeCollectionMethod *ChangeFeeCollectionMethod `json:"ChangeFeeCollectionMethod,omitempty"`
	NetRemitInfo              *NetRemitInfo              `json:"NetRemitInfo,omitempty"`
	Type_                     string                     `json:"@type"`
	// The reporting number.
	Id                   string      `json:"id,omitempty"`
	DocumentOverridesRef string      `json:"DocumentOverridesRef,omitempty"`
	Identifier           *Identifier `json:"Identifier,omitempty"`
}

type DocumentOverridesId struct {
	Type_ string `json:"@type"`
	// The reporting number.
	Id                   string      `json:"id,omitempty"`
	DocumentOverridesRef string      `json:"DocumentOverridesRef,omitempty"`
	Identifier           *Identifier `json:"Identifier,omitempty"`
}

type DocumentOverridesResponse struct {
	// Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
	TransactionId string `json:"transactionId,omitempty"`
	// Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TraceId                string                   `json:"traceId,omitempty"`
	Result                 *Result                  `json:"Result,omitempty"`
	Identifier             *Identifier              `json:"Identifier,omitempty"`
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"`
	DocumentOverrides      *DocumentOverridesId     `json:"DocumentOverrides,omitempty"`
}

type DocumentOverridesResponseWrapper struct {
	DocumentOverridesResponse *DocumentOverridesResponse `json:"DocumentOverridesResponse,omitempty"`
}

// Electronic email addresses, in IETF specified format.
type Email struct {
	Value string `json:"value,omitempty"`
	// Electronic email addresses, in IETF specified format.
	Id string `json:"id,omitempty"`
	// Type of the e-mail
	EmailType string `json:"emailType,omitempty"`
	// Assigned Type: c-1100:StringText
	Comment string `json:"comment,omitempty"`
	// Mime media type
	PreferredFormat string            `json:"preferredFormat,omitempty"`
	ShareMarketing  *YesNoInheritEnum `json:"shareMarketing,omitempty"`
	ShareSync       *YesNoInheritEnum `json:"shareSync,omitempty"`
	OptOutInd       *YesNoInheritEnum `json:"optOutInd,omitempty"`
	OptInStatus     *OptInStatusEnum  `json:"optInStatus,omitempty"`
	// The datetime of receiving the opt in notice
	OptInDate time.Time `json:"optInDate,omitempty"`
	// The datetime the opt out notice was received
	OptOutDate time.Time `json:"optOutDate,omitempty"`
	// If true, this is a valid email address that has been system verified via a successful email transmission.
	ValidInd bool `json:"validInd,omitempty"`
	// If true then the email address came from the provisioning process
	ProvisionedInd bool `json:"provisionedInd,omitempty"`
}

// Note this field is deprecated in Payment schema and should be passed in FormOfPaymentPaymentCardExtendPayment schema
type ExtendedPayment struct {
	// The number of installment payments to be charged by the payment card provider
	NumberOfInstallments int32 `json:"NumberOfInstallments"`
	// For Pagos Parceledos, specify the first installment amount. This will be the same currency as the payment
	FirstInstallment float32 `json:"FirstInstallment,omitempty"`
	// For Pagos Parceledos, specify the remaining amount to be charged that will be spread across the number of installments. This will be the same currency as the payment
	RemainingAmount float32 `json:"RemainingAmount,omitempty"`
	// For Pagos Parceledos the OTATOCode
	OTATOCode string `json:"OTATOCode,omitempty"`
}

type Emd struct {
	PersonName         *PersonName       `json:"PersonName"`
	ReservationLocator []SupplierLocator `json:"ReservationLocator,omitempty"`
	AgencyInfo         *AgencyInfo       `json:"AgencyInfo"`
	EMDSegment         []EmdSegment      `json:"EMDSegment"`
	TotalAmount        *TotalAmount      `json:"TotalAmount,omitempty"`
	FormOfPayment      *FormOfPayment    `json:"FormOfPayment"`
	// The BSP ESAC code assign for a void or refund transaction\\nThe BSP E
	ESAC                   string        `json:"ESAC,omitempty"`
	AssociatedTicketNumber *TicketNumber `json:"AssociatedTicketNumber,omitempty"`
	Restrictions           []string      `json:"Restrictions,omitempty"`
	Type_                  string        `json:"@type"`
	Id                     string        `json:"id,omitempty"`
	EMDRef                 string        `json:"EMDRef,omitempty"`
	Identifier             *Identifier   `json:"Identifier,omitempty"`
}

// A description of the ancillary with two description codes
type EmdDescription struct {
	Value string `json:"value,omitempty"`
	// A description of the ancillary with two description codes
	Code string `json:"code,omitempty"`
	// EMD number sub code
	SubCode string `json:"subCode,omitempty"`
	// Code context
	CodeContext string `json:"codeContext,omitempty"`
}

type EmdQueryUpdateEmd struct {
	Type_ string `json:"@type"`
	// Assigned Type: c-1100:AgencyCodeIATA
	AgencyCode string `json:"agencyCode,omitempty"`
	// Assigned Type: c-1100:StringTiny
	Status string `json:"status"`
	// The date the EMD was issued
	DateOfIssue string `json:"dateOfIssue,omitempty"`
}

type EmdSegment struct {
	Type_ string `json:"@type,omitempty"`
	// Sequence of EMDSegment
	Sequence int32 `json:"sequence"`
	// The quantity of the ancillary available on this EMDSegment
	Quantity       int32           `json:"quantity,omitempty"`
	EMDDescription *EmdDescription `json:"EMDDescription,omitempty"`
	Amount         *Amount         `json:"Amount,omitempty"`
	Status         *EmdStatusEnum  `json:"Status,omitempty"`
	// The date of service the service is available for
	DateOfService string `json:"DateOfService,omitempty"`
	// The airline the EMD should be presented to to supply the service
	PresentTo string `json:"PresentTo,omitempty"`
	// The location the EMD should be presented to to supply the service
	PresentAt string `json:"PresentAt,omitempty"`
	// The routing the service is valid on
	Routing string `json:"Routing,omitempty"`
}

type Emdid struct {
	Type_      string      `json:"@type"`
	Id         string      `json:"id,omitempty"`
	EMDRef     string      `json:"EMDRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type ErrorDetail struct {
	// The identifier of the source system sending the error or warning
	SourceID string `json:"SourceID"`
	// The error or warning code returned by the source airline or host system
	SourceCode string `json:"SourceCode,omitempty"`
	// The error or warning message as it is returned by the source airline or host system
	SourceDescription string `json:"SourceDescription,omitempty"`
	Type_             string `json:"@type"`
	// Http standard response code
	StatusCode int32 `json:"StatusCode,omitempty"`
	// The Travelport standardized error or warning message
	Message       string          `json:"Message,omitempty"`
	NameValuePair []NameValuePair `json:"NameValuePair,omitempty"`
}

type ErrorWarning struct {
	Type_ string `json:"@type"`
	// Http standard response code
	StatusCode int32 `json:"StatusCode,omitempty"`
	// The Travelport standardized error or warning message
	Message       string          `json:"Message,omitempty"`
	NameValuePair []NameValuePair `json:"NameValuePair,omitempty"`
}

type ErrorWarningDetail struct {
	// The identifier of the source system sending the error or warning
	SourceID string `json:"SourceID"`
	// The error or warning code returned by the source airline or host system
	SourceCode string `json:"SourceCode,omitempty"`
	// The error or warning message as it is returned by the source airline or host system
	SourceDescription string `json:"SourceDescription,omitempty"`
	Type_             string `json:"@type"`
	// Http standard response code
	StatusCode int32 `json:"StatusCode,omitempty"`
	// The Travelport standardized error or warning message
	Message       string          `json:"Message,omitempty"`
	NameValuePair []NameValuePair `json:"NameValuePair,omitempty"`
}

type FareQualifierEnum struct {
	Value *FareQualifierEnumBase `json:"value,omitempty"`
}

type Fee struct {
	Type_ string `json:"@type,omitempty"`
	// Fee code
	FeeCode string `json:"feeCode,omitempty"`
	// Identifies the reporting authority.
	ReportingAuthority string `json:"reportingAuthority,omitempty"`
	// Fee purpose
	Purpose string `json:"purpose,omitempty"`
	// Fee description
	Description        string              `json:"description,omitempty"`
	FeeApplication     *ApplicationEnum    `json:"feeApplication,omitempty"`
	FeeFrequency       *FrequencyEnum      `json:"feeFrequency,omitempty"`
	FeeAmountOrPercent *FeeAmountOrPercent `json:"FeeAmountOrPercent"`
	Tax                []Tax               `json:"Tax,omitempty"`
	// If the fee is included in the Base Price
	IncludedinBaseInd bool `json:"includedinBaseInd,omitempty"`
}

type ModelError struct {
	Type_ string `json:"@type"`
	// Http standard response code
	StatusCode int32 `json:"StatusCode,omitempty"`
	// The Travelport standardized error or warning message
	Message       string          `json:"Message,omitempty"`
	NameValuePair []NameValuePair `json:"NameValuePair,omitempty"`
}

type FeeAmountOrPercent struct {
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type FeeAmountOrPercentAmount struct {
	Amount      *CurrencyAmount `json:"Amount,omitempty"`
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type FeeAmountOrPercentPercent struct {
	// Percent amount of commission
	Percent     float32         `json:"Percent,omitempty"`
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type Fees struct {
	Type_ string `json:"@type"`
	// A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	TotalFees float32 `json:"TotalFees,omitempty"`
}

type FeesDetail struct {
	Fee   []Fee  `json:"Fee,omitempty"`
	Type_ string `json:"@type"`
	// A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	TotalFees float32 `json:"TotalFees,omitempty"`
}

type FormOfPayment struct {
	Value *FormOfPaymentTypeEnum `json:"value,omitempty"`
	// Payment document number
	DocumentNumber string `json:"documentNumber,omitempty"`
	// Encrypted value
	EncryptedValue string `json:"encryptedValue,omitempty"`
	// Document issuer
	DocumentIssuer string            `json:"documentIssuer,omitempty"`
	DocumentType   *DocumentTypeEnum `json:"documentType,omitempty"`
}

// The base amount of a ticket price or net price that is filed in local currency
type FiledAmount struct {
	// Filed amount value
	Value float32 `json:"value,omitempty"`
	// Filed amount currency code
	CurrencyCode string `json:"currencyCode,omitempty"`
	// Filed amount currency code authority
	CodeAuthority string `json:"codeAuthority"`
	// ISO 4217 standard has a different number of decimals
	DecimalPlace int32 `json:"decimalPlace"`
	// ISO 4217 standard decimal authority
	DecimalAuthority string `json:"decimalAuthority,omitempty"`
}

type FlightId struct {
	Type_ string `json:"@type"`
	// Internal ID
	Id string `json:"id,omitempty"`
	// Reference to a Flight object eslewhere in the message
	FlightRef  string      `json:"FlightRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type FlightProduct struct {
	Type_ string `json:"@type,omitempty"`
	// The Segment sequence
	SegmentSequence []int32 `json:"segmentSequence"`
	// The class of service
	ClassOfService string        `json:"classOfService,omitempty"`
	Cabin          *CabinAirEnum `json:"cabin,omitempty"`
	// Fare basis code
	FareBasisCode string        `json:"fareBasisCode,omitempty"`
	FareType      *FareTypeEnum `json:"fareType,omitempty"`
	// The car code
	CarCode string `json:"carCode,omitempty"`
	// The value code
	ValueCode                  string                       `json:"valueCode,omitempty"`
	Brand                      *BrandId                     `json:"Brand,omitempty"`
	CustomerLoyaltyCredit      []CustomerLoyaltyCredit      `json:"CustomerLoyaltyCredit,omitempty"`
	ClassOfServiceAvailability []ClassOfServiceAvailability `json:"ClassOfServiceAvailability,omitempty"`
	FareQualifier              *FareQualifierEnum           `json:"FareQualifier,omitempty"`
	StopoverPriced             *YesNoUnknownEnum            `json:"stopoverPriced,omitempty"`
	// The ticket designator
	TicketDesignator string `json:"ticketDesignator,omitempty"`
	// The ATPCO fare type code
	FareTypeCode string `json:"fareTypeCode,omitempty"`
}

type FlightSegment struct {
	Type_ string `json:"@type,omitempty"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Segment sequence
	Sequence int32 `json:"sequence"`
	// The actual duration (in minutes) between
	ConnectionDuration string                 `json:"connectionDuration,omitempty"`
	Flight             *FlightId              `json:"Flight"`
	OperationalStatus  *OperationalStatusEnum `json:"OperationalStatus,omitempty"`
	// If present and true, the Segments in this Connection must be sold and cancelled together.
	BoundFlightsInd bool         `json:"boundFlightsInd,omitempty"`
	CO2Actual       *Measurement `json:"CO2Actual,omitempty"`
}

type FormOfPaymentAgencyAccount struct {
	// The agency Id
	AgencyId         string      `json:"agencyId,omitempty"`
	Type_            string      `json:"@type"`
	Id               string      `json:"id,omitempty"`
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"`
	Identifier       *Identifier `json:"Identifier,omitempty"`
}

type FormOfPaymentBsp struct {
	// The account number for the Form of payment BSP
	AccountNumber    string      `json:"accountNumber,omitempty"`
	Type_            string      `json:"@type"`
	Id               string      `json:"id,omitempty"`
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"`
	Identifier       *Identifier `json:"Identifier,omitempty"`
}

type FormOfPaymentCash struct {
	// This indicates that the Cash payment should not be refunded
	AgentNonRefundableInd bool        `json:"agentNonRefundableInd,omitempty"`
	Type_                 string      `json:"@type"`
	Id                    string      `json:"id,omitempty"`
	FormOfPaymentRef      string      `json:"FormOfPaymentRef,omitempty"`
	Identifier            *Identifier `json:"Identifier,omitempty"`
}

type FormOfPaymentDocument struct {
	DocumentNumber   *DocumentNumber `json:"DocumentNumber,omitempty"`
	Type_            string          `json:"@type"`
	Id               string          `json:"id,omitempty"`
	FormOfPaymentRef string          `json:"FormOfPaymentRef,omitempty"`
	Identifier       *Identifier     `json:"Identifier,omitempty"`
}

type FormOfPaymentForfeit struct {
	// If true, this form of payment instruction is to forfeit residual amounts specified in an Offer. Used in conjunction with Payment to specify which amounts to be forfeited
	ForfeitInd       bool        `json:"forfeitInd,omitempty"`
	Type_            string      `json:"@type"`
	Id               string      `json:"id,omitempty"`
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"`
	Identifier       *Identifier `json:"Identifier,omitempty"`
}

type FormOfPaymentId struct {
	Type_            string      `json:"@type"`
	Id               string      `json:"id,omitempty"`
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"`
	Identifier       *Identifier `json:"Identifier,omitempty"`
}

type FormOfPaymentIdentifier struct {
	Type_            string      `json:"@type,omitempty"`
	Id               string      `json:"id,omitempty"`
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"`
	Identifier       *Identifier `json:"Identifier,omitempty"`
}

type FormOfPaymentInvoice struct {
	// The invoice number applicable to this form of payment. Send null or empty string if no invoice number specified.
	InvoiceNumber    string      `json:"InvoiceNumber,omitempty"`
	Type_            string      `json:"@type"`
	Id               string      `json:"id,omitempty"`
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"`
	Identifier       *Identifier `json:"Identifier,omitempty"`
}

type FormOfPaymentPaymentCard struct {
	PaymentCard *PaymentCard `json:"PaymentCard,omitempty"`
	// If true, the payment card will not go through card authorization process
	InhibitPaymentCardAuthorizationInd bool             `json:"inhibitPaymentCardAuthorizationInd,omitempty"`
	ExtendedPayment                    *ExtendedPayment `json:"ExtendedPayment,omitempty"`
	Type_                              string           `json:"@type"`
	Id                                 string           `json:"id,omitempty"`
	FormOfPaymentRef                   string           `json:"FormOfPaymentRef,omitempty"`
	Identifier                         *Identifier      `json:"Identifier,omitempty"`
}

type FormOfPaymentVirtualPaymentAccount struct {
	Type_            string      `json:"@type"`
	Id               string      `json:"id,omitempty"`
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"`
	Identifier       *Identifier `json:"Identifier,omitempty"`
	Supplier         string      `json:"Supplier,omitempty"`
	AccountID        string      `json:"AccountID,omitempty"`
	// The alternate agency email to be used for correspondence with this virtual payment
	AlternateEmailAddress []Email `json:"AlternateEmailAddress,omitempty"`
	// Optional text to be sent to the supplier
	PaymentComment string `json:"PaymentComment,omitempty"`
	// Hotel fax number to be used if the hotel fax is unknown or not provided in Property details
	AlternateHotelFax []Telephone `json:"AlternateHotelFax,omitempty"`
	// The maximum amount that the supplier may charge to the payment card including room rate and any incidentals specified
	MaximumChargeableAmount []CurrencyAmount `json:"MaximumChargeableAmount,omitempty"`
	// List of incidentals that are permitted to be charged to the virtual payment card.
	IncidentalCharges []string `json:"IncidentalCharges,omitempty"`
}

type FormOfPaymentWaiverCode struct {
	WaiverCode       *WaiverCode `json:"WaiverCode,omitempty"`
	Type_            string      `json:"@type"`
	Id               string      `json:"id,omitempty"`
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"`
	Identifier       *Identifier `json:"Identifier,omitempty"`
}

// Used to specify the geographic coordinates of a location
type GeoLocation struct {
	// The measure of the angular distance on a meridan north or south equator
	Latitude float64 `json:"latitude"`
	// The measure of the angular distance on a meridan east or west equator
	Longitude float64 `json:"longitude"`
	// The height or an item, typically measured above sea level
	Altitude               float32               `json:"altitude,omitempty"`
	AltitudeUnitOfDistance *UnitOfDistanceEnum   `json:"altitudeUnitOfDistance,omitempty"`
	PositionAccuracy       *PositionAccuracyEnum `json:"positionAccuracy,omitempty"`
	// link for embedded map showing location
	MapURL string `json:"mapURL,omitempty"`
	// The URL to the format for the latitude and longitude for this location.
	FormatURL string `json:"formatURL,omitempty"`
}

// Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
type Identifier struct {
	Value string `json:"value,omitempty"`
	// Name of the authoritative system that created this Guid
	Authority string `json:"authority,omitempty"`
}

type IdentifierRef struct {
	Value string `json:"value,omitempty"`
	// A locally referenced ID
	Id string `json:"id,omitempty"`
	// Descriptive text used to identify the contents of a target object
	Description string `json:"description,omitempty"`
	// Uniform Resource Identifier
	Uris []string `json:"uris,omitempty"`
}

type Guarantee struct {
	Type_ string `json:"@type,omitempty"`
	// Guarantee code
	Code                   string             `json:"code,omitempty"`
	GuaranteeType          *GuaranteeTypeEnum `json:"guaranteeType,omitempty"`
	CredentialsRequiredInd bool               `json:"credentialsRequiredInd,omitempty"`
}

type GuestCount struct {
	Type_ string `json:"@type,omitempty"`
	// The age of the guest
	Age int32 `json:"age,omitempty"`
	// The number of guests in one AgeQualifyingCode or Count.
	Count int32 `json:"count,omitempty"`
	// Enter 10 for an adult or 08 for a child
	AgeQualifyingCode string `json:"ageQualifyingCode,omitempty"`
}

type GuestCounts struct {
	Type_      string       `json:"@type,omitempty"`
	GuestCount []GuestCount `json:"GuestCount"`
}

type GuestRoomInfo struct {
	Type_ string `json:"@type,omitempty"`
	// OTA code for this room type.
	Code string `json:"code,omitempty"`
	// Number of guest rooms with the guest room code.
	Number int32 `json:"number,omitempty"`
	// description of the guest room code.
	Description string `json:"description,omitempty"`
}

type HotelPenalty struct {
	Type_        string            `json:"@type"`
	SubjectToTax *YesNoUnknownEnum `json:"subjectToTax,omitempty"`
}

type HotelPenaltyAmount struct {
	IncludesTax  *YesNoUnknownEnum `json:"includesTax,omitempty"`
	Amount       []CurrencyAmount  `json:"Amount"`
	Type_        string            `json:"@type"`
	SubjectToTax *YesNoUnknownEnum `json:"subjectToTax,omitempty"`
}

type HotelPenaltyNights struct {
	// The number of nights that will be charged as a penalty
	Nights       int32             `json:"Nights"`
	Type_        string            `json:"@type"`
	SubjectToTax *YesNoUnknownEnum `json:"subjectToTax,omitempty"`
}

type HotelPenaltyPercent struct {
	AppliesTo *PercentAppliesTo `json:"appliesTo"`
	// A percentage charged as a Penalty
	Percent float32 `json:"Percent"`
	// The number of nights the percentage needs to be applied to determine cancel penalty amount
	Nights       int32             `json:"Nights,omitempty"`
	Amount       []CurrencyAmount  `json:"Amount,omitempty"`
	Type_        string            `json:"@type"`
	SubjectToTax *YesNoUnknownEnum `json:"subjectToTax,omitempty"`
}

type HotelSearchCriterion struct {
	Type_ string `json:"@type,omitempty"`
	// Number of rooms requested
	NumberOfRooms      int32               `json:"numberOfRooms,omitempty"`
	PropertyRequest    []PropertyRequest   `json:"PropertyRequest"`
	RoomStayCandidates *RoomStayCandidates `json:"RoomStayCandidates,omitempty"`
	RateCandidates     *RateCandidates     `json:"RateCandidates,omitempty"`
}

// URL of the image
type Image struct {
	Value string `json:"value,omitempty"`
	// Deprecated and replaced by Image Size
	DimensionCategory string `json:"dimensionCategory,omitempty"`
	// Width of image
	Width int32 `json:"width,omitempty"`
	// Height
	Height int32 `json:"height,omitempty"`
	// Image title
	Caption string `json:"caption,omitempty"`
	// deprecated and replaced by pictureOf
	PictureCategory int32          `json:"pictureCategory,omitempty"`
	ImageSize       *ImageSizeEnum `json:"imageSize,omitempty"`
	PictureOf       *PictureofEnum `json:"pictureOf,omitempty"`
}

type ListPaymentCardIssuerEnum struct {
	Value *ListPaymentCardIssuerEnumBase `json:"value,omitempty"`
}

// Contains the locator (PNR or external locator) or cancellation number for the reservation, order, or offer
type Locator struct {
	Value string `json:"value,omitempty"`
	// Specifies the type of reservation ID
	LocatorType string `json:"locatorType,omitempty"`
	// Specifies a unique identifier to indicate the source system which generated the resid
	Source string `json:"source,omitempty"`
	// Specifies the context of the source
	SourceContext string `json:"sourceContext,omitempty"`
	// Used for codes
	OtaType string `json:"otaType,omitempty"`
	// Reservation Creation date
	CreationDate string `json:"creationDate,omitempty"`
}

type MagneticStripe struct {
	Type_ string `json:"@type,omitempty"`
	// Note: This contains a key required to retrieve the full payment instrument details compliant with PCI DSS standards.
	EncryptionKey string `json:"encryptionKey,omitempty"`
	// Developer: This contains a reference to the key generation method being used - this is NOT the key value.
	EncryptionKeyMethod string `json:"encryptionKeyMethod,omitempty"`
	// OpenTravel Best Practice: Encryption Method: When using the OpenTravel Encryption element, it is RECOMMENDED that all trading partners be informed of all encryption methods being used in advance of implementation to ensure message processing compatibility.
	EncryptionMethod string `json:"encryptionMethod,omitempty"`
	// Encrypted value
	EncryptedValue string `json:"encryptedValue,omitempty"`
	// Masked Value
	Mask string `json:"mask,omitempty"`
	// Token value
	Token string `json:"token,omitempty"`
	// Developer: This contains a provider ID if multiple providers are used for secure information exchange.
	TokenProviderID      string                       `json:"tokenProviderID,omitempty"`
	AuthenticationMethod *EncryptionTokenTypeAuthEnum `json:"authenticationMethod,omitempty"`
	// Don't use this unless it is REALLY ok to not use encryption. Non-secure (plain text) value.
	PlainText    string        `json:"PlainText,omitempty"`
	ErrorWarning *ErrorWarning `json:"ErrorWarning,omitempty"`
}

// Maximum Available Rate for this Property, including the authority who provided the rate if it is different from the PropertyInfo authority.
type MaximumAvailableRate struct {
	Value float32 `json:"value,omitempty"`
	// An ISO 4217 alpha character code that specifies a money unit
	Code string `json:"code,omitempty"`
	// Minor units are a mechanism for expressing the relationship between a major currency unit and its corresponding minor currency unit.
	MinorUnit int32 `json:"minorUnit,omitempty"`
	// Name of the authoritative system that provided this rate
	Authority string `json:"authority,omitempty"`
}

// Indicates if a meal is included
type MealsIncluded struct {
	BreakfastInd bool `json:"breakfastInd,omitempty"`
	LunchInd     bool `json:"lunchInd,omitempty"`
	DinnerInd    bool `json:"dinnerInd,omitempty"`
}

// Used for dimensional units (width, height, depth) or weight
type Measurement struct {
	Value           float32              `json:"value,omitempty"`
	MeasurementType *MeasurementTypeEnum `json:"measurementType,omitempty"`
	Unit            *UnitOfMeasureEnum   `json:"unit,omitempty"`
}

type MeetingRoom struct {
	Type_ string `json:"@type,omitempty"`
	// The name of the meeting room
	Name string `json:"name"`
	// OTA code for this room type.
	Codes      []string    `json:"codes,omitempty"`
	Capacity   int32       `json:"capacity,omitempty"`
	Size       int32       `json:"size,omitempty"`
	UnitOfSize *UnitOfSize `json:"unitOfSize,omitempty"`
}

type MeetingRooms struct {
	Type_       string        `json:"@type,omitempty"`
	Number      int32         `json:"number,omitempty"`
	MeetingRoom []MeetingRoom `json:"MeetingRoom,omitempty"`
}

// Used for data stored in Name Value pairs
type NameValuePair struct {
	Value string `json:"value,omitempty"`
	// Optional internally referenced id
	Id string `json:"id,omitempty"`
	// Key
	Name string `json:"name"`
}

type NetRemitInfo struct {
	Type_ string `json:"@type,omitempty"`
	// The CAR code applied to this product for use in net remit
	CarCode string `json:"CarCode,omitempty"`
	// The Value code applied to this product for use in net remit
	ValueCode string `json:"ValueCode,omitempty"`
	// The actual selling fare which will override the Offer base fare on the document
	ActualSellingFare float32      `json:"ActualSellingFare,omitempty"`
	NetBaseAmount     *FiledAmount `json:"NetBaseAmount,omitempty"`
}

// A URL that describes a step that can be applied to the resource containing the next step structure.
type NextStep struct {
	Value string `json:"value,omitempty"`
	// Identifier for the Next Step
	Id string `json:"id,omitempty"`
	// The action this next step is intended to achieve
	Action string              `json:"action"`
	Method *NextStepMethodEnum `json:"method"`
	// Additional clarification for the next step
	Description string `json:"description,omitempty"`
}

type NextSteps struct {
	// The base portion of the uri in order to shorten the uri's in the individual steps
	BaseURI string `json:"baseURI"`
	// Optional internally referenced id
	Id       string     `json:"id,omitempty"`
	NextStep []NextStep `json:"NextStep"`
}

type NightlyRate struct {
	Type_ string `json:"@type,omitempty"`
	// Start date
	StartDate string `json:"startDate"`
	// Number of nights this rate applies
	Nights int32   `json:"nights,omitempty"`
	Amount *Amount `json:"Amount"`
}

type Offer struct {
	// A reference to the Offer this offer is sold in conjunction with
	ParentOfferRef         string                   `json:"parentOfferRef,omitempty"`
	Product                []ProductIdentifier      `json:"Product"`
	Price                  *Price                   `json:"Price"`
	TermsAndConditionsFull []TermsAndConditionsFull `json:"TermsAndConditionsFull"`
	// If true, the Offer is passive for booking purposes.
	PassiveOfferInd bool   `json:"passiveOfferInd,omitempty"`
	Type_           string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message
	OfferRef   string      `json:"offerRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type OfferId struct {
	Type_ string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message
	OfferRef   string      `json:"offerRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type OfferIdentifier struct {
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message
	OfferRef   string      `json:"offerRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type OfferLink struct {
	Type_       string       `json:"@type,omitempty"`
	OfferRef    string       `json:"OfferRef,omitempty"`
	ParentOffer *ParentOffer `json:"ParentOffer,omitempty"`
}

type OfferQueryBuildFromCatalogOffering struct {
	Type_                               string                               `json:"@type"`
	BuildFromCatalogOfferingHospitality *BuildFromCatalogOfferingHospitality `json:"BuildFromCatalogOfferingHospitality,omitempty"`
}

type OfferQueryBuildFromCatalogOfferingWrapper struct {
	OfferQueryBuildFromCatalogOffering *OfferQueryBuildFromCatalogOffering `json:"OfferQueryBuildFromCatalogOffering,omitempty"`
}

type OfferQueryBuildFromCatalogOfferingsHospitality struct {
	Type_                            string                            `json:"@type"`
	BuildFromCatalogOfferingsRequest *BuildFromCatalogOfferingsRequest `json:"BuildFromCatalogOfferingsRequest,omitempty"`
}

type OfferQueryBuildFromCatalogOfferingsHospitalityWrapper struct {
	OfferQueryBuildFromCatalogOfferingsHospitality *OfferQueryBuildFromCatalogOfferingsHospitality `json:"OfferQueryBuildFromCatalogOfferingsHospitality,omitempty"`
}

type OfferStatus struct {
	Type_ string `json:"@type"`
}

type OfferStatusHospitality struct {
	Status *OfferStatusEnum `json:"Status,omitempty"`
	Type_  string           `json:"@type"`
}

type OperationTimes struct {
	Type_      string          `json:"@type,omitempty"`
	DaysOfWeek []DayOfWeekEnum `json:"daysOfWeek,omitempty"`
	OpenTime   string          `json:"openTime,omitempty"`
	CloseTime  string          `json:"closeTime,omitempty"`
}

type OperationalStatusEnum struct {
	Value *OperationalStatusEnumBase `json:"value,omitempty"`
}

type OrganizationLoyaltyProgram struct {
	// The supplier of the loyalty program
	Supplier string `json:"Supplier"`
	// Loyalty Identifier
	LoyaltyIdentifier string      `json:"LoyaltyIdentifier"`
	Type_             string      `json:"@type"`
	Id                string      `json:"id,omitempty"`
	Identifier        *Identifier `json:"Identifier,omitempty"`
}

type OrganizationLoyaltyProgramId struct {
	Type_      string      `json:"@type"`
	Id         string      `json:"id,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

// code
type OtaCodeWithDescription struct {
	Value string `json:"value,omitempty"`
	// OTA code description
	Description string `json:"description,omitempty"`
}

type ParentOffer struct {
	Type_      string `json:"@type"`
	OfferRef   string `json:"OfferRef,omitempty"`
	ProductRef string `json:"ProductRef,omitempty"`
}

type PassengerFlight struct {
	Type_ string `json:"@type,omitempty"`
	// Number of passengers of the specified passenger type code
	PassengerQuantity int32 `json:"passengerQuantity,omitempty"`
	// Passenger type code
	PassengerTypeCode string          `json:"passengerTypeCode,omitempty"`
	FlightProduct     []FlightProduct `json:"FlightProduct,omitempty"`
}

type Password struct {
	Type_ string `json:"@type,omitempty"`
	// Note: This contains a key required to retrieve the full payment instrument details compliant with PCI DSS standards.
	EncryptionKey string `json:"encryptionKey,omitempty"`
	// Developer: This contains a reference to the key generation method being used - this is NOT the key value.
	EncryptionKeyMethod string `json:"encryptionKeyMethod,omitempty"`
	// OpenTravel Best Practice: Encryption Method: When using the OpenTravel Encryption element, it is RECOMMENDED that all trading partners be informed of all encryption methods being used in advance of implementation to ensure message processing compatibility.
	EncryptionMethod string `json:"encryptionMethod,omitempty"`
	// Encrypted value
	EncryptedValue string `json:"encryptedValue,omitempty"`
	// Masked Value
	Mask string `json:"mask,omitempty"`
	// Token value
	Token string `json:"token,omitempty"`
	// Developer: This contains a provider ID if multiple providers are used for secure information exchange.
	TokenProviderID      string                       `json:"tokenProviderID,omitempty"`
	AuthenticationMethod *EncryptionTokenTypeAuthEnum `json:"authenticationMethod,omitempty"`
	// Don't use this unless it is REALLY ok to not use encryption. Non-secure (plain text) value.
	PlainText    string        `json:"PlainText,omitempty"`
	ErrorWarning *ErrorWarning `json:"ErrorWarning,omitempty"`
}

type Payment struct {
	Amount                  *CurrencyAmount          `json:"Amount"`
	FormOfPaymentIdentifier *FormOfPaymentIdentifier `json:"FormOfPaymentIdentifier,omitempty"`
	OfferIdentifier         []OfferIdentifier        `json:"OfferIdentifier,omitempty"`
	Fees                    *Fees                    `json:"Fees,omitempty"`
	Taxes                   *Taxes                   `json:"Taxes,omitempty"`
	TravelerIdentifierRef   []TravelerIdentifierRef  `json:"TravelerIdentifierRef,omitempty"`
	BaseAmount              *CurrencyAmount          `json:"BaseAmount,omitempty"`
	// If true, the payment is a deposit on the referenced Offer
	DepositInd                 bool                         `json:"depositInd,omitempty"`
	ExtendedPayment            *ExtendedPayment             `json:"ExtendedPayment,omitempty"`
	AgencyServiceFeeIdentifier []AgencyServiceFeeIdentifier `json:"AgencyServiceFeeIdentifier,omitempty"`
	// If true, the payment is a guarantee for the referenced Offer
	GuaranteeInd bool        `json:"guaranteeInd,omitempty"`
	Type_        string      `json:"@type"`
	Id           string      `json:"id,omitempty"`
	PaymentRef   string      `json:"PaymentRef,omitempty"`
	Identifier   *Identifier `json:"Identifier,omitempty"`
}

type PaymentCard struct {
	Type_ string `json:"@type"`
	// Payment card reference ID.
	Id string `json:"id,omitempty"`
	// Indicated starting date.
	EffectiveDate string `json:"effectiveDate,omitempty"`
	// The expiration date value
	ExpireDate string `json:"expireDate,omitempty"`
	// The approval code value
	ApprovalCode string               `json:"approvalCode,omitempty"`
	PrivacyGroup *Privacy             `json:"PrivacyGroup,omitempty"`
	CardType     *PaymentCardTypeEnum `json:"CardType,omitempty"`
	// Specifies the two character code (MC, VI, AX, etc) for the payment card (open enumeration)
	CardCode   string                 `json:"CardCode,omitempty"`
	CardBrand  *PaymentCardTypeIssuer `json:"CardBrand,omitempty"`
	CardIssuer *PaymentCardTypeIssuer `json:"CardIssuer,omitempty"`
	// Name as displayed on Payment Card
	CardHolderName string           `json:"CardHolderName,omitempty"`
	CardNumber     *CardNumber      `json:"CardNumber,omitempty"`
	SeriesCode     *SeriesCode      `json:"SeriesCode,omitempty"`
	MagneticStripe []MagneticStripe `json:"MagneticStripe,omitempty"`
	// Implementer: If true, all or a portion of this data is secure, via tokenization, encryption and/or masking.
	SecureInd bool `json:"secureInd,omitempty"`
}

type PaymentCardDetail struct {
	// The country code ISO
	CountryOfIssue string `json:"countryOfIssue,omitempty"`
	// The company card reference
	CompanyCardReference string `json:"companyCardReference,omitempty"`
	// The bank name value
	BankName string `json:"bankName,omitempty"`
	// The bank country code ISO
	BankCountryCode string `json:"bankCountryCode,omitempty"`
	// The bank state code ISO
	BankStateCode       string               `json:"bankStateCode,omitempty"`
	CardHolderId        *Identifier          `json:"CardHolderId,omitempty"`
	PersonName          *PersonName          `json:"PersonName,omitempty"`
	Address             *Address             `json:"Address,omitempty"`
	Telephone           []Telephone          `json:"Telephone,omitempty"`
	Email               []Email              `json:"Email,omitempty"`
	CustomerLoyalty     []CustomerLoyalty    `json:"CustomerLoyalty,omitempty"`
	SignatureOnFile     *SignatureOnFile     `json:"SignatureOnFile,omitempty"`
	ThreeDomainSecurity *ThreeDomainSecurity `json:"ThreeDomainSecurity,omitempty"`
	// Implementer: If true, the credit card company is requested to delay the date on which the amount of this transaction is applied to the customer's account.
	ExtendedPaymentInd bool `json:"extendedPaymentInd,omitempty"`
	// True if this payment card has been issued through Enett
	EnettInd bool `json:"enettInd,omitempty"`
	// If true, then the payment card holder is not one of the travelers in the reservation
	ThirdPartyInd bool `json:"thirdPartyInd,omitempty"`
	// If true, override airline restriction on the payment card
	AcceptanceOverrideInd bool   `json:"acceptanceOverrideInd,omitempty"`
	Type_                 string `json:"@type"`
	// Payment card reference ID.
	Id string `json:"id,omitempty"`
	// Indicated starting date.
	EffectiveDate string `json:"effectiveDate,omitempty"`
	// The expiration date value
	ExpireDate string `json:"expireDate,omitempty"`
	// The approval code value
	ApprovalCode string               `json:"approvalCode,omitempty"`
	PrivacyGroup *Privacy             `json:"PrivacyGroup,omitempty"`
	CardType     *PaymentCardTypeEnum `json:"CardType,omitempty"`
	// Specifies the two character code (MC, VI, AX, etc) for the payment card (open enumeration)
	CardCode   string                 `json:"CardCode,omitempty"`
	CardBrand  *PaymentCardTypeIssuer `json:"CardBrand,omitempty"`
	CardIssuer *PaymentCardTypeIssuer `json:"CardIssuer,omitempty"`
	// Name as displayed on Payment Card
	CardHolderName string           `json:"CardHolderName,omitempty"`
	CardNumber     *CardNumber      `json:"CardNumber,omitempty"`
	SeriesCode     *SeriesCode      `json:"SeriesCode,omitempty"`
	MagneticStripe []MagneticStripe `json:"MagneticStripe,omitempty"`
	// Implementer: If true, all or a portion of this data is secure, via tokenization, encryption and/or masking.
	SecureInd bool `json:"secureInd,omitempty"`
}

// This object contains Cards details for Payment
type PaymentCardTypeIssuer struct {
	PaymentCardIssuers          *ListPaymentCardIssuerEnum `json:"paymentCardIssuers,omitempty"`
	PaymentCardIssuersExtension string                     `json:"paymentCardIssuersExtension,omitempty"`
	// Assigned Type: c-1100:NumberDoubleDigit
	IssueNumber int32 `json:"issueNumber,omitempty"`
}

type PaymentId struct {
	Type_      string      `json:"@type"`
	Id         string      `json:"id,omitempty"`
	PaymentRef string      `json:"PaymentRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type PaymentIdentifier struct {
	Type_      string      `json:"@type,omitempty"`
	Id         string      `json:"id,omitempty"`
	PaymentRef string      `json:"PaymentRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type PaidTaxes struct {
	Type_ string `json:"@type"`
	// A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	TotalTaxes float32 `json:"TotalTaxes,omitempty"`
}

type PaidTaxesDetail struct {
	Tax        []Tax       `json:"Tax,omitempty"`
	TaxPercent *TaxPercent `json:"TaxPercent,omitempty"`
	Type_      string      `json:"@type"`
	// A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	TotalTaxes float32 `json:"TotalTaxes,omitempty"`
}

type Penalty struct {
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type PenaltyAmount struct {
	Amount      *CurrencyAmount `json:"Amount,omitempty"`
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type PenaltyPercent struct {
	// Percent amount of commission
	Percent     float32         `json:"Percent,omitempty"`
	Type_       string          `json:"@type"`
	Application *CommissionEnum `json:"application,omitempty"`
}

type PersonName struct {
	Type_ string `json:"@type"`
	// Salutation of honorific (e.g. Mr., Mrs., Ms., Miss, Dr.)
	Prefix string `json:"Prefix,omitempty"`
	// Given name, first name or names.
	Given string `json:"Given,omitempty"`
	// The middle name of the person name.
	Middle string `json:"Middle,omitempty"`
	// Family name, last name.
	Surname string `json:"Surname"`
}

type PetPolicy struct {
	Type_   string            `json:"@type,omitempty"`
	Allowed *YesNoUnknownEnum `json:"allowed"`
	// Pet policy code
	PolicyCode  string                   `json:"policyCode,omitempty"`
	Description *TextTitleAndDescription `json:"Description,omitempty"`
}

type Preference struct {
	AppliesTo          *AppliesTo          `json:"AppliesTo,omitempty"`
	TravelerIdentifier *TravelerIdentifier `json:"TravelerIdentifier,omitempty"`
	Type_              string              `json:"@type"`
	Id                 string              `json:"id,omitempty"`
}

type PreferenceId struct {
	Type_ string `json:"@type"`
	Id    string `json:"id,omitempty"`
}

type Price struct {
	Type_ string `json:"@type"`
	// Internally referenced id
	Id           string        `json:"id,omitempty"`
	CurrencyCode *CurrencyCode `json:"CurrencyCode,omitempty"`
	// The total amount not including taxes and/or fees
	Base float32 `json:"Base,omitempty"`
	// The total of the taxes included in the total price
	TotalTaxes float32 `json:"TotalTaxes,omitempty"`
	// The total of the fees included in the total price
	TotalFees float32 `json:"TotalFees,omitempty"`
	// The total price of the product in the currency indicated
	TotalPrice          float32              `json:"TotalPrice,omitempty"`
	VendorCurrencyTotal *VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"`
}

type PriceBreakdown struct {
	Type_      string      `json:"@type"`
	Amount     *Amount     `json:"Amount,omitempty"`
	Commission *Commission `json:"Commission,omitempty"`
}

type PriceBreakdownHospitality struct {
	RoomPricingType    *PricingEnum             `json:"roomPricingType,omitempty"`
	Description        string                   `json:"Description,omitempty"`
	NightlyRate        []NightlyRate            `json:"NightlyRate,omitempty"`
	AverageNightlyRate []CurrencyAmount         `json:"AverageNightlyRate,omitempty"`
	AmenitySurcharges  *AmenitySurchargesDetail `json:"AmenitySurcharges,omitempty"`
	// If present and true, indicates the nightly price changes one or more times during the stay
	PriceChangesDuringStayInd bool        `json:"priceChangesDuringStayInd,omitempty"`
	Type_                     string      `json:"@type"`
	Amount                    *Amount     `json:"Amount,omitempty"`
	Commission                *Commission `json:"Commission,omitempty"`
}

type PriceDetail struct {
	PriceBreakdown []PriceBreakdown `json:"PriceBreakdown,omitempty"`
	Type_          string           `json:"@type"`
	// Internally referenced id
	Id           string        `json:"id,omitempty"`
	CurrencyCode *CurrencyCode `json:"CurrencyCode,omitempty"`
	// The total amount not including taxes and/or fees
	Base float32 `json:"Base,omitempty"`
	// The total of the taxes included in the total price
	TotalTaxes float32 `json:"TotalTaxes,omitempty"`
	// The total of the fees included in the total price
	TotalFees float32 `json:"TotalFees,omitempty"`
	// The total price of the product in the currency indicated
	TotalPrice          float32              `json:"TotalPrice,omitempty"`
	VendorCurrencyTotal *VendorCurrencyTotal `json:"VendorCurrencyTotal,omitempty"`
}

type PrimaryContact struct {
	ShareWith *ShareWithEnum `json:"shareWith,omitempty"`
	// Primary contact shared with supplier
	ShareWithSupplier  []string            `json:"shareWithSupplier,omitempty"`
	Email              *Email              `json:"Email,omitempty"`
	Telephone          *Telephone          `json:"Telephone,omitempty"`
	TravelerIdentifier *TravelerIdentifier `json:"TravelerIdentifier,omitempty"`
	// If true, the passenger has refused to provide emergency contact details
	ContactInformationRefusedInd bool   `json:"contactInformationRefusedInd,omitempty"`
	Type_                        string `json:"@type"`
	Id                           string `json:"id,omitempty"`
}

type PrimaryContactId struct {
	Type_ string `json:"@type"`
	Id    string `json:"id,omitempty"`
}

// Confidential details for marketing purpose
type Privacy struct {
	// Optional internally referenced id
	Id             string            `json:"id,omitempty"`
	ShareMarketing *YesNoInheritEnum `json:"shareMarketing,omitempty"`
	ShareSync      *YesNoInheritEnum `json:"shareSync,omitempty"`
	OptOutInd      *YesNoInheritEnum `json:"optOutInd,omitempty"`
	OptInStatus    *OptInStatusEnum  `json:"optInStatus,omitempty"`
	// The datetime of receiving the opt in notice
	OptInDate time.Time `json:"optInDate,omitempty"`
	// The datetime the opt out notice was received
	OptOutDate time.Time `json:"optOutDate,omitempty"`
}

type Product struct {
	// The quantity of the product
	Quantity int32  `json:"Quantity,omitempty"`
	Type_    string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message
	ProductRef string      `json:"productRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type ProductAir struct {
	// Total duration of all Segments that are part of this ProductAir
	TotalDuration   string            `json:"totalDuration,omitempty"`
	FlightSegment   []FlightSegment   `json:"FlightSegment"`
	PassengerFlight []PassengerFlight `json:"PassengerFlight"`
	// The quantity of the product
	Quantity int32 `json:"Quantity,omitempty"`
}

type ProductHospitality struct {
	// Booking code retrieved from the Availability response.
	BookingCode string `json:"bookingCode,omitempty"`
	// Total number of guests
	Guests int32 `json:"guests,omitempty"`
	// More rates token
	MoreRatesToken string            `json:"moreRatesToken,omitempty"`
	AdaCompliant   *YesNoUnknownEnum `json:"adaCompliant,omitempty"`
	// The name of the hotel property
	PropertyName string       `json:"propertyName,omitempty"`
	PropertyKey  *PropertyKey `json:"PropertyKey"`
	RoomType     *RoomType    `json:"RoomType,omitempty"`
	DateRange    *DateRange   `json:"DateRange,omitempty"`
	// The quantity of the product
	Quantity int32 `json:"Quantity,omitempty"`
}

type ProductHospitalityOffer struct {
	// Refers to PropertyDates object in the ReferenceList_PropertyDates
	PropertyDatesRef string `json:"propertyDatesRef,omitempty"`
	// Booking code retrieved from the Availability response.
	BookingCode string    `json:"bookingCode,omitempty"`
	RoomType    *RoomType `json:"RoomType,omitempty"`
	// The quantity of the product
	Quantity int32 `json:"Quantity,omitempty"`
}

type ProductIdentifier struct { // TODO : also ProductId
	Type_ string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message
	ProductRef string      `json:"productRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type ProductOptions struct {
	// NonnegativeInteger
	Sequence int32               `json:"sequence,omitempty"`
	Product  []ProductIdentifier `json:"Product"`
	Type_    string              `json:"@type"`
	// Local indentifier within a given message for this object.
	Id         string      `json:"id,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
	// Used to reference another instance of this object in the same message
	ProductOptionsRef string `json:"ProductOptionsRef,omitempty"`
}

type ProductOptionsId struct {
	Type_ string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id         string      `json:"id,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
	// Used to reference another instance of this object in the same message
	ProductOptionsRef string `json:"ProductOptionsRef,omitempty"`
}

type ProductRateCodeInfo struct {
	Type_ string `json:"@type,omitempty"`
	// Product reference
	ProductRef   string        `json:"ProductRef,omitempty"`
	RateCodeInfo *RateCodeInfo `json:"RateCodeInfo"`
}

type Properties struct {
	// Total number or properties returned for the request
	TotalProperties int32 `json:"totalProperties,omitempty"`
	// Number of properties per page
	PropertiesPerPage int32          `json:"propertiesPerPage,omitempty"`
	PropertyInfo      []PropertyInfo `json:"PropertyInfo"`
	Type_             string         `json:"@type"`
	Identifier        *Identifier    `json:"Identifier,omitempty"`
}

type PropertiesId struct {
	Type_      string      `json:"@type"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type PropertiesQueryPrecisionSearch struct {
	Type_     string              `json:"@type"`
	SortOrder *HotelSortOrderEnum `json:"SortOrder,omitempty"`
	// Check In Date
	CheckInDate string `json:"CheckInDate"`
	// Check Out Date
	CheckOutDate string `json:"CheckOutDate"`
	// The permitted property chain code(s) to be returned for this request
	ChainCodes []string `json:"ChainCodes,omitempty"`
	// The preferred name of the property
	HotelName string `json:"HotelName,omitempty"`
	// You can use requested currency to request conversion rate information. The response will return the currencyRateConversion object which will contain conversion rate of the requested currency.
	RequestedCurrency string              `json:"RequestedCurrency,omitempty"`
	ImageSize         *ImageSizeEnum      `json:"ImageSize,omitempty"`
	RoomStayCandidate []RoomStayCandidate `json:"RoomStayCandidate,omitempty"`
	RateCandidates    *RateCandidates     `json:"RateCandidates,omitempty"`
	SearchBy          *SearchBy           `json:"SearchBy"`
	// If true, all property images of the size requested will be returned. If blank or false the best single property image will be returned.
	ReturnAllImagesInd  bool           `json:"returnAllImagesInd,omitempty"`
	PropertyAmenityCode []string       `json:"PropertyAmenityCode,omitempty"`
	MealsIncluded       *MealsIncluded `json:"MealsIncluded,omitempty"`
	// If true, return Properties with at least one refundable rate.
	RefundableInd bool `json:"RefundableInd,omitempty"`
	// If true, return Properties with at least one commissionable rate.
	CommissionableInd bool `json:"CommissionableInd,omitempty"`
	// If true, return Properties with at least one rate for a smoking room.
	SmokingInd bool `json:"SmokingInd,omitempty"`
}

type PropertiesQueryPrecisionSearchWrapper struct {
	PropertiesQuerySearch *PropertiesQueryPrecisionSearch `json:"PropertiesQuerySearch,omitempty"`
}

type PropertiesQuerySearch struct {
	Type_     string              `json:"@type"`
	SortOrder *HotelSortOrderEnum `json:"SortOrder,omitempty"`
	// Check In Date
	CheckInDate string `json:"CheckInDate"`
	// Check Out Date
	CheckOutDate string `json:"CheckOutDate"`
	// The permitted property chain code(s) to be returned for this request
	ChainCodes []string `json:"ChainCodes,omitempty"`
	// The preferred name of the property
	HotelName string `json:"HotelName,omitempty"`
	// You can use requested currency to request conversion rate information. The response will return the currencyRateConversion object which will contain conversion rate of the requested currency.
	RequestedCurrency string              `json:"RequestedCurrency,omitempty"`
	ImageSize         *ImageSizeEnum      `json:"ImageSize,omitempty"`
	RoomStayCandidate []RoomStayCandidate `json:"RoomStayCandidate,omitempty"`
	RateCandidates    *RateCandidates     `json:"RateCandidates,omitempty"`
	SearchBy          *SearchBy           `json:"SearchBy"`
	// If true, all property images of the size requested will be returned. If blank or false the best single property image will be returned.
	ReturnAllImagesInd  bool     `json:"returnAllImagesInd,omitempty"`
	PropertyAmenityCode []string `json:"PropertyAmenityCode,omitempty"`
}

type PropertiesQuerySearchWrapper struct {
	PropertiesQuerySearch *PropertiesQuerySearch `json:"PropertiesQuerySearch,omitempty"`
}

type PropertiesQuerySpecificPrecisionPropertyList struct {
	Type_ string `json:"@type"`
	// Checkin date
	CheckinDate string `json:"checkinDate"`
	// Checkout date
	CheckoutDate string `json:"checkoutDate"`
	// Number of travelers. Must be a numeric value between 1 and 9.
	NumberOfGuests int32 `json:"numberOfGuests"`
	// You can use requested currency to request conversion rate information. The response will return the currencyRateConversion object which will contain conversion rate of the requested currency.
	RequestedCurrency string `json:"requestedCurrency,omitempty"`
	// Minimum rate
	MinimumRate float32 `json:"minimumRate,omitempty"`
	// Maximum rate
	MaximumRate float32 `json:"maximumRate,omitempty"`
	// Number of rooms
	NumberOfRooms      int32               `json:"numberOfRooms,omitempty"`
	PropertyKey        []PropertyKey       `json:"PropertyKey"`
	RateCandidates     *RateCandidates     `json:"RateCandidates,omitempty"`
	ImageSize          *ImageSizeEnum      `json:"imageSize,omitempty"`
	RoomStayCandidates *RoomStayCandidates `json:"RoomStayCandidates,omitempty"`
	// If true, all property images of the size requested will be returned. If blank or false the best single property image will be returned.
	ReturnAllImagesInd bool           `json:"returnAllImagesInd,omitempty"`
	MealsIncluded      *MealsIncluded `json:"mealsIncluded,omitempty"`
	// If true, return Properties with at least one refundable rate.
	RefundableInd bool `json:"refundableInd,omitempty"`
	// If true, return Properties with at least one commissionable rate.
	CommissionableInd bool `json:"commissionableInd,omitempty"`
	// If true, return Properties with at least one rate for a smoking room.
	SmokingInd bool `json:"smokingInd,omitempty"`
}

type PropertiesQuerySpecificPrecisionPropertyListWrapper struct {
	PropertiesQuerySpecificPrecisionPropertyList *PropertiesQuerySpecificPrecisionPropertyList `json:"PropertiesQuerySpecificPrecisionPropertyList,omitempty"`
}

type PropertiesQuerySpecificPropertyList struct {
	Type_ string `json:"@type"`
	// Checkin date
	CheckinDate string `json:"checkinDate"`
	// Checkout date
	CheckoutDate string `json:"checkoutDate"`
	// Number of travelers. Must be a numeric value between 1 and 9.
	NumberOfGuests int32 `json:"numberOfGuests"`
	// You can use requested currency to request conversion rate information. The response will return the currencyRateConversion object which will contain conversion rate of the requested currency.
	RequestedCurrency string `json:"requestedCurrency,omitempty"`
	// Minimum rate
	MinimumRate float32 `json:"minimumRate,omitempty"`
	// Maximum rate
	MaximumRate float32 `json:"maximumRate,omitempty"`
	// Number of rooms
	NumberOfRooms      int32               `json:"numberOfRooms,omitempty"`
	PropertyKey        []PropertyKey       `json:"PropertyKey"`
	RateCandidates     *RateCandidates     `json:"RateCandidates,omitempty"`
	ImageSize          *ImageSizeEnum      `json:"imageSize,omitempty"`
	RoomStayCandidates *RoomStayCandidates `json:"RoomStayCandidates,omitempty"`
	// If true, all property images of the size requested will be returned. If blank or false the best single property image will be returned.
	ReturnAllImagesInd bool `json:"returnAllImagesInd,omitempty"`
}

type PropertiesQuerySpecificPropertyListWrapper struct {
	PropertiesQuerySpecificPropertyList *PropertiesQuerySpecificPropertyList `json:"PropertiesQuerySpecificPropertyList,omitempty"`
}

type Property struct {
	// The property name
	Name                 string                   `json:"name"`
	Rating               []Rating                 `json:"Rating,omitempty"`
	GeoLocation          *GeoLocation             `json:"GeoLocation,omitempty"`
	Image                []Image                  `json:"Image,omitempty"`
	Description          []string                 `json:"Description,omitempty"`
	BusinessService      []Service                `json:"BusinessService,omitempty"`
	AccessibilityFeature []OtaCodeWithDescription `json:"AccessibilityFeature,omitempty"`
	GuestRoomInfo        []GuestRoomInfo          `json:"GuestRoomInfo,omitempty"`
	Type_                string                   `json:"@type"`
	// Local reference id.
	Id          string       `json:"id,omitempty"`
	Identifier  *Identifier  `json:"Identifier,omitempty"`
	PropertyKey *PropertyKey `json:"PropertyKey"`
}

type PropertyAmenity struct {
	Type_ string `json:"@type,omitempty"`
	// Type of amenity.
	Description string `json:"description,omitempty"`
	// Location of the property
	Location string `json:"location,omitempty"`
	// Name of the property
	Name           string          `json:"Name,omitempty"`
	OperationTimes *OperationTimes `json:"OperationTimes,omitempty"`
	Inclusion      []string        `json:"Inclusion,omitempty"`
	// To represent if the Amenity is included in the rate
	IncludedInd bool `json:"includedInd,omitempty"`
	// To represent if the Amenity attracts a surcharge
	SurchargeInd bool `json:"surchargeInd,omitempty"`
	// OTA code used to describe the property amenity.
	Code string `json:"code,omitempty"`
}

type PropertyDates struct {
	Type_        string                  `json:"@type,omitempty"`
	Availability *AvailabilityStatusEnum `json:"availability,omitempty"`
	AdaCompliant *YesNoUnknownEnum       `json:"adaCompliant,omitempty"`
	Id           string                  `json:"id"`
	// More rates token
	MoreRatesToken string `json:"moreRatesToken,omitempty"`
	// The name of the hotel property
	PropertyName string       `json:"propertyName,omitempty"`
	PropertyKey  *PropertyKey `json:"PropertyKey"`
	DateRange    *DateRange   `json:"DateRange"`
}

type PropertyDetail struct {
	// The OTA code of the property detail
	ClassTypeCode string `json:"classTypeCode,omitempty"`
	// Segment category code
	SegmentCatagoryCode string `json:"segmentCatagoryCode,omitempty"`
	// Location category code
	LocationCatagoryCode string            `json:"locationCatagoryCode,omitempty"`
	ComplimentaryParking *YesNoUnknownEnum `json:"complimentaryParking,omitempty"`
	Address              *Address          `json:"Address,omitempty"`
	Telephone            []string          `json:"Telephone,omitempty"`
	PropertyAmenity      []PropertyAmenity `json:"PropertyAmenity,omitempty"`
	PetPolicy            *PetPolicy        `json:"PetPolicy,omitempty"`
	Restaurant           []Restaurant      `json:"Restaurant,omitempty"`
	Attraction           []Attraction      `json:"Attraction,omitempty"`
	DrivingDirections    *TextFree         `json:"DrivingDirections,omitempty"`
	MeetingRooms         *MeetingRooms     `json:"MeetingRooms,omitempty"`
	VirtualTour          *VirtualTour      `json:"VirtualTour,omitempty"`
	CheckInOutPolicy     *CheckInOutPolicy `json:"CheckInOutPolicy,omitempty"`
	// The property name
	Name                 string                   `json:"name"`
	Rating               []Rating                 `json:"Rating,omitempty"`
	GeoLocation          *GeoLocation             `json:"GeoLocation,omitempty"`
	Image                []Image                  `json:"Image,omitempty"`
	Description          []string                 `json:"Description,omitempty"`
	BusinessService      []Service                `json:"BusinessService,omitempty"`
	AccessibilityFeature []OtaCodeWithDescription `json:"AccessibilityFeature,omitempty"`
	GuestRoomInfo        []GuestRoomInfo          `json:"GuestRoomInfo,omitempty"`
}

type PropertyId struct {
	Type_ string `json:"@type"`
	// Local reference id.
	Id          string       `json:"id,omitempty"`
	Identifier  *Identifier  `json:"Identifier,omitempty"`
	PropertyKey *PropertyKey `json:"PropertyKey"`
}

type PropertyInfo struct {
	Availability        *AvailabilityStatusEnum `json:"availability,omitempty"`
	Distance            *Distance               `json:"Distance"`
	Property            *PropertyDetail         `json:"Property"`
	LowestAvailableRate *CurrencyAmount         `json:"LowestAvailableRate"`
	NextSteps           *NextSteps              `json:"NextSteps,omitempty"`
	// If present and true then this property was added to the list based on criteria other than those in the request
	FeaturedPropertyInd  bool                  `json:"featuredPropertyInd,omitempty"`
	MaximumAvailableRate *MaximumAvailableRate `json:"MaximumAvailableRate,omitempty"`
	Type_                string                `json:"@type,omitempty"`
	Id                   string                `json:"id,omitempty"`
	Identifier           *Identifier           `json:"Identifier,omitempty"`
}

type PropertyInfoId struct {
	Type_      string      `json:"@type,omitempty"`
	Id         string      `json:"id,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type PropertyKey struct {
	Type_ string `json:"@type,omitempty"`
	// Chain code for the property.
	ChainCode string `json:"chainCode"`
	// Code for the property within the hotel chain.
	PropertyCode string `json:"propertyCode"`
}

// The name of the Rail Discount
type RailDiscountCard struct {
	Value string `json:"value,omitempty"`
	// Code of the Supplier
	SupplierCode string `json:"supplierCode"`
	// ReferenceNumber
	ReferenceNumber string `json:"referenceNumber,omitempty"`
}

type RateCandidate struct {
	Type_ string `json:"@type"`
	// rate candidate priority
	Priority int32 `json:"priority,omitempty"`
	// The rateCode to be applied to the request
	RateCode     string            `json:"rateCode,omitempty"`
	RateCategory *RateCategoryEnum `json:"rateCategory,omitempty"`
	// The hotel chain code
	ChainCode string `json:"chainCode,omitempty"`
	// The hotel chain code
	PropertyCode string `json:"propertyCode,omitempty"`
}

type RateCandidateDetail struct {
	// ID of the rate plan associated with the negotiated rate.
	RateID          string           `json:"rateID,omitempty"`
	CustomerLoyalty *CustomerLoyalty `json:"CustomerLoyalty,omitempty"`
	Type_           string           `json:"@type"`
	// rate candidate priority
	Priority int32 `json:"priority,omitempty"`
	// The rateCode to be applied to the request
	RateCode     string            `json:"rateCode,omitempty"`
	RateCategory *RateCategoryEnum `json:"rateCategory,omitempty"`
	// The hotel chain code
	ChainCode string `json:"chainCode,omitempty"`
	// The hotel chain code
	PropertyCode string `json:"propertyCode,omitempty"`
}

type RateCandidates struct {
	Type_         string          `json:"@type"`
	RateCandidate []RateCandidate `json:"RateCandidate"`
	// If true, only prepay rates will be returned
	PrePayRatesOnlyInd bool `json:"prePayRatesOnlyInd,omitempty"`
	// If true, only postpay rates will be returned
	PostPayRatesOnlyInd bool `json:"postPayRatesOnlyInd,omitempty"`
}

type RateCandidatesDetail struct {
	// Minimum number rate plans requested in response
	NumberOfRatePlans int32           `json:"numberOfRatePlans,omitempty"`
	Type_             string          `json:"@type"`
	RateCandidate     []RateCandidate `json:"RateCandidate"`
	// If true, only prepay rates will be returned
	PrePayRatesOnlyInd bool `json:"prePayRatesOnlyInd,omitempty"`
	// If true, only postpay rates will be returned
	PostPayRatesOnlyInd bool `json:"postPayRatesOnlyInd,omitempty"`
}

// Rate Code
type RateCodeInfo struct {
	Value string `json:"value,omitempty"`
	// Rate code name
	RateName string `json:"rateName,omitempty"`
	// Identifier for the rate code
	RateID       string            `json:"rateID,omitempty"`
	RateCategory *RateCategoryEnum `json:"rateCategory,omitempty"`
}

// The actual award or rating received by the facility.
type Rating struct {
	// Rating used to classify hotels according to the quality
	Value float32 `json:"value,omitempty"`
	// The provider who has granted the quality rating
	Provider string `json:"provider,omitempty"`
}

type Receipt struct {
	// Receipt date time
	DateTime time.Time `json:"dateTime,omitempty"`
	// List of offers which links with the receipt
	OfferRef []string `json:"OfferRef,omitempty"`
	// Reference of product
	ProductRef string `json:"ProductRef,omitempty"`
	Type_      string `json:"@type"`
	// The verification number.
	Id         string      `json:"id,omitempty"`
	ReceiptRef string      `json:"ReceiptRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type ReceiptCancellation struct {
	Cancellation *Cancellation `json:"Cancellation,omitempty"`
	// Receipt date time
	DateTime time.Time `json:"dateTime,omitempty"`
	// List of offers which links with the receipt
	OfferRef []string `json:"OfferRef,omitempty"`
	// Reference of product
	ProductRef string `json:"ProductRef,omitempty"`
}

type ReceiptConfirmation struct {
	Confirmation *Confirmation `json:"Confirmation,omitempty"`
	// The segmentSequenceList the ReceiptConfirmation applies to
	SegmentSequenceList []int32 `json:"SegmentSequenceList,omitempty"`
	// Receipt date time
	DateTime time.Time `json:"dateTime,omitempty"`
	// List of offers which links with the receipt
	OfferRef []string `json:"OfferRef,omitempty"`
	// Reference of product
	ProductRef string `json:"ProductRef,omitempty"`
}

type ReceiptId struct {
	Type_ string `json:"@type"`
	// The verification number.
	Id         string      `json:"id,omitempty"`
	ReceiptRef string      `json:"ReceiptRef,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type ReceiptPayment struct {
	PaymentIdentifier *PaymentIdentifier `json:"PaymentIdentifier,omitempty"`
	Document          []Document         `json:"Document,omitempty"`
	// if true, the receipt is for a deposit or prepayment
	DepositInd bool `json:"depositInd,omitempty"`
	// if true, the receipt is for a guarantee only. Guarantee rules are in accordance with the Offer TermsAndConditions.
	GuaranteeInd bool `json:"guaranteeInd,omitempty"`
	// Receipt date time
	DateTime time.Time `json:"dateTime,omitempty"`
	// List of offers which links with the receipt
	OfferRef []string `json:"OfferRef,omitempty"`
	// Reference of product
	ProductRef string `json:"ProductRef,omitempty"`
}

type Reservation struct {
	Offer              []Offer              `json:"Offer,omitempty"`
	Traveler           []Traveler           `json:"Traveler,omitempty"`
	TravelerProduct    []TravelerProduct    `json:"TravelerProduct,omitempty"`
	FormOfPayment      []FormOfPaymentId    `json:"FormOfPayment,omitempty"`
	Payment            []Payment            `json:"Payment,omitempty"`
	Receipt            []Receipt            `json:"Receipt,omitempty"`
	OfferLink          []OfferLink          `json:"OfferLink,omitempty"`
	ReservationComment []ReservationComment `json:"ReservationComment,omitempty"`
	PrimaryContact     []PrimaryContact     `json:"PrimaryContact,omitempty"`
	TravelAgency       *TravelAgency        `json:"TravelAgency,omitempty"`
	// A name assigned to a Reservation containing an offer with Passengerflight/Flight Quantity equal to or greater than 10
	GroupName                  string                       `json:"GroupName,omitempty"`
	SpecialService             []SpecialService             `json:"SpecialService,omitempty"`
	Preference                 *Preference                  `json:"Preference,omitempty"`
	OrganizationLoyaltyProgram []OrganizationLoyaltyProgram `json:"OrganizationLoyaltyProgram,omitempty"`
	ShoppingCart               *ShoppingCart                `json:"ShoppingCart,omitempty"`
	ReservationDisplaySequence *ReservationDisplaySequence  `json:"ReservationDisplaySequence,omitempty"`
	AgencyServiceFee           []AgencyServiceFee           `json:"AgencyServiceFee,omitempty"`
	// The auto delete date represents the date that the Reservation will be kept active. Also known as retention segment or retention date.
	AutoDeleteDate string `json:"autoDeleteDate,omitempty"`
	// The notification date represents the date that the Reservation should be reviewed. Also known as ticket time limit date.
	NotificationDate string `json:"notificationDate,omitempty"`
	Type_            string `json:"@type"`
	// Internal ID
	Id         string      `json:"id,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type ReservationBuild struct {
	Type_               string                 `json:"@type"`
	Traveler            []TravelerId           `json:"Traveler"`
	FormOfPayment       []FormOfPaymentId      `json:"FormOfPayment,omitempty"`
	Payment             []PaymentId            `json:"Payment,omitempty"`
	ReservationComment  []ReservationCommentId `json:"ReservationComment,omitempty"`
	PrimaryContact      []PrimaryContactId     `json:"PrimaryContact,omitempty"`
	SpecialService      []SpecialServiceId     `json:"SpecialService,omitempty"`
	Accounting          *AccountingId          `json:"Accounting,omitempty"`
	DocumentOverrides   []DocumentOverridesId  `json:"DocumentOverrides,omitempty"`
	Preference          []PreferenceId         `json:"Preference,omitempty"`
	ReceiptConfirmation *ReceiptConfirmation   `json:"ReceiptConfirmation,omitempty"`
	TravelAgency        *TravelAgency          `json:"TravelAgency,omitempty"`
}

type ReservationBuildFromCatalogOffering struct {
	BuildFromCatalogOfferingHospitality *BuildFromCatalogOfferingHospitality `json:"BuildFromCatalogOfferingHospitality,omitempty"`
	Type_                               string                               `json:"@type"`
	Traveler                            []TravelerId                         `json:"Traveler"`
	FormOfPayment                       []FormOfPaymentId                    `json:"FormOfPayment,omitempty"`
	Payment                             []PaymentId                          `json:"Payment,omitempty"`
	ReservationComment                  []ReservationCommentId               `json:"ReservationComment,omitempty"`
	PrimaryContact                      []PrimaryContactId                   `json:"PrimaryContact,omitempty"`
	SpecialService                      []SpecialServiceId                   `json:"SpecialService,omitempty"`
	Accounting                          *AccountingId                        `json:"Accounting,omitempty"`
	DocumentOverrides                   []DocumentOverridesId                `json:"DocumentOverrides,omitempty"`
	Preference                          []PreferenceId                       `json:"Preference,omitempty"`
	ReceiptConfirmation                 *ReceiptConfirmation                 `json:"ReceiptConfirmation,omitempty"`
	TravelAgency                        *TravelAgency                        `json:"TravelAgency,omitempty"`
}

type ReservationComment struct {
	CommentSource *CommentSourceEnum `json:"commentSource,omitempty"`
	ShareWith     *ShareWithEnum     `json:"shareWith,omitempty"`
	// Reservation comment shared with supplier
	ShareWithSupplier []string    `json:"shareWithSupplier,omitempty"`
	Comment           []Comment   `json:"Comment,omitempty"`
	AppliesTo         []AppliesTo `json:"AppliesTo,omitempty"`
	Type_             string      `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
}

type ReservationCommentId struct {
	Type_ string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
}

type ReservationDetail struct {
	Accounting         *Accounting          `json:"Accounting,omitempty"`
	DocumentOverrides  []DocumentOverrides  `json:"DocumentOverrides,omitempty"`
	Offer              []Offer              `json:"Offer,omitempty"`
	Traveler           []Traveler           `json:"Traveler,omitempty"`
	TravelerProduct    []TravelerProduct    `json:"TravelerProduct,omitempty"`
	FormOfPayment      []FormOfPaymentId    `json:"FormOfPayment,omitempty"`
	Payment            []Payment            `json:"Payment,omitempty"`
	Receipt            []Receipt            `json:"Receipt,omitempty"`
	OfferLink          []OfferLink          `json:"OfferLink,omitempty"`
	ReservationComment []ReservationComment `json:"ReservationComment,omitempty"`
	PrimaryContact     []PrimaryContact     `json:"PrimaryContact,omitempty"`
	TravelAgency       *TravelAgency        `json:"TravelAgency,omitempty"`
	// A name assigned to a Reservation containing an offer with Passengerflight/Flight Quantity equal to or greater than 10
	GroupName                  string                       `json:"GroupName,omitempty"`
	SpecialService             []SpecialService             `json:"SpecialService,omitempty"`
	Preference                 *Preference                  `json:"Preference,omitempty"`
	OrganizationLoyaltyProgram []OrganizationLoyaltyProgram `json:"OrganizationLoyaltyProgram,omitempty"`
	ShoppingCart               *ShoppingCart                `json:"ShoppingCart,omitempty"`
	ReservationDisplaySequence *ReservationDisplaySequence  `json:"ReservationDisplaySequence,omitempty"`
	AgencyServiceFee           []AgencyServiceFee           `json:"AgencyServiceFee,omitempty"`
	// The auto delete date represents the date that the Reservation will be kept active. Also known as retention segment or retention date.
	AutoDeleteDate string `json:"autoDeleteDate,omitempty"`
	// The notification date represents the date that the Reservation should be reviewed. Also known as ticket time limit date.
	NotificationDate string `json:"notificationDate,omitempty"`
}

type ReservationDetailWrapper struct {
	ReservationDetail *ReservationDetail `json:"ReservationDetail,omitempty"`
}

type ReservationDisplaySequence struct {
	Type_           string            `json:"@type"`
	DisplaySequence []DisplaySequence `json:"DisplaySequence,omitempty"`
	// The sequence of the autoDeleteDate (retention segment) within the Reservation
	AutoDeleteDateSequence int32 `json:"autoDeleteDateSequence,omitempty"`
}

type ReservationId struct {
	Type_ string `json:"@type"`
	// Internal ID
	Id         string      `json:"id,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type ReservationQueryBuild struct {
	Type_            string            `json:"@type"`
	ReservationBuild *ReservationBuild `json:"ReservationBuild"`
}

type ReservationQueryBuildWrapper struct {
	ReservationQueryBuild *ReservationQueryBuild `json:"ReservationQueryBuild,omitempty"`
}

type ReferenceListPropertyDates struct {
	PropertyDates []PropertyDates `json:"PropertyDates"`
	Type_         string          `json:"@type"`
	// Uniquely identifies for the Reference List
	Id string `json:"id,omitempty"`
}

type ReferenceList struct {
	Type_ string `json:"@type"`
	// Uniquely identifies for the Reference List
	Id string `json:"id,omitempty"`
}

type Restaurant struct {
	Type_ string `json:"@type,omitempty"`
	// The name of the restaurant
	Name string `json:"name"`
	// An OTA code to define the cuisine type
	CuisineCodes []string `json:"cuisineCodes,omitempty"`
	// An OTA proximity code
	ProximityCode string    `json:"proximityCode,omitempty"`
	Distance      *Distance `json:"Distance,omitempty"`
}

type Restrictions struct {
	Type_                 string                  `json:"@type,omitempty"`
	TravelerIdentifierRef []TravelerIdentifierRef `json:"TravelerIdentifierRef,omitempty"`
	Restriction           []string                `json:"Restriction"`
	DocumentType          *DocumentTypeEnum       `json:"DocumentType,omitempty"`
}

type Result struct {
	Type_   string            `json:"@type,omitempty"`
	Status  *ResultStatusEnum `json:"status,omitempty"`
	Error_  []ModelError      `json:"Error,omitempty"`
	Warning []Warning         `json:"Warning,omitempty"`
}

type RoomAmenity struct {
	Type_ string `json:"@type,omitempty"`
	// description of the room amenity
	Description string `json:"description,omitempty"`
	// quantity of amenity
	Quantity int32 `json:"quantity,omitempty"`
	// Room Amenity Name
	Name      string   `json:"Name,omitempty"`
	Inclusion []string `json:"Inclusion,omitempty"`
	// To represent if the Amenity is included in the rate
	IncludedInd bool `json:"includedInd,omitempty"`
	// To represent if the Amenity attracts a surcharge
	SurchargeInd bool `json:"surchargeInd,omitempty"`
	// OTA code used to describe the room amenity. This is optional in the Properties Search request but mandatory in the response
	Code string `json:"code,omitempty"`
}

type RoomCharacteristics struct {
	Type_ string `json:"@type,omitempty"`
	// Type code
	TypeCode string `json:"typeCode,omitempty"`
	// Free text describing the view.
	ViewCode string `json:"viewCode,omitempty"`
	// Category of the room.
	Category         string             `json:"category,omitempty"`
	SmokingAllowed   *YesNoUnknownEnum  `json:"smokingAllowed,omitempty"`
	WifiIncluded     *YesNoUnknownEnum  `json:"wifiIncluded,omitempty"`
	BedConfiguration []BedConfiguration `json:"BedConfiguration,omitempty"`
	NonSmokingInd    bool               `json:"nonSmokingInd,omitempty"`
}

type RoomOccupancy struct {
	Type_ string `json:"@type,omitempty"`
	// The minimum occupancy
	MinOccupancy int32 `json:"minOccupancy,omitempty"`
	// The maximum number of room occupants.
	MaxOccupancy  int32           `json:"maxOccupancy,omitempty"`
	AgeQualifying []AgeQualifying `json:"AgeQualifying,omitempty"`
}

type RoomStayCandidate struct {
	GuestCounts *GuestCounts  `json:"GuestCounts"`
	RoomAmenity []RoomAmenity `json:"RoomAmenity,omitempty"`
}

type RoomStayCandidates struct {
	RoomStayCandidate []RoomStayCandidate `json:"RoomStayCandidate"`
}

type RoomType struct {
	Type_               string                   `json:"@type"`
	RoomCharacteristics *RoomCharacteristics     `json:"RoomCharacteristics,omitempty"`
	Description         *TextTitleAndDescription `json:"Description,omitempty"`
	RoomAmenity         []RoomAmenity            `json:"RoomAmenity,omitempty"`
}

type RoomTypeDetail struct {
	// The number of rooms that have been combined to create this room type.
	NumberOfUnits int32 `json:"numberOfUnits,omitempty"`
	// TODO-(Should this be Guarantee?)Denotes the form of guarantee for this room.
	ReqdGuaranteeType string             `json:"reqdGuaranteeType,omitempty"`
	AdditionalDetails *AdditionalDetails `json:"AdditionalDetails,omitempty"`
	RoomOccupancy     []RoomOccupancy    `json:"RoomOccupancy,omitempty"`
	// Indicates the room is a sleeping room when true.
	RoomInd bool `json:"roomInd,omitempty"`
	// Indicates the room is converted when true.
	ConvertedInd bool `json:"convertedInd,omitempty"`
	// Indicates the room is an alternate room type to the requested room type when true.
	AlternateInd        bool                     `json:"alternateInd,omitempty"`
	Type_               string                   `json:"@type"`
	RoomCharacteristics *RoomCharacteristics     `json:"RoomCharacteristics,omitempty"`
	Description         *TextTitleAndDescription `json:"Description,omitempty"`
	RoomAmenity         []RoomAmenity            `json:"RoomAmenity,omitempty"`
}

type SearchAddress struct {
	// City (e.g., Dublin), town, or postal station (i.e., a postal service territory, often used in a military address).
	City      string     `json:"City,omitempty"`
	StateProv *StateProv `json:"StateProv,omitempty"`
	// Country Code
	Country string `json:"Country,omitempty"`
}

type SearchBy struct {
	Type_        string    `json:"@type"`
	SearchRadius *Distance `json:"SearchRadius,omitempty"`
}

type SearchByAddress struct {
	SearchAddress *SearchAddress `json:"SearchAddress"`
	Type_         string         `json:"@type"`
	SearchRadius  *Distance      `json:"SearchRadius,omitempty"`
}

type SearchByAirport struct {
	// Properties located near this IATA airport code
	SearchAirport string    `json:"SearchAirport"`
	Type_         string    `json:"@type"`
	SearchRadius  *Distance `json:"SearchRadius,omitempty"`
}

type SearchByCity struct {
	// Properties located near this IATA city code
	SearchCity   string    `json:"SearchCity"`
	Type_        string    `json:"@type"`
	SearchRadius *Distance `json:"SearchRadius,omitempty"`
}

type SearchByGeoLocation struct {
	// The measure of the angular distance on a meridan north or south equator
	Latitude float64 `json:"Latitude"`
	// The measure of the angular distance on a meridan east or west equator
	Longitude    float64   `json:"Longitude"`
	Type_        string    `json:"@type"`
	SearchRadius *Distance `json:"SearchRadius,omitempty"`
}

type SearchControlConsoleChannelId struct {
	Value string `json:"value,omitempty"`
	// Assigned Type: c-1100:StringTiny
	SccType string `json:"sccType,omitempty"`
}

type SeriesCode struct {
	Type_ string `json:"@type,omitempty"`
	// Note: This contains a key required to retrieve the full payment instrument details compliant with PCI DSS standards.
	EncryptionKey string `json:"encryptionKey,omitempty"`
	// Developer: This contains a reference to the key generation method being used - this is NOT the key value.
	EncryptionKeyMethod string `json:"encryptionKeyMethod,omitempty"`
	// OpenTravel Best Practice: Encryption Method: When using the OpenTravel Encryption element, it is RECOMMENDED that all trading partners be informed of all encryption methods being used in advance of implementation to ensure message processing compatibility.
	EncryptionMethod string `json:"encryptionMethod,omitempty"`
	// Encrypted value
	EncryptedValue string `json:"encryptedValue,omitempty"`
	// Masked Value
	Mask string `json:"mask,omitempty"`
	// Token value
	Token string `json:"token,omitempty"`
	// Developer: This contains a provider ID if multiple providers are used for secure information exchange.
	TokenProviderID      string                       `json:"tokenProviderID,omitempty"`
	AuthenticationMethod *EncryptionTokenTypeAuthEnum `json:"authenticationMethod,omitempty"`
	// Don't use this unless it is REALLY ok to not use encryption. Non-secure (plain text) value.
	PlainText    string        `json:"PlainText,omitempty"`
	ErrorWarning *ErrorWarning `json:"ErrorWarning,omitempty"`
}

// A business service offered by the property
type Service struct {
	// OTA code for this service type.
	Code string `json:"code,omitempty"`
	// Description of the service
	Description string `json:"description,omitempty"`
	// Proximity code
	ProximityCode string `json:"proximityCode,omitempty"`
	// If present and true this service exists
	ExistsInd bool `json:"existsInd,omitempty"`
	// If present and true this service is included with no charge
	IncludedInd bool `json:"includedInd,omitempty"`
}

type ShoppingCart struct {
	Type_   string       `json:"@type,omitempty"`
	Product []ProductAir `json:"Product,omitempty"`
}

type ShoppingCartProductStatus struct {
	Type_ string `json:"@type"`
}

type ShoppingCartProductStatusAir struct {
	StatusAir []StatusAir `json:"StatusAir,omitempty"`
	Type_     string      `json:"@type"`
}

type SignatureOnFile struct {
	Type_               string               `json:"@type,omitempty"`
	DateEffectiveExpire *DateEffectiveExpire `json:"Date_EffectiveExpire,omitempty"`
	// When true, indicates a signature has been obtained.
	SignatureOnFileInd bool `json:"signatureOnFileInd,omitempty"`
}

type SpecialService struct {
	AppliesTo *AppliesTo `json:"AppliesTo,omitempty"`
	Status    *Status    `json:"Status,omitempty"`
	// The type of service animal accompanying the Traveler. If no service animal leave blank.
	ServiceAnimalType  string              `json:"ServiceAnimalType,omitempty"`
	TravelerIdentifier *TravelerIdentifier `json:"TravelerIdentifier,omitempty"`
	Type_              string              `json:"@type"`
	// Internal Id
	Id         string      `json:"id,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

type SpecialServiceId struct {
	Type_ string `json:"@type"`
	// Internal Id
	Id         string      `json:"id,omitempty"`
	Identifier *Identifier `json:"Identifier,omitempty"`
}

// The standard code or abbreviation for the state, province, or region with optional name
type StateProv struct {
	Value string `json:"value,omitempty"`
	// State,province or region name or code needed to identify location
	Name string `json:"name,omitempty"`
}

type Status struct {
	Value *ConfirmationStatusEnum `json:"value,omitempty"`
	// Supplier status text
	SupplierText string `json:"supplierText,omitempty"`
}

type StatusAir struct {
	Value *OfferStatusEnum `json:"value,omitempty"`
	// The flightRefs the status is applicable to within the Offer
	FlightRefs []string `json:"flightRefs,omitempty"`
	// Status code
	Code string `json:"code,omitempty"`
	// If true, the flight is considered to be past date
	PastDateInd bool `json:"pastDateInd,omitempty"`
}

// The fee amount with feecode and reporting informtion
type Surcharge struct {
	Value float32 `json:"value,omitempty"`
	// Sur charge currency code
	CurrencyCode string `json:"currencyCode,omitempty"`
	// Sur charge code
	SurchargeCode string `json:"surchargeCode,omitempty"`
	// Sur charge reporting authority
	ReportingAuthority string `json:"reportingAuthority,omitempty"`
	// Sur charge purpose
	Purpose string `json:"purpose,omitempty"`
	// Description
	Description          string           `json:"description,omitempty"`
	SurchargeApplication *ApplicationEnum `json:"surchargeApplication,omitempty"`
	SurchargeFrequency   *FrequencyEnum   `json:"surchargeFrequency,omitempty"`
	// Surcharge code authority
	CodeAuthority string `json:"codeAuthority,omitempty"`
	// Decimal place for the currency unit
	DecimalPlace int32 `json:"decimalPlace,omitempty"`
	// Currency code decimal authority
	DecimalAuthority string `json:"decimalAuthority,omitempty"`
}

type Surcharges struct {
	Type_ string `json:"@type"`
	// A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	TotalSurcharges float32 `json:"TotalSurcharges,omitempty"`
	// if true, the surcharge amounts are approximate
	ApproximateInd bool `json:"approximateInd,omitempty"`
}

type SurchargesDetail struct {
	Surcharge []Surcharge `json:"Surcharge"`
	Type_     string      `json:"@type"`
	// A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	TotalSurcharges float32 `json:"TotalSurcharges,omitempty"`
	// if true, the surcharge amounts are approximate
	ApproximateInd bool `json:"approximateInd,omitempty"`
}

// The supplier and the supplier's locator code for a product
type SupplierLocator struct {
	Value string `json:"value,omitempty"`
	// Supplier Code
	SupplierCode string `json:"supplierCode,omitempty"`
	// Name of the supplier
	SupplierName string `json:"supplierName,omitempty"`
}

type Tax struct {
	Value float32 `json:"value,omitempty"`
	// Currency code of the city.
	CurrencyCode string `json:"currencyCode,omitempty"`
	// Tax code of the city
	TaxCode string `json:"taxCode,omitempty"`
	// Identifies the reporting authority such as airport code for XF taxes.
	ReportingAuthority string `json:"reportingAuthority,omitempty"`
	// purpose
	Purpose string `json:"purpose,omitempty"`
	// additional information
	Description    string            `json:"description,omitempty"`
	IncludedInBase *YesNoUnknownEnum `json:"includedInBase,omitempty"`
	// Code Authority
	CodeAuthority string `json:"codeAuthority,omitempty"`
	// Allowed number of decimals.
	DecimalPlace int32 `json:"decimalPlace,omitempty"`
	// Decimal Authority
	DecimalAuthority string `json:"decimalAuthority,omitempty"`
	// If true, this tax is exempt
	ExemptInd bool `json:"exemptInd,omitempty"`
}

type TaxBreakdown struct {
	// The airport location the tax applies to
	AirportCode  string        `json:"AirportCode"`
	CurrencyCode *CurrencyCode `json:"CurrencyCode,omitempty"`
	// The amount of the tax applied
	Amount float32 `json:"Amount,omitempty"`
}

type TaxInfo struct {
	// The tax code
	TaxCode      string        `json:"TaxCode"`
	CurrencyCode *CurrencyCode `json:"CurrencyCode,omitempty"`
	// The amount of the tax applied
	Amount float32 `json:"Amount"`
	// The breakdown of the tax for this tax code
	TaxBreakdown []TaxBreakdown `json:"TaxBreakdown"`
}

type TaxPercent struct {
	Value float32 `json:"value,omitempty"`
	// Tax code
	TaxCode string `json:"taxCode,omitempty"`
	// Tax reporting authority
	ReportingAuthority string `json:"reportingAuthority,omitempty"`
	// Purpose of tax
	Purpose string `json:"purpose,omitempty"`
	// Description
	Description    string            `json:"description,omitempty"`
	IncludedInBase *YesNoUnknownEnum `json:"includedInBase,omitempty"`
}

type Taxes struct {
	Type_ string `json:"@type"`
	// A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	TotalTaxes float32   `json:"TotalTaxes,omitempty"`
	TaxInfo    []TaxInfo `json:"TaxInfo,omitempty"`
}

type TaxesDetail struct {
	Tax        []Tax       `json:"Tax,omitempty"`
	TaxPercent *TaxPercent `json:"TaxPercent,omitempty"`
	Type_      string      `json:"@type"`
	// A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	TotalTaxes float32   `json:"TotalTaxes,omitempty"`
	TaxInfo    []TaxInfo `json:"TaxInfo,omitempty"`
}

type Telephone struct {
	Type_ string `json:"@type"`
	// TelephoneCountry AccessCode
	CountryAccessCode string `json:"countryAccessCode,omitempty"`
	// Telephone Area CityCode
	AreaCityCode string `json:"areaCityCode,omitempty"`
	// Mobile/Telephone Number
	PhoneNumber string `json:"phoneNumber"`
	// Telephone extension number
	Extension string `json:"extension,omitempty"`
	// UOptional internally referenced id
	Id string `json:"id,omitempty"`
	// City Code
	CityCode string             `json:"cityCode,omitempty"`
	Role     *EnumTelephoneRole `json:"role,omitempty"`
}

type TermsAndConditions struct {
	// The data and time the offer will expire
	ExpiryDate      time.Time         `json:"ExpiryDate,omitempty"`
	CustomerLoyalty []CustomerLoyalty `json:"CustomerLoyalty,omitempty"`
	Type_           string            `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message.
	TermsAndConditionsRef string      `json:"termsAndConditionsRef,omitempty"`
	Identifier            *Identifier `json:"Identifier,omitempty"`
}

type TermsAndConditionsFull struct {
	// The data and time the offer will expire
	ExpiryDate      time.Time         `json:"ExpiryDate,omitempty"`
	CustomerLoyalty []CustomerLoyalty `json:"CustomerLoyalty,omitempty"`
	TextBlock       []TextBlock       `json:"TextBlock,omitempty"`
	Type_           string            `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message.
	TermsAndConditionsRef string      `json:"termsAndConditionsRef,omitempty"`
	Identifier            *Identifier `json:"Identifier,omitempty"`
}

type TermsAndConditionsFullHospitality struct {
	Guarantee           []Guarantee           `json:"Guarantee,omitempty"`
	CancelPenalty       []CancelPenalty       `json:"CancelPenalty,omitempty"`
	AcceptedCreditCard  []AcceptedCreditCard  `json:"AcceptedCreditCard,omitempty"`
	Description         []string              `json:"Description,omitempty"`
	MealsIncluded       *MealsIncluded        `json:"MealsIncluded,omitempty"`
	ProductRateCodeInfo []ProductRateCodeInfo `json:"ProductRateCodeInfo,omitempty"`
	CheckInOutPolicy    *CheckInOutPolicy     `json:"CheckInOutPolicy,omitempty"`
	DepositPolicy       *DepositPolicy        `json:"DepositPolicy,omitempty"`
	RatePaymentInfo     *RatePaymentEnum      `json:"RatePaymentInfo,omitempty"`
	// The data and time the offer will expire
	ExpiryDate      time.Time         `json:"ExpiryDate,omitempty"`
	CustomerLoyalty []CustomerLoyalty `json:"CustomerLoyalty,omitempty"`
	TextBlock       []TextBlock       `json:"TextBlock,omitempty"`
}

type TermsAndConditionsFullId struct {
	Type_ string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message.
	TermsAndConditionsRef string      `json:"termsAndConditionsRef,omitempty"`
	Identifier            *Identifier `json:"Identifier,omitempty"`
}

type TermsAndConditionsHospitality struct {
	Guarantee           []Guarantee           `json:"Guarantee,omitempty"`
	CancelPenalty       []CancelPenalty       `json:"CancelPenalty,omitempty"`
	AcceptedCreditCard  []AcceptedCreditCard  `json:"AcceptedCreditCard,omitempty"`
	Description         []string              `json:"Description,omitempty"`
	MealsIncluded       *MealsIncluded        `json:"MealsIncluded,omitempty"`
	ProductRateCodeInfo []ProductRateCodeInfo `json:"ProductRateCodeInfo,omitempty"`
	CheckInOutPolicy    *CheckInOutPolicy     `json:"CheckInOutPolicy,omitempty"`
	DepositPolicy       *DepositPolicy        `json:"DepositPolicy,omitempty"`
	RatePaymentInfo     *RatePaymentEnum      `json:"RatePaymentInfo,omitempty"`
	// The data and time the offer will expire
	ExpiryDate      time.Time         `json:"ExpiryDate,omitempty"`
	CustomerLoyalty []CustomerLoyalty `json:"CustomerLoyalty,omitempty"`
}

type TermsAndConditionsId struct {
	Type_ string `json:"@type"`
	// Local indentifier within a given message for this object.
	Id string `json:"id,omitempty"`
	// Used to reference another instance of this object in the same message.
	TermsAndConditionsRef string      `json:"termsAndConditionsRef,omitempty"`
	Identifier            *Identifier `json:"Identifier,omitempty"`
}

type TextBlock struct {
	Type_ string `json:"@type"`
	// Title
	Title string `json:"title,omitempty"`
	// Internally referenced id
	Id            string          `json:"id,omitempty"`
	TextFormatted []TextFormatted `json:"TextFormatted"`
}

// Provides text and indicates whether it is formatted or not.
type TextFormatted struct {
	Value string `json:"value,omitempty"`
	// The language in which the text is provided.
	Language   string          `json:"language,omitempty"`
	TextFormat *TextFormatEnum `json:"textFormat,omitempty"`
}

// Textual information to provide descriptions and/or additional information.
type TextFree struct {
	Value string `json:"value,omitempty"`
	// Language of the text.
	Language string `json:"language,omitempty"`
}

// Descriptive text
type TextTitleAndDescription struct {
	Value string `json:"value,omitempty"`
	// Language of the text
	Languages []string `json:"languages,omitempty"`
	// Title of the Text
	Title string `json:"title,omitempty"`
}

type ThreeDomainSecurity struct {
	Type_                      string                      `json:"@type,omitempty"`
	ThreeDomainSecurityGateway *ThreeDomainSecurityGateway `json:"ThreeDomainSecurityGateway"`
	ThreeDomainSecurityResults *ThreeDomainSecurityResults `json:"ThreeDomainSecurityResults"`
}

type ThreeDomainSecurityGateway struct {
	Type_ string `json:"@type,omitempty"`
	// The eCI value
	ECI string `json:"eCI,omitempty"`
	// The merchant ID value
	MerchantID string `json:"merchantID,omitempty"`
	// The processor ID value
	ProcessorID string `json:"processorID,omitempty"`
	// Transaction URL.
	URL                        string                      `json:"uRL,omitempty"`
	AuthenticationVerification *AuthenticationVerification `json:"AuthenticationVerification,omitempty"`
	Password                   *Password                   `json:"Password,omitempty"`
}

type ThreeDomainSecurityResults struct {
	Type_ string `json:"@type,omitempty"`
	// The cAVV value
	CAVV string `json:"cAVV,omitempty"`
	// The pAResStatus value
	PAResStatus string `json:"pAResStatus,omitempty"`
	// The signature Verification value
	SignatureVerfication string `json:"signatureVerfication,omitempty"`
	// The transaction ID
	TransactionID string `json:"transactionID,omitempty"`
	// Merchants must ensure that each Payer Authentication Request (PAReq) contains a unique combination of account ID and XID
	XID string `json:"xID,omitempty"`
	// Electronic Commerce Indicator - 3-D secure data, contact your authenticator for rules and downline processing.
	ECI string `json:"eCI,omitempty"`
	// Universal Card Authentication Field™ MasterCard only UCAF is the mechanism that is used to transmit the AAV from the merchant to issuer for authentication purposes during the authorization process
	UCAFIndicator string `json:"uCAFIndicator,omitempty"`
}

type TicketDesignators struct {
	Type_                 string                  `json:"@type"`
	TravelerIdentifierRef []TravelerIdentifierRef `json:"TravelerIdentifierRef,omitempty"`
	TicketDesignator      string                  `json:"TicketDesignator"`
}

// The ticketNumber that will be used as partial payment for this Offer/Offering
type TicketNumber struct {
	// Ticket number
	Value string `json:"value,omitempty"`
	// Ticket issuer
	TicketIssuer string `json:"ticketIssuer,omitempty"`
}

type TotalAmount struct {
	Type_          string              `json:"@type,omitempty"`
	CurrencySource *CurrencySourceEnum `json:"currencySource,omitempty"`
	CurrencyCode   *CurrencyCode       `json:"CurrencyCode,omitempty"`
	// The price prior to all applicable taxes of a product such as the rate for a room or fare for a flight.
	Base  float32 `json:"Base,omitempty"`
	Taxes *Taxes  `json:"Taxes,omitempty"`
	Fees  *Fees   `json:"Fees,omitempty"`
	// Specifies the total price including base + taxes + fees
	Total float32 `json:"Total,omitempty"`
	// True if this amount has been converted from the original amount
	ApproximateInd bool `json:"approximateInd,omitempty"`
}

// Tour code
type TourCode struct {
	Value        string            `json:"value,omitempty"`
	TourCodeType *TourCodeTypeEnum `json:"tourCodeType,omitempty"`
}

type TourCodes struct {
	Type_                 string                  `json:"@type,omitempty"`
	TravelerIdentifierRef []TravelerIdentifierRef `json:"TravelerIdentifierRef,omitempty"`
	TourCode              *TourCode               `json:"TourCode"`
}

type TravelAgency struct {
	OrganizationType *OrganizationTypeEnum `json:"organizationType,omitempty"`
	OrganizationName *CompanyName          `json:"OrganizationName"`
	// A reference assigned by the Travel Agency to identify the corporate organization
	CorporateCode string   `json:"CorporateCode,omitempty"`
	ProfileName   []string `json:"ProfileName,omitempty"`
	Type_         string   `json:"@type"`
	// Simple xsd id, not for external use
	Id string `json:"id,omitempty"`
	// An organization that has a name and a structure and members and directly works in the travel industry
	TravelOrganizationRef string      `json:"TravelOrganizationRef,omitempty"`
	Identifier            *Identifier `json:"Identifier,omitempty"`
}

type TravelAgencyDetail struct {
	Telephone        []Telephone           `json:"Telephone,omitempty"`
	Address          []Address             `json:"Address,omitempty"`
	Email            []Email               `json:"Email,omitempty"`
	OrganizationType *OrganizationTypeEnum `json:"organizationType,omitempty"`
	OrganizationName *CompanyName          `json:"OrganizationName"`
	// A reference assigned by the Travel Agency to identify the corporate organization
	CorporateCode string   `json:"CorporateCode,omitempty"`
	ProfileName   []string `json:"ProfileName,omitempty"`
}

type TravelAgencyId struct {
	Type_ string `json:"@type"`
	// Simple xsd id, not for external use
	Id string `json:"id,omitempty"`
	// An organization that has a name and a structure and members and directly works in the travel industry
	TravelOrganizationRef string      `json:"TravelOrganizationRef,omitempty"`
	Identifier            *Identifier `json:"Identifier,omitempty"`
}

type TravelDocument struct {
	Type_ string `json:"@type"`
	// Document number value
	DocNumber string           `json:"docNumber"`
	DocType   *DocTypeCodeEnum `json:"docType,omitempty"`
	// Date of Issue
	IssueDate string `json:"issueDate,omitempty"`
	// Date of expiration
	ExpireDate string `json:"expireDate,omitempty"`
	// State Province Code value
	StateProvCode string `json:"stateProvCode,omitempty"`
	// Place of issue value
	PlaceOfIssue string `json:"placeOfIssue,omitempty"`
	// Issue country on Country Code ISO
	IssueCountry string `json:"issueCountry,omitempty"`
	// The date of birth of the document holder
	BirthDate string `json:"birthDate,omitempty"`
	// Birth country on Country Code ISO value
	BirthCountry string `json:"birthCountry,omitempty"`
	// Birth place value
	BirthPlace string `json:"birthPlace,omitempty"`
	// Residence value
	Residence string `json:"residence,omitempty"`
	// Locally referenced id
	Id     string      `json:"id,omitempty"`
	Gender *GenderEnum `json:"Gender"`
	// Specifies a 2 character country code as defined in ISO3166.
	Nationality string      `json:"Nationality,omitempty"`
	PersonName  *PersonName `json:"PersonName"`
}

type Traveler struct {
	// Date of Birth YYYY-MM-DD
	BirthDate  string      `json:"birthDate,omitempty"`
	Gender     *GenderEnum `json:"gender,omitempty"`
	PersonName *PersonName `json:"PersonName"`
	Address    []Address   `json:"Address,omitempty"`
	Telephone  []Telephone `json:"Telephone,omitempty"`
	Email      []Email     `json:"Email,omitempty"`
	// Passenger type code
	PassengerTypeCode string `json:"passengerTypeCode,omitempty"`
	// Nationality on country code ISO
	Nationality            string             `json:"nationality,omitempty"`
	CustomerLoyalty        []CustomerLoyalty  `json:"CustomerLoyalty,omitempty"`
	AlternateContact       []AlternateContact `json:"AlternateContact,omitempty"`
	TravelDocument         []TravelDocument   `json:"TravelDocument,omitempty"`
	Comments               *Comments          `json:"Comments,omitempty"`
	RailDiscountCard       []RailDiscountCard `json:"RailDiscountCard,omitempty"`
	AccompaniedByInfantInd bool               `json:"accompaniedByInfantInd,omitempty"`
	Type_                  string             `json:"@type"`
	Id                     string             `json:"id,omitempty"`
	TravelerRef            string             `json:"TravelerRef,omitempty"`
	Identifier             *Identifier        `json:"Identifier,omitempty"`
}

type TravelerId struct {
	Type_       string      `json:"@type"`
	Id          string      `json:"id,omitempty"`
	TravelerRef string      `json:"TravelerRef,omitempty"`
	Identifier  *Identifier `json:"Identifier,omitempty"`
}

type TravelerIdentifier struct {
	Id          string      `json:"id,omitempty"`
	TravelerRef string      `json:"TravelerRef,omitempty"`
	Identifier  *Identifier `json:"Identifier,omitempty"`
}

type TravelerIdentifierRef struct {
	Value string `json:"value,omitempty"`
	// A locally referenced ID
	Id string `json:"id,omitempty"`
	// Descriptive text used to identify the contents of a target object
	Description string `json:"description,omitempty"`
	// Uniform Resource Identifier
	Uris []string `json:"uris,omitempty"`
	// Traveler identifier
	Name string `json:"name,omitempty"`
	// Passenger Type code
	PassengerTypeCode string `json:"passengerTypeCode,omitempty"`
}

type TravelerProduct struct {
	// A pointer to the traveler id
	TravelerRef string `json:"TravelerRef,omitempty"`
	// A pointer to the Offer id
	OfferRef string `json:"OfferRef,omitempty"`
	// A pointer to the product id
	ProductRef             string                  `json:"ProductRef,omitempty"`
	ConfirmationStatusEnum *ConfirmationStatusEnum `json:"ConfirmationStatusEnum,omitempty"`
	Type_                  string                  `json:"@type"`
	Id                     string                  `json:"id,omitempty"`
}

type TravelerProductId struct {
	Type_ string `json:"@type"`
	Id    string `json:"id,omitempty"`
}

type VendorCurrencyTotal struct {
	Type_          string              `json:"@type,omitempty"`
	CurrencySource *CurrencySourceEnum `json:"currencySource,omitempty"`
	CurrencyCode   *CurrencyCode       `json:"CurrencyCode,omitempty"`
	// The price prior to all applicable taxes of a product such as the rate for a room or fare for a flight.
	Base  float32 `json:"Base,omitempty"`
	Taxes *Taxes  `json:"Taxes,omitempty"`
	Fees  *Fees   `json:"Fees,omitempty"`
	// Specifies the total price including base + taxes + fees
	Total float32 `json:"Total,omitempty"`
	// True if this amount has been converted from the original amount
	ApproximateInd bool `json:"approximateInd,omitempty"`
}

type VirtualTour struct {
	Type_       string    `json:"@type,omitempty"`
	Url         string    `json:"url"`
	Description *TextFree `json:"Description,omitempty"`
}

type WaiverCode struct {
	Value string `json:"value,omitempty"`
	// A code assigned to identify the reason for disruption
	ReasonCode int32 `json:"reasonCode,omitempty"`
}

type Warning struct {
	Type_ string `json:"@type"`
	// Http standard response code
	StatusCode int32 `json:"StatusCode,omitempty"`
	// The Travelport standardized error or warning message
	Message       string          `json:"Message,omitempty"`
	NameValuePair []NameValuePair `json:"NameValuePair,omitempty"`
}

type WarningDetail struct {
	// The identifier of the source system sending the error or warning
	SourceID string `json:"SourceID"`
	// The error or warning code returned by the source airline or host system
	SourceCode string `json:"SourceCode,omitempty"`
	// The error or warning message as it is returned by the source airline or host system
	SourceDescription string `json:"SourceDescription,omitempty"`
	Type_             string `json:"@type"`
	// Http standard response code
	StatusCode int32 `json:"StatusCode,omitempty"`
	// The Travelport standardized error or warning message
	Message       string          `json:"Message,omitempty"`
	NameValuePair []NameValuePair `json:"NameValuePair,omitempty"`
}
