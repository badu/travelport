package travelport

// ApplicationEnum : Application values like perperson , peroom
type ApplicationEnum string

// List of ApplicationEnum
const (
	PER_PERSON_ApplicationEnum        ApplicationEnum = "PerPerson"
	PER_ROOM_ApplicationEnum          ApplicationEnum = "PerRoom"
	PER_ACCOMMODATION_ApplicationEnum ApplicationEnum = "PerAccommodation"
	PER_HOUSE_ApplicationEnum         ApplicationEnum = "PerHouse"
	PER_APARTMENT_ApplicationEnum     ApplicationEnum = "PerApartment"
	PER_ADULT_ApplicationEnum         ApplicationEnum = "PerAdult"
	PER_CHILD_ApplicationEnum         ApplicationEnum = "PerChild"
)

// AvailabilityStatusEnum : Identifies the availability status of an item.
type AvailabilityStatusEnum string

// List of AvailabilityStatusEnum
const (
	OPEN_AvailabilityStatusEnum                         AvailabilityStatusEnum = "Open"
	CLOSE_AvailabilityStatusEnum                        AvailabilityStatusEnum = "Close"
	CLOSED_ON_ARRIVAL_AvailabilityStatusEnum            AvailabilityStatusEnum = "ClosedOnArrival"
	CLOSED_ON_ARRIVAL_ON_REQUEST_AvailabilityStatusEnum AvailabilityStatusEnum = "ClosedOnArrivalOnRequest"
	ON_REQUEST_AvailabilityStatusEnum                   AvailabilityStatusEnum = "OnRequest"
	REMOVE_CLOSE_ONLY_AvailabilityStatusEnum            AvailabilityStatusEnum = "RemoveCloseOnly"
	OTHER_AvailabilityStatusEnum                        AvailabilityStatusEnum = "Other"
)

// CabinAirEnum : Specifies the cabin type (e.g. first, business, economy).
type CabinAirEnum string

// List of CabinAirEnum
const (
	PREMIUM_FIRST_CabinAirEnum   CabinAirEnum = "PremiumFirst"
	FIRST_CabinAirEnum           CabinAirEnum = "First"
	BUSINESS_CabinAirEnum        CabinAirEnum = "Business"
	PREMIUM_ECONOMY_CabinAirEnum CabinAirEnum = "PremiumEconomy"
	ECONOMY_CabinAirEnum         CabinAirEnum = "Economy"
)

type CommissionApplyEnum string

// List of CommissionApplyEnum
const (
	BASE_CommissionApplyEnum CommissionApplyEnum = "Base"
	FEE_CommissionApplyEnum  CommissionApplyEnum = "Fee"
)

// CommissionEnum : Type of commission
type CommissionEnum string

// List of CommissionEnum
const (
	FULL_CommissionEnum           CommissionEnum = "Full"
	PARTIAL_CommissionEnum        CommissionEnum = "Partial"
	NON_PAYING_CommissionEnum     CommissionEnum = "Non-paying"
	NO_SHOW_CommissionEnum        CommissionEnum = "No-show"
	ADJUSTMENT_CommissionEnum     CommissionEnum = "Adjustment"
	COMMISSIONABLE_CommissionEnum CommissionEnum = "Commissionable"
)

type ChangeFeeMethodEnumBase string

// List of ChangeFeeMethodEnum_Base
const (
	EMD_ChangeFeeMethodEnumBase     ChangeFeeMethodEnumBase = "EMD"
	MCO_ChangeFeeMethodEnumBase     ChangeFeeMethodEnumBase = "MCO"
	TAX_ChangeFeeMethodEnumBase     ChangeFeeMethodEnumBase = "Tax"
	UNKNOWN_ChangeFeeMethodEnumBase ChangeFeeMethodEnumBase = "Unknown"
)

// CurrencySourceEnum : The system requesting or returning the currency code specified in the attribute
type CurrencySourceEnum string

// List of CurrencySourceEnum
const (
	SUPPLIER_CurrencySourceEnum  CurrencySourceEnum = "Supplier"
	CHARGED_CurrencySourceEnum   CurrencySourceEnum = "Charged"
	REQUESTED_CurrencySourceEnum CurrencySourceEnum = "Requested"
)

type CommentSourceEnum string

// List of CommentSourceEnum
const (
	AGENCY_CommentSourceEnum   CommentSourceEnum = "Agency"
	SUPPLIER_CommentSourceEnum CommentSourceEnum = "Supplier"
	TRAVELER_CommentSourceEnum CommentSourceEnum = "Traveler"
)

// ConfirmationStatusEnum : Status returned in a response for a two or more phase commitment process
type ConfirmationStatusEnum string

// List of ConfirmationStatusEnum
const (
	PENDING_ConfirmationStatusEnum   ConfirmationStatusEnum = "Pending"
	CONFIRMED_ConfirmationStatusEnum ConfirmationStatusEnum = "Confirmed"
	CANCELLED_ConfirmationStatusEnum ConfirmationStatusEnum = "Cancelled"
	REJECTED_ConfirmationStatusEnum  ConfirmationStatusEnum = "Rejected"
)

// DayOfWeekEnum : The names of the days of the week.
type DayOfWeekEnum string

