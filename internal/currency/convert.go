package currency

import (
	"errors"
	"os"
	"strconv"
)

func Convert(data []Currency) (float32, error) {
	args := os.Args

	if len(args) > 3 || len(args) < 3 {
		return 0, errors.New("Arguments error")
	}

	argValue, argName := args[1], args[2]

	value, err := strconv.ParseFloat(argValue, 32)
	if err != nil {
		return 0, errors.New("Failed to parse value from CLI")
	}

	cur, err := findCurrency(data, argName)
	if err != nil {
		return 0, err
	}

	return float32(value) / float32(cur.Units) * cur.ExchangeRate, nil
}

func findCurrency(data []Currency, arg string) (Currency, error) {
	for i := range data {
		if data[i].CharCode == arg {
			return data[i], nil
		}
	}

	return Currency{}, errors.New("Failed to find currency")
}
