package airasia

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
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
	resp := AirasiaSearchResponse{}
	data, err := os.ReadFile(c.baseURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}

	if strings.ToLower(resp.Status) != "ok" {
		err = errors.New("external server error")
		return
	}
	//random number 0-99
	successRate := rand.Intn(100)
	//if equal or over the threshold simulate error with empty response
	if successRate >= constant.AirasiaAPISuccessRate {
		err = errors.New("error timeout")
		return
	}

	for _, val := range resp.Flights {
		res = append(res, val.ToAggregateFlight())
	}

	simulate := times.SimulateTime(constant.AirasiaAPIMinTime, constant.AirasiaAPIMaxTime)
	fmt.Println("Airasia API Time:", simulate)
	time.Sleep(simulate)

	return
}
