package batik_air

import (
	"fmt"
	"ota/constant"
	"ota/data"
	"ota/internal/entity"
	"ota/pkg/format"
	"strings"
	"time"
)

func (f *BatikAirSearchData) ToAggregateFlight() entity.FlightResult {
	//handle baggage
	bags := strings.Split(f.BaggageInfo, ",")
	if len(bags) != 2 {
		bags = []string{"", ""}
	}
	for i, b := range bags {
		bags[i] = strings.ReplaceAll(b, "cabin", "")
		bags[i] = strings.ReplaceAll(b, "checked", "")
	}

	arriveTime := parseDateTime(f.ArrivalDateTime)
	departTime := parseDateTime(f.DepartureDateTime)

	temp := entity.FlightResult{
		ID:       fmt.Sprintf("%s_%s", f.FlightNumber, f.AirlineName),
		Provider: f.AirlineName,
		Airline: entity.AirlineResult{
			Name: f.AirlineName,
			Code: f.AirlineIATA,
		},
		FlightNumber: f.FlightNumber,
		Departure: entity.Schedule{
			Airport:   f.Origin,
			City:      data.GetCityName(f.Origin),
			DateTime:  departTime,
			Timestamp: format.DateTimeToUnix(departTime),
		},
		Arrival: entity.Schedule{
			Airport:   f.Destination,
			City:      data.GetCityName(f.Destination),
			DateTime:  arriveTime,
			Timestamp: format.DateTimeToUnix(arriveTime),
		},
		Duration: entity.Duration{
			TotalMinutes: format.FormattedToMinutes(f.TravelTime),
			Formatted:    f.TravelTime,
		},
		Stops: f.NumberOfStops,
		Price: entity.Price{
			Amount:    f.Fare.TotalPrice,
			Currency:  f.Fare.CurrencyCode,
			Formatted: fmt.Sprintf("%s%s", constant.Rupiah, format.CurrencyComma(f.Fare.TotalPrice)),
		},
		AvailableSeats: f.SeatsAvailable,
		CabinClass:     f.Fare.Class,
		Aircraft:       &f.AircraftModel,
		Amenities:      f.OnboardServices,
		Baggage: entity.Baggage{
			CarryOn: strings.TrimSpace(bags[0]),
			Checked: strings.TrimSpace(bags[1]),
		},
	}

	return temp
}

func parseDateTime(date string) string {
	tm, err := time.Parse(constant.BatikAirDateTimeFormat, date)
	if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
		return ""
	}

	return tm.Format(time.RFC3339)
}
