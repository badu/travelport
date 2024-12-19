package travelport

// AllSomeNoneEnum is an enum. Used to indicate if all, some, or none of the ticket can be exchanged or         refunded
type AllSomeNoneEnum string

const (
	AllSomeNoneEnumAll  AllSomeNoneEnum = "All"
	AllSomeNoneEnumNone AllSomeNoneEnum = "None"
	AllSomeNoneEnumSome AllSomeNoneEnum = "Some"
)

// ApplicableLevelEnum is an enum.
type ApplicableLevelEnum string

const (
	ApplicableLevelEnumItinerary        ApplicableLevelEnum = "Itinerary"
	ApplicableLevelEnumOther            ApplicableLevelEnum = "Other"
	ApplicableLevelEnumPassengerOD      ApplicableLevelEnum = "PassengerOD"
	ApplicableLevelEnumPassengerSegment ApplicableLevelEnum = "PassengerSegment"
	ApplicableLevelEnumPassengers       ApplicableLevelEnum = "Passengers"
	ApplicableLevelEnumSegment          ApplicableLevelEnum = "Segment"
)

// ApplicationEnum is an enum. Application values like perperson , peroom
type ApplicationEnum string

const (
	ApplicationEnumPerAccommodation ApplicationEnum = "PerAccommodation"
	ApplicationEnumPerAdult         ApplicationEnum = "PerAdult"
	ApplicationEnumPerApartment     ApplicationEnum = "PerApartment"
	ApplicationEnumPerChild         ApplicationEnum = "PerChild"
	ApplicationEnumPerHouse         ApplicationEnum = "PerHouse"
	ApplicationEnumPerPerson        ApplicationEnum = "PerPerson"
	ApplicationEnumPerRoom          ApplicationEnum = "PerRoom"
)

// AvailabilitySourceCodeENUM is an enum. A code representing the source of the flight availability
type AvailabilitySourceCodeENUM string

const (
	AvailabilitySourceCodeENUMA AvailabilitySourceCodeENUM = "A"
	AvailabilitySourceCodeENUMB AvailabilitySourceCodeENUM = "B"
	AvailabilitySourceCodeENUMC AvailabilitySourceCodeENUM = "C"
	AvailabilitySourceCodeENUMD AvailabilitySourceCodeENUM = "D"
	AvailabilitySourceCodeENUME AvailabilitySourceCodeENUM = "E"
	AvailabilitySourceCodeENUMF AvailabilitySourceCodeENUM = "F"
	AvailabilitySourceCodeENUMG AvailabilitySourceCodeENUM = "G"
	AvailabilitySourceCodeENUMH AvailabilitySourceCodeENUM = "H"
	AvailabilitySourceCodeENUMI AvailabilitySourceCodeENUM = "I"
	AvailabilitySourceCodeENUMJ AvailabilitySourceCodeENUM = "J"
	AvailabilitySourceCodeENUMK AvailabilitySourceCodeENUM = "K"
	AvailabilitySourceCodeENUML AvailabilitySourceCodeENUM = "L"
	AvailabilitySourceCodeENUMM AvailabilitySourceCodeENUM = "M"
	AvailabilitySourceCodeENUMN AvailabilitySourceCodeENUM = "N"
	AvailabilitySourceCodeENUMO AvailabilitySourceCodeENUM = "O"
	AvailabilitySourceCodeENUMP AvailabilitySourceCodeENUM = "P"
	AvailabilitySourceCodeENUMQ AvailabilitySourceCodeENUM = "Q"
	AvailabilitySourceCodeENUMS AvailabilitySourceCodeENUM = "S"
	AvailabilitySourceCodeENUMT AvailabilitySourceCodeENUM = "T"
	AvailabilitySourceCodeENUMU AvailabilitySourceCodeENUM = "U"
	AvailabilitySourceCodeENUMX AvailabilitySourceCodeENUM = "X"
	AvailabilitySourceCodeENUMY AvailabilitySourceCodeENUM = "Y"
	AvailabilitySourceCodeENUMZ AvailabilitySourceCodeENUM = "Z"
)

// AvailabilityStatusENUMBase is an enum.
type AvailabilityStatusENUMBase string

const (
	AvailabilityStatusENUMBaseAvailable      AvailabilityStatusENUMBase = "Available"
	AvailabilityStatusENUMBaseNotAvailable   AvailabilityStatusENUMBase = "NotAvailable"
	AvailabilityStatusENUMBaseRequest        AvailabilityStatusENUMBase = "Request"
	AvailabilityStatusENUMBaseWaitlist       AvailabilityStatusENUMBase = "Waitlist"
	AvailabilityStatusENUMBaseWaitlistClosed AvailabilityStatusENUMBase = "WaitlistClosed"
)

// BaggageTypeEnum is an enum.
type BaggageTypeEnum string

const (
	BaggageTypeEnumAdditionalBags   BaggageTypeEnum = "AdditionalBags"
	BaggageTypeEnumBaggageEmbargo   BaggageTypeEnum = "BaggageEmbargo"
	BaggageTypeEnumCarryOn          BaggageTypeEnum = "CarryOn"
	BaggageTypeEnumFirstCheckedBag  BaggageTypeEnum = "FirstCheckedBag"
	BaggageTypeEnumSecondCheckedBag BaggageTypeEnum = "SecondCheckedBag"
)

// BatteryTypeEnum is an enum. The type of battery that is used in the device
type BatteryTypeEnum string

const (
	BatteryTypeEnumLithiumION           BatteryTypeEnum = "LithiumION"
	BatteryTypeEnumNoBatteryManualPower BatteryTypeEnum = "NoBatteryManualPower"
	BatteryTypeEnumNonSpillableBattery  BatteryTypeEnum = "NonSpillableBattery"
	BatteryTypeEnumWetCellBattery       BatteryTypeEnum = "WetCellBattery"
)

// BrandClassificationEnum is an enum. The Travelport classification used for a category of ancillaries such as Seat, Bags, etc. This is an initial list that will be added to.
type BrandClassificationEnum string

const (
	BrandClassificationEnumCarryOn        BrandClassificationEnum = "CarryOn"
	BrandClassificationEnumCheckedBag     BrandClassificationEnum = "CheckedBag"
	BrandClassificationEnumLieFlatSeat    BrandClassificationEnum = "LieFlatSeat"
	BrandClassificationEnumMeals          BrandClassificationEnum = "Meals"
	BrandClassificationEnumOther          BrandClassificationEnum = "Other"
	BrandClassificationEnumPersonalItem   BrandClassificationEnum = "PersonalItem"
	BrandClassificationEnumPremiumSeat    BrandClassificationEnum = "PremiumSeat"
	BrandClassificationEnumRebooking      BrandClassificationEnum = "Rebooking"
	BrandClassificationEnumRefund         BrandClassificationEnum = "Refund"
	BrandClassificationEnumSeatAssignment BrandClassificationEnum = "SeatAssignment"
	BrandClassificationEnumWiFi           BrandClassificationEnum = "WiFi"
)

// BrandInclusionEnum is an enum. The indicator as to how the brand and the product are associated.
type BrandInclusionEnum string

const (
	BrandInclusionEnumChargeable BrandInclusionEnum = "Chargeable"
	BrandInclusionEnumIncluded   BrandInclusionEnum = "Included"
	BrandInclusionEnumNotOffered BrandInclusionEnum = "Not Offered"
)

// CabinAirEnum is an enum. Specifies the cabin type (e.g. first, business, economy).
type CabinAirEnum string

const (
	CabinAirEnumBusiness       CabinAirEnum = "Business"
	CabinAirEnumEconomy        CabinAirEnum = "Economy"
	CabinAirEnumFirst          CabinAirEnum = "First"
	CabinAirEnumPremiumEconomy CabinAirEnum = "PremiumEconomy"
	CabinAirEnumPremiumFirst   CabinAirEnum = "PremiumFirst"
)

// CabinPreferenceTypeEnum is an enum.
type CabinPreferenceTypeEnum string

const (
	CabinPreferenceTypeEnumPermitted            CabinPreferenceTypeEnum = "Permitted"
	CabinPreferenceTypeEnumPreferred            CabinPreferenceTypeEnum = "Preferred"
	CabinPreferenceTypeEnumPreferredWithUpgrade CabinPreferenceTypeEnum = "PreferredWithUpgrade"
	CabinPreferenceTypeEnumProhibited           CabinPreferenceTypeEnum = "Prohibited"
)

