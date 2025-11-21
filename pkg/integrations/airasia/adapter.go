package airasia

import (
	"fmt"
	"ota/constant"
	"ota/data"
	"ota/internal/entity"
	"ota/pkg/format"
	"strings"
)

func (f *AirasiaFlight) ToAggregateFlight() entity.FlightResult {
	//handle baggage
	bags := strings.Split(f.BaggageNote, ",")
	if len(bags) != 2 {
		bags = []string{"", ""}
	}
	bags[1] = strings.ReplaceAll(bags[1], "checked bags", "")

	minutes := int(f.DurationHours * 60)

	for _, s := range f.Stops {
		minutes += s.WaitTimeMinutes
	}

	temp := entity.FlightResult{
		ID:       fmt.Sprintf("%s_%s", f.FlightCode, f.Airline),
		Provider: f.Airline,
		Airline: entity.AirlineResult{
			Name: f.Airline,
			Code: format.FlightCodeToCode(f.FlightCode),
		},
		FlightNumber: f.FlightCode,
		Departure: entity.Schedule{
			Airport:   f.FromAirport,
			City:      data.GetCityName(f.FromAirport),
			DateTime:  f.DepartTime,
			Timestamp: format.DateTimeToUnix(f.DepartTime),
		},
		Arrival: entity.Schedule{
			Airport:   f.ToAirport,
			City:      data.GetCityName(f.ToAirport),
			DateTime:  f.ArriveTime,
			Timestamp: format.DateTimeToUnix(f.ArriveTime),
		},
		Duration: entity.Duration{
			TotalMinutes: minutes,
			Formatted:    format.MinutesToFormatted(minutes),
		},
		Stops: len(f.Stops),
		Price: entity.Price{
			Amount:    f.PriceIDR,
			Currency:  constant.IDR,
			Formatted: fmt.Sprintf("%s%s", constant.Rupiah, format.CurrencyComma(f.PriceIDR)),
		},
		AvailableSeats: f.Seats,
		CabinClass:     f.CabinClass,
		Baggage: entity.Baggage{
			CarryOn: strings.TrimSpace(bags[0]),
			Checked: strings.TrimSpace(bags[1]),
		},
		Amenities: []string{},
	}
	return temp
}
