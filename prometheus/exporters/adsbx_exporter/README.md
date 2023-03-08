# ADS-Bx exporter

This folder contains a simple Golang exporter that fetches ADS-B data from <https://www.adsbexchange.com> and
exposes them as [Prometheus](https://prometheus.io/) metrics on `http://localhost:19100/metrics`. It currently exposes the following metrics:

-   `adsbx_mlat_feeders` - The number of MLAT feeders connected to ADS-B Exchange.

## Development

If you want to contribute, please check the [CONTRIBUTING](../../../CONTRIBUTING.md) Guide in the main folder. Please test your changes using [promtool](https://prometheus.io/docs/prometheus/latest/configuration/unit_testing_rules/) before creating a pull request.
