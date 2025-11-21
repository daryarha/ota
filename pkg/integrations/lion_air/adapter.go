package lion_air

import (
	"fmt"
	"ota/constant"
	"ota/internal/entity"
	"ota/pkg/format"
	"strings"
	"time"
)

func (f *LionAirFlight) ToAggregateFlight() entity.FlightResult {

	arrivalDateTime := parseDateTime(f.Schedule.Arrival, f.Schedule.ArrivalTimezone)
	departureDateTime := parseDateTime(f.Schedule.Departure, f.Schedule.DepartureTimezone)
	minutes := f.FlightTime
	for _, l := range f.Layovers {
		minutes += l.DurationMinutes
	}
	temp := entity.FlightResult{
		ID:       fmt.Sprintf("%s_%s", f.ID, f.Carrier.Name),
		Provider: f.Carrier.Name,
		Airline: entity.AirlineResult{
			Name: f.Carrier.Name,
			Code: f.Carrier.Iata,
		},
		FlightNumber: f.ID,
		Departure: entity.Schedule{
			Airport:   f.Route.From.Code,
			City:      f.Route.From.City,
			DateTime:  departureDateTime,
			Timestamp: format.DateTimeToUnix(departureDateTime),
		},
		Arrival: entity.Schedule{
			Airport:   f.Route.To.Code,
			City:      f.Route.To.City,
			DateTime:  arrivalDateTime,
			Timestamp: format.DateTimeToUnix(arrivalDateTime),
		},
		Duration: entity.Duration{
			TotalMinutes: minutes,
			Formatted:    format.MinutesToFormatted(minutes),
		},
		Stops: f.StopCount,
		Price: entity.Price{
			Amount:    f.Pricing.Total,
			Currency:  f.Pricing.Currency,
			Formatted: fmt.Sprintf("%s%s", constant.Rupiah, format.CurrencyComma(f.Pricing.Total)),
		},
		AvailableSeats: f.SeatsLeft,
		CabinClass:     strings.ToLower(f.Pricing.FareType),
		Aircraft:       &f.PlaneType,
		Amenities:      f.Services.ToAmenities(),
		Baggage: entity.Baggage{
			CarryOn: fmt.Sprintf("%s", f.Services.BaggageAllowance.Cabin),
			Checked: fmt.Sprintf("%s", f.Services.BaggageAllowance.Hold),
		},
	}

	return temp
}

func parseDateTime(str string, loc string) string {
	tz, err := time.LoadLocation(loc)
	if err != nil || tz == nil {
		return ""
	}
	tm, err := time.ParseInLocation(constant.LionAirDateTimeFormat, str, tz)
	if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
	}
	return tm.Format(time.RFC3339)
}

func (f LionAirServices) ToAmenities() []string {
	amenities := []string{}
	if f.MealsIncluded {
		amenities = append(amenities, "meals")
	}
	if f.WifiAvailable {
		amenities = append(amenities, "wifi")
	}
	return amenities
}
