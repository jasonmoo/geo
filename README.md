# geo 

Just a simple go wrapper for the Google Geocoding API (https://developers.google.com/maps/documentation/geocoding/)

Usage:

	package main

	import 	(
		"fmt"
		"github.com/jasonmoo/geo"	
	)

	add, err := geo.Geocode("555 w 18th st, ny, ny")
	
	fmt.Println(add)
	
	add2, err2 := geo.ReverseGeocode("40.7453721,-74.0078293")
	
	fmt.Println(add2)
	
	>> &geo.Address{Lat:40.7453721, Lng:-74.0078293, Address:"555 W 18th St, New York, NY 10011, USA", Response:(*geo.Response)(0x11185d60)}
	
