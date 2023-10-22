# ADS-B-dashboards

This repository contains several useful ADS-B Grafana dashboards.

## Content

-   [aggregators_dashboard](grafana/dashboards/aggregators_dashboard.json): A dashboard that shows exciting statistics about several ADS-B aggregators.

## Instructions

This repository contains a [docker-compose](docker-compose.yml) file that can run the dashboards locally. You can use the following steps to do this:

1.  Clone the repository.
2.  Install [Docker](https://www.docker.com/).
3.  Run `docker compose up -d` inside the repository folder;.
4.  Checkout the dashboards at `localhost:3000`.

## Docker-compose components

The `docker-compose` file orchestrates three containers, each serving a distinct purpose:

-   **json_exporter**: Here resides a container running a Prometheus [Json Exporter](https://github.com/prometheus-community/json_exporter) accessible on port `7979`.
-   **adsbx_exporter**: This container is dedicated to executing a [custom ADSBx exporter](prometheus/exporters/adsbx_exporter/main.go) and exposes its metrics on port `19100`.
-   **Grafana**: Hosting [Grafana](https://grafana.com/), this container makes its services available on port `3000`.

Moreover, within this repository, a [Portainer container](portainer/README.md) awaits, ready to facilitate the management of the other containers.

## Contributing

Contributions of any kind are welcome 💙! Please check out the [contributing guidelines](contributing.md).

[![CC0](https://i.creativecommons.org/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)
