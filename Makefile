all: hmrc_exchange_rates_usd.csv hmrc_exchange_rates_eur.csv

hmrc-exchange-rates: main.go
	go build

hmrc_exchange_rates.csv: hmrc-exchange-rates
	./hmrc-exchange-rates > hmrc_exchange_rates.csv

hmrc_exchange_rates_usd.csv: hmrc_exchange_rates.csv
	awk -F ',' "NR == 1 || \$$6 == \"USD\" {print}" hmrc_exchange_rates.csv > hmrc_exchange_rates_usd.csv

hmrc_exchange_rates_eur.csv: hmrc_exchange_rates.csv
	awk -F ',' "NR == 1 || \$$6 == \"EUR\" {print}" hmrc_exchange_rates.csv > hmrc_exchange_rates_eur.csv

clean:
	rm -f hmrc-exchange-rates hmrc_exchange_rates.csv hmrc_exchange_rates_usd.csv hmrc_exchange_rates_eur.csv

.PHONY: clean