// List of DayOfWeekEnum
const (
	SUNDAY_DayOfWeekEnum    DayOfWeekEnum = "Sunday"
	MONDAY_DayOfWeekEnum    DayOfWeekEnum = "Monday"
	TUESDAY_DayOfWeekEnum   DayOfWeekEnum = "Tuesday"
	WEDNESDAY_DayOfWeekEnum DayOfWeekEnum = "Wednesday"
	THURSDAY_DayOfWeekEnum  DayOfWeekEnum = "Thursday"
	FRIDAY_DayOfWeekEnum    DayOfWeekEnum = "Friday"
	SATURDAY_DayOfWeekEnum  DayOfWeekEnum = "Saturday"
)

type DestinationEnum string

// List of DestinationEnum
const (
	UNITED_STATES_OF_AMERICA_DestinationEnum                       DestinationEnum = "United States of America"
	MEXICO__CENTRAL_AMERICA__CANAL_ZONE_COSTA_RICA_DestinationEnum DestinationEnum = "Mexico / Central America / Canal Zone/ Costa Rica"
	ISLANDS_AND_COUNTRIES_OF_THE_CARIBBEAN_DestinationEnum         DestinationEnum = "Islands and Countries of the Caribbean"
	SOUTH_AMERICA_DestinationEnum                                  DestinationEnum = "South America"
	EUROPE_DestinationEnum                                         DestinationEnum = "Europe"
	AFRICA_DestinationEnum                                         DestinationEnum = "Africa"
	MIDDLE_EAST__WESTERN_ASIA_DestinationEnum                      DestinationEnum = "Middle East / Western Asia"
	ASIA_DestinationEnum                                           DestinationEnum = "Asia"
	AUSTRALIA__NEW_ZEALAND__PACIFIC_ISLANDS_DestinationEnum        DestinationEnum = "Australia / New Zealand / Pacific Islands"
	CANADA_AND_GREENLAND_DestinationEnum                           DestinationEnum = "Canada and Greenland"
)

// DocTypeCodeEnum : Codes from OTA DOC - Document Type
type DocTypeCodeEnum string

// List of DocTypeCodeEnum
const (
	VISA_DocTypeCodeEnum                                DocTypeCodeEnum = "Visa"
	PASSPORT_DocTypeCodeEnum                            DocTypeCodeEnum = "Passport"
	MILITARY_IDENTIFICATION_DocTypeCodeEnum             DocTypeCodeEnum = "MilitaryIdentification"
	DRIVERS_LICENSE_DocTypeCodeEnum                     DocTypeCodeEnum = "DriversLicense"
	NATIONAL_IDENTITY_DOCUMENT_DocTypeCodeEnum          DocTypeCodeEnum = "NationalIdentityDocument"
	VACCINATION_CERTIFICATE_DocTypeCodeEnum             DocTypeCodeEnum = "VaccinationCertificate"
	ALIEN_REGISTRATION_NUMBER_DocTypeCodeEnum           DocTypeCodeEnum = "AlienRegistrationNumber"
	INSURANCE_POLICY_NUMBER_DocTypeCodeEnum             DocTypeCodeEnum = "InsurancePolicyNumber"
	TAX_EXEMPTION_NUMBER_DocTypeCodeEnum                DocTypeCodeEnum = "TaxExemptionNumber"
	VEHICLE_REGISTRATION_LICENSE_NUMBER_DocTypeCodeEnum DocTypeCodeEnum = "VehicleRegistrationLicenseNumber"
	BODER_CROSSING_CARD_DocTypeCodeEnum                 DocTypeCodeEnum = "BoderCrossingCard"
	REFUGEE_TRAVEL_DOCUMENT_DocTypeCodeEnum             DocTypeCodeEnum = "RefugeeTravelDocument"
	PILOTS_LICENSE_DocTypeCodeEnum                      DocTypeCodeEnum = "PilotsLicense"
	PERMANENT_RESIDENT_CARD_DocTypeCodeEnum             DocTypeCodeEnum = "PermanentResidentCard"
	REDRESS_NUMBER_DocTypeCodeEnum                      DocTypeCodeEnum = "RedressNumber"
	KNOWN_TRAVELER_NUMBER_DocTypeCodeEnum               DocTypeCodeEnum = "KnownTravelerNumber"
	NON_STANDARD_DocTypeCodeEnum                        DocTypeCodeEnum = "Non-Standard"
	MERCHANT_NUMBER_DocTypeCodeEnum                     DocTypeCodeEnum = "MerchantNumber"
	AIR_NEXUS_CARD_DocTypeCodeEnum                      DocTypeCodeEnum = "AirNexusCard"
	CREW_MEMBER_CERTIFICATE_DocTypeCodeEnum             DocTypeCodeEnum = "CrewMemberCertificate"
	PASSPORT_CARD_DocTypeCodeEnum                       DocTypeCodeEnum = "PassportCard"
	NATURALIZATION_CERTIFICATE_DocTypeCodeEnum          DocTypeCodeEnum = "NaturalizationCertificate"
	TICKET_NUMBER_DocTypeCodeEnum                       DocTypeCodeEnum = "TicketNumber"
	LARGE_FAMILY_DISCOUNT_CARD_DocTypeCodeEnum          DocTypeCodeEnum = "LargeFamilyDiscountCard"
)

type DocumentTypeEnum string

