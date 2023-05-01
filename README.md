# `hmrc-exchange-rates`

HMRC publishes monthly exchange rates in XML format [[Source](http://www.hmrc.gov.uk/softwaredevelopers/2023-exrates.html)]. Unfortunately, the format in which those exchange rates are published is rather unwieldy.

`hmrc-exchange-rates` downloads exchange rates from HMRC and produces a simple CSV file, containing monthly exchange rates for all major currencies between February 2015 and the most recent month for which HMRC has published rates.

## Usage

Run `make all`, which will
1. build the binary,
2. download monthly exchange rates between February 2015 and the most recent month for which HMRC has published rates,
3. generate `hmrc_exchange_rates.csv`, `hmrc_exchange_rates_usd.csv`, and `hmrc_exchange_rates_eur.csv`.

```
alexandergugel@192 hmrc-exchange-rates % make all
go build
./hmrc-exchange-rates > hmrc_exchange_rates.csv
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0215.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0315.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0415.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0515.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0615.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0715.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0815.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0915.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-1015.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-1115.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-1215.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0116.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0216.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0316.XML...
2023/05/01 13:59:06 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0416.XML...
2023/05/01 13:59:07 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0516.XML...
2023/05/01 13:59:07 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0616.XML...
2023/05/01 13:59:07 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0716.XML...
2023/05/01 13:59:07 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0816.XML...
2023/05/01 13:59:07 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0916.XML...
2023/05/01 13:59:07 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-1016.XML...
2023/05/01 13:59:07 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-1116.XML...
[...]
2023/05/01 13:59:13 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0423.XML...
2023/05/01 13:59:13 fetching http://www.hmrc.gov.uk/softwaredevelopers/rates/exrates-monthly-0523.XML...
awk -F ',' "NR == 1 || \$6 == \"USD\" {print}" hmrc_exchange_rates.csv > hmrc_exchange_rates_usd.csv
awk -F ',' "NR == 1 || \$6 == \"EUR\" {print}" hmrc_exchange_rates.csv > hmrc_exchange_rates_eur.csv
alexandergugel@192 hmrc-exchange-rates %
```

- `hmrc_exchange_rates.csv` contains monthly exchange rates for **all** major currencies.
- `hmrc_exchange_rates_{usd,eur}.csv` are filtered versions of `hmrc_exchange_rates.csv`, only containing USD / EUR exchange rates respectively.

The program requires a somewhat recent version of Go. I've tested things with go1.19.2, but older versions should work as well.

## License

This program is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
