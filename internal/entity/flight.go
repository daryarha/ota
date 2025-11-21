package entity

type FlightResponse struct {
	SearchCriteria SearchCriteriaResponse `json:"search_criteria"`
	Metadata       MetadataResponse       `json:"metadata"`
	Flights        []FlightResult         `json:"flights"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

type SearchCriteriaResponse struct {
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departure_date"`
	Passengers    int    `json:"passengers"`
	CabinClass    string `json:"cabin_class"`
}

type SearchCriteriaRequest struct {
	Origin        string  `json:"origin"`
	Destination   string  `json:"destination"`
	DepartureDate string  `json:"departureDate"`
	ReturnDate    *string `json:"returnDate"`
	Passengers    int     `json:"passengers"`
	CabinClass    string  `json:"cabinClass"`
	Filter        Filter  `json:"filter"`
	Sort          Sort    `json:"sort"`
}

type Sort struct {
	Field  string `json:"field"`
	SortBy string `json:"sortBy"`
}

type Filter struct {
	PriceRange    Range    `json:"priceRange"`
	Stops         *int     `json:"stops"`
	Airlines      []string `json:"airlines"`
	DurationRange Range    `json:"durationRange"`
}

type Range struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

type MetadataResponse struct {
	TotalResults       int  `json:"total_results"`
	ProvidersQueried   int  `json:"providers_queried"`
	ProvidersSucceeded int  `json:"providers_succeeded"`
	ProvidersFailed    int  `json:"providers_failed"`
	SearchTimeMs       int  `json:"search_time_ms"`
	CacheHit           bool `json:"cache_hit"`
}

type FlightResult struct {
	ID             string        `json:"id"`
	Provider       string        `json:"provider"`
	Airline        AirlineResult `json:"airline"`
	FlightNumber   string        `json:"flight_number"`
	Departure      Schedule      `json:"departure"`
	Arrival        Schedule      `json:"arrival"`
	Duration       Duration      `json:"duration"`
	Stops          int           `json:"stops"`
	Price          Price         `json:"price"`
	AvailableSeats int           `json:"available_seats"`
	CabinClass     string        `json:"cabin_class"`
	Aircraft       *string       `json:"aircraft"`
	Amenities      []string      `json:"amenities"`
	Baggage        Baggage       `json:"baggage"`
}

type AirlineResult struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Schedule struct {
	Airport   string `json:"airport"`
	City      string `json:"city"`
	DateTime  string `json:"date_time"`
	Timestamp int64  `json:"timestamp"`
}

type Duration struct {
	TotalMinutes int    `json:"total_minutes"`
	Formatted    string `json:"formatted"`
}

type Price struct {
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	Formatted string `json:"formatted"`
}

type Baggage struct {
	CarryOn string `json:"carry_on"`
	Checked string `json:"checked"`
}