// List of DocumentTypeEnum
const (
	EMD_DocumentTypeEnum    DocumentTypeEnum = "EMD"
	MCO_DocumentTypeEnum    DocumentTypeEnum = "MCO"
	TICKET_DocumentTypeEnum DocumentTypeEnum = "Ticket"
)

// DurationUnitEnum : Defines the Units that can be applied to Stay restrictions.
type DurationUnitEnum string

// List of DurationUnitEnum
const (
	MINUTES_DurationUnitEnum DurationUnitEnum = "Minutes"
	HOURS_DurationUnitEnum   DurationUnitEnum = "Hours"
	DAYS_DurationUnitEnum    DurationUnitEnum = "Days"
	MONTHS_DurationUnitEnum  DurationUnitEnum = "Months"
	MON_DurationUnitEnum     DurationUnitEnum = "MON"
	TUES_DurationUnitEnum    DurationUnitEnum = "TUES"
	WED_DurationUnitEnum     DurationUnitEnum = "WED"
	THU_DurationUnitEnum     DurationUnitEnum = "THU"
	FRI_DurationUnitEnum     DurationUnitEnum = "FRI"
	SAT_DurationUnitEnum     DurationUnitEnum = "SAT"
	SUN_DurationUnitEnum     DurationUnitEnum = "SUN"
)

// EncryptionTokenTypeAuthEnum : Type of Authentication
type EncryptionTokenTypeAuthEnum string

// List of EncryptionTokenTypeAuthEnum
const (
	SECURITY_CODE_EncryptionTokenTypeAuthEnum   EncryptionTokenTypeAuthEnum = "SecurityCode"
	MAGNETIC_STRIPE_EncryptionTokenTypeAuthEnum EncryptionTokenTypeAuthEnum = "MagneticStripe"
)

type EnumAddressRole string

// List of Enum_AddressRole
const (
	HOME_EnumAddressRole        EnumAddressRole = "Home"
	BUSINESS_EnumAddressRole    EnumAddressRole = "Business"
	MAILING_EnumAddressRole     EnumAddressRole = "Mailing"
	DELIVERY_EnumAddressRole    EnumAddressRole = "Delivery"
	DESTINATION_EnumAddressRole EnumAddressRole = "Destination"
	OTHER_EnumAddressRole       EnumAddressRole = "Other"
	BILLING_EnumAddressRole     EnumAddressRole = "Billing"
)

type EnumTelephoneRole string

// List of Enum_TelephoneRole
const (
	MOBILE_EnumTelephoneRole EnumTelephoneRole = "Mobile"
	HOME_EnumTelephoneRole   EnumTelephoneRole = "Home"
	WORK_EnumTelephoneRole   EnumTelephoneRole = "Work"
	OFFICE_EnumTelephoneRole EnumTelephoneRole = "Office"
	FAX_EnumTelephoneRole    EnumTelephoneRole = "Fax"
	OTHER_EnumTelephoneRole  EnumTelephoneRole = "Other"
)

type EmdStatusEnum string

// List of EMDStatusENUM
const (
	OPEN_EmdStatusEnum   EmdStatusEnum = "Open"
	REFUND_EmdStatusEnum EmdStatusEnum = "Refund"
	USED_EmdStatusEnum   EmdStatusEnum = "Used"
	VOID_EmdStatusEnum   EmdStatusEnum = "Void"
)

type FareQualifierEnumBase string

// List of FareQualifierENUM_Base
const (
	CONSOLIDATOR_FareQualifierEnumBase               FareQualifierEnumBase = "Consolidator"
	GOVERNMENT_FareQualifierEnumBase                 FareQualifierEnumBase = "Government"
	MARINE_FareQualifierEnumBase                     FareQualifierEnumBase = "Marine"
	MILITARY_FareQualifierEnumBase                   FareQualifierEnumBase = "Military"
	REWARD_FareQualifierEnumBase                     FareQualifierEnumBase = "Reward"
	STAND_BY_FareQualifierEnumBase                   FareQualifierEnumBase = "StandBy"
	STAFF_FareQualifierEnumBase                      FareQualifierEnumBase = "Staff"
	STUDENT_FareQualifierEnumBase                    FareQualifierEnumBase = "Student"
	TOUR_FareQualifierEnumBase                       FareQualifierEnumBase = "Tour"
	YOUTH_FareQualifierEnumBase                      FareQualifierEnumBase = "Youth"
	VIST_FRIENDS_AND_RELATIVES_FareQualifierEnumBase FareQualifierEnumBase = "VistFriendsAndRelatives"
)

// FareTypeEnum : Defines the type of fares to return (Only public fares, Only private fares, Only agency private fares, Only
type FareTypeEnum string

// List of FareTypeEnum
const (
	PUBLIC_FARE_FareTypeEnum          FareTypeEnum = "PublicFare"
	AGENCY_PRIVATE_FARE_FareTypeEnum  FareTypeEnum = "AgencyPrivateFare"
	AIRLINE_PRIVATE_FARE_FareTypeEnum FareTypeEnum = "AirlinePrivateFare"
	NET_FARE_FareTypeEnum             FareTypeEnum = "NetFare"
)

// FormOfPaymentTypeEnum : The list of valid forms of payment.
type FormOfPaymentTypeEnum string