// CarrierPreferenceTypeEnum is an enum.
type CarrierPreferenceTypeEnum string

const (
	CarrierPreferenceTypeEnumPermitted  CarrierPreferenceTypeEnum = "Permitted"
	CarrierPreferenceTypeEnumPreferred  CarrierPreferenceTypeEnum = "Preferred"
	CarrierPreferenceTypeEnumProhibited CarrierPreferenceTypeEnum = "Prohibited"
)

// ChangeFeeMethodEnumBase is an enum.
type ChangeFeeMethodEnumBase string

const (
	ChangeFeeMethodEnumBaseEMD     ChangeFeeMethodEnumBase = "EMD"
	ChangeFeeMethodEnumBaseMCO     ChangeFeeMethodEnumBase = "MCO"
	ChangeFeeMethodEnumBaseTax     ChangeFeeMethodEnumBase = "Tax"
	ChangeFeeMethodEnumBaseUnknown ChangeFeeMethodEnumBase = "Unknown"
)

// ChangeTypeENUM is an enum.
type ChangeTypeENUM string

const (
	ChangeTypeENUMChangeable      ChangeTypeENUM = "Changeable"
	ChangeTypeENUMNonChangeable   ChangeTypeENUM = "NonChangeable"
	ChangeTypeENUMPenaltyToChange ChangeTypeENUM = "PenaltyToChange"
)

// CityOrAirportEnum is an enum. Clarification of how the airport or city code is used
type CityOrAirportEnum string

const (
	CityOrAirportEnumAirportOnly   CityOrAirportEnum = "Airport Only"
	CityOrAirportEnumCityOnly      CityOrAirportEnum = "City Only"
	CityOrAirportEnumCityOrAirport CityOrAirportEnum = "City or Airport"
	CityOrAirportEnumUseDefault    CityOrAirportEnum = "Use Default"
)

// ClassOfServicePreferenceTypeEnum is an enum.
type ClassOfServicePreferenceTypeEnum string

const (
	ClassOfServicePreferenceTypeEnumPermitted  ClassOfServicePreferenceTypeEnum = "Permitted"
	ClassOfServicePreferenceTypeEnumPreferred  ClassOfServicePreferenceTypeEnum = "Preferred"
	ClassOfServicePreferenceTypeEnumProhibited ClassOfServicePreferenceTypeEnum = "Prohibited"
)

// CommentSourceEnum is an enum.
type CommentSourceEnum string

const (
	CommentSourceEnumAgency   CommentSourceEnum = "Agency"
	CommentSourceEnumSupplier CommentSourceEnum = "Supplier"
	CommentSourceEnumTraveler CommentSourceEnum = "Traveler"
)

// CommissionApplyEnum is an enum.
type CommissionApplyEnum string

const (
	CommissionApplyEnumBase CommissionApplyEnum = "Base"
	CommissionApplyEnumFee  CommissionApplyEnum = "Fee"
)

// CommissionEnum is an enum. Type of commission
type CommissionEnum string

const (
	CommissionEnumAdjustment     CommissionEnum = "Adjustment"
	CommissionEnumCommissionable CommissionEnum = "Commissionable"
	CommissionEnumFull           CommissionEnum = "Full"
	CommissionEnumNoShow         CommissionEnum = "No-show"
	CommissionEnumNonPaying      CommissionEnum = "Non-paying"
	CommissionEnumPartial        CommissionEnum = "Partial"
)

// ConfirmationStatusEnum is an enum. Status returned in a response for a two or more phase commitment process
type ConfirmationStatusEnum string

const (
	ConfirmationStatusEnumCancelled ConfirmationStatusEnum = "Cancelled"
	ConfirmationStatusEnumConfirmed ConfirmationStatusEnum = "Confirmed"
	ConfirmationStatusEnumPending   ConfirmationStatusEnum = "Pending"
	ConfirmationStatusEnumRejected  ConfirmationStatusEnum = "Rejected"
)

// ConnectionPointPreferenceTypeENUM is an enum. Preference type - preferred, permitted or prohibited. Preferred is not permitted as a preference type and will be ignored
type ConnectionPointPreferenceTypeENUM string

const (
	ConnectionPointPreferenceTypeENUMPermitted  ConnectionPointPreferenceTypeENUM = "Permitted"
	ConnectionPointPreferenceTypeENUMPreferred  ConnectionPointPreferenceTypeENUM = "Preferred"
	ConnectionPointPreferenceTypeENUMProhibited ConnectionPointPreferenceTypeENUM = "Prohibited"
)

// ConnectionTypeEnum is an enum.
type ConnectionTypeEnum string

const (
	ConnectionTypeEnumDoubleConnection ConnectionTypeEnum = "DoubleConnection"
	ConnectionTypeEnumNonStopDirect    ConnectionTypeEnum = "NonStopDirect"
	ConnectionTypeEnumSingleConnection ConnectionTypeEnum = "SingleConnection"
	ConnectionTypeEnumStopDirect       ConnectionTypeEnum = "StopDirect"
	ConnectionTypeEnumTripleConnection ConnectionTypeEnum = "TripleConnection"
)

// ContentSourceEnum is an enum. The source of the content to be returned in CatalogOfferings
type ContentSourceEnum string

const (
	ContentSourceEnumGDS ContentSourceEnum = "GDS"
	ContentSourceEnumNDC ContentSourceEnum = "NDC"
)

// CurrencySourceEnum is an enum. The system requesting or returning the currency code specified in the attribute
type CurrencySourceEnum string

const (
	CurrencySourceEnumCharged   CurrencySourceEnum = "Charged"
	CurrencySourceEnumRequested CurrencySourceEnum = "Requested"
	CurrencySourceEnumSupplier  CurrencySourceEnum = "Supplier"
)

// DayOfWeekEnum is an enum. The names of the days of the week.
type DayOfWeekEnum string

const (
	DayOfWeekEnumFriday    DayOfWeekEnum = "Friday"
	DayOfWeekEnumMonday    DayOfWeekEnum = "Monday"
	DayOfWeekEnumSaturday  DayOfWeekEnum = "Saturday"
	DayOfWeekEnumSunday    DayOfWeekEnum = "Sunday"
	DayOfWeekEnumThursday  DayOfWeekEnum = "Thursday"
	DayOfWeekEnumTuesday   DayOfWeekEnum = "Tuesday"
	DayOfWeekEnumWednesday DayOfWeekEnum = "Wednesday"
)

// DestinationEnum is an enum.
type DestinationEnum string

const (
	DestinationEnumAfrica                                 DestinationEnum = "Africa"
	DestinationEnumAsia                                   DestinationEnum = "Asia"
	DestinationEnumAustraliaNewZealandPacificIslands      DestinationEnum = "Australia / New Zealand / Pacific Islands"
	DestinationEnumCanadaAndGreenland                     DestinationEnum = "Canada and Greenland"
	DestinationEnumEurope                                 DestinationEnum = "Europe"
	DestinationEnumIslandsAndCountriesOfTheCaribbean      DestinationEnum = "Islands and Countries of the Caribbean"
	DestinationEnumMexicoCentralAmericaCanalZoneCostaRica DestinationEnum = "Mexico / Central America / Canal Zone/ Costa Rica"
	DestinationEnumMiddleEastWesternAsia                  DestinationEnum = "Middle East / Western Asia"
	DestinationEnumSouthAmerica                           DestinationEnum = "South America"
	DestinationEnumUnitedStatesOfAmerica                  DestinationEnum = "United States of America"
)

// DocTypeCodeEnum is an enum. Codes from OTA DOC - Document Type
type DocTypeCodeEnum string

