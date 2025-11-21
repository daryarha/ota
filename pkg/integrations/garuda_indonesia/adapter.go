package garuda_indonesia

import (
	"fmt"
	"ota/constant"
	"ota/internal/entity"
	"ota/pkg/format"
)

func (f *GarudaIndonesiaSearchData) ToAggregateFlight() entity.FlightResult {
	temp := entity.FlightResult{
		ID:       fmt.Sprintf("%s_%s", f.FlightID, f.Airline),
		Provider: f.Airline,
		Airline: entity.AirlineResult{
			Name: f.Airline,
			Code: f.AirlineCode,
		},
		FlightNumber: f.FlightID,
		Departure: entity.Schedule{
			Airport:   f.Departure.Airport,
			City:      f.Departure.City,
			DateTime:  f.Departure.Time,
			Timestamp: format.DateTimeToUnix(f.Departure.Time),
		},
		Arrival: entity.Schedule{
			Airport:   f.Arrival.Airport,
			City:      f.Arrival.City,
			DateTime:  f.Arrival.Time,
			Timestamp: format.DateTimeToUnix(f.Arrival.Time),
		},
		Duration: entity.Duration{
			TotalMinutes: f.DurationMinutes,
			Formatted:    format.MinutesToFormatted(f.DurationMinutes),
		},
		Stops: f.Stops,
		Price: entity.Price{
			Amount:    f.Price.Amount,
			Currency:  f.Price.Currency,
			Formatted: fmt.Sprintf("%s%s", constant.Rupiah, format.CurrencyComma(f.Price.Amount)),
		},
		AvailableSeats: f.AvailableSeats,
		CabinClass:     f.FareClass,
		Aircraft:       &f.Aircraft,
		Amenities:      f.Amenities,
		Baggage: entity.Baggage{
			CarryOn: fmt.Sprintf("%d kg", f.Baggage.CarryOn),
			Checked: fmt.Sprintf("%d kg", f.Baggage.Checked),
		},
	}
	return temp
}
