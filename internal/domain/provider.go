package domain

import "ota/internal/entity"

type ProviderInterface interface {
	GetFlightSearch() (res []entity.FlightResult, err error)
}