const (
	DocTypeCodeEnumAirNexusCard                     DocTypeCodeEnum = "AirNexusCard"
	DocTypeCodeEnumAlienRegistrationNumber          DocTypeCodeEnum = "AlienRegistrationNumber"
	DocTypeCodeEnumBoderCrossingCard                DocTypeCodeEnum = "BoderCrossingCard"
	DocTypeCodeEnumCrewMemberCertificate            DocTypeCodeEnum = "CrewMemberCertificate"
	DocTypeCodeEnumDriversLicense                   DocTypeCodeEnum = "DriversLicense"
	DocTypeCodeEnumInsurancePolicyNumber            DocTypeCodeEnum = "InsurancePolicyNumber"
	DocTypeCodeEnumKnownTravelerNumber              DocTypeCodeEnum = "KnownTravelerNumber"
	DocTypeCodeEnumLargeFamilyDiscountCard          DocTypeCodeEnum = "LargeFamilyDiscountCard"
	DocTypeCodeEnumMerchantNumber                   DocTypeCodeEnum = "MerchantNumber"
	DocTypeCodeEnumMilitaryIdentification           DocTypeCodeEnum = "MilitaryIdentification"
	DocTypeCodeEnumNationalIdentityDocument         DocTypeCodeEnum = "NationalIdentityDocument"
	DocTypeCodeEnumNaturalizationCertificate        DocTypeCodeEnum = "NaturalizationCertificate"
	DocTypeCodeEnumNonStandard                      DocTypeCodeEnum = "Non-Standard"
	DocTypeCodeEnumPassport                         DocTypeCodeEnum = "Passport"
	DocTypeCodeEnumPassportCard                     DocTypeCodeEnum = "PassportCard"
	DocTypeCodeEnumPermanentResidentCard            DocTypeCodeEnum = "PermanentResidentCard"
	DocTypeCodeEnumPilotsLicense                    DocTypeCodeEnum = "PilotsLicense"
	DocTypeCodeEnumRedressNumber                    DocTypeCodeEnum = "RedressNumber"
	DocTypeCodeEnumRefugeeTravelDocument            DocTypeCodeEnum = "RefugeeTravelDocument"
	DocTypeCodeEnumTaxExemptionNumber               DocTypeCodeEnum = "TaxExemptionNumber"
	DocTypeCodeEnumTicketNumber                     DocTypeCodeEnum = "TicketNumber"
	DocTypeCodeEnumVaccinationCertificate           DocTypeCodeEnum = "VaccinationCertificate"
	DocTypeCodeEnumVehicleRegistrationLicenseNumber DocTypeCodeEnum = "VehicleRegistrationLicenseNumber"
	DocTypeCodeEnumVisa                             DocTypeCodeEnum = "Visa"
)

// DocumentTypeEnum is an enum. Document type like EMD, MCO
type DocumentTypeEnum string

const (
	DocumentTypeEnumEMD    DocumentTypeEnum = "EMD"
	DocumentTypeEnumMCO    DocumentTypeEnum = "MCO"
	DocumentTypeEnumTicket DocumentTypeEnum = "Ticket"
)

// DurationUnitEnum is an enum. Defines the Units that can be applied to Stay restrictions.
type DurationUnitEnum string

const (
	DurationUnitEnumDays    DurationUnitEnum = "Days"
	DurationUnitEnumFRI     DurationUnitEnum = "FRI"
	DurationUnitEnumHours   DurationUnitEnum = "Hours"
	DurationUnitEnumMON     DurationUnitEnum = "MON"
	DurationUnitEnumMinutes DurationUnitEnum = "Minutes"
	DurationUnitEnumMonths  DurationUnitEnum = "Months"
	DurationUnitEnumSAT     DurationUnitEnum = "SAT"
	DurationUnitEnumSUN     DurationUnitEnum = "SUN"
	DurationUnitEnumTHU     DurationUnitEnum = "THU"
	DurationUnitEnumTUES    DurationUnitEnum = "TUES"
	DurationUnitEnumWED     DurationUnitEnum = "WED"
)

// EMDStatusENUM is an enum.
type EMDStatusENUM string

const (
	EMDStatusENUMOpen   EMDStatusENUM = "Open"
	EMDStatusENUMRefund EMDStatusENUM = "Refund"
	EMDStatusENUMUsed   EMDStatusENUM = "Used"
	EMDStatusENUMVoid   EMDStatusENUM = "Void"
)

// EncryptionTokenTypeAuthEnum is an enum. Type of Authentication
type EncryptionTokenTypeAuthEnum string

const (
	EncryptionTokenTypeAuthEnumMagneticStripe EncryptionTokenTypeAuthEnum = "MagneticStripe"
	EncryptionTokenTypeAuthEnumSecurityCode   EncryptionTokenTypeAuthEnum = "SecurityCode"
)

// EnumAddressRole is an enum.
type EnumAddressRole string

const (
	EnumAddressRoleBilling     EnumAddressRole = "Billing"
	EnumAddressRoleBusiness    EnumAddressRole = "Business"
	EnumAddressRoleDelivery    EnumAddressRole = "Delivery"
	EnumAddressRoleDestination EnumAddressRole = "Destination"
	EnumAddressRoleHome        EnumAddressRole = "Home"
	EnumAddressRoleMailing     EnumAddressRole = "Mailing"
	EnumAddressRoleOther       EnumAddressRole = "Other"
)

// EnumTelephoneRole is an enum.
type EnumTelephoneRole string

const (
	EnumTelephoneRoleFax    EnumTelephoneRole = "Fax"
	EnumTelephoneRoleHome   EnumTelephoneRole = "Home"
	EnumTelephoneRoleMobile EnumTelephoneRole = "Mobile"
	EnumTelephoneRoleOffice EnumTelephoneRole = "Office"
	EnumTelephoneRoleOther  EnumTelephoneRole = "Other"
	EnumTelephoneRoleWork   EnumTelephoneRole = "Work"
)

// ExcludeGroundTypeEnum is an enum.
type ExcludeGroundTypeEnum string

const (
	ExcludeGroundTypeEnumAll   ExcludeGroundTypeEnum = "All"
	ExcludeGroundTypeEnumTrain ExcludeGroundTypeEnum = "Train"
)

// FareQualifierENUMBase is an enum.
type FareQualifierENUMBase string

const (
	FareQualifierENUMBaseConsolidator            FareQualifierENUMBase = "Consolidator"
	FareQualifierENUMBaseGovernment              FareQualifierENUMBase = "Government"
	FareQualifierENUMBaseMarine                  FareQualifierENUMBase = "Marine"
	FareQualifierENUMBaseMilitary                FareQualifierENUMBase = "Military"
	FareQualifierENUMBaseReward                  FareQualifierENUMBase = "Reward"
	FareQualifierENUMBaseStaff                   FareQualifierENUMBase = "Staff"
	FareQualifierENUMBaseStandBy                 FareQualifierENUMBase = "StandBy"
	FareQualifierENUMBaseStudent                 FareQualifierENUMBase = "Student"
	FareQualifierENUMBaseTour                    FareQualifierENUMBase = "Tour"
	FareQualifierENUMBaseVistFriendsAndRelatives FareQualifierENUMBase = "VistFriendsAndRelatives"
	FareQualifierENUMBaseYouth                   FareQualifierENUMBase = "Youth"
)

// FareRuleCategoryEnum is an enum.
type FareRuleCategoryEnum string

