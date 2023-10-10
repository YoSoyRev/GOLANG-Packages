package main

import "main/countries"

func main() {
	countries.Add("Mexico")
	countries.Add("Rusia")
	countries.Add("Bielorusia")
	countries.Add("Tlaxcala")

	countries.SetMyCountry("Tlaxcala")

	countries.List()
}
