# ADS-B-dashboards

This repository contains several useful ADS-B Grafana dashboards.

## Content

- [aggregators_dashboard](grafana/dashboards/aggregators_dashboard.json): A dashboard that shows general statistics about several ADS-B aggregators.

## Instructions

This repository contains [docker-compose](docker-compose.yml) files that can be used to run the dashboards. You can use the following steps to do this:

1.  Clone the repository.
2.  Install [Docker](https://www.docker.com/).
3.  Run `docker compose up -d` inside the repository folder.
4.  Checkout the dashboards at `localhost:3000`.

## Docker-compose components

The `docker-compose` file orchestrates ~~three~~two containers, each serving a distinct purpose:

**prometheus**: A [Prometheus](https://prometheus.io/) container for storing tracked data. Accessible on port `9090`. - **json_exporter**: A Prometheus [Json Exporter](https://github.com/prometheus-community/json_exporter) used for collecting API data. Accessible on port `7979`.

- ~~**adsbx_exporter**: A [custom ADSBx exporter](prometheus/exporters/adsbx_exporter/main.go) container dedicated to collecting adsbx API data. Accessible on port `19100`.~~ - No longer works due to recent [adsb-x](https://adsbexchange.com/) api changes.
- **Grafana**: A [Grafana](https://grafana.com/) container used for displaying tracked data. Accessible on port `3000`.

Additionally, this repository contains a [Portainer container](portainer/README.md) that can be used to facilitate the management of the other containers.

## Contributing

Contributions of any kind are welcome 💙! Please check out the [contributing guidelines](contributing.md).

[![CC0](https://i.creativecommons.org/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)
