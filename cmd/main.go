package main

import (
	"fmt"
	"net/http"
	"ota/constant"
	"ota/internal/domain"
	"ota/internal/route"
	"ota/pkg/config"
	"ota/pkg/integrations/airasia"
	"ota/pkg/integrations/batik_air"
	"ota/pkg/integrations/garuda_indonesia"
	"ota/pkg/integrations/lion_air"
)

func main() {
	conf, err := config.GetConfiguration()
	if err != nil {
		fmt.Printf("error get config, err: %v", err)
	}

	airasiaClient := airasia.NewClient(conf.GetHost(constant.Airasia))
	lionAirClient := lion_air.NewClient(conf.GetHost(constant.LionAir))
	batikAirClient := batik_air.NewClient(conf.GetHost(constant.BatikAir))
	garudaIndonesiaClient := garuda_indonesia.NewClient(conf.GetHost(constant.GarudaIndonesia))

	providers := []domain.ProviderInterface{airasiaClient, lionAirClient, batikAirClient, garudaIndonesiaClient}

	routes := route.RegisterRoutes(providers)

	fmt.Println("starting server in:", conf.Server.Host)
	http.ListenAndServe(conf.Server.Host, routes)

}
