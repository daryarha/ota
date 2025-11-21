## Table of Contents
- [Installation](#installation)
- [Parameter](#parameters)
- [Usage](#usage)

## Installation

1. Use go version 1.24.0 (if possible, to make sure that version I use same with what you run)
2. Clone the repo:
   ```bash
   git clone git@github.com:daryarha/ota.git
3. Update the module first:
   ```bash
   go mod tidy
   go mod vendor
3. Run the app:
   ```bash
   go run ./cmd/main.go
   # or
   go run cmd/main.go


## Parameters

| Parameter | Type | Description | Default | Possible Values |
|----------|------|-------------|---------|----------------|
| origin | string | Origin airport IATA code | "" | "CGK" |
| destination | string | Destination airport IATA code | "" | "DPS" |
| departureDate | string | Departure date (YYYY-MM-DD) | "" | "2025-12-15" |
| passengers | int | Number of passengers | 1 | any integer ≥ 1 |
| cabinClass | string | Cabin class | "" | "economy", "Y" |
| filter | object | Filtering options | {} | Filter |
| filter.priceRange | object | Price range | {min:0,max:0} | Range |
| filter.priceRange.min | int | Min price | 0 | ≥ 0 |
| filter.priceRange.max | int | Max price | 0 | ≥ 0 |
| filter.stops | int (optional) | Number of stops | null | 0, 1, 2, null |
| filter.airlines | string[] | Allowed airlines | [] | ["garuda indonesia", "batik air", "lion air", "airasia"] |
| filter.durationRange | object | Duration range (minutes) | {min:0,max:0} | Range |
| filter.durationRange.min | int | Minimum duration | 0 | ≥ 0 |
| filter.durationRange.max | int | Maximum duration | 0 | ≥ 0 |
| sort | object | Sorting options | {} | Sort |
| sort.field | string | Field to sort by | "" | "price", "duration", "departure", "arrival" |
| sort.sortBy | string | Sorting order | "desc" | "asc", "desc" |


## Usage
1. Use this Curl to hit the API:
    ```bash
    curl --location --request GET 'localhost:8080/search' \
    --header 'Content-Type: application/json' \
    --data '{
        "origin": "CGK",
        "destination": "DPS",
        "departureDate": "2025-12-15",
        "returnDate": null,
        "passengers": 1,
        "cabinClass": "",
        "filter": {
            "priceRange": {
                "min": 0,
                "max": 0
            },
            "stops": 0,
            "airlines": [],
            "durationRange": {
                "min": 0,
                "max": 0
            }
        },
        "sort": {
            "field": "",
            "sortby": ""
        }
    }'