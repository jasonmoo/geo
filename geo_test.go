package geo

import (
	"testing"
	"fmt"
)

func TestGeocode(t *testing.T) {

	add, err := Geocode("555 west 18th street, manhattan, ny")
	if err != nil {
		t.Error(err)
	}
	if add.Address != "555 W 18th St, New York, NY 10011, USA" {
		t.Errorf("Incorrect address returned (%v)", add)
	}
	if add.Lat != 40.7453721 || add.Lng != -74.0078293 {
		t.Errorf("Incorrect coords returned (%v)", add)
	}
	fmt.Println(add)

	add, err = Geocode("Bartlett, Il")
	if err != nil {
		t.Error(err)
	}
	if add.Address != "Bartlett, IL, USA" {
		t.Errorf("Incorrect address returned (%v)", add)
	}
	if add.Lat != 41.9950276 || add.Lng != -88.1856301 {
		t.Errorf("Incorrect coords returned (%v)", add)
	}
	fmt.Println(add)

}

func TestReverseGeocode(t *testing.T) {

	add, err := ReverseGeocode("40.7453721,-74.0078293")
	if err != nil {
		t.Error(err)
	}
	if add.Address != "555 W 18th St, Manhattan, NY 10011, USA" {
		t.Errorf("Incorrect address returned (%v)", add)
	}
	if add.Lat != 40.7453721 || add.Lng != -74.0078293 {
		t.Errorf("Incorrect coords returned (%v)", add)
	}
	fmt.Println(add)

	add, err = Geocode("41.9950276,-88.1856301")
	if err != nil {
		t.Error(err)
	}
	if add.Address != "Village Hall, Bartlett, IL 60103, USA" {
		t.Errorf("Incorrect address returned (%v)", add)
	}
	if add.Lat != 41.9950276 || add.Lng != -88.1856301 {
		t.Errorf("Incorrect coords returned (%v)", add)
	}
	fmt.Println(add)

}