// List of FormOfPaymentTypeEnum
const (
	AGENCY_ACCOUNT_FormOfPaymentTypeEnum FormOfPaymentTypeEnum = "AgencyAccount"
	BSP_FormOfPaymentTypeEnum            FormOfPaymentTypeEnum = "BSP"
	CASH_FormOfPaymentTypeEnum           FormOfPaymentTypeEnum = "Cash"
	DOCUMENT_FormOfPaymentTypeEnum       FormOfPaymentTypeEnum = "Document"
	INVOICE_FormOfPaymentTypeEnum        FormOfPaymentTypeEnum = "Invoice"
	PAYMENT_CARD_FormOfPaymentTypeEnum   FormOfPaymentTypeEnum = "PaymentCard"
	WAIVER_CODE_FormOfPaymentTypeEnum    FormOfPaymentTypeEnum = "WaiverCode"
)

// FrequencyEnum : Stay frequency like PerNight, PerDay
type FrequencyEnum string

// List of FrequencyEnum
const (
	PER_NIGHT_FrequencyEnum  FrequencyEnum = "PerNight"
	PER_DAY_FrequencyEnum    FrequencyEnum = "PerDay"
	PER_STAY_FrequencyEnum   FrequencyEnum = "PerStay"
	PER_WEEK_FrequencyEnum   FrequencyEnum = "PerWeek"
	ROUND_TRIP_FrequencyEnum FrequencyEnum = "RoundTrip"
	ONE_WAY_FrequencyEnum    FrequencyEnum = "OneWay"
)

// GenderEnum : Gender Type Male, Female etc. This field is not used by Hotel APIs and will be ignored
type GenderEnum string

// List of GenderEnum
const (
	MALE_GenderEnum        GenderEnum = "Male"
	FEMALE_GenderEnum      GenderEnum = "Female"
	UNKNOWN_GenderEnum     GenderEnum = "Unknown"
	UNDISCLOSED_GenderEnum GenderEnum = "Undisclosed"
)

// GuaranteeTypeEnum : An enumerated type defining the guarantee to be applied to this reservation.
type GuaranteeTypeEnum string

// List of GuaranteeTypeEnum
const (
	GUARANTEE_REQUIRED_GuaranteeTypeEnum      GuaranteeTypeEnum = "GuaranteeRequired"
	CCDCVOUCHER_GuaranteeTypeEnum             GuaranteeTypeEnum = "CC/DC/Voucher"
	PROFILE_GuaranteeTypeEnum                 GuaranteeTypeEnum = "Profile"
	NO_GUARANTEES_ACCEPTED_GuaranteeTypeEnum  GuaranteeTypeEnum = "NoGuaranteesAccepted"
	GUARANTEES_ACCEPTED_GuaranteeTypeEnum     GuaranteeTypeEnum = "GuaranteesAccepted"
	DEPOSIT_REQUIRED_GuaranteeTypeEnum        GuaranteeTypeEnum = "DepositRequired"
	GUARANTEES_NOT_REQUIRED_GuaranteeTypeEnum GuaranteeTypeEnum = "GuaranteesNotRequired"
	DEPOSIT_NOT_REQUIRED_GuaranteeTypeEnum    GuaranteeTypeEnum = "DepositNotRequired"
	PREPAY_REQUIRED_GuaranteeTypeEnum         GuaranteeTypeEnum = "PrepayRequired"
	PREPAY_NOT_REQUIRED_GuaranteeTypeEnum     GuaranteeTypeEnum = "PrepayNotRequired"
	NO_DEPOSITS_ACCEPTED_GuaranteeTypeEnum    GuaranteeTypeEnum = "NoDepositsAccepted"
)

type HotelAggregatorEnum string

// List of HotelAggregatorEnum
const (
	TRAVELPORT_HotelAggregatorEnum HotelAggregatorEnum = "Travelport"
	AGODA_HotelAggregatorEnum      HotelAggregatorEnum = "Agoda"
	BOOKING_HotelAggregatorEnum    HotelAggregatorEnum = "Booking"
	EXPEDIA_HotelAggregatorEnum    HotelAggregatorEnum = "Expedia"
	BONOTEL_HotelAggregatorEnum    HotelAggregatorEnum = "Bonotel"
)

// HotelSortOrderEnum : The method to be used in sorting hotel properties
type HotelSortOrderEnum string

// List of HotelSortOrderEnum
const (
	STAR_RATING_HotelSortOrderEnum HotelSortOrderEnum = "StarRating"
	PROXIMITY_HotelSortOrderEnum   HotelSortOrderEnum = "Proximity"
)

type IdentifierTypeEnum string

// List of IdentifierTypeENUM
const (
	RESERVATION_IdentifierTypeEnum      IdentifierTypeEnum = "Reservation"
	LOCATOR_IdentifierTypeEnum          IdentifierTypeEnum = "Locator"
	SUPPLIER_LOCATOR_IdentifierTypeEnum IdentifierTypeEnum = "SupplierLocator"
	DOCUMENT_NUMBER_IdentifierTypeEnum  IdentifierTypeEnum = "DocumentNumber"
)

// MeasurementTypeEnum : The type of measurement such as width, height, weight
type MeasurementTypeEnum string

// List of MeasurementTypeEnum
const (
	WIDTH_MeasurementTypeEnum             MeasurementTypeEnum = "Width"
	HEIGHT_MeasurementTypeEnum            MeasurementTypeEnum = "Height"
	DEPTH_MeasurementTypeEnum             MeasurementTypeEnum = "Depth"
	WEIGHT_MeasurementTypeEnum            MeasurementTypeEnum = "Weight"
	OVERALL_DIMENSION_MeasurementTypeEnum MeasurementTypeEnum = "OverallDimension"
)

