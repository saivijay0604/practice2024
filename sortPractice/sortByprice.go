package sortPractice

import (
	"fmt"
	"sort"
)

type Flight struct {
	Destination string
	Price       int
}

func SortByPrice(flights []Flight) []Flight {
	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price > flights[j].Price
	})
	return flights
}

func PrintFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Airline: %s, Price: %d\n", flight.Destination, flight.Price)
	}
}
