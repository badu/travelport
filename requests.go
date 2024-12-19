package travelport

// DocumentOverrides is an object.
type DocumentOverrides struct {
	Type                      string                     `json:"@type"`                               // Type:
	ChangeFeeCollectionMethod *ChangeFeeCollectionMethod `json:"ChangeFeeCollectionMethod,omitempty"` // ChangeFeeCollectionMethod:
	Commissions               []Commissions              `json:"Commissions,omitempty"`               // Commissions:
	DestinationPurpose        []DestinationPurpose       `json:"DestinationPurpose,omitempty"`        // DestinationPurpose:
	DocumentOverridesRef      string                     `json:"DocumentOverridesRef,omitempty"`      // DocumentOverridesRef:
	Identifier                *Identifier                `json:"Identifier,omitempty"`                // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	NetRemitInfo              *NetRemitInfo              `json:"NetRemitInfo,omitempty"`              // NetRemitInfo:
	OfferIdentifier           *OfferIdentifier           `json:"OfferIdentifier,omitempty"`           // OfferIdentifier:
	ProductIdentifier         *ProductIdentifier         `json:"ProductIdentifier,omitempty"`         // ProductIdentifier:
	Restrictions              []Restrictions             `json:"Restrictions,omitempty"`              // Restrictions:
	TourCodes                 []TourCodes                `json:"TourCodes,omitempty"`                 // TourCodes:
	Id                        string                     `json:"id,omitempty"`                        // Id: The reporting number.
}

// EMDQueryUpdateEMD is an object.
type EMDQueryUpdateEMD struct {
	Type        string `json:"@type"`                 // Type:
	AgencyCode  string `json:"agencyCode,omitempty"`  // AgencyCode: Assigned Type: c-1100:AgencyCodeIATA
	DateOfIssue string `json:"dateOfIssue,omitempty"` // DateOfIssue: The date the EMD was issued
	Status      string `json:"status"`                // Status: Assigned Type: c-1100:StringTiny
}

// OfferQueryBuildFromCatalogOfferings is an object.
type OfferQueryBuildFromCatalogOfferings struct {
	Type                             string                            `json:"@type"`                                      // Type:
	BuildFromCatalogOfferingsRequest *BuildFromCatalogOfferingsRequest `json:"BuildFromCatalogOfferingsRequest,omitempty"` // BuildFromCatalogOfferingsRequest:
	CabinPreference                  *CabinPreference                  `json:"CabinPreference,omitempty"`                  // CabinPreference:
	FareRuleCategory                 []FareRuleCategoryEnum            `json:"FareRuleCategory,omitempty"`                 // FareRuleCategory:
	FareRuleType                     FareRuleEnum                      `json:"FareRuleType,omitempty"`                     // FareRuleType:
	PaymentCriteria                  *PaymentCriteria                  `json:"PaymentCriteria,omitempty"`                  // PaymentCriteria:
	LowFareFinderInd                 bool                              `json:"lowFareFinderInd,omitempty"`                 // LowFareFinderInd: If present and true, this is a low fare finder request
	ReCheckInventoryInd              bool                              `json:"reCheckInventoryInd,omitempty"`              // ReCheckInventoryInd: If present and true, then the host system will perform a sell inventory check.
	ReturnBrandedFaresInd            bool                              `json:"returnBrandedFaresInd,omitempty"`            // ReturnBrandedFaresInd: If present and true, branded fares are returned
	ValidateInventoryInd             bool                              `json:"validateInventoryInd,omitempty"`             // ValidateInventoryInd: If true, the flight inventory will be checked during the price step
}