// ListPaymentCardIssuerEnumBase : Source: OpenTravel
type ListPaymentCardIssuerEnumBase string

// List of ListPaymentCardIssuerEnum_Base
const (
	VISA_ListPaymentCardIssuerEnumBase                 ListPaymentCardIssuerEnumBase = "VISA"
	US_AIRWAYS_ListPaymentCardIssuerEnumBase           ListPaymentCardIssuerEnumBase = "USAirways"
	UNITED_AIRLINES_ListPaymentCardIssuerEnumBase      ListPaymentCardIssuerEnumBase = "UnitedAirlines"
	STARWOOD_HOTELS_ListPaymentCardIssuerEnumBase      ListPaymentCardIssuerEnumBase = "StarwoodHotels"
	SOUTHWEST_AIRLINES_ListPaymentCardIssuerEnumBase   ListPaymentCardIssuerEnumBase = "SouthwestAirlines"
	RITZ_CARLTON_ListPaymentCardIssuerEnumBase         ListPaymentCardIssuerEnumBase = "RitzCarlton"
	MASTERCARD_ListPaymentCardIssuerEnumBase           ListPaymentCardIssuerEnumBase = "Mastercard"
	MARIOTT_ListPaymentCardIssuerEnumBase              ListPaymentCardIssuerEnumBase = "Mariott"
	HYATT_ListPaymentCardIssuerEnumBase                ListPaymentCardIssuerEnumBase = "Hyatt"
	HILTON_ListPaymentCardIssuerEnumBase               ListPaymentCardIssuerEnumBase = "Hilton"
	EUROCARD_ListPaymentCardIssuerEnumBase             ListPaymentCardIssuerEnumBase = "Eurocard"
	DISNEY_ListPaymentCardIssuerEnumBase               ListPaymentCardIssuerEnumBase = "Disney"
	DISCOVER_CARD_ListPaymentCardIssuerEnumBase        ListPaymentCardIssuerEnumBase = "DiscoverCard"
	DELTA_AIRLINES_ListPaymentCardIssuerEnumBase       ListPaymentCardIssuerEnumBase = "DeltaAirlines"
	CONTINENTAL_AIRLINES_ListPaymentCardIssuerEnumBase ListPaymentCardIssuerEnumBase = "ContinentalAirlines"
	CITIBANK_ListPaymentCardIssuerEnumBase             ListPaymentCardIssuerEnumBase = "Citibank"
	CHASE_ListPaymentCardIssuerEnumBase                ListPaymentCardIssuerEnumBase = "Chase"
	CAPITAL_ONE_ListPaymentCardIssuerEnumBase          ListPaymentCardIssuerEnumBase = "CapitalOne"
	BRITISH_AIRWAYS_ListPaymentCardIssuerEnumBase      ListPaymentCardIssuerEnumBase = "BritishAirways"
	BANK_OF_AMERICA_ListPaymentCardIssuerEnumBase      ListPaymentCardIssuerEnumBase = "BankOfAmerica"
	AMERICAN_EXPRESS_ListPaymentCardIssuerEnumBase     ListPaymentCardIssuerEnumBase = "AmericanExpress"
)

// ImageSizeEnum : Indicates the size of the image. Hospitality APIs no longer support thumbnail
type ImageSizeEnum string

// List of ImageSizeEnum
const (
	LARGE_ImageSizeEnum       ImageSizeEnum = "Large"
	MEDIUM_ImageSizeEnum      ImageSizeEnum = "Medium"
	SMALL_ImageSizeEnum       ImageSizeEnum = "Small"
	THUMBNAIL_ImageSizeEnum   ImageSizeEnum = "Thumbnail"
	EXTRA_LARGE_ImageSizeEnum ImageSizeEnum = "ExtraLarge"
)

// NextStepMethodEnum : Describes the set of potential methods that can be taken after an operation.
type NextStepMethodEnum string

// List of NextStepMethodEnum
const (
	GET_NextStepMethodEnum    NextStepMethodEnum = "GET"
	DELETE_NextStepMethodEnum NextStepMethodEnum = "DELETE"
	PUT_NextStepMethodEnum    NextStepMethodEnum = "PUT"
	POST_NextStepMethodEnum   NextStepMethodEnum = "POST"
)

type PurposeEnum string

// List of PurposeEnum
const (
	BUSINESS_PurposeEnum        PurposeEnum = "Business"
	PLEASURE_PurposeEnum        PurposeEnum = "Pleasure"
	CHARTER_SERVICE_PurposeEnum PurposeEnum = "Charter Service"
)

// ResultStatusEnum : The status of an error or warning
type ResultStatusEnum string

// List of ResultStatusEnum
const (
	NOT_PROCESSED_ResultStatusEnum ResultStatusEnum = "Not processed"
	INCOMPLETE_ResultStatusEnum    ResultStatusEnum = "Incomplete"
	COMPLETE_ResultStatusEnum      ResultStatusEnum = "Complete"
	UNKNOWN_ResultStatusEnum       ResultStatusEnum = "Unknown"
)

type TourCodeTypeEnum string