const (
	FareRuleCategoryEnumAccompaniedTravel             FareRuleCategoryEnum = "AccompaniedTravel"
	FareRuleCategoryEnumAdvanceReservationsTicketing  FareRuleCategoryEnum = "AdvanceReservationsTicketing"
	FareRuleCategoryEnumAgentDiscounts                FareRuleCategoryEnum = "AgentDiscounts"
	FareRuleCategoryEnumAllOtherDiscounts             FareRuleCategoryEnum = "AllOtherDiscounts"
	FareRuleCategoryEnumApplicationAndOtherConditions FareRuleCategoryEnum = "ApplicationAndOtherConditions"
	FareRuleCategoryEnumBlackoutDates                 FareRuleCategoryEnum = "BlackoutDates"
	FareRuleCategoryEnumChildrenDiscounts             FareRuleCategoryEnum = "ChildrenDiscounts"
	FareRuleCategoryEnumCombinations                  FareRuleCategoryEnum = "Combinations"
	FareRuleCategoryEnumDayTime                       FareRuleCategoryEnum = "DayTime"
	FareRuleCategoryEnumDeposits                      FareRuleCategoryEnum = "Deposits"
	FareRuleCategoryEnumEligibility                   FareRuleCategoryEnum = "Eligibility"
	FareRuleCategoryEnumFareByRule                    FareRuleCategoryEnum = "FareByRule"
	FareRuleCategoryEnumFlightApplication             FareRuleCategoryEnum = "FlightApplication"
	FareRuleCategoryEnumGroups                        FareRuleCategoryEnum = "Groups"
	FareRuleCategoryEnumHIPMileageExeptions           FareRuleCategoryEnum = "HIPMileageExeptions"
	FareRuleCategoryEnumMaximumStay                   FareRuleCategoryEnum = "MaximumStay"
	FareRuleCategoryEnumMinimumStay                   FareRuleCategoryEnum = "MinimumStay"
	FareRuleCategoryEnumMiscellaneousProvisions       FareRuleCategoryEnum = "MiscellaneousProvisions"
	FareRuleCategoryEnumNegotiatedFares               FareRuleCategoryEnum = "NegotiatedFares"
	FareRuleCategoryEnumPenalties                     FareRuleCategoryEnum = "Penalties"
	FareRuleCategoryEnumSalesRestrictions             FareRuleCategoryEnum = "SalesRestrictions"
	FareRuleCategoryEnumSeasonality                   FareRuleCategoryEnum = "Seasonality"
	FareRuleCategoryEnumStopovers                     FareRuleCategoryEnum = "Stopovers"
	FareRuleCategoryEnumSurcharges                    FareRuleCategoryEnum = "Surcharges"
	FareRuleCategoryEnumTicketEndorsements            FareRuleCategoryEnum = "TicketEndorsements"
	FareRuleCategoryEnumTourConductorDiscounts        FareRuleCategoryEnum = "TourConductorDiscounts"
	FareRuleCategoryEnumTours                         FareRuleCategoryEnum = "Tours"
	FareRuleCategoryEnumTransfers                     FareRuleCategoryEnum = "Transfers"
	FareRuleCategoryEnumTravelRestrictions            FareRuleCategoryEnum = "TravelRestrictions"
	FareRuleCategoryEnumVisitAnotherCountry           FareRuleCategoryEnum = "VisitAnotherCountry"
	FareRuleCategoryEnumVoluntaryChanges              FareRuleCategoryEnum = "VoluntaryChanges"
	FareRuleCategoryEnumVoluntaryRefunds              FareRuleCategoryEnum = "VoluntaryRefunds"
)

// FareRuleEnum is an enum.
type FareRuleEnum string

const (
	FareRuleEnumLongText   FareRuleEnum = "LongText"
	FareRuleEnumShortText  FareRuleEnum = "ShortText"
	FareRuleEnumStructured FareRuleEnum = "Structured"
)

// FareTypeEnum is an enum. Defines the type of fares to return (Only public fares, Only private fares, Only agency private fares, Only
type FareTypeEnum string

const (
	FareTypeEnumAgencyPrivateFare  FareTypeEnum = "AgencyPrivateFare"
	FareTypeEnumAirlinePrivateFare FareTypeEnum = "AirlinePrivateFare"
	FareTypeEnumNetFare            FareTypeEnum = "NetFare"
	FareTypeEnumPublicFare         FareTypeEnum = "PublicFare"
)

// FaresFilterEnum is an enum. Defines the type of fares to return (Only public fares, Only private fares, Only agency private fares, Only
type FaresFilterEnum string

const (
	FaresFilterEnumAgencyPrivateFaresOnly  FaresFilterEnum = "AgencyPrivateFaresOnly"
	FaresFilterEnumAirlinePrivateFaresOnly FaresFilterEnum = "AirlinePrivateFaresOnly"
	FaresFilterEnumAllFares                FaresFilterEnum = "AllFares"
	FaresFilterEnumNetFaresOnly            FaresFilterEnum = "NetFaresOnly"
	FaresFilterEnumPrivateFaresOnly        FaresFilterEnum = "PrivateFaresOnly"
	FaresFilterEnumPublicAndPrivateFares   FaresFilterEnum = "PublicAndPrivateFares"
	FaresFilterEnumPublicFaresOnly         FaresFilterEnum = "PublicFaresOnly"
)

// FormOfPaymentTypeEnum is an enum. The list of valid forms of payment.
type FormOfPaymentTypeEnum string

const (
	FormOfPaymentTypeEnumAgencyAccount FormOfPaymentTypeEnum = "AgencyAccount"
	FormOfPaymentTypeEnumBSP           FormOfPaymentTypeEnum = "BSP"
	FormOfPaymentTypeEnumCash          FormOfPaymentTypeEnum = "Cash"
	FormOfPaymentTypeEnumDocument      FormOfPaymentTypeEnum = "Document"
	FormOfPaymentTypeEnumInvoice       FormOfPaymentTypeEnum = "Invoice"
	FormOfPaymentTypeEnumPaymentCard   FormOfPaymentTypeEnum = "PaymentCard"
	FormOfPaymentTypeEnumWaiverCode    FormOfPaymentTypeEnum = "WaiverCode"
)

// FrequencyEnum is an enum. Stay frequency like PerNight, PerDay
type FrequencyEnum string

const (
	FrequencyEnumOneWay    FrequencyEnum = "OneWay"
	FrequencyEnumPerDay    FrequencyEnum = "PerDay"
	FrequencyEnumPerNight  FrequencyEnum = "PerNight"
	FrequencyEnumPerStay   FrequencyEnum = "PerStay"
	FrequencyEnumPerWeek   FrequencyEnum = "PerWeek"
	FrequencyEnumRoundTrip FrequencyEnum = "RoundTrip"
)

// GenderEnum is an enum. Gender Type Male, Female etc. This field is not used by Hotel APIs and will be ignored
type GenderEnum string

const (
	GenderEnumFemale      GenderEnum = "Female"
	GenderEnumMale        GenderEnum = "Male"
	GenderEnumUndisclosed GenderEnum = "Undisclosed"
	GenderEnumUnknown     GenderEnum = "Unknown"
)

// GeoPoliticalAreaLevelEnum is an enum. Represents the type of geopolitical area (country, Continent, State etc)
type GeoPoliticalAreaLevelEnum string

const (
	GeoPoliticalAreaLevelEnumAirport        GeoPoliticalAreaLevelEnum = "Airport"
	GeoPoliticalAreaLevelEnumCity           GeoPoliticalAreaLevelEnum = "City"
	GeoPoliticalAreaLevelEnumContinent      GeoPoliticalAreaLevelEnum = "Continent"
	GeoPoliticalAreaLevelEnumContinentGroup GeoPoliticalAreaLevelEnum = "Continent Group"
	GeoPoliticalAreaLevelEnumCountry        GeoPoliticalAreaLevelEnum = "Country"
	GeoPoliticalAreaLevelEnumGlobalArea     GeoPoliticalAreaLevelEnum = "Global Area"
	GeoPoliticalAreaLevelEnumStateProvince  GeoPoliticalAreaLevelEnum = "StateProvince"
	GeoPoliticalAreaLevelEnumWorld          GeoPoliticalAreaLevelEnum = "World"
)

// GuaranteeTypeEnum is an enum. An enumerated type defining the guarantee to be applied to this reservation.
type GuaranteeTypeEnum string

const (
	GuaranteeTypeEnumCCDCVoucher           GuaranteeTypeEnum = "CC/DC/Voucher"
	GuaranteeTypeEnumDepositNotRequired    GuaranteeTypeEnum = "DepositNotRequired"
	GuaranteeTypeEnumDepositRequired       GuaranteeTypeEnum = "DepositRequired"
	GuaranteeTypeEnumGuaranteeRequired     GuaranteeTypeEnum = "GuaranteeRequired"
	GuaranteeTypeEnumGuaranteesAccepted    GuaranteeTypeEnum = "GuaranteesAccepted"
	GuaranteeTypeEnumGuaranteesNotRequired GuaranteeTypeEnum = "GuaranteesNotRequired"
	GuaranteeTypeEnumNoDepositsAccepted    GuaranteeTypeEnum = "NoDepositsAccepted"
	GuaranteeTypeEnumNoGuaranteesAccepted  GuaranteeTypeEnum = "NoGuaranteesAccepted"
	GuaranteeTypeEnumPrepayNotRequired     GuaranteeTypeEnum = "PrepayNotRequired"
	GuaranteeTypeEnumPrepayRequired        GuaranteeTypeEnum = "PrepayRequired"
	GuaranteeTypeEnumProfile               GuaranteeTypeEnum = "Profile"
)

