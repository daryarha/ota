package usecase

import (
	"encoding/json"
	"ota/internal/cache"
	"ota/internal/domain"
	"ota/internal/entity"
	"sync"
	"time"
)

type SearchUsecase struct {
	providers []domain.ProviderInterface
	cache     *cache.InMemCache
}

func NewSearchUsecase(providers []domain.ProviderInterface, cache *cache.InMemCache) SearchUsecase {
	return SearchUsecase{
		providers: providers,
		cache:     cache,
	}
}

func (s SearchUsecase) SearchFlight(req entity.SearchCriteriaRequest) entity.FlightResponse {
	startTime := time.Now()
	key := s.cacheKey(req)
	//get cache data if exist
	if data, ok := s.cache.Get(key); ok {
		var resp entity.FlightResponse
		json.Unmarshal([]byte(data), &resp)
		resp.Metadata.CacheHit = true
		return resp
	}
	//get data concurrently from multiple external apis
	res := s.getFlights()
	//filter data
	filtered := s.filter(res, req)
	//sort data
	s.sort(filtered, req.Sort)
	endTime := time.Since(startTime)
	filtered.Metadata.SearchTimeMs = int(endTime.Milliseconds())

	//cache data
	data, _ := json.Marshal(filtered)
	s.cache.Set(key, string(data))

	return filtered
}

func (s SearchUsecase) cacheKey(req entity.SearchCriteriaRequest) string {
	b, _ := json.Marshal(req)
	return string(b)
}

func (s SearchUsecase) getFlights() (res entity.FlightResponse) {
	var (
		//for response
		flights       = []entity.FlightResult{}
		totalProvider = len(s.providers)
		totalFailed   = 0
		//for go routine
		wg      sync.WaitGroup
		resChan = make(chan []entity.FlightResult, totalProvider)
		errChan = make(chan error, totalProvider)
	)

	for _, p := range s.providers {
		wg.Add(1)
		go func(provider domain.ProviderInterface) {
			defer wg.Done()
			flights, err := provider.GetFlightSearch()
			if err != nil {
				errChan <- err
				return
			}

			resChan <- flights
		}(p)
	}

	wg.Wait()
	close(resChan)
	close(errChan)
	for r := range resChan {
		flights = append(flights, r...)
	}
	for _ = range errChan {
		totalFailed++
	}

	metadata := entity.MetadataResponse{
		ProvidersQueried:   totalProvider,
		ProvidersFailed:    totalFailed,
		ProvidersSucceeded: totalProvider - totalFailed,
	}

	res = entity.FlightResponse{
		Metadata: metadata,
		Flights:  flights,
	}
	return
}
