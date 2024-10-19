package main

import (
	"encoding/json"
	nord "epri/nordpool"
	"fmt"
	"math"
)

func PrettyPrint[K any](data K) {
	s, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(s))
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func changeCalc(data *nord.NordPoolResponse, fn calc) {
	for entry := range data.MultiAreaEntries {
		dat := data.MultiAreaEntries[entry]
		for k, v := range dat.EntryPerArea {
			dat.EntryPerArea[k] = fn(v)
		}
	}
}
