package travelport

import "net/http"

type BaseResponse struct {
	// Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
	TransactionId string `json:"transactionId,omitempty"`
	// Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TraceId                string                   `json:"traceId,omitempty"`
	Result                 *Result                  `json:"Result,omitempty"`
	Identifier             *Identifier              `json:"Identifier,omitempty"`
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"`
}

type ErrorResponse struct {
	// Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
	TransactionId string `json:"transactionId,omitempty"`
	// Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TraceId                string                   `json:"traceId,omitempty"`
	Result                 *Result                  `json:"Result,omitempty"`
	Identifier             *Identifier              `json:"Identifier,omitempty"`
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"`
}

type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the swagger operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

type EmdListResponse struct {
	EMDID []Emd `json:"EMDID,omitempty"`
}

type EmdListResponseWrapper struct {
	EMDListResponse *EmdListResponse `json:"EMDListResponse,omitempty"`
}

type CatalogOfferingsHospitalityResponse struct {
	// Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
	TransactionId string `json:"transactionId,omitempty"`
	// Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TraceId                string                   `json:"traceId,omitempty"`
	Result                 *Result                  `json:"Result,omitempty"`
	Identifier             *Identifier              `json:"Identifier,omitempty"`
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"`
	Type_                  string                   `json:"@type,omitempty"`
	CatalogOfferings       *CatalogOfferings        `json:"CatalogOfferings,omitempty"`
}

type CatalogOfferingsHospitalityResponseWrapper struct {
	CatalogOfferingsHospitalityResponse *CatalogOfferingsHospitalityResponse `json:"CatalogOfferingsHospitalityResponse,omitempty"`
}

type OfferHospitalityResponse struct {
	// Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
	TransactionId string `json:"transactionId,omitempty"`
	// Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TraceId                string                   `json:"traceId,omitempty"`
	Result                 *Result                  `json:"Result,omitempty"`
	Identifier             *Identifier              `json:"Identifier,omitempty"`
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"`
	Offer                  *Offer                   `json:"Offer,omitempty"`
}

type OfferHospitalityResponseWrapper struct {
	OfferHospitalityResponse *OfferHospitalityResponse `json:"OfferHospitalityResponse,omitempty"`
}

type PropertiesResponse struct {
	// Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
	TransactionId string `json:"transactionId,omitempty"`
	// Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TraceId                string                   `json:"traceId,omitempty"`
	Result                 *Result                  `json:"Result,omitempty"`
	Identifier             *Identifier              `json:"Identifier,omitempty"`
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"`
	Properties             *Properties              `json:"Properties,omitempty"`
}

type PropertiesResponseWrapper struct {
	PropertiesResponse *PropertiesResponse `json:"PropertiesResponse,omitempty"`
}

type ReservationResponse struct {
	// Unique transaction, correlation or tracking id for a single request and reply i.e. for a single transaction. Should be a 128 bit GUID format. Also know as E2ETrackingId.
	TransactionId string `json:"transactionId,omitempty"`
	// Optional ID for internal child transactions created for processing a single request (single transaction). Should be a 128 bit GUID format. Also known as ChildTrackingId.
	TraceId                string                   `json:"traceId,omitempty"`
	Result                 *Result                  `json:"Result,omitempty"`
	Identifier             *Identifier              `json:"Identifier,omitempty"`
	NextSteps              *NextSteps               `json:"NextSteps,omitempty"`
	ReferenceList          []ReferenceList          `json:"ReferenceList,omitempty"`
	CurrencyRateConversion []CurrencyRateConversion `json:"CurrencyRateConversion,omitempty"`
	Reservation            *ReservationId           `json:"Reservation,omitempty"`
}

type ReservationResponseWrapper struct {
	ReservationResponse *ReservationResponse `json:"ReservationResponse,omitempty"`
}
