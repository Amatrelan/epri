package nordpool

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUrl(t *testing.T) {
	country := FI
	currency := EUR

	generated := getUrl(country, currency, "2024-10-13")
	want := "https://dataportal-api.nordpoolgroup.com/api/DayAheadPrices?market=DayAhead&deliveryArea=FI&currency=EUR&date=2024-10-13"

	if generated != want {
		t.Error("not same string")
	}
}

func TestJsonParse(t *testing.T) {
	assert := assert.New(t)
	dat, err := os.ReadFile("../test/data.json")
	if err != nil {
		t.Fatal(err)
	}

	var resp NordPoolResponse
	json.Unmarshal(dat, &resp)

	assert.Equal(resp.Currency, EUR, "should be euro")

	ti, _ := time.Parse("2006-01-02T15:04:05Z0700", "2024-10-16T11:24:39.1821423Z")
	assert.Equal(ti, resp.UpdatedAt)

	assert.Equal(resp.Market, "DayAhead")
	assert.Equal(resp.DeliveryAreas, []Country{FI})
	start, _ := time.Parse("2006-01-02T15:04:05Z0700", "2024-10-16T22:00:00Z")
	end, _ := time.Parse("2006-01-02T15:04:05Z0700", "2024-10-16T23:00:00Z")
	assert.Equal(resp.MultiAreaEntries[0], AreaEntry{
		EntryPerArea:  map[Country]float64{FI: 0.01},
		DeliveryStart: start,
		DeliveryEnd:   end,
	})
}
