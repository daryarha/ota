package route

import (
	"net/http"
	"ota/constant"
	"ota/internal/cache"
	"ota/internal/domain"
	"ota/internal/handler"
	"ota/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(providers []domain.ProviderInterface) http.Handler {
	inmem := cache.NewInMemCache(constant.InMemTTL)
	searchUsecase := usecase.NewSearchUsecase(providers, inmem)
	searchHandler := handler.NewSearchHandler(searchUsecase)

	r := chi.NewRouter()

	r.Route("/search", func(r chi.Router) {
		r.Get("/", searchHandler.Flight)
	})
	return r
}
