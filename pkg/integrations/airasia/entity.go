package airasia

type AirasiaSearchResponse struct {
	Status  string          `json:"status"`
	Flights []AirasiaFlight `json:"flights"`
}

type AirasiaFlight struct {
	FlightCode    string        `json:"flight_code"`
	Airline       string        `json:"airline"`
	FromAirport   string        `json:"from_airport"`
	ToAirport     string        `json:"to_airport"`
	DepartTime    string        `json:"depart_time"`
	ArriveTime    string        `json:"arrive_time"`
	DurationHours float32       `json:"duration_hours"`
	DirectFlight  bool          `json:"direct_flight"`
	PriceIDR      int64         `json:"price_idr"`
	Seats         int           `json:"seats"`
	CabinClass    string        `json:"cabin_class"`
	BaggageNote   string        `json:"baggage_note"`
	Stops         []AirasiaStop `json:"stops"`
}

type AirasiaStop struct {
	Airport         string `json:"airport"`
	WaitTimeMinutes int    `json:"wait_time_minutes"`
}
