# `hmrc-exchange-rates`

HMRC publishes monthly exchange rates in XML format [[Source](http://www.hmrc.gov.uk/softwaredevelopers/2023-exrates.html)]. Unfortunately, the format in which those exchange rates are published is rather unwieldy.

`hmrc-exchange-rates` downloads exchange rates from HMRC and produces a single, easy-to-use CSV file, containing monthly exchange rates for all major currencies between February 2015 and the most recent month for which HMRC has published rates.

**How is this different from the CSV files published by HMRC?**

HMRC publishes **one CSV file per month**, containing exchange rates against all major currencies [[Example](https://assets.publishing.service.gov.uk/government/uploads/system/uploads/attachment_data/file/1120505/exrates-monthly-1222.csv/preview)]. If you need to look up monthly exchange rates over three years, you'd need to download and normalize 12 * 3 = 36 CSV files. This is annoying.

This script produces a **single CSV** containing monthly exchange rates against **all** major currencies, starting February 2015.

The resulting CSV looks something like this:

```CSV
period_start,period_end,country_name,country_code,currency_name,currency_code,rate_new
2015-02-01,2015-02-28,Argentina,AR,Peso ,ARS,13.02
2015-02-01,2015-02-28,Australia,AU,Dollar ,AUD,1.8385
2015-02-01,2015-02-28,Austria,AT,Euro ,EUR,1.3008
2015-02-01,2015-02-28,Belgium,BE,Euro ,EUR,1.3008
2015-02-01,2015-02-28,Brazil,BR,Real ,BRL,3.9326
2015-02-01,2015-02-28,Canada,CA,Dollar,CAD,1.8251
2015-02-01,2015-02-28,Cyprus,CY,Euro ,EUR,1.3008
2015-02-01,2015-02-28,Denmark,DK,Krone ,DKK,9.664
2015-02-01,2015-02-28,Estonia,EE,Euro,EUR,1.3008
2015-02-01,2015-02-28,Finland,FI,Euro ,EUR,1.3008
2015-02-01,2015-02-28,France,FR,Euro ,EUR,1.3008
2015-02-01,2015-02-28,Germany,DE,Euro ,EUR,1.3008
[...]
2023-05-01,2023-05-31,Vanuatu,VU,Vatu ,VUV,148.3112
2023-05-01,2023-05-31,Venezuela,VE,Bolivar Fuerte ,VEF,319866.39
2023-05-01,2023-05-31,Vietnam,VN,Dong ,VND,29259.5291
2023-05-01,2023-05-31,Wallis & Futuna Islands,WF,CFP Franc ,XPF,135.5738
2023-05-01,2023-05-31,Western Samoa,WS,Tala ,WST,3.3924
2023-05-01,2023-05-31,Yemen (Rep of),YE,Rial ,YER,311.4628
2023-05-01,2023-05-31,Zambia,ZM,Kwacha ,ZMW,21.5483
2023-05-01,2023-05-31,Zimbabwe,ZW,Dollar ,ZWL,450.5671
```

## Usage

Run `make all`, which will
1. build the binary,
2. download monthly exchange rates between February 2015 and the most recent month for which HMRC has published rates,
3. generate `hmrc_exchange_rates.csv`, `hmrc_exchange_rates_usd.csv`, and `hmrc_exchange_rates_eur.csv`.

```sh
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
