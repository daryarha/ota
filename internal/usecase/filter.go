package usecase

import (
	"ota/internal/entity"
	"strings"
)

func (s SearchUsecase) filter(res entity.FlightResponse, req entity.SearchCriteriaRequest) (filteredRes entity.FlightResponse) {
	filteredFlights := []entity.FlightResult{}
	for _, r := range res.Flights {
		//if arrival first before departure skip
		if r.Arrival.Timestamp < r.Departure.Timestamp {
			continue
		}

		//filter by price range
		if req.Filter.PriceRange.Max > 0 {
			//if outside range skip
			if r.Price.Amount < req.Filter.PriceRange.Min || r.Price.Amount > req.Filter.PriceRange.Max {
				continue
			}
		}

		//filter if number stops exist or not, if exist check equal or not
		if req.Filter.Stops != nil {
			stops := *req.Filter.Stops
			if stops != r.Stops {
				continue
			}
		}

		//filter duration range
		if req.Filter.DurationRange.Max > 0 {
			//if outside range skip
			if int64(r.Duration.TotalMinutes) < req.Filter.DurationRange.Min || int64(r.Duration.TotalMinutes) > req.Filter.DurationRange.Max {
				continue
			}
		}

		if len(req.Filter.Airlines) > 0 {
			isAdd := false
			for _, al := range req.Filter.Airlines {
				if strings.EqualFold(al, r.Provider) {
					isAdd = true
					break
				}
			}

			if !isAdd {
				continue
			}
		}

		if req.DepartureDate != "" {
			//if req date is not substring of departure skip
			if !strings.Contains(r.Departure.DateTime, req.DepartureDate) {
				continue
			}
		}

		if req.Origin != "" {
			//if origin is not substring of airport origin
			if !strings.Contains(r.Departure.Airport, req.Origin) {
				continue
			}
		}

		if req.Destination != "" {
			//if destination is not substring of airport destination
			if !strings.Contains(r.Arrival.Airport, req.Destination) {
				continue
			}
		}

		if req.CabinClass != "" {
			//if req cabinclass is not substring of cabinclass
			if !strings.Contains(r.CabinClass, req.CabinClass) {
				continue
			}
		}

		if req.Passengers > 0 {
			//if available seats is lower than total passengers needed skip
			if r.AvailableSeats < req.Passengers {
				continue
			}
		}

		filteredFlights = append(filteredFlights, r)
	}

	filteredRes = entity.FlightResponse{
		Metadata: entity.MetadataResponse{
			TotalResults:       len(filteredFlights),
			ProvidersQueried:   res.Metadata.ProvidersQueried,
			ProvidersSucceeded: res.Metadata.ProvidersSucceeded,
			ProvidersFailed:    res.Metadata.ProvidersFailed,
		},
		Flights: filteredFlights,
		SearchCriteria: entity.SearchCriteriaResponse{
			Origin:        req.Origin,
			Destination:   req.Destination,
			DepartureDate: req.DepartureDate,
			Passengers:    req.Passengers,
			CabinClass:    req.CabinClass,
		},
	}

	return
}
