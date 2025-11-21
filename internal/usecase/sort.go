package usecase

import (
	"ota/data"
	"ota/internal/entity"
	"sort"
)

func (s *SearchUsecase) sort(res entity.FlightResponse, req entity.Sort) {
	sort.Slice(res.Flights, func(i, j int) bool {
		switch req.Field {
		case data.SortFieldPrice:
			if req.SortBy == data.SortByAsc {
				return res.Flights[i].Price.Amount < res.Flights[j].Price.Amount
			}
			return res.Flights[i].Price.Amount > res.Flights[j].Price.Amount
		case data.SortFieldDuration:
			if req.SortBy == data.SortByAsc {
				return res.Flights[i].Duration.TotalMinutes < res.Flights[j].Duration.TotalMinutes
			}
			return res.Flights[i].Duration.TotalMinutes > res.Flights[j].Duration.TotalMinutes
		case data.SortFieldDeparture:
			if req.SortBy == data.SortByAsc {
				return res.Flights[i].Departure.Timestamp < res.Flights[j].Departure.Timestamp
			}
			return res.Flights[i].Departure.Timestamp > res.Flights[j].Departure.Timestamp
		case data.SortFieldArrival:
			if req.SortBy == data.SortByAsc {
				return res.Flights[i].Arrival.Timestamp < res.Flights[j].Arrival.Timestamp
			}
			return res.Flights[i].Arrival.Timestamp > res.Flights[j].Arrival.Timestamp
		default:
			if res.Flights[i].Price.Amount != res.Flights[j].Price.Amount {
				return res.Flights[i].Price.Amount < res.Flights[j].Price.Amount
			}
			if res.Flights[i].Duration.TotalMinutes != res.Flights[j].Duration.TotalMinutes {
				return res.Flights[i].Duration.TotalMinutes < res.Flights[j].Duration.TotalMinutes
			}
			return res.Flights[i].Stops < res.Flights[j].Stops
		}
	})
}
