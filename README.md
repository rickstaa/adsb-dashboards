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

The `docker-compose` file spins up four containers:

-   **Portainer**: A container that runs [Portainer](https://www.portainer.io/install) which can be used to manage the other containers (port `9443`).
-   **Prometheus**: A container that runs [Prometheus](https://prometheus.io/) (port `9090`).
-   **json_exporter**: A Container runs a Prometheus [Json Exporter](https://github.com/prometheus-community/json_exporter) (port `7979`).
-   **Grafana**: A container runs [Grafana](https://grafana.com/) (port `3000`).

## Contributing

Contributions of any kind are welcome 💙! Please check out the [contributing guidelines](contributing.md).

[![CC0](https://i.creativecommons.org/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)