// IdentifierTypeENUM is an enum.
type IdentifierTypeENUM string

const (
	IdentifierTypeENUMDocumentNumber  IdentifierTypeENUM = "DocumentNumber"
	IdentifierTypeENUMLocator         IdentifierTypeENUM = "Locator"
	IdentifierTypeENUMReservation     IdentifierTypeENUM = "Reservation"
	IdentifierTypeENUMSupplierLocator IdentifierTypeENUM = "SupplierLocator"
)

// ImageSizeEnum is an enum. Indicates the size of the image. Hospitality APIs no longer support thumbnail
type ImageSizeEnum string

const (
	ImageSizeEnumExtraLarge ImageSizeEnum = "ExtraLarge"
	ImageSizeEnumLarge      ImageSizeEnum = "Large"
	ImageSizeEnumMedium     ImageSizeEnum = "Medium"
	ImageSizeEnumSmall      ImageSizeEnum = "Small"
	ImageSizeEnumThumbnail  ImageSizeEnum = "Thumbnail"
)

// ListPaymentCardIssuerEnumBase is an enum. Source: OpenTravel
type ListPaymentCardIssuerEnumBase string

const (
	ListPaymentCardIssuerEnumBaseAmericanExpress     ListPaymentCardIssuerEnumBase = "AmericanExpress"
	ListPaymentCardIssuerEnumBaseBankOfAmerica       ListPaymentCardIssuerEnumBase = "BankOfAmerica"
	ListPaymentCardIssuerEnumBaseBritishAirways      ListPaymentCardIssuerEnumBase = "BritishAirways"
	ListPaymentCardIssuerEnumBaseCapitalOne          ListPaymentCardIssuerEnumBase = "CapitalOne"
	ListPaymentCardIssuerEnumBaseChase               ListPaymentCardIssuerEnumBase = "Chase"
	ListPaymentCardIssuerEnumBaseCitibank            ListPaymentCardIssuerEnumBase = "Citibank"
	ListPaymentCardIssuerEnumBaseContinentalAirlines ListPaymentCardIssuerEnumBase = "ContinentalAirlines"
	ListPaymentCardIssuerEnumBaseDeltaAirlines       ListPaymentCardIssuerEnumBase = "DeltaAirlines"
	ListPaymentCardIssuerEnumBaseDiscoverCard        ListPaymentCardIssuerEnumBase = "DiscoverCard"
	ListPaymentCardIssuerEnumBaseDisney              ListPaymentCardIssuerEnumBase = "Disney"
	ListPaymentCardIssuerEnumBaseEurocard            ListPaymentCardIssuerEnumBase = "Eurocard"
	ListPaymentCardIssuerEnumBaseHilton              ListPaymentCardIssuerEnumBase = "Hilton"
	ListPaymentCardIssuerEnumBaseHyatt               ListPaymentCardIssuerEnumBase = "Hyatt"
	ListPaymentCardIssuerEnumBaseMariott             ListPaymentCardIssuerEnumBase = "Mariott"
	ListPaymentCardIssuerEnumBaseMastercard          ListPaymentCardIssuerEnumBase = "Mastercard"
	ListPaymentCardIssuerEnumBaseRitzCarlton         ListPaymentCardIssuerEnumBase = "RitzCarlton"
	ListPaymentCardIssuerEnumBaseSouthwestAirlines   ListPaymentCardIssuerEnumBase = "SouthwestAirlines"
	ListPaymentCardIssuerEnumBaseStarwoodHotels      ListPaymentCardIssuerEnumBase = "StarwoodHotels"
	ListPaymentCardIssuerEnumBaseUSAirways           ListPaymentCardIssuerEnumBase = "USAirways"
	ListPaymentCardIssuerEnumBaseUnitedAirlines      ListPaymentCardIssuerEnumBase = "UnitedAirlines"
	ListPaymentCardIssuerEnumBaseVISA                ListPaymentCardIssuerEnumBase = "VISA"
)

// MeasurementTypeEnum is an enum. The type of measurement such as width, height, weight
type MeasurementTypeEnum string

const (
	MeasurementTypeEnumDepth            MeasurementTypeEnum = "Depth"
	MeasurementTypeEnumHeight           MeasurementTypeEnum = "Height"
	MeasurementTypeEnumOverallDimension MeasurementTypeEnum = "OverallDimension"
	MeasurementTypeEnumWeight           MeasurementTypeEnum = "Weight"
	MeasurementTypeEnumWidth            MeasurementTypeEnum = "Width"
)

// NameTypeEnum is an enum. OTA Code
type NameTypeEnum string

const (
	NameTypeEnumAlternate NameTypeEnum = "Alternate"
	NameTypeEnumFormer    NameTypeEnum = "Former"
	NameTypeEnumMaiden    NameTypeEnum = "Maiden"
	NameTypeEnumNickname  NameTypeEnum = "Nickname"
)

// NextStepMethodEnum is an enum. Describes the set of potential methods that can be taken after an operation.
type NextStepMethodEnum string

const (
	NextStepMethodEnumDELETE NextStepMethodEnum = "DELETE"
	NextStepMethodEnumGET    NextStepMethodEnum = "GET"
	NextStepMethodEnumPOST   NextStepMethodEnum = "POST"
	NextStepMethodEnumPUT    NextStepMethodEnum = "PUT"
)

// OfferStatusEnum is an enum. Offer Status like confirmed ,Pending etc
type OfferStatusEnum string

const (
	OfferStatusEnumCancelled  OfferStatusEnum = "Cancelled"
	OfferStatusEnumConfirmed  OfferStatusEnum = "Confirmed"
	OfferStatusEnumModifed    OfferStatusEnum = "Modifed"
	OfferStatusEnumPending    OfferStatusEnum = "Pending"
	OfferStatusEnumRejected   OfferStatusEnum = "Rejected"
	OfferStatusEnumWaitlisted OfferStatusEnum = "Waitlisted"
)

// OfferTypeENUM is an enum.
type OfferTypeENUM string

const (
	OfferTypeENUMAgencyCalculatedExchange OfferTypeENUM = "AgencyCalculatedExchange"
	OfferTypeENUMAgencyCalculatedRefund   OfferTypeENUM = "AgencyCalculatedRefund"
)

// OperationalStatusENUMBase is an enum.
type OperationalStatusENUMBase string

const (
	OperationalStatusENUMBaseFlightBoarding               OperationalStatusENUMBase = "FlightBoarding"
	OperationalStatusENUMBaseFlightCancelled              OperationalStatusENUMBase = "FlightCancelled"
	OperationalStatusENUMBaseFlightDeparted               OperationalStatusENUMBase = "FlightDeparted"
	OperationalStatusENUMBaseFlightPastScheduledDeparture OperationalStatusENUMBase = "FlightPastScheduledDeparture"
	OperationalStatusENUMBaseNotAvailableUseSearch        OperationalStatusENUMBase = "NotAvailableUseSearch"
)

// OptInStatusEnum is an enum. Used to indicate marketing preferences, OptIn, OptOut
type OptInStatusEnum string

const (
	OptInStatusEnumOptedIn  OptInStatusEnum = "OptedIn"
	OptInStatusEnumOptedOut OptInStatusEnum = "OptedOut"
	OptInStatusEnumUnknown  OptInStatusEnum = "Unknown"
)

// OrganizationCodeTypeEnum is an enum. Defines the type of code given to the Organization to obtain discounts or additional benefits
type OrganizationCodeTypeEnum string

const (
	OrganizationCodeTypeEnumAccount                    OrganizationCodeTypeEnum = "Account"
	OrganizationCodeTypeEnumOrganizationLoyaltyProgram OrganizationCodeTypeEnum = "OrganizationLoyaltyProgram"
	OrganizationCodeTypeEnumTicketDesignator           OrganizationCodeTypeEnum = "TicketDesignator"
	OrganizationCodeTypeEnumTour                       OrganizationCodeTypeEnum = "Tour"
)

