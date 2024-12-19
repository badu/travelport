package travelport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Client struct {
	Username     string
	Password     string
	ClientID     string
	ClientSecret string

	baseURL      string
	accessGroup  string
	token        string
	dumpRequest  bool
	dumpResponse bool
	client       *http.Client
}

// OAuth endpoints:
// Pre-production: https://oauth.pp.travelport.com/oauth/oauth20/token
// Production: https://oauth.travelport.com/oauth/oauth20/token

// OAuth payload should be "x-www-form-urlencoded",username, password, client_id and client_secret

// OAuth response contains JSON properties "access_token", "token_type" ("Bearer"), "expires_in" (should refresh after this), "refresh_token" (to refresh), "id_token"
func NewClient() *Client {
	return &Client{}
}

// Hotel API
// Pre-production: https://api.apim-a.adc.pp.travelport.io/hotel/
// Production: https://api.apim-a.adc.prod.travelport.io/hotel/

// Use document override to send remarks such as tour code, commission, or endorsements/restrictions.Document override remarks are returned in the PNR retrieve only when the detailViewInd query parameter is set to true. Document override remarks can be added to an existing PNR but cannot be modified or deleted; see PNR Modify Guide.
func (c *Client) CreateDocumentOverrides(ctx context.Context, req *DocumentOverrides) (*DocumentOverridesResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/Reservation/{ReservationResource_Identifier}/documentoverrides", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *DocumentOverridesResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) DeleteDocumentOverrides(ctx context.Context, reservationResourceIdentifier, id string) error {
	endpointURL := fmt.Sprintf("%s/Reservation/%s/documentoverrides/%s", c.baseURL, reservationResourceIdentifier, id)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpointURL, http.NoBody)
	if err != nil {
		return err
	}

	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		return nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return err
		}
		return fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) UpdateDocumentOverrides(ctx context.Context, req *DocumentOverrides) (*DocumentOverridesResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/Reservation/{ReservationResource_Identifier}/documentoverrides/{id}", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *DocumentOverridesResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) EMDGetByLocator(ctx context.Context) (*EMDListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/emds/getbylocator", c.baseURL)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *EMDListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Display an EMD to retrieve EMD details such as the amount paid and agency ticketing information.
