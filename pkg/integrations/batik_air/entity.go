package batik_air

type BatikAirSearchResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Results []BatikAirSearchData `json:"results"`
}

type BatikAirSearchData struct {
	FlightNumber      string       `json:"flightNumber"`
	AirlineName       string       `json:"airlineName"`
	AirlineIATA       string       `json:"airlineIATA"`
	Origin            string       `json:"origin"`
	Destination       string       `json:"destination"`
	DepartureDateTime string       `json:"departureDateTime"`
	ArrivalDateTime   string       `json:"arrivalDateTime"`
	TravelTime        string       `json:"travelTime"`
	NumberOfStops     int          `json:"numberOfStops"`
	Fare              BatikAirFare `json:"fare"`
	SeatsAvailable    int          `json:"seatsAvailable"`
	AircraftModel     string       `json:"aircraftModel"`
	BaggageInfo       string       `json:"baggageInfo"`
	OnboardServices   []string     `json:"onboardServices"`
}

type BatikAirFare struct {
	BasePrice    int64  `json:"basePrice"`
	Taxes        int64  `json:"taxes"`
	TotalPrice   int64  `json:"totalPrice"`
	CurrencyCode string `json:"currencyCode"`
	Class        string `json:"class"`
}