// OrganizationTypeEnum is an enum. The type of organization such as an Agency, Branch, Company, Supplier, Provider
type OrganizationTypeEnum string

const (
	OrganizationTypeEnumAgencyBranch     OrganizationTypeEnum = "AgencyBranch"
	OrganizationTypeEnumIdDocumentIssuer OrganizationTypeEnum = "IdDocumentIssuer"
	OrganizationTypeEnumLoyaltyProgram   OrganizationTypeEnum = "LoyaltyProgram"
	OrganizationTypeEnumRegulatory       OrganizationTypeEnum = "Regulatory"
	OrganizationTypeEnumTravelAgency     OrganizationTypeEnum = "TravelAgency"
	OrganizationTypeEnumTravelProvider   OrganizationTypeEnum = "TravelProvider"
	OrganizationTypeEnumTravelSupplier   OrganizationTypeEnum = "TravelSupplier"
)

// PaymentCardTypeEnum is an enum. Credit, Debit, etc.
type PaymentCardTypeEnum string

const (
	PaymentCardTypeEnumCredit PaymentCardTypeEnum = "Credit"
	PaymentCardTypeEnumDebit  PaymentCardTypeEnum = "Debit"
	PaymentCardTypeEnumGift   PaymentCardTypeEnum = "Gift"
)

// PenaltyAppliesToEnum is an enum. Penalty applicable list
type PenaltyAppliesToEnum string

const (
	PenaltyAppliesToEnumOneWay               PenaltyAppliesToEnum = "OneWay"
	PenaltyAppliesToEnumPerCoupon            PenaltyAppliesToEnum = "PerCoupon"
	PenaltyAppliesToEnumPerDirectionOfTravel PenaltyAppliesToEnum = "PerDirectionOfTravel"
	PenaltyAppliesToEnumPerTicket            PenaltyAppliesToEnum = "PerTicket"
	PenaltyAppliesToEnumRoundTrip            PenaltyAppliesToEnum = "RoundTrip"
)

// PenaltyTypeEnum is an enum.
type PenaltyTypeEnum string

const (
	PenaltyTypeEnumAfterDeparture      PenaltyTypeEnum = "AfterDeparture"
	PenaltyTypeEnumAnytime             PenaltyTypeEnum = "Anytime"
	PenaltyTypeEnumBeforeDeparture     PenaltyTypeEnum = "BeforeDeparture"
	PenaltyTypeEnumExchangeNotRequired PenaltyTypeEnum = "ExchangeNotRequired"
	PenaltyTypeEnumExchangeRequired    PenaltyTypeEnum = "ExchangeRequired"
	PenaltyTypeEnumMaximum             PenaltyTypeEnum = "Maximum"
	PenaltyTypeEnumMinimum             PenaltyTypeEnum = "Minimum"
	PenaltyTypeEnumNoShow              PenaltyTypeEnum = "NoShow"
)

// PercentAppliesTo is an enum. The increment the percent applies to. Default value is Amount
type PercentAppliesTo string

const (
	PercentAppliesToAmount PercentAppliesTo = "Amount"
	PercentAppliesToNights PercentAppliesTo = "Nights"
	PercentAppliesToStay   PercentAppliesTo = "Stay"
)

// PictureofEnum is an enum.
type PictureofEnum string

const (
	PictureofEnumBallroom       PictureofEnum = "Ballroom"
	PictureofEnumBar            PictureofEnum = "Bar"
	PictureofEnumBeach          PictureofEnum = "Beach"
	PictureofEnumBusinessCentre PictureofEnum = "BusinessCentre"
	PictureofEnumConferenceRoom PictureofEnum = "ConferenceRoom"
	PictureofEnumExterior       PictureofEnum = "Exterior"
	PictureofEnumGolf           PictureofEnum = "Golf"
	PictureofEnumGuestRoom      PictureofEnum = "GuestRoom"
	PictureofEnumHealthClub     PictureofEnum = "HealthClub"
	PictureofEnumLobby          PictureofEnum = "Lobby"
	PictureofEnumPool           PictureofEnum = "Pool"
	PictureofEnumPopertyAmenity PictureofEnum = "PopertyAmenity"
	PictureofEnumRecreational   PictureofEnum = "Recreational"
	PictureofEnumRestaurant     PictureofEnum = "Restaurant"
	PictureofEnumRoomAmenity    PictureofEnum = "RoomAmenity"
	PictureofEnumSpa            PictureofEnum = "Spa"
	PictureofEnumSuite          PictureofEnum = "Suite"
)

// PricingEnum is an enum. An enumerated type that defines how a service is priced.
type PricingEnum string

const (
	PricingEnumPerNight          PricingEnum = "Per night"
	PricingEnumPerPerson         PricingEnum = "Per person"
	PricingEnumPerPersonPerNight PricingEnum = "Per person per night"
	PricingEnumPerStay           PricingEnum = "Per stay"
	PricingEnumPerUse            PricingEnum = "Per use"
)

// PurposeEnum is an enum.
type PurposeEnum string

const (
	PurposeEnumBusiness       PurposeEnum = "Business"
	PurposeEnumCharterService PurposeEnum = "Charter Service"
	PurposeEnumPleasure       PurposeEnum = "Pleasure"
)

// RateAvailabilityEnum is an enum. Options are available to sell, need to call, or closed
type RateAvailabilityEnum string

const (
	RateAvailabilityEnumAvailable  RateAvailabilityEnum = "Available"
	RateAvailabilityEnumClosed     RateAvailabilityEnum = "Closed"
	RateAvailabilityEnumNeedToCall RateAvailabilityEnum = "NeedToCall"
)

// RateCategoryEnum is an enum. Rate Category
type RateCategoryEnum string

const (
	RateCategoryEnumAll                       RateCategoryEnum = "All"
	RateCategoryEnumAssociation               RateCategoryEnum = "Association"
	RateCategoryEnumBusiness                  RateCategoryEnum = "Business"
	RateCategoryEnumBusinessStandard          RateCategoryEnum = "BusinessStandard"
	RateCategoryEnumClub                      RateCategoryEnum = "Club"
	RateCategoryEnumConsortiums               RateCategoryEnum = "Consortiums"
	RateCategoryEnumConvention                RateCategoryEnum = "Convention"
	RateCategoryEnumCorporate                 RateCategoryEnum = "Corporate"
	RateCategoryEnumCredential                RateCategoryEnum = "Credential"
	RateCategoryEnumDiscount                  RateCategoryEnum = "Discount"
	RateCategoryEnumEmployee                  RateCategoryEnum = "Employee"
	RateCategoryEnumFamilyPlan                RateCategoryEnum = "FamilyPlan"
	RateCategoryEnumFullInclusive             RateCategoryEnum = "FullInclusive"
	RateCategoryEnumGovernment                RateCategoryEnum = "Government"
	RateCategoryEnumInclusive                 RateCategoryEnum = "Inclusive"
	RateCategoryEnumIndustryTravelAgentRate   RateCategoryEnum = "Industry/TravelAgentRate"
	RateCategoryEnumLeisure                   RateCategoryEnum = "Leisure"
	RateCategoryEnumMilitary                  RateCategoryEnum = "Military"
	RateCategoryEnumMonthly                   RateCategoryEnum = "Monthly"
	RateCategoryEnumMultLevelNegotiatedSecure RateCategoryEnum = "MultLevel/Negotiated/Secure"
	RateCategoryEnumMultiDayPackage           RateCategoryEnum = "Multi-DayPackage"
	RateCategoryEnumOther                     RateCategoryEnum = "Other"
	RateCategoryEnumPackage                   RateCategoryEnum = "Package"
	RateCategoryEnumPrePaid                   RateCategoryEnum = "PrePaid"
	RateCategoryEnumPromotional               RateCategoryEnum = "Promotional"
	RateCategoryEnumRackGeneral               RateCategoryEnum = "RackGeneral"
	RateCategoryEnumSeniorCitizen             RateCategoryEnum = "SeniorCitizen"
	RateCategoryEnumStandard                  RateCategoryEnum = "Standard"
	RateCategoryEnumTour                      RateCategoryEnum = "Tour"
	RateCategoryEnumVIP                       RateCategoryEnum = "VIP"
	RateCategoryEnumWeekend                   RateCategoryEnum = "Weekend"
	RateCategoryEnumWeekly                    RateCategoryEnum = "Weekly"
)

