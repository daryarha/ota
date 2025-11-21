package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ota/constant"
	"ota/data"
	"ota/internal/entity"
	"ota/internal/usecase"
	"time"
)

type Search struct {
	searchUC usecase.SearchUsecase
}

func NewSearchHandler(uc usecase.SearchUsecase) Search {
	return Search{
		searchUC: uc,
	}
}

func (s Search) Flight(w http.ResponseWriter, r *http.Request) {
	searchReq := entity.SearchCriteriaRequest{}
	json.NewDecoder(r.Body).Decode(&searchReq)
	invalid, param := s.IsInvalid(searchReq)
	if invalid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.ErrorResponse{
			ErrorMessage: fmt.Sprintf(constant.ErrorInvalidParam, param),
		})
		return
	}

	res := s.searchUC.SearchFlight(searchReq)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (s Search) IsInvalid(req entity.SearchCriteriaRequest) (bool, string) {
	if req.DepartureDate != "" {
		_, err := time.Parse(time.DateOnly, req.DepartureDate)
		if err != nil {
			return true, constant.DepartureDateInvalidFormat
		}
	}

	if req.ReturnDate != nil {
		_, err := time.Parse(time.DateOnly, *req.ReturnDate)
		if err != nil {
			return true, constant.ReturnDateInvalidFormat
		}
	}

	if req.Passengers <= 0 {
		return true, constant.PassengersZeroNegative
	}

	if req.Sort.Field != "" {
		if !data.IsSortFieldValid(req.Sort.Field) {
			return true, constant.SortFieldInvalidName
		}
	}

	if req.Filter.PriceRange.Max < 0 || req.Filter.PriceRange.Min < 0 {
		return true, constant.FilterPriceRangeMinMaxNegative
	}

	if req.Filter.PriceRange.Max < req.Filter.PriceRange.Min {
		return true, constant.FilterPriceRangeMaxLowerThanMin
	}

	if req.Filter.DurationRange.Max < 0 || req.Filter.DurationRange.Min < 0 {
		return true, constant.FilterDurationRangeMinMaxNegative
	}

	if req.Filter.DurationRange.Max < req.Filter.DurationRange.Min {
		return true, constant.FilterDurationRangeMaxLowerThanMin
	}

	if len(req.Filter.Airlines) > 0 {
		for _, al := range req.Filter.Airlines {
			if !data.IsAirlineFieldValid(al) {
				return true, constant.FilterAirlineInvalidName
			}
		}
	}

	return false, ""
}
