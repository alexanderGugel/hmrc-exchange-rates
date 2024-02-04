package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type ExchangeRate struct {
	XMLName      xml.Name `xml:"exchangeRate"`
	CountryName  string   `xml:"countryName"`
	CountryCode  string   `xml:"countryCode"`
	CurrencyName string   `xml:"currencyName"`
	CurrencyCode string   `xml:"currencyCode"`
	RateNew      string   `xml:"rateNew"`
}

type ExchangeRateMonthList struct {
	XMLName      xml.Name       `xml:"exchangeRateMonthList"`
	Period       string         `xml:"Period,attr"`
	ExchangeRate []ExchangeRate `xml:"exchangeRate"`
}

func main() {
	w := csv.NewWriter(os.Stdout)

	if err := w.Write([]string{
		"period_start",
		"period_end",
		"country_name",
		"country_code",
		"currency_name",
		"currency_code",
		"rate_new",
	}); err != nil {
		log.Fatal(err)
		return
	}

	t := time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC)

	for t.Before(time.Now()) {
		year, month, _ := t.Date()
		url := fmt.Sprintf("https://www.trade-tariff.service.gov.uk/api/v2/exchange_rates/files/monthly_xml_%d-%d.xml", year, month)
		log.Printf("fetching %s...\n", url)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return
		}

		list := ExchangeRateMonthList{}
		if err := xml.Unmarshal(body, &list); err != nil {
			log.Fatal(err)
			return
		}

		periodSplit := strings.Split(list.Period, " to ")
		periodStart, err := time.Parse("02/Jan/2006", periodSplit[0])
		if err != nil {
			log.Fatal(err)
		}
		periodEnd, err := time.Parse("02/Jan/2006", periodSplit[1])
		if err != nil {
			log.Fatal(err)
		}

		for _, exchangeRate := range list.ExchangeRate {
			if err := w.Write([]string{
				periodStart.Format("2006-01-02"),
				periodEnd.Format("2006-01-02"),
				exchangeRate.CountryName,
				exchangeRate.CountryCode,
				exchangeRate.CurrencyName,
				exchangeRate.CurrencyCode,
				exchangeRate.RateNew,
			}); err != nil {
				log.Fatal(err)
				return
			}
		}

		t = t.AddDate(0, 1, 0)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
