package geo

import (
	"fmt"
	"testing"
)

var (
	Addresses = []struct {
		Address  string
		Lat, Lng float64
	}{{
		Address: "323 South Albany Street, Ithaca, NY 14850, USA",
		Lat:     42.435901,
		Lng:     -76.501238,
	}}
)

func TestGeocode(t *testing.T) {

	for _, test_address := range Addresses {
		// t.Logf("checking %s\n", test_address)
		addy, err := Geocode(test_address.Address)
		if err != nil {
			t.Error(err)
		}
		if addy.Address != test_address.Address {
			t.Errorf("Expected: %s, Got: %s", test_address.Address, addy.Address)
		}
		if addy.Lat != test_address.Lat || addy.Lng != test_address.Lng {
			t.Errorf("Expected: %f:%f, Got: %f:%f", test_address.Lat, test_address.Lng, addy.Lat, addy.Lng)
		}
	}

}

func TestReverseGeocode(t *testing.T) {

	for _, test_address := range Addresses {
		// t.Logf("checking %s\n", test_address)
		latlng := fmt.Sprintf("%f,%f", test_address.Lat, test_address.Lng)
		addy, err := ReverseGeocode(latlng)
		if err != nil {
			t.Error(err)
		}
		if addy.Address != test_address.Address {
			t.Errorf("Expected: %s, Got: %s", test_address.Address, addy.Address)
		}
		if addy.Lat != test_address.Lat || addy.Lng != test_address.Lng {
			t.Errorf("Expected: %f:%f, Got: %f:%f", test_address.Lat, test_address.Lng, addy.Lat, addy.Lng)
		}
	}

}
