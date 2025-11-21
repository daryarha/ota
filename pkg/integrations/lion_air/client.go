package lion_air

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"ota/constant"
	"ota/internal/entity"
	"ota/pkg/times"
	"time"
)

type Client struct {
	baseURL string
}

func NewClient(baseURL string) Client {
	return Client{
		baseURL: baseURL,
	}
}

func (c Client) GetFlightSearch() (res []entity.FlightResult, err error) {
	resp := LionAirSearchResponse{}
	data, err := os.ReadFile(c.baseURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}

	if !resp.Success {
		err = errors.New("external server error")
		return
	}

	for _, val := range resp.Data.AvailableFlights {
		res = append(res, val.ToAggregateFlight())
	}

	simulate := times.SimulateTime(constant.LionAirAPIMinTime, constant.LionAirAPIMaxTime)
	fmt.Println("Lion API Time:", simulate)
	time.Sleep(simulate)

	return
}
