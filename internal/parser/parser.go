package parser

import (
	"currency-converter/internal/currency"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func GetData() []currency.Currency {
	c := colly.NewCollector()

	var currencyData []currency.Currency

	c.OnHTML(`tr`, func(e *colly.HTMLElement) {
		lines := e.ChildTexts("td")

		if len(lines) == 0 {
			return
		}

		units, err := strconv.Atoi(lines[2])

		if err != nil {
			log.Fatal(err)
		}

		lines[4] = strings.Replace(lines[4], ",", ".", 1)

		exchangeRate, err := strconv.ParseFloat(lines[4], 32)

		if err != nil {
			log.Fatal(err)
		}

		currencyData = append(currencyData, currency.Currency{
			DigitalCode:  lines[0],
			CharCode:     lines[1],
			Units:        units,
			ValuteName:   lines[3],
			ExchangeRate: float32(exchangeRate),
		})
	})

	c.Visit("https://www.cbr.ru/currency_base/daily/")

	return currencyData
}
