package nordpool

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

// https://dataportal-api.nordpoolgroup.com/api/DayAheadPrices?market=DayAhead&deliveryArea=FI&currency=EUR&date=2024-10-13
const BASE_URL = "https://dataportal-api.nordpoolgroup.com/api/DayAheadPrices?market=DayAhead&deliveryArea=%s&currency=%s&date=%s"

type Country string

const (
	FI  Country = "FI"
	DK1 Country = "DK1"
	DK2 Country = "DK2"
	NO1 Country = "NO1"
	NO2 Country = "NO2"
	NO3 Country = "NO3"
	NO4 Country = "NO4"
	NO5 Country = "NO5"
	SE1 Country = "SE1"
	SE2 Country = "SE2"
	SE3 Country = "SE3"
	SE4 Country = "SE4"
)

type Currency string

const (
	EUR Currency = "EUR"
	NOK Currency = "NOK"
	DKK Currency = "DKK"
	SEK Currency = "SEK"
	PLN Currency = "PLN"
)

type AreaEntry struct {
	EntryPerArea  map[Country]float64 `json:"entryPerArea"`
	DeliveryStart time.Time           `json:"deliveryStart"`
	DeliveryEnd   time.Time           `json:"deliveryEnd"`
}

type NordPoolResponse struct {
	DeliveryDate     string      `json:"deliveryDateCET"`
	UpdatedAt        time.Time   `json:"updatedAt"`
	Market           string      `json:"market"`
	Currency         Currency    `json:"currency"`
	DeliveryAreas    []Country   `json:"deliveryAreas"`
	MultiAreaEntries []AreaEntry `json:"multiAreaEntries"`
	Version          int32       `json:"version"`
}

func getUrl(country Country, currency Currency, d time.Time) string {
	date := fmt.Sprintf("%d-%d-%d", d.Year(), d.Month(), d.Day())
	return fmt.Sprintf(BASE_URL, country, currency, date)
}

func GetData(country Country, currency Currency, date time.Time) NordPoolResponse {
	url := getUrl(country, currency, date)
	r, err := http.Get(url)
	if err != nil {
		slog.Error("http response", "err", err)
		panic(err)
	}
	defer r.Body.Close()

	dat, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("reading response", "err", err)
		panic(err)
	}

	var resp NordPoolResponse
	json.Unmarshal(dat, &resp)

	return resp
}
