package geocodefarm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	forwardEndpoint = "https://api.geocode.farm/forward/"
	reverseEndpoint = "https://api.geocode.farm/reverse/"
	userAgent       = "GeocodeFarmGoSDK/4.0"
)

// Client handles requests to the Geocode.Farm API.
type Client struct {
	APIKey string
	Client *http.Client
}

// NewClient creates a new Geocode.Farm client.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		Client: &http.Client{Timeout: 10 * time.Second},
	}
}

// Response represents the top-level response.
type Response struct {
	Success     bool                   `json:"success"`
	StatusCode  int                    `json:"status_code"`
	Lat         *string                `json:"lat,omitempty"`
	Lon         *string                `json:"lon,omitempty"`
	Accuracy    *string                `json:"accuracy,omitempty"`
	FullAddress *string                `json:"full_address,omitempty"`
	Result      map[string]interface{} `json:"result,omitempty"`
	Error       string                 `json:"error,omitempty"`
}

// Forward performs forward geocoding on an address.
func (c *Client) Forward(address string) (*Response, error) {
	params := url.Values{}
	params.Set("key", c.APIKey)
	params.Set("addr", address)

	fullURL := forwardEndpoint + "?" + params.Encode()
	return c.makeRequest(fullURL, "forward")
}

// Reverse performs reverse geocoding on coordinates.
func (c *Client) Reverse(lat, lon float64) (*Response, error) {
	params := url.Values{}
	params.Set("key", c.APIKey)
	params.Set("lat", fmt.Sprintf("%f", lat))
	params.Set("lon", fmt.Sprintf("%f", lon))

	fullURL := reverseEndpoint + "?" + params.Encode()
	return c.makeRequest(fullURL, "reverse")
}

func (c *Client) makeRequest(url string, mode string) (*Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.Client.Do(req)
	if err != nil {
		return &Response{Success: false, StatusCode: 0, Error: "Request failed"}, nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var parsed map[string]interface{}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return &Response{Success: false, StatusCode: resp.StatusCode, Error: "Invalid JSON"}, nil
	}

	status := safeStr(parsed, "STATUS.status")
	if status != "SUCCESS" {
		return &Response{Success: false, StatusCode: resp.StatusCode, Error: "API failure: " + status}, nil
	}

	var resData map[string]interface{}
	switch mode {
	case "reverse":
		if results, ok := parsed["RESULTS"].(map[string]interface{}); ok {
			if resultList, ok := results["result"].([]interface{}); ok && len(resultList) > 0 {
				resData = resultList[0].(map[string]interface{})
			}
		}
	case "forward":
		if results, ok := parsed["RESULTS"].(map[string]interface{}); ok {
			if result, ok := results["result"].(map[string]interface{}); ok {
				resData = result
			}
		}
	}

	return &Response{
		Success:     true,
		StatusCode:  resp.StatusCode,
		Lat:         strPtr(safeStr(resData, "latitude", "coordinates.lat")),
		Lon:         strPtr(safeStr(resData, "longitude", "coordinates.lon")),
		Accuracy:    strPtr(safeStr(resData, "accuracy")),
		FullAddress: strPtr(safeStr(resData, "formatted_address", "address.full_address")),
		Result:      resData,
	}, nil
}

func safeStr(m map[string]interface{}, keys ...string) string {
	for _, k := range keys {
		parts := []string{k}
		if sub, ok := m[k].(map[string]interface{}); ok && len(keys) > 1 {
			m = sub
			continue
		}
		if val, ok := m[k]; ok {
			if s, ok := val.(string); ok {
				return s
			}
		}
	}
	return ""
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
