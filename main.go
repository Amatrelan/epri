package main

import (
	nord "epri/nordpool"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

const CONFIG_FOLDER = "$XDG_CONFIG_HOME/epri"

type calc func(float64) float64

func main() {
	loadConfig()

	location := nord.Country(viper.GetString("location"))
	currency := nord.Currency(viper.GetString("currency"))
	tax := (1 + viper.GetFloat64("tax")/1000)

	data := fromNordPool(nord.GetData(location, currency, time.Now()), tax)

	columns := []table.Column{
		{Title: "Date", Width: 5},
		{Title: "Time", Width: 5},
		{Title: "price/kWh", Width: 10},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(data),
		table.WithFocused(true),
		table.WithHeight(25),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240")).BorderBottom(true).Bold(true)
	s.Selected = s.Selected.Foreground(lipgloss.Color("229")).Background(lipgloss.Color("57")).Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		slog.Error("running program:", "err", err)
		os.Exit(1)
	}
}

func currencyToSign(c nord.Currency) string {
	subunit := viper.GetBool("subunit")
	switch c {
	case nord.EUR:
		if subunit {
			return "¢"
		} else {
			return "€"
		}
	case nord.PLN:
		if subunit {
			return "gr"
		} else {
			return "zł"
		}
	case nord.NOK:
		if subunit {
			return "øre"
		} else {
			return "kr"
		}
	case nord.DKK:
		if subunit {
			return "øre"
		} else {
			return "kr"
		}
	case nord.SEK:
		if subunit {
			return "øre"
		} else {
			return "kr"
		}
	default:
		return ""
	}
}

func fromNordPool(data nord.NordPoolResponse, tax float64) []table.Row {
	changeCalc(&data, func(v float64) float64 {
		return v / 1000
	})

	if viper.GetBool("subunit") {
		changeCalc(&data, func(v float64) float64 {
			return v * 100
		})
	}

	var out []table.Row
	for i := range data.MultiAreaEntries {
		elem := data.MultiAreaEntries[i]
		delivery_time := elem.DeliveryStart.Local()
		date := fmt.Sprintf("%0d.%0d", delivery_time.Day(), delivery_time.Month())
		time := fmt.Sprintf("%0d:%02d", delivery_time.Hour(), delivery_time.Minute())

		location := nord.Country(viper.GetString("location"))

		listed := elem.EntryPerArea[location]
		if listed > 0 {
			listed = listed * tax
		}
		amount := fmt.Sprintf("%.2f%s", listed, currencyToSign(data.Currency))
		out = append(out, table.Row{
			date,
			time,
			amount,
		})
	}

	return out
}

func configPath() string {
	return strings.ReplaceAll(CONFIG_FOLDER, "$XDG_CONFIG_HOME", os.Getenv("XDG_CONFIG_HOME"))
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(CONFIG_FOLDER)
	viper.AddConfigPath(".")
	viper.SetDefault("location", nord.FI)
	viper.SetDefault("currency", nord.EUR)
	viper.SetDefault("tax", 25.5)
	viper.SetDefault("subunit", true)
	viper.ReadInConfig()

	os.MkdirAll(configPath(), os.ModePerm)
	viper.SafeWriteConfig()
}
