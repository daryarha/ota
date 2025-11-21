package data

import "strings"

const (
	Airasia         = "airasia"
	GarudaIndonesia = "garuda indonesia"
	BatikAir        = "batik air"
	LionAir         = "lion air"
)

var (
	mapAirlineField = map[string]bool{
		Airasia:         true,
		GarudaIndonesia: true,
		BatikAir:        true,
		LionAir:         true,
	}
)

func IsAirlineFieldValid(field string) bool {
	return mapAirlineField[strings.ToLower(field)]
}