// List of TourCodeTypeEnum
const (
	BULK_TOUR_TourCodeTypeEnum      TourCodeTypeEnum = "Bulk Tour"
	INCLUSIVE_TOUR_TourCodeTypeEnum TourCodeTypeEnum = "Inclusive Tour"
)

// YesNoUnknownEnum : Yes , No , Unknown
type YesNoUnknownEnum string

// List of YesNoUnknownEnum
const (
	YES_YesNoUnknownEnum     YesNoUnknownEnum = "Yes"
	NO_YesNoUnknownEnum      YesNoUnknownEnum = "No"
	UNKNOWN_YesNoUnknownEnum YesNoUnknownEnum = "Unknown"
)

// OfferStatusEnum : Offer Status like confirmed ,Pending etc
type OfferStatusEnum string

// List of OfferStatusEnum
const (
	CONFIRMED_OfferStatusEnum  OfferStatusEnum = "Confirmed"
	CANCELLED_OfferStatusEnum  OfferStatusEnum = "Cancelled"
	PENDING_OfferStatusEnum    OfferStatusEnum = "Pending"
	MODIFED_OfferStatusEnum    OfferStatusEnum = "Modifed"
	REJECTED_OfferStatusEnum   OfferStatusEnum = "Rejected"
	WAITLISTED_OfferStatusEnum OfferStatusEnum = "Waitlisted"
)

type OperationalStatusEnumBase string

// List of OperationalStatusENUM_Base
const (
	FLIGHT_BOARDING_OperationalStatusEnumBase                 OperationalStatusEnumBase = "FlightBoarding"
	FLIGHT_CANCELLED_OperationalStatusEnumBase                OperationalStatusEnumBase = "FlightCancelled"
	FLIGHT_DEPARTED_OperationalStatusEnumBase                 OperationalStatusEnumBase = "FlightDeparted"
	FLIGHT_PAST_SCHEDULED_DEPARTURE_OperationalStatusEnumBase OperationalStatusEnumBase = "FlightPastScheduledDeparture"
	NOT_AVAILABLE_USE_SEARCH_OperationalStatusEnumBase        OperationalStatusEnumBase = "NotAvailableUseSearch"
)

// OptInStatusEnum : Used to indicate marketing preferences, OptIn, OptOut
type OptInStatusEnum string

// List of OptInStatusEnum
const (
	OPTED_IN_OptInStatusEnum  OptInStatusEnum = "OptedIn"
	OPTED_OUT_OptInStatusEnum OptInStatusEnum = "OptedOut"
	UNKNOWN_OptInStatusEnum   OptInStatusEnum = "Unknown"
)

// OrganizationTypeEnum : The type of organization such as an Agency, Branch, Company, Supplier, Provider
type OrganizationTypeEnum string

// List of OrganizationTypeEnum
const (
	TRAVEL_AGENCY_OrganizationTypeEnum      OrganizationTypeEnum = "TravelAgency"
	AGENCY_BRANCH_OrganizationTypeEnum      OrganizationTypeEnum = "AgencyBranch"
	LOYALTY_PROGRAM_OrganizationTypeEnum    OrganizationTypeEnum = "LoyaltyProgram"
	ID_DOCUMENT_ISSUER_OrganizationTypeEnum OrganizationTypeEnum = "IdDocumentIssuer"
	TRAVEL_SUPPLIER_OrganizationTypeEnum    OrganizationTypeEnum = "TravelSupplier"
	TRAVEL_PROVIDER_OrganizationTypeEnum    OrganizationTypeEnum = "TravelProvider"
	REGULATORY_OrganizationTypeEnum         OrganizationTypeEnum = "Regulatory"
)

// PaymentCardTypeEnum : Credit, Debit, etc.
type PaymentCardTypeEnum string

// List of PaymentCardTypeEnum
const (
	CREDIT_PaymentCardTypeEnum PaymentCardTypeEnum = "Credit"
	DEBIT_PaymentCardTypeEnum  PaymentCardTypeEnum = "Debit"
	GIFT_PaymentCardTypeEnum   PaymentCardTypeEnum = "Gift"
)

// PercentAppliesTo : The increment the percent applies to. Default value is Amount
type PercentAppliesTo string

// List of PercentAppliesTo
const (
	NIGHTS_PercentAppliesTo PercentAppliesTo = "Nights"
	STAY_PercentAppliesTo   PercentAppliesTo = "Stay"
	AMOUNT_PercentAppliesTo PercentAppliesTo = "Amount"
)

type PictureofEnum string

// List of PictureofEnum
const (
	EXTERIOR_PictureofEnum        PictureofEnum = "Exterior"
	LOBBY_PictureofEnum           PictureofEnum = "Lobby"
	POOL_PictureofEnum            PictureofEnum = "Pool"
	RESTAURANT_PictureofEnum      PictureofEnum = "Restaurant"
	HEALTH_CLUB_PictureofEnum     PictureofEnum = "HealthClub"
	GUEST_ROOM_PictureofEnum      PictureofEnum = "GuestRoom"
	SUITE_PictureofEnum           PictureofEnum = "Suite"
	CONFERENCE_ROOM_PictureofEnum PictureofEnum = "ConferenceRoom"
	BALLROOM_PictureofEnum        PictureofEnum = "Ballroom"
	GOLF_PictureofEnum            PictureofEnum = "Golf"
	BEACH_PictureofEnum           PictureofEnum = "Beach"
	SPA_PictureofEnum             PictureofEnum = "Spa"
	BAR_PictureofEnum             PictureofEnum = "Bar"
	RECREATIONAL_PictureofEnum    PictureofEnum = "Recreational"
	ROOM_AMENITY_PictureofEnum    PictureofEnum = "RoomAmenity"
	POPERTY_AMENITY_PictureofEnum PictureofEnum = "PopertyAmenity"
	BUSINESS_CENTRE_PictureofEnum PictureofEnum = "BusinessCentre"
)