func (c *Client) GetEMD(ctx context.Context, identifier string) (*EMDListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/emds/%s", c.baseURL, identifier)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *EMDListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Void an EMD to cancel it. You can also use EMD void with GDS Exchange Ticketing API to refund an EMD back to the FOP.
func (c *Client) UpdateEMD(ctx context.Context, req *EMDQueryUpdateEMD) (*EMDListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/emds/{Identifier}", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *EMDListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Use the Add Offer reference payload request to add an offer to the reservation workbench as part of the booking workflow. The reference payload request sends identifiers from the Exchage Search response instead of full itinerary details.
func (c *Client) TripChangeBuildFromCatalogOfferings(ctx context.Context, req *OfferQueryBuildFromCatalogOfferings) (*OfferListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/offer/reservationworkbench/{ReservationResource_Identifier}/offers/buildfromcatalogofferings", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *OfferListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:

		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) CreateAgencyCalculatedExchange(ctx context.Context, req *Offer) (*OfferListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/offer/reservationworkbench/{ReservationResource_Identifier}/offers", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *OfferListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:

		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Delete a Form Of Payment
func (c *Client) DeleteFormOfPayment(ctx context.Context, reservationResourceIdentifier, identifier string) error {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/%s/formofpayment/%s", c.baseURL, reservationResourceIdentifier, identifier)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpointURL, http.NoBody)
	if err != nil {
		return err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		return nil
	default:
		return fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Update a Form Of Payment with new information
func (c *Client) UpdateFormOfPayment(ctx context.Context, req *FormOfPaymentID) (*FormOfPaymentResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/{ReservationResource_Identifier}/formofpayment/{Identifier}", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *FormOfPaymentResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) AddFormOfPayments(ctx context.Context, req *FormOfPaymentListRequest) (*FormOfPaymentListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/{ReservationResource_Identifier}/formofpayments/list", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *FormOfPaymentListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// You can send an Add Form of Payment (FOP) request in either a booking or ticketing workbench session. FOPs of cash and credit are supported. FOPs of agent invoice and non-standard credit card are supported for GDS only.
func (c *Client) AddFormOfPayment(ctx context.Context, req *FormOfPaymentID) (*FormOfPaymentResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/{ReservationResource_Identifier}/formofpayment", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *FormOfPaymentResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// The Add Payment step takes place in a ticketing workbench session and sends the payment. It references both the form of payment to be used for the payment and the offer to pay for.
func (c *Client) AddPayment(ctx context.Context, req *Payment) (*PaymentResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/{ReservationResource_Identifier}/payments", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *PaymentResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) UpdatePayment(ctx context.Context, req *Payment) (*PaymentResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/{ReservationResource_Identifier}/payments", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *PaymentResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) AddPayments(ctx context.Context, req *PaymentListRequest) (*PaymentListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/{ReservationResource_Identifier}/payments/list", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *PaymentListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) DeletePayment(ctx context.Context, reservationResourceIdentifier, id string) error {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/%s/payments/%s", c.baseURL, reservationResourceIdentifier, id)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpointURL, http.NoBody)
	if err != nil {
		return err
	}

	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:

		return nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return err
		}
		return fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) CancelPayment(ctx context.Context, reservationResourceIdentifier, id string) (*PaymentResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/%s/payments/%s", c.baseURL, reservationResourceIdentifier, id)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *PaymentResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Cancel an Offer by modifying the Reservation
func (c *Client) CancelReservationOffer(ctx context.Context, reservationIdentifier string) (*ReservationResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/book/reservations/%s/canceloffer", c.baseURL, reservationIdentifier)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Create a reservation on the core or with the vendor/provider.
func (c *Client) CreateReservation(ctx context.Context, req *ReservationDetailWrapper) (*ReservationResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// As an alternative to the booking workflow that takes place in a workbench session, you can send all booking details and commit a single payload to create a booking. The single payload book request does not support any of the optional steps in the booking workflow, such as adding seats or ancillaries.
func (c *Client) BuildReservation(ctx context.Context, req *ReservationQueryBuildWrapper) (*ReservationResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations/build", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) Divide(ctx context.Context, req *ReservationQueryDivide) (*ReservationListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations/divide", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) FindReservation(ctx context.Context, req *ReservationQuerySearchCriteriaReservation) (*ReservationListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations/find", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:

		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// To be deprecated and replaced by Get by Identifier using identifier Type "Locator"
func (c *Client) GetReservationByLocator(ctx context.Context) (*ReservationResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations/getbylocator", c.baseURL)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Retrieve details about a held booking, or PNR. While a PNR refers to a held booking that has not been ticketed, the PNR code persists after ticketing to provide the booking records. Once a PNR has been ticketed, you can still use PNR Retrieve to return both booking and ticketing details. A Ticket Display request can also be used to retrieve any ticketed itinerary.
func (c *Client) RetrieveReservation(ctx context.Context, identifier string) (*ReservationResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations/%s", c.baseURL, identifier)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// After all required and any optional steps in a booking workbench session, send a POST request with the workbench identifier to commit the workbench. The resulting actions depend on whether payment is present in the workbench. If no Add Payment request has been sent, committing the workbench books the itinerary and generates a PNR. If an Add Payment request has not been sent, committing the workbench tickets the itinerary and generates ticket number/s.
func (c *Client) CommitReservation(ctx context.Context, req *ReservationQueryCommitReservation) (*ReservationResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations/{Identifier}", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

func (c *Client) UpdateReservation(ctx context.Context, req *ReservationDetailWrapper) (*ReservationResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations/{Identifier}", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Process all unprocessesed offers and create ticket receipts.
func (c *Client) BuildReceiptsfromLocator(ctx context.Context, reservationResourceIdentifier string) (*ReceiptListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations/%s/receipts/buildfromlocator", c.baseURL, reservationResourceIdentifier)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReceiptListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Process all un-processesed payments and create a list of payment receipts.
func (c *Client) BuildReceiptsFromPayment(ctx context.Context, req *ReceiptQueryBuildFromPayment) (*ReceiptListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations/{ReservationResource_Identifier}/receipts/buildfrompayment", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReceiptListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Get a list of ticket receipts for a reservation.
func (c *Client) GetReceipts(ctx context.Context, reservationResourceIdentifier string) (*ReceiptListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations/%s/receipts", c.baseURL, reservationResourceIdentifier)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReceiptListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Create a set of offer cancelation receipts for every offer in the reservation.
func (c *Client) CancelReservation(ctx context.Context, reservationResourceIdentifier string) (*ReceiptListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservations/%s/receipts", c.baseURL, reservationResourceIdentifier)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReceiptListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Use this request to initiate a workbench for a new reservation. This prerequisite step for booking creates the workbench session in which all booking details are added together to create a PNR at commit.
func (c *Client) CreateReservationWorkbench(ctx context.Context, req *ReservationID) (*ReservationResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservationworkbench", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Initiate a post-commit workbench to create a session for ticketing or updating an existing reservation. This is a prerequisite step for any transaction that modifies, updates, or tickets any PNR.
func (c *Client) CreateReservationWorkbenchFromIdentifier(ctx context.Context, identifier string) (*ReservationResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/buildfromidentifier/%s", c.baseURL, identifier)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// Initiate a post-commit workbench to create a session for ticketing or updating an existing reservation. This is a prerequisite step for any transaction that modifies, updates, or tickets any PNR.
func (c *Client) CreateReservationWorkbenchFromLocator(ctx context.Context) (*ReservationResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/buildfromlocator", c.baseURL)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// At any point in the booking session, you can retrieve the workbench. The response returns all details added to the workbench at that point.
func (c *Client) RetrieveReservationWorkbench(ctx context.Context, identifier string) (*ReservationResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/%s", c.baseURL, identifier)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *ReservationResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// At any point in a booking or ticketing workflow, if necessary, you can discard the workbench and any information in it.
func (c *Client) IgnoreReservationWorkbench(ctx context.Context, identifier string) error {
	endpointURL := fmt.Sprintf("%s/reservationworkbench/%s", c.baseURL, identifier)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpointURL, http.NoBody)
	if err != nil {
		return err
	}

	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)
	httpReq.Header.Set("TravelportPlusSessionID", c.token)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		return nil
	default:
		return fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// The Eligibility API is the first step in the GDS exchange workflow. Eligibility returns information about whether a ticket may have value in an exchange or refund scenario, and the range of potential exchange and refund fees. This information relates only to fees and not any fare or tax difference for the itinerary.
func (c *Client) GetEligibility(ctx context.Context) (*TicketChangeEligibilityListResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/ticketchangeeligibilities", c.baseURL)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("TravelportPlusSessionID", c.token)
	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *TicketChangeEligibilityListResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// The Product Specific Search function of the Exchange Search API is the second step in the GDS exchange workflow, after Eligibility. It supports searching for a specified itinerary and in addition can provide alternate options. The response details any differences in base fare, taxes, fees, and total price between the currently ticketed itinerary and the possible new itinerary.
func (c *Client) ProductSpecificSearch(ctx context.Context, req *CatalogOfferingsQueryProductSpecificSearch) (*CatalogOfferingsAirChangeResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/catalogofferingsairchange/productspecificsearch", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("TravelportPlusSessionID", c.token)
	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *CatalogOfferingsAirChangeResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}

// The Exchange Search API is the second step in the GDS exchange workflow, after Eligibility. It supports searching for an alternate itinerary for a possible exchange on a currently ticketed GDS itinerary. The response details any differences in base fare, taxes, fees, and total price between the currently ticketed itinerary and the possible new itinerary.
func (c *Client) CreateExchangeSearch(ctx context.Context, req *CatalogOfferingsQueryAirChange) (*CatalogOfferingsAirChangeResponseWrapper, error) {
	endpointURL := fmt.Sprintf("%s/catalogofferingsairchange", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("TravelportPlusSessionID", c.token)
	httpReq.Header.Set("XAUTH_TRAVELPORT_ACCESSGROUP", c.accessGroup)

	if c.dumpRequest {
		dump, err := httputil.DumpRequest(httpReq, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if c.dumpResponse {
		dump, err := httputil.DumpResponse(httpResp, true)
		if err == nil {
			fmt.Println(string(dump))
		}
	}
	switch httpResp.StatusCode {
	case 200, 201:
		var result *CatalogOfferingsAirChangeResponseWrapper
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		var result *ErrorResponse
		err = json.NewDecoder(httpResp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("bad status code in response %d", httpResp.StatusCode)
	}
}
