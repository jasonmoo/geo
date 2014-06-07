package geo

import (
	"fmt"
	"net/http"
	"strings"

	"encoding/json"
	"errors"
	"net/url"
)

const (
	StatusOk             = "OK"
	StatusZeroResults    = "ZERO_RESULTS"
	StatusOverQueryLimit = "OVER_QUERY_LIMIT"
	StatusRequestDenied  = "REQUEST_DENIED"
	StatusInvalidRequest = "INVALID_REQUEST"
)

var (
	RemoteServerError = errors.New("Unable to contact the goog.")
	BodyReadError     = errors.New("Unable to read the response body.")
)

type (
	Address struct {
		Lat      float64   `json:"lat"`
		Lng      float64   `json:"lng"`
		Address  string    `json:"address"`
		Response *Response `json:"response"`
	}

	Response struct {
		Status  string   `json:"status"`
		Results []Result `json:"results"`
	}

	Result struct {
		Types             []string           `json:"types"`
		FormattedAddress  string             `json:"formatted_address"`
		AddressComponents []AddressComponent `json:"address_components"`
		Geometry          GeometryData       `json:"geometry"`
	}

	AddressComponent struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	}

	GeometryData struct {
		Location     LatLng `json:"location"`
		LocationType string `json:"location_type"`
		Viewport     struct {
			Southwest LatLng `json:"southwest"`
			Northeast LatLng `json:"northeast"`
		} `json:"viewport"`
		Bounds struct {
			Southwest LatLng `json:"southwest"`
			Northeast LatLng `json:"northeast"`
		} `json:"bounds"`
	}

	LatLng struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
)

func (a *Address) String() string {
	return fmt.Sprintf("%s (lat: %3.7f, lng: %3.7f)", a.Address, a.Lat, a.Lng)
}

func Geocode(q string) (*Address, error) {
	return fetch("https://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=" + url.QueryEscape(strings.TrimSpace(q)))
}
func ReverseGeocode(ll string) (*Address, error) {
	return fetch("https://maps.googleapis.com/maps/api/geocode/json?sensor=false&latlng=" + url.QueryEscape(strings.TrimSpace(ll)))
}

func fetch(url string) (*Address, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, RemoteServerError
	}

	defer resp.Body.Close()

	var g = new(Response)
	err = json.NewDecoder(resp.Body).Decode(g)

	if err != nil {
		return nil, err
	}

	if g.Status != StatusOk {
		return nil, fmt.Errorf("Geocoder service error!  (%s)", g.Status)
	}

	return &Address{
		Lat:      g.Results[0].Geometry.Location.Lat,
		Lng:      g.Results[0].Geometry.Location.Lng,
		Address:  g.Results[0].FormattedAddress,
		Response: g,
	}, nil

}