// RatePaymentEnum is an enum. Payment Rate
type RatePaymentEnum string

const (
	RatePaymentEnumPostPay RatePaymentEnum = "PostPay"
	RatePaymentEnumPrePay  RatePaymentEnum = "PrePay"
	RatePaymentEnumUnknown RatePaymentEnum = "Unknown"
)

// RatePeriodEnum is an enum. The time period for a rate such as daily, weekly, monthly
type RatePeriodEnum string

const (
	RatePeriodEnumBundle       RatePeriodEnum = "Bundle"
	RatePeriodEnumDay          RatePeriodEnum = "Day"
	RatePeriodEnumExtraDay     RatePeriodEnum = "ExtraDay"
	RatePeriodEnumExtraHour    RatePeriodEnum = "ExtraHour"
	RatePeriodEnumExtraMonth   RatePeriodEnum = "ExtraMonth"
	RatePeriodEnumExtraWeek    RatePeriodEnum = "ExtraWeek"
	RatePeriodEnumHour         RatePeriodEnum = "Hour"
	RatePeriodEnumMonth        RatePeriodEnum = "Month"
	RatePeriodEnumOther        RatePeriodEnum = "Other"
	RatePeriodEnumPackage      RatePeriodEnum = "Package"
	RatePeriodEnumRentalPeriod RatePeriodEnum = "RentalPeriod"
	RatePeriodEnumTotal        RatePeriodEnum = "Total"
	RatePeriodEnumWeek         RatePeriodEnum = "Week"
	RatePeriodEnumWeekend      RatePeriodEnum = "Weekend"
	RatePeriodEnumYear         RatePeriodEnum = "Year"
)

// RateQualifierEnum is an enum. A closed enumeration of possible rate qualifiers for vehicle rental
type RateQualifierEnum string

const (
	RateQualifierEnumDeposit     RateQualifierEnum = "Deposit"
	RateQualifierEnumGuarantee   RateQualifierEnum = "Guarantee"
	RateQualifierEnumOther       RateQualifierEnum = "Other"
	RateQualifierEnumPostPayment RateQualifierEnum = "PostPayment"
	RateQualifierEnumPrePayment  RateQualifierEnum = "PrePayment"
)

// RefundMethodEnumBase is an enum.
type RefundMethodEnumBase string

const (
	RefundMethodEnumBaseEMD                 RefundMethodEnumBase = "EMD"
	RefundMethodEnumBaseMCO                 RefundMethodEnumBase = "MCO"
	RefundMethodEnumBaseRefundToOriginalFOP RefundMethodEnumBase = "RefundToOriginalFOP"
	RefundMethodEnumBaseUnknown             RefundMethodEnumBase = "Unknown"
)

// RefundTypeEnum is an enum.
type RefundTypeEnum string

const (
	RefundTypeEnumNonRefundable RefundTypeEnum = "NonRefundable"
	RefundTypeEnumPartialRefund RefundTypeEnum = "PartialRefund"
	RefundTypeEnumRefundable    RefundTypeEnum = "Refundable"
)

// RefundabilityEnumBase is an enum.
type RefundabilityEnumBase string

const (
	RefundabilityEnumBaseNonRefundable RefundabilityEnumBase = "NonRefundable"
	RefundabilityEnumBaseRefundable    RefundabilityEnumBase = "Refundable"
	RefundabilityEnumBaseReusable      RefundabilityEnumBase = "Reusable"
)

// ResultStatusEnum is an enum. The status of an error or warning
type ResultStatusEnum string

const (
	ResultStatusEnumComplete     ResultStatusEnum = "Complete"
	ResultStatusEnumIncomplete   ResultStatusEnum = "Incomplete"
	ResultStatusEnumNotProcessed ResultStatusEnum = "Not processed"
	ResultStatusEnumUnknown      ResultStatusEnum = "Unknown"
)

// ScheduleChangeRepriceEnum is an enum.
type ScheduleChangeRepriceEnum string

const (
	ScheduleChangeRepriceEnumAcceptOfferPriceDifference ScheduleChangeRepriceEnum = "AcceptOfferPriceDifference"
	ScheduleChangeRepriceEnumRetainOfferPrice           ScheduleChangeRepriceEnum = "RetainOfferPrice"
)

// SearchRepresentationENUM is an enum. Customize search result set as leg or journey based
type SearchRepresentationENUM string

const (
	SearchRepresentationENUMJourney SearchRepresentationENUM = "Journey"
	SearchRepresentationENUMLeg     SearchRepresentationENUM = "Leg"
)

// SearchTypeEnum is an enum.
type SearchTypeEnum string

const (
	SearchTypeEnumMetaSearch    SearchTypeEnum = "MetaSearch"
	SearchTypeEnumOfferSearch   SearchTypeEnum = "OfferSearch"
	SearchTypeEnumProductSearch SearchTypeEnum = "ProductSearch"
)

// SeatLocationEnum is an enum. Window, aisle, middle, etc.
type SeatLocationEnum string

const (
	SeatLocationEnumAisle  SeatLocationEnum = "Aisle"
	SeatLocationEnumMiddle SeatLocationEnum = "Middle"
	SeatLocationEnumWindow SeatLocationEnum = "Window"
)

// ShareWithEnum is an enum. Share with like Supplier,agency etc
type ShareWithEnum string

const (
	ShareWithEnumAgency   ShareWithEnum = "Agency"
	ShareWithEnumSupplier ShareWithEnum = "Supplier"
	ShareWithEnumTraveler ShareWithEnum = "Traveler"
)

// SpecialMealTypeEnum is an enum. Special Meal Type
type SpecialMealTypeEnum string

const (
	SpecialMealTypeEnumBaby               SpecialMealTypeEnum = "Baby"
	SpecialMealTypeEnumBland              SpecialMealTypeEnum = "Bland"
	SpecialMealTypeEnumChild              SpecialMealTypeEnum = "Child"
	SpecialMealTypeEnumDiabetic           SpecialMealTypeEnum = "Diabetic"
	SpecialMealTypeEnumFruitPlatter       SpecialMealTypeEnum = "FruitPlatter"
	SpecialMealTypeEnumGlutenIntolerant   SpecialMealTypeEnum = "GlutenIntolerant"
	SpecialMealTypeEnumHindu              SpecialMealTypeEnum = "Hindu"
	SpecialMealTypeEnumJain               SpecialMealTypeEnum = "Jain"
	SpecialMealTypeEnumKosher             SpecialMealTypeEnum = "Kosher"
	SpecialMealTypeEnumLowCalorie         SpecialMealTypeEnum = "LowCalorie"
	SpecialMealTypeEnumLowFat             SpecialMealTypeEnum = "LowFat"
	SpecialMealTypeEnumLowSalt            SpecialMealTypeEnum = "LowSalt"
	SpecialMealTypeEnumMuslim             SpecialMealTypeEnum = "Muslim"
	SpecialMealTypeEnumNonLactose         SpecialMealTypeEnum = "NonLactose"
	SpecialMealTypeEnumNone               SpecialMealTypeEnum = "None"
	SpecialMealTypeEnumSeafood            SpecialMealTypeEnum = "Seafood"
	SpecialMealTypeEnumVegan              SpecialMealTypeEnum = "Vegan"
	SpecialMealTypeEnumVegetarianHindu    SpecialMealTypeEnum = "VegetarianHindu"
	SpecialMealTypeEnumVegetarianLactoOvo SpecialMealTypeEnum = "VegetarianLactoOvo"
	SpecialMealTypeEnumVegetarianOriental SpecialMealTypeEnum = "VegetarianOriental"
	SpecialMealTypeEnumVegetarianRaw      SpecialMealTypeEnum = "VegetarianRaw"
)

// TextFormatEnum is an enum. Describes the format of text such as plain text or html
type TextFormatEnum string

const (
	TextFormatEnumHTML      TextFormatEnum = "HTML"
	TextFormatEnumPlainText TextFormatEnum = "PlainText"
)

// TourCodeTypeEnum is an enum.
type TourCodeTypeEnum string

