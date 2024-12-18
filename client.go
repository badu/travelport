package travelport

type Client struct {
	Username     string
	Password     string
	BaseURL      string
	ClientID     string
	ClientSecret string
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