// PositionAccuracyEnum : Specifies the level of accuracy for the position
type PositionAccuracyEnum string

// List of PositionAccuracyEnum
const (
	ZIP9_CODE_PositionAccuracyEnum    PositionAccuracyEnum = "Zip9Code"
	ZIP7_CODE_PositionAccuracyEnum    PositionAccuracyEnum = "Zip7Code"
	ZIP5_CODE_PositionAccuracyEnum    PositionAccuracyEnum = "Zip5Code"
	STREET_PositionAccuracyEnum       PositionAccuracyEnum = "Street"
	STATE_PositionAccuracyEnum        PositionAccuracyEnum = "State"
	PROPERTY_PositionAccuracyEnum     PositionAccuracyEnum = "Property"
	INTERSECTION_PositionAccuracyEnum PositionAccuracyEnum = "Intersection"
	EXACT_PositionAccuracyEnum        PositionAccuracyEnum = "Exact"
	COUNTY_PositionAccuracyEnum       PositionAccuracyEnum = "County"
	CITY_PositionAccuracyEnum         PositionAccuracyEnum = "City"
	BLOCK_PositionAccuracyEnum        PositionAccuracyEnum = "Block"
)

// PricingEnum : An enumerated type that defines how a service is priced.
type PricingEnum string

// List of PricingEnum
const (
	STAY_PricingEnum             PricingEnum = "Per stay"
	PERSON_PricingEnum           PricingEnum = "Per person"
	NIGHT_PricingEnum            PricingEnum = "Per night"
	PERSON_PER_NIGHT_PricingEnum PricingEnum = "Per person per night"
	USE_PricingEnum              PricingEnum = "Per use"
)

// RateCategoryEnum : Rate Category
type RateCategoryEnum string

// List of RateCategoryEnum
const (
	ALL_RateCategoryEnum                        RateCategoryEnum = "All"
	ASSOCIATION_RateCategoryEnum                RateCategoryEnum = "Association"
	BUSINESS_RateCategoryEnum                   RateCategoryEnum = "Business"
	BUSINESS_STANDARD_RateCategoryEnum          RateCategoryEnum = "BusinessStandard"
	CLUB_RateCategoryEnum                       RateCategoryEnum = "Club"
	CONVENTION_RateCategoryEnum                 RateCategoryEnum = "Convention"
	CORPORATE_RateCategoryEnum                  RateCategoryEnum = "Corporate"
	CONSORTIUMS_RateCategoryEnum                RateCategoryEnum = "Consortiums"
	DISCOUNT_RateCategoryEnum                   RateCategoryEnum = "Discount"
	CREDENTIAL_RateCategoryEnum                 RateCategoryEnum = "Credential"
	EMPLOYEE_RateCategoryEnum                   RateCategoryEnum = "Employee"
	FAMILY_PLAN_RateCategoryEnum                RateCategoryEnum = "FamilyPlan"
	FULL_INCLUSIVE_RateCategoryEnum             RateCategoryEnum = "FullInclusive"
	GOVERNMENT_RateCategoryEnum                 RateCategoryEnum = "Government"
	INCLUSIVE_RateCategoryEnum                  RateCategoryEnum = "Inclusive"
	INDUSTRYTRAVEL_AGENT_RATE_RateCategoryEnum  RateCategoryEnum = "Industry/TravelAgentRate"
	LEISURE_RateCategoryEnum                    RateCategoryEnum = "Leisure"
	MILITARY_RateCategoryEnum                   RateCategoryEnum = "Military"
	MONTHLY_RateCategoryEnum                    RateCategoryEnum = "Monthly"
	MULTI_DAY_PACKAGE_RateCategoryEnum          RateCategoryEnum = "Multi-DayPackage"
	MULT_LEVELNEGOTIATEDSECURE_RateCategoryEnum RateCategoryEnum = "MultLevel/Negotiated/Secure"
	OTHER_RateCategoryEnum                      RateCategoryEnum = "Other"
	PACKAGE__RateCategoryEnum                   RateCategoryEnum = "Package"
	PRE_PAID_RateCategoryEnum                   RateCategoryEnum = "PrePaid"
	PROMOTIONAL_RateCategoryEnum                RateCategoryEnum = "Promotional"
	RACK_GENERAL_RateCategoryEnum               RateCategoryEnum = "RackGeneral"
	SENIOR_CITIZEN_RateCategoryEnum             RateCategoryEnum = "SeniorCitizen"
	STANDARD_RateCategoryEnum                   RateCategoryEnum = "Standard"
	TOUR_RateCategoryEnum                       RateCategoryEnum = "Tour"
	VIP_RateCategoryEnum                        RateCategoryEnum = "VIP"
	WEEKEND_RateCategoryEnum                    RateCategoryEnum = "Weekend"
	WEEKLY_RateCategoryEnum                     RateCategoryEnum = "Weekly"
)

