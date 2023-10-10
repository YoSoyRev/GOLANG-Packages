package countries

import (
	"errors"
	"fmt"
)

var myCountry string
var collection []string
var errNotFound error = errors.New("Country Not found")

// Add countries to the list .
func Add(country string) {
	collection = append(collection, country)
}

// SetMyCountry select a country as favorite country.
func SetMyCountry(country string) error {
	if !isInCollection(country) {
		return errNotFound
	}

	myCountry = country
	return nil
}
// List shows a sorted list of our previously added countries.
func List() {
	for i, c := range collection {
		myCountryMsg := ""
		if c == myCountry {
			myCountryMsg = "[my country]"
		}
		fmt.Println(i+1, c, myCountryMsg)
	}
}

func isInCollection(country string) bool {
	for _, c := range collection {
		if c == country {
			myCountry = country
			return true
		}
	}

	return false
}
