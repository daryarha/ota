package data

import "strings"

const (
	SortByAsc          = "asc"
	SortFieldPrice     = "price"
	SortFieldDuration  = "duration"
	SortFieldDeparture = "departure"
	SortFieldArrival   = "arrival"
)

var (
	mapSortField = map[string]bool{
		SortFieldPrice:     true,
		SortFieldDuration:  true,
		SortFieldDeparture: true,
		SortFieldArrival:   true,
	}
)

func IsSortFieldValid(field string) bool {
	return mapSortField[strings.ToLower(field)]
}
