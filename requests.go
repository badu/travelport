package travelport

type BuildFromCatalogOfferingsRequest struct {
	Type_                       string                        `json:"@type"`
	CatalogOfferingsIdentifier  *CatalogOfferingsIdentifier   `json:"CatalogOfferingsIdentifier"`
	CatalogOfferingIdentifier   *CatalogOfferingIdentifier    `json:"CatalogOfferingIdentifier"`
	ProductIdentifier           []ProductIdentifier           `json:"ProductIdentifier"`
	AncillaryOfferingIdentifier []AncillaryOfferingIdentifier `json:"AncillaryOfferingIdentifier,omitempty"`
}

type CatalogOfferingsQueryRequest struct {
	Type_                   string                               `json:"@type"`
	CatalogOfferingsRequest []CatalogOfferingsRequestHospitality `json:"CatalogOfferingsRequest"`
}

type CatalogOfferingsQueryRequestHospitalityWrapper struct {
	CatalogOfferingsQueryRequest *CatalogOfferingsQueryRequest `json:"CatalogOfferingsQueryRequest,omitempty"`
}

type CatalogOfferingsRequest struct {
	Type_                         string                         `json:"@type,omitempty"`
	SearchControlConsoleChannelID *SearchControlConsoleChannelId `json:"SearchControlConsoleChannelID,omitempty"`
}

type CatalogOfferingsRequestHospitality struct {
	// You can use requested currency to request conversion rate information. The response will return the currencyRateConversion object which will contain conversion rate of the requested currency.
	RequestedCurrency string `json:"requestedCurrency,omitempty"`
	// Maximum time (in milliseconds) to wait for provider responses before returning a response to the consumer of this service
	MaxResponseWaitTime  int32                 `json:"maxResponseWaitTime,omitempty"`
	StayDates            *DateOrDateWindows    `json:"StayDates"`
	HotelSearchCriterion *HotelSearchCriterion `json:"HotelSearchCriterion,omitempty"`
	MinimumAmount        *CurrencyAmount       `json:"MinimumAmount,omitempty"`
	MaximumAmount        *CurrencyAmount       `json:"MaximumAmount,omitempty"`
	// Used to specify that a verbose response is to be returned.  Verbose responses repeat the Property information in each Product and do not return the reference list.
	VerboseResponseInd            bool                           `json:"verboseResponseInd,omitempty"`
	Type_                         string                         `json:"@type,omitempty"`
	SearchControlConsoleChannelID *SearchControlConsoleChannelId `json:"SearchControlConsoleChannelID,omitempty"`
}

type OfferQueryHospitalityRequest struct {
	Type_        string `json:"@type"`
	CheckinDate  string `json:"checkinDate"`
	CheckoutDate string `json:"checkoutDate"`
	// The number of guests
	NumberOfGuests int32 `json:"numberOfGuests"`
	// The booking code
	BookingCode string `json:"bookingCode,omitempty"`
	// stored currency
	StoredCurrency string `json:"storedCurrency,omitempty"`
	// stored amount
	StoredAmount float32 `json:"storedAmount,omitempty"`
	// You can use requested currency to request conversion rate information. The response will return the currencyRateConversion object which will contain conversion rate of the requested currency.
	RequestedCurrency  string               `json:"requestedCurrency,omitempty"`
	RoomStayCandidates *RoomStayCandidates  `json:"RoomStayCandidates,omitempty"`
	PropertyKey        *PropertyKey         `json:"PropertyKey"`
	HotelAggregator    *HotelAggregatorEnum `json:"HotelAggregator,omitempty"`
	RateCandidate      *RateCandidate       `json:"RateCandidate,omitempty"`
}

type OfferQueryHospitalityRequestWrapper struct {
	OfferQueryHospitalityRequest *OfferQueryHospitalityRequest `json:"OfferQueryHospitalityRequest,omitempty"`
}

type PropertyRequest struct {
	Type_ string `json:"@type,omitempty"`
	// More rates token
	MoreRatesToken string       `json:"moreRatesToken,omitempty"`
	PropertyKey    *PropertyKey `json:"PropertyKey"`
}