const (
	TourCodeTypeEnumBulkTour      TourCodeTypeEnum = "Bulk Tour"
	TourCodeTypeEnumInclusiveTour TourCodeTypeEnum = "Inclusive Tour"
)

// TransportationCategoryEnum is an enum. Transportation Category
type TransportationCategoryEnum string

const (
	TransportationCategoryEnumAlternate    TransportationCategoryEnum = "Alternate"
	TransportationCategoryEnumBicycle      TransportationCategoryEnum = "Bicycle"
	TransportationCategoryEnumBoat         TransportationCategoryEnum = "Boat"
	TransportationCategoryEnumBus          TransportationCategoryEnum = "Bus"
	TransportationCategoryEnumCableCar     TransportationCategoryEnum = "CableCar"
	TransportationCategoryEnumCar          TransportationCategoryEnum = "Car"
	TransportationCategoryEnumCarRushHour  TransportationCategoryEnum = "Car-RushHour"
	TransportationCategoryEnumCarriage     TransportationCategoryEnum = "Carriage"
	TransportationCategoryEnumCourtesyCar  TransportationCategoryEnum = "Courtesy Car"
	TransportationCategoryEnumExpressTrain TransportationCategoryEnum = "ExpressTrain"
	TransportationCategoryEnumFerry        TransportationCategoryEnum = "Ferry"
	TransportationCategoryEnumHelicopter   TransportationCategoryEnum = "Helicopter"
	TransportationCategoryEnumLimousine    TransportationCategoryEnum = "Limousine"
	TransportationCategoryEnumMetro        TransportationCategoryEnum = "Metro"
	TransportationCategoryEnumMonorail     TransportationCategoryEnum = "Monorail"
	TransportationCategoryEnumMotorbike    TransportationCategoryEnum = "Motorbike"
	TransportationCategoryEnumOther        TransportationCategoryEnum = "Other"
	TransportationCategoryEnumPackAnimal   TransportationCategoryEnum = "Pack Animal"
	TransportationCategoryEnumPlane        TransportationCategoryEnum = "Plane"
	TransportationCategoryEnumPublic       TransportationCategoryEnum = "Public"
	TransportationCategoryEnumRentalCar    TransportationCategoryEnum = "Rental Car"
	TransportationCategoryEnumRickshaw     TransportationCategoryEnum = "Rickshaw"
	TransportationCategoryEnumSedanChair   TransportationCategoryEnum = "SedanChair"
	TransportationCategoryEnumShuttle      TransportationCategoryEnum = "Shuttle"
	TransportationCategoryEnumSubway       TransportationCategoryEnum = "Subway"
	TransportationCategoryEnumTaxi         TransportationCategoryEnum = "Taxi"
	TransportationCategoryEnumTaxiRushHour TransportationCategoryEnum = "Taxi-RushHour"
	TransportationCategoryEnumTrain        TransportationCategoryEnum = "Train"
	TransportationCategoryEnumTrolley      TransportationCategoryEnum = "Trolley"
	TransportationCategoryEnumTube         TransportationCategoryEnum = "Tube"
	TransportationCategoryEnumWalk         TransportationCategoryEnum = "Walk"
	TransportationCategoryEnumWaterTaxi    TransportationCategoryEnum = "WaterTaxi"
)

// TravelerGeographicTypeEnum is an enum.
type TravelerGeographicTypeEnum string

const (
	TravelerGeographicTypeEnumCity          TravelerGeographicTypeEnum = "City"
	TravelerGeographicTypeEnumCountry       TravelerGeographicTypeEnum = "Country"
	TravelerGeographicTypeEnumStateProvince TravelerGeographicTypeEnum = "StateProvince"
)

// UnitOfDistanceEnum is an enum. Miles, Kilometers, etc.
type UnitOfDistanceEnum string

const (
	UnitOfDistanceEnumKilometers UnitOfDistanceEnum = "Kilometers"
	UnitOfDistanceEnumMiles      UnitOfDistanceEnum = "Miles"
)

// UnitOfMeasureEnum is an enum. The unit of measure in a code format. Refer to OpenTravel Code List Unit of Measure Code (UOM).
type UnitOfMeasureEnum string

const (
	UnitOfMeasureEnumAcre             UnitOfMeasureEnum = "Acre"
	UnitOfMeasureEnumBlock            UnitOfMeasureEnum = "Block"
	UnitOfMeasureEnumCentimeters      UnitOfMeasureEnum = "Centimeters"
	UnitOfMeasureEnumCubicMeters      UnitOfMeasureEnum = "Cubic meters"
	UnitOfMeasureEnumFeet             UnitOfMeasureEnum = "Feet"
	UnitOfMeasureEnumGallons          UnitOfMeasureEnum = "Gallons"
	UnitOfMeasureEnumGigabytes        UnitOfMeasureEnum = "Gigabytes"
	UnitOfMeasureEnumGram             UnitOfMeasureEnum = "Gram"
	UnitOfMeasureEnumHectare          UnitOfMeasureEnum = "Hectare"
	UnitOfMeasureEnumInches           UnitOfMeasureEnum = "Inches"
	UnitOfMeasureEnumKilograms        UnitOfMeasureEnum = "Kilograms"
	UnitOfMeasureEnumKilometers       UnitOfMeasureEnum = "Kilometers"
	UnitOfMeasureEnumKilowatts        UnitOfMeasureEnum = "Kilowatts"
	UnitOfMeasureEnumLiters           UnitOfMeasureEnum = "Liters"
	UnitOfMeasureEnumMegabytes        UnitOfMeasureEnum = "Megabytes"
	UnitOfMeasureEnumMeters           UnitOfMeasureEnum = "Meters"
	UnitOfMeasureEnumMiles            UnitOfMeasureEnum = "Miles"
	UnitOfMeasureEnumMillimeters      UnitOfMeasureEnum = "Millimeters"
	UnitOfMeasureEnumOunce            UnitOfMeasureEnum = "Ounce"
	UnitOfMeasureEnumPixels           UnitOfMeasureEnum = "Pixels"
	UnitOfMeasureEnumPounds           UnitOfMeasureEnum = "Pounds"
	UnitOfMeasureEnumSquareCentimeter UnitOfMeasureEnum = "Square centimeter"
	UnitOfMeasureEnumSquareFeet       UnitOfMeasureEnum = "Square feet"
	UnitOfMeasureEnumSquareInch       UnitOfMeasureEnum = "Square inch"
	UnitOfMeasureEnumSquareMeters     UnitOfMeasureEnum = "Square meters"
	UnitOfMeasureEnumSquareMillimeter UnitOfMeasureEnum = "Square millimeter"
	UnitOfMeasureEnumSquareYard       UnitOfMeasureEnum = "Square yard"
	UnitOfMeasureEnumYards            UnitOfMeasureEnum = "Yards"
)

// WaiverEnum is an enum. Type of Waiver like Death of Pessanger,illness Of Passenger, Death of Immediate Family Member  etc
type WaiverEnum string

const (
	WaiverEnumDeathOfImmediateFamilyMember   WaiverEnum = "DeathOfImmediateFamilyMember"
	WaiverEnumDeathOfPassenger               WaiverEnum = "DeathOfPassenger"
	WaiverEnumIllnessOfImmediateFamilyMember WaiverEnum = "IllnessOfImmediateFamilyMember"
	WaiverEnumIllnessOfPassenger             WaiverEnum = "IllnessOfPassenger"
	WaiverEnumScheduleChange                 WaiverEnum = "ScheduleChange"
	WaiverEnumTicketUpgrade                  WaiverEnum = "TicketUpgrade"
)

// YesNoInheritEnum is an enum. Used to indicate marketing preferences, Yes, No, Inherit
type YesNoInheritEnum string

const (
	YesNoInheritEnumInherit YesNoInheritEnum = "Inherit"
	YesNoInheritEnumNo      YesNoInheritEnum = "No"
	YesNoInheritEnumYes     YesNoInheritEnum = "Yes"
)

// YesNoUnknownEnum is an enum. Yes , No , Unknown
type YesNoUnknownEnum string

const (
	YesNoUnknownEnumNo      YesNoUnknownEnum = "No"
	YesNoUnknownEnumUnknown YesNoUnknownEnum = "Unknown"
	YesNoUnknownEnumYes     YesNoUnknownEnum = "Yes"
)