// RatePaymentEnum : Payment Rate
type RatePaymentEnum string

// List of RatePaymentEnum
const (
	PRE_PAY_RatePaymentEnum  RatePaymentEnum = "PrePay"
	POST_PAY_RatePaymentEnum RatePaymentEnum = "PostPay"
	UNKNOWN_RatePaymentEnum  RatePaymentEnum = "Unknown"
)

// ShareWithEnum : Share with like Supplier,agency etc
type ShareWithEnum string

// List of ShareWithEnum
const (
	SUPPLIER_ShareWithEnum ShareWithEnum = "Supplier"
	AGENCY_ShareWithEnum   ShareWithEnum = "Agency"
	TRAVELER_ShareWithEnum ShareWithEnum = "Traveler"
)

// TextFormatEnum : Describes the format of text such as plain text or html
type TextFormatEnum string

// List of TextFormatEnum
const (
	PLAIN_TEXT_TextFormatEnum TextFormatEnum = "PlainText"
	HTML_TextFormatEnum       TextFormatEnum = "HTML"
)

// UnitOfDistanceEnum : Miles, Kilometers, etc.
type UnitOfDistanceEnum string

// List of UnitOfDistanceEnum
const (
	MILES_UnitOfDistanceEnum      UnitOfDistanceEnum = "Miles"
	KILOMETERS_UnitOfDistanceEnum UnitOfDistanceEnum = "Kilometers"
)

// UnitOfMeasureEnum : The unit of measure in a code format. Refer to OpenTravel Code List Unit of Measure Code (UOM).
type UnitOfMeasureEnum string

// List of UnitOfMeasureEnum
const (
	MILES_UnitOfMeasureEnum             UnitOfMeasureEnum = "Miles"
	KILOMETERS_UnitOfMeasureEnum        UnitOfMeasureEnum = "Kilometers"
	METERS_UnitOfMeasureEnum            UnitOfMeasureEnum = "Meters"
	MILLIMETERS_UnitOfMeasureEnum       UnitOfMeasureEnum = "Millimeters"
	CENTIMETERS_UnitOfMeasureEnum       UnitOfMeasureEnum = "Centimeters"
	YARDS_UnitOfMeasureEnum             UnitOfMeasureEnum = "Yards"
	FEET_UnitOfMeasureEnum              UnitOfMeasureEnum = "Feet"
	INCHES_UnitOfMeasureEnum            UnitOfMeasureEnum = "Inches"
	PIXELS_UnitOfMeasureEnum            UnitOfMeasureEnum = "Pixels"
	BLOCK_UnitOfMeasureEnum             UnitOfMeasureEnum = "Block"
	MEGABYTES_UnitOfMeasureEnum         UnitOfMeasureEnum = "Megabytes"
	GIGABYTES_UnitOfMeasureEnum         UnitOfMeasureEnum = "Gigabytes"
	SQUARE_FEET_UnitOfMeasureEnum       UnitOfMeasureEnum = "Square feet"
	SQUARE_METERS_UnitOfMeasureEnum     UnitOfMeasureEnum = "Square meters"
	POUNDS_UnitOfMeasureEnum            UnitOfMeasureEnum = "Pounds"
	KILOGRAMS_UnitOfMeasureEnum         UnitOfMeasureEnum = "Kilograms"
	SQUARE_INCH_UnitOfMeasureEnum       UnitOfMeasureEnum = "Square inch"
	SQUARE_YARD_UnitOfMeasureEnum       UnitOfMeasureEnum = "Square yard"
	ACRE_UnitOfMeasureEnum              UnitOfMeasureEnum = "Acre"
	SQUARE_MILLIMETER_UnitOfMeasureEnum UnitOfMeasureEnum = "Square millimeter"
	SQUARE_CENTIMETER_UnitOfMeasureEnum UnitOfMeasureEnum = "Square centimeter"
	HECTARE_UnitOfMeasureEnum           UnitOfMeasureEnum = "Hectare"
	OUNCE_UnitOfMeasureEnum             UnitOfMeasureEnum = "Ounce"
	GRAM_UnitOfMeasureEnum              UnitOfMeasureEnum = "Gram"
	GALLONS_UnitOfMeasureEnum           UnitOfMeasureEnum = "Gallons"
	LITERS_UnitOfMeasureEnum            UnitOfMeasureEnum = "Liters"
	KILOWATTS_UnitOfMeasureEnum         UnitOfMeasureEnum = "Kilowatts"
	CUBIC_METERS_UnitOfMeasureEnum      UnitOfMeasureEnum = "Cubic meters"
)

// UnitOfSize : List of units of size i.e Square Feet, Square Meters
type UnitOfSize string

// List of UnitOfSize
const (
	FEET_UnitOfSize   UnitOfSize = "Square Feet"
	METERS_UnitOfSize UnitOfSize = "Square Meters"
)

// YesNoInheritEnum : Used to indicate marketing preferences, Yes, No, Inherit
type YesNoInheritEnum string

// List of YesNoInheritEnum
const (
	YES_YesNoInheritEnum     YesNoInheritEnum = "Yes"
	NO_YesNoInheritEnum      YesNoInheritEnum = "No"
	INHERIT_YesNoInheritEnum YesNoInheritEnum = "Inherit"
)
