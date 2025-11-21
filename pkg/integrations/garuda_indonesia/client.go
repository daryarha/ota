package garuda_indonesia

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"ota/constant"
	"ota/internal/entity"
	"ota/pkg/times"
	"strings"
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
	resp := GarudaIndonesiaSearchResponse{}
	data, err := os.ReadFile(c.baseURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}

	if strings.ToLower(resp.Status) != "success" {
		err = errors.New("external server error")
		return
	}

	for _, val := range resp.Flights {
		res = append(res, val.ToAggregateFlight())
	}

	simulate := times.SimulateTime(constant.GarudaIndonesiaAPIMinTime, constant.GarudaIndonesiaAPIMaxTime)
	fmt.Println("Garuda API Time:", simulate)
	time.Sleep(simulate)

	return
}
