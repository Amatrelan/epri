# epri

This is simple tool to fetch Nordpool price data.

I started this to learn go and wanted to make small simple product so I can see end.

## run

There isn't anything fancy in this.
`go run .`

This will generate config file in `$XDG_CONFIG_HOME/epri/config.toml` or
`~/epri/config.toml` if you don't have `$XDG_CONFIG_HOME`.

This behaviour can be overridden with env variable `EPRI_CONFIG` (this is needed for for non *nix OS)
