package garuda_indonesia

type GarudaIndonesiaSearchResponse struct {
	Status  string                      `json:"status"`
	Flights []GarudaIndonesiaSearchData `json:"flights"`
}

type GarudaIndonesiaSearchData struct {
	FlightID        string                  `json:"flight_id"`
	Airline         string                  `json:"airline"`
	AirlineCode     string                  `json:"airline_code"`
	Departure       GarudaIndonesiaSchedule `json:"departure"`
	Arrival         GarudaIndonesiaSchedule `json:"arrival"`
	DurationMinutes int                     `json:"duration_minutes"`
	Stops           int                     `json:"stops"`
	Aircraft        string                  `json:"aircraft"`
	Price           GarudaIndonesiaPrice    `json:"price"`
	AvailableSeats  int                     `json:"available_seats"`
	FareClass       string                  `json:"fare_class"`
	Baggage         GarudaIndonesiaBaggage  `json:"baggage"`
	Amenities       []string                `json:"amenities"`
}

type GarudaIndonesiaSchedule struct {
	Airport  string `json:"airport"`
	City     string `json:"city"`
	Time     string `json:"time"`
	Terminal string `json:"terminal"`
}

type GarudaIndonesiaPrice struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type GarudaIndonesiaBaggage struct {
	CarryOn int `json:"carry_on"`
	Checked int `json:"checked"`
}