// Offer is an object.
type Offer struct {
	Type                   string                   `json:"@type"`                     // Type:
	Identifier             *Identifier              `json:"Identifier,omitempty"`      // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Price                  Price                    `json:"Price"`                     // Price:
	Product                []ProductID              `json:"Product"`                   // Product:
	TermsAndConditionsFull []TermsAndConditionsFull `json:"TermsAndConditionsFull"`    // TermsAndConditionsFull:
	Id                     string                   `json:"id,omitempty"`              // Id: Local indentifier within a given message for this object.
	OfferRef               string                   `json:"offerRef,omitempty"`        // OfferRef: Used to reference another instance of this object in the same message
	ParentOfferRef         string                   `json:"parentOfferRef,omitempty"`  // ParentOfferRef: A reference to the Offer this offer is sold in conjunction with
	PassiveOfferInd        bool                     `json:"passiveOfferInd,omitempty"` // PassiveOfferInd: If true, the Offer is passive for booking purposes.
}

// FormOfPaymentID is an object.
type FormOfPaymentID struct {
	Type             string      `json:"@type"`                      // Type:
	FormOfPaymentRef string      `json:"FormOfPaymentRef,omitempty"` // FormOfPaymentRef:
	Identifier       *Identifier `json:"Identifier,omitempty"`       // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id               string      `json:"id,omitempty"`               // Id:
}

// FormOfPaymentListRequest is an object.
type FormOfPaymentListRequest struct {
	FormOfPayment []FormOfPaymentID `json:"FormOfPayment"` // FormOfPayment:
}

// Payment is an object.
type Payment struct {
	Type                       string                       `json:"@type,omitempty"`                      // Type:
	AgencyServiceFeeIdentifier []AgencyServiceFeeIdentifier `json:"AgencyServiceFeeIdentifier,omitempty"` // AgencyServiceFeeIdentifier:
	Amount                     CurrencyAmount               `json:"Amount"`                               // Amount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	BaseAmount                 *CurrencyAmount              `json:"BaseAmount,omitempty"`                 // BaseAmount: A monetary amount, up to 4 decimal places. Decimal place needs to be included.
	ExtendedPayment            *ExtendedPayment             `json:"ExtendedPayment,omitempty"`            // ExtendedPayment: Note this field is deprecated in Payment schema and should be passed in FormOfPaymentPaymentCardExtendPayment schema
	Fees                       *Fees                        `json:"Fees,omitempty"`                       // Fees:
	FormOfPaymentIdentifier    *FormOfPaymentIdentifier     `json:"FormOfPaymentIdentifier,omitempty"`    // FormOfPaymentIdentifier:
	Identifier                 *Identifier                  `json:"Identifier,omitempty"`                 // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	OfferIdentifier            []OfferIdentifier            `json:"OfferIdentifier,omitempty"`            // OfferIdentifier:
	PaymentRef                 string                       `json:"PaymentRef,omitempty"`                 // PaymentRef:
	Taxes                      *Taxes                       `json:"Taxes,omitempty"`                      // Taxes:
	TravelerIdentifierRef      []TravelerIdentifierRef      `json:"TravelerIdentifierRef,omitempty"`      // TravelerIdentifierRef:
	DepositInd                 bool                         `json:"depositInd,omitempty"`                 // DepositInd: If true, the payment is a deposit on the referenced Offer
	GuaranteeInd               bool                         `json:"guaranteeInd,omitempty"`               // GuaranteeInd: If true, the payment is a guarantee for the referenced Offer
	Id                         string                       `json:"id,omitempty"`                         // Id:
}

// PaymentListRequest is an object.
type PaymentListRequest struct {
	Payment []Payment `json:"Payment"` // Payment:
}

// ReservationDetailWrapper is an object.
type ReservationDetailWrapper struct {
	ReservationDetail *ReservationDetail `json:"ReservationDetail,omitempty"` // ReservationDetail:
}

// ReservationQueryBuildWrapper is an object.
type ReservationQueryBuildWrapper struct {
	ReservationQueryBuild *ReservationQueryBuild `json:"ReservationQueryBuild,omitempty"` // ReservationQueryBuild:
}

// ReservationQueryDivide is an object.
type ReservationQueryDivide struct {
	ReservationIdentifier ReservationIdentifier `json:"ReservationIdentifier"` // ReservationIdentifier:
	TravelerIdentifier    []TravelerIdentifier  `json:"TravelerIdentifier"`    // TravelerIdentifier:
}

