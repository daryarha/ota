package lion_air

type LionAirSearchResponse struct {
	Success bool              `json:"success"`
	Data    LionAirSearchData `json:"data"`
}

type LionAirSearchData struct {
	AvailableFlights []LionAirFlight `json:"available_flights"`
}

type LionAirFlight struct {
	ID         string           `json:"id"`
	Carrier    LionAirCarrier   `json:"carrier"`
	Route      LionAirRoute     `json:"route"`
	Schedule   LionAirSchedule  `json:"schedule"`
	FlightTime int              `json:"flight_time"`
	IsDirect   bool             `json:"is_direct"`
	Pricing    LionAirPrice     `json:"pricing"`
	SeatsLeft  int              `json:"seats_left"`
	PlaneType  string           `json:"plane_type"`
	Services   LionAirServices  `json:"services"`
	StopCount  int              `json:"stop_count"`
	Layovers   []LionAirLayover `json:"layovers"`
}

type LionAirLayover struct {
	Airport         string `json:"airport"`
	DurationMinutes int    `json:"duration_minutes"`
}

type LionAirCarrier struct {
	Name string `json:"name"`
	Iata string `json:"iata"`
}

type LionAirRoute struct {
	From LionAirAirport `json:"from"`
	To   LionAirAirport `json:"to"`
}

type LionAirAirport struct {
	Code string `json:"code"`
	Name string `json:"name"`
	City string `json:"city"`
}

type LionAirSchedule struct {
	Departure         string `json:"departure"`
	DepartureTimezone string `json:"departure_timezone"`
	Arrival           string `json:"arrival"`
	ArrivalTimezone   string `json:"arrival_timezone"`
}

type LionAirPrice struct {
	Total    int64  `json:"total"`
	Currency string `json:"currency"`
	FareType string `json:"fare_type"`
}

type LionAirServices struct {
	WifiAvailable    bool           `json:"wifi_available"`
	MealsIncluded    bool           `json:"meals_included"`
	BaggageAllowance LionAirBaggage `json:"baggage_allowance"`
}

type LionAirBaggage struct {
	Cabin string `json:"cabin"`
	Hold  string `json:"hold"`
}
