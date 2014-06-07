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

type Address struct {
	Lat, Lng float64
	Address  string
	Response *Response
}

type Response struct {
	Status  string
	Results []Result
}

type Result struct {
	Types              []string
	Formatted_address  string
	Address_components []AddressComponent
	Geometry           GeometryData
}

type AddressComponent struct {
	Long_name  string
	Short_name string
	Types      []string
}

type GeometryData struct {
	Location      LatLng
	Location_type string
	Viewport      struct {
		Southwest, Northeast LatLng
	}
	Bounds struct {
		Southwest, Northeast LatLng
	}
}

type LatLng struct {
	Lat, Lng float64
}

func (a *Address) String() string {
	return fmt.Sprintf("%s (lat: %3.7f, lng: %3.7f)", a.Address, a.Lat, a.Lng)
}

func Geocode(q string) (*Address, error) {
	return fetch("https://maps.googleapis.com/maps/api/geocode/json?sensor=false&address="+url.QueryEscape(strings.TrimSpace(q)))
}
func ReverseGeocode(ll string) (*Address, error) {
	return fetch("https://maps.googleapis.com/maps/api/geocode/json?sensor=false&latlng="+url.QueryEscape(strings.TrimSpace(ll)))
}


func fetch(url string) (add *Address, error error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, RemoteServerError
	}

	defer resp.Body.Close()

	var g Response
	err = json.NewDecoder(resp.Body).Decode(&g)

	if err != nil {
		return nil, err
	}

	if g.Status != StatusOk {
		return nil, errors.New(fmt.Sprintf("Geocoder service error!  (%s)", g.Status))
	}

	// fmt.Printf("%#v", g)
	// return nil, nil

	return &Address{
		g.Results[0].Geometry.Location.Lat,
		g.Results[0].Geometry.Location.Lng,
		g.Results[0].Formatted_address,
		&g}, nil

}