// ReservationQuerySearchCriteriaReservation is an object.
type ReservationQuerySearchCriteriaReservation struct {
	Type               string      `json:"@type"`                        // Type:
	Arrival            string      `json:"Arrival,omitempty"`            // Arrival: The city or airport code a flight is arriving at in the Reservation
	Departure          string      `json:"Departure,omitempty"`          // Departure: The city or airport code a flight is departing from in the Reservation
	DepartureDate      string      `json:"DepartureDate,omitempty"`      // DepartureDate: Local date of flight departure
	DepartureDateRange *DateRange  `json:"DepartureDateRange,omitempty"` // DepartureDateRange: Specifies the begin and end date of an event
	PersonName         *PersonName `json:"PersonName,omitempty"`         // PersonName:
	DetailViewInd      bool        `json:"detailViewInd,omitempty"`      // DetailViewInd: If true, ReservationDetail will be returned
}

// ReservationQueryCommitReservation is an object.
type ReservationQueryCommitReservation struct {
	Notification                    []Notification            `json:"Notification,omitempty"`                    // Notification:
	ErrorWhenOfferPriceCancelledInd bool                      `json:"errorWhenOfferPriceCancelledInd,omitempty"` // ErrorWhenOfferPriceCancelledInd: If true, and the OfferPrice is invalidated, error will be returned and Reservation commit will not be processed
	ScheduleChangeAcceptedInd       bool                      `json:"scheduleChangeAcceptedInd,omitempty"`       // ScheduleChangeAcceptedInd: if true, the schedule change is accepted by the consumer
	ScheduleChangeReprice           ScheduleChangeRepriceEnum `json:"scheduleChangeReprice,omitempty"`           // ScheduleChangeReprice:
}

// ReceiptQueryBuildFromPayment is an object.
type ReceiptQueryBuildFromPayment struct {
	Payment *PaymentID `json:"Payment,omitempty"` // Payment:
}

// ReservationID is an object.
type ReservationID struct {
	Type       string      `json:"@type"`                // Type:
	Identifier *Identifier `json:"Identifier,omitempty"` // Identifier: Identifier provides the ability to create a globally unique identifier. For the identifier to be globally unique it must have a system provided identifier and the system must be identified using a global naming authority. System identification uses the domain naming system (DNS) to assure they are globally unique and should be an URL. The system provided ID will typically be a primary or surrogate key in a database.
	Id         string      `json:"id,omitempty"`         // Id: Internal ID
}

// CatalogOfferingsQueryProductSpecificSearch is an object.
type CatalogOfferingsQueryProductSpecificSearch struct {
	Type                              string                            `json:"@type"`                             // Type:
	BuildFromProductsRequestAirChange BuildFromProductsRequestAirChange `json:"BuildFromProductsRequestAirChange"` // BuildFromProductsRequestAirChange:
	CabinPreference                   []CabinPreference                 `json:"CabinPreference,omitempty"`         // CabinPreference:
	PassengerCriteria                 []PassengerCriteria               `json:"PassengerCriteria"`                 // PassengerCriteria:
	PaymentCriteria                   PaymentCriteria                   `json:"PaymentCriteria"`                   // PaymentCriteria:
	DetailView                        bool                              `json:"detailView,omitempty"`              // DetailView: If true, detail view will be returned in response
	LowFareFinderInd                  bool                              `json:"lowFareFinderInd,omitempty"`        // LowFareFinderInd: If true, the lowest fare will be returned regardless of the booking class sent
	ReturnBrandedfaresInd             bool                              `json:"returnBrandedfaresInd,omitempty"`   // ReturnBrandedfaresInd: If true, branded fares will be returned in the response
}

// CatalogOfferingsQueryAirChange is an object.
type CatalogOfferingsQueryAirChange struct {
	Type                             string                                      `json:"@type"`                            // Type:
	CatalogOfferingsAirChangeRequest CatalogOfferingsAirChangeRequestReservation `json:"CatalogOfferingsAirChangeRequest"` // CatalogOfferingsAirChangeRequest:
}
