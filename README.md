# epri

This is simple tool to fetch Nordpool price data.

I started this to learn go and wanted to make small simple product so I can see end.

## run

There isn't anything fancy in this.
`go run .`

## config

This will generate config file in `$XDG_CONFIG_HOME/epri/config.toml` or
`~/epri/config.toml` if you don't have `$XDG_CONFIG_HOME`.

This behaviour can be overridden with env variable `EPRI_CONFIG` (this is needed for for non *nix OS)

```toml
currency = 'EUR'
location = 'FI'
subunit = true
tax = 25.5
```

`currency` possible values:

- `EUR`
- `NOK`
- `DKK`
- `SEK`
- `PLN`

`location` possible values:

- `FI`
- `DK1`
- `DK2`
- `NO1`
- `NO2`
- `NO3`
- `NO4`
- `NO5`
- `SE1`
- `SE2`
- `SE3`
- `SE4`

`subunit` possible values:

with this you can manage will you show price in full unit, or in subunit (ex: Euro, or Cent)

- `true`
- `false`

`tax`:
You can add to this tax percentage.
