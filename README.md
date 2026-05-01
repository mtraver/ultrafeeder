# Unmarshal ultrafeeder's/readsb's `aircraft.json` data

[![GoDoc](https://godoc.org/github.com/mtraver/ultrafeeder?status.svg)](https://godoc.org/github.com/mtraver/ultrafeeder)

[readsb](https://github.com/wiedehopf/readsb) (via tar1090) and the [SDR Enthusiasts ultrafeeder container](https://github.com/sdr-enthusiasts/docker-adsb-ultrafeeder) generate data about the currently known aircraft and make it available at `/data/aircraft.json`. This package contains Go types representing `aircraft.json`.

Documentation on the format: https://github.com/wiedehopf/readsb/blob/dev/README-json.md
