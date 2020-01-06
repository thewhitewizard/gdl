package gdl

import (
	"net/http"
	"net/url"
	"errors"
	"encoding/json"
)

var (
		// ErrUnsetAPIKey comment
	ErrUnsetAPIKey = errors.New("unset API key")

	// ErrJCDecauxInternal comment
	ErrGDLInternal = errors.New("internal gdl error (status code 500)")

	// ErrUnsupportedStatusCode comment
	ErrUnsupportedStatusCode = errors.New("unsupported status code")

	// ErrBadRequest comment
	ErrBadRequest = errors.New("bad request")

	// ErrNotFound comment
	ErrNotFound = errors.New("not found")

	// ErrUnauthorized comment
	ErrUnauthorized = errors.New("unauthorized, verify api token")
)

const (
	// BaseURL
	BaseURL = "https://download.data.grandlyon.com"
)

type GDLClient struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL
}

func New(httpClient *http.Client) (*GDLClient, error) {
	
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, err := url.Parse(BaseURL)
	if err != nil {
		return nil, err
	}

	c := &GDLClient{
		client:    httpClient,
		BaseURL:   baseURL,
	}
	return c, err
}



func (c *GDLClient) GetParkings() ([]Parking, error) {
	req, err := c.NewRequest("GET", "files/rdata/lpa_mobilite.donnees/parking_temps_reel.json", "")
	if err != nil {
		return nil, err
	}

	body, err := c.Perform(req)
	if err != nil {
		return nil, err
	}
	parkings := new([]Parking)
	err = json.Unmarshal(body, &parkings)

	return *parkings, err
}


 