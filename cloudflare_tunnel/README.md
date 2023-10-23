# Cloudflare tunnel

This folder contains a docker-compose file for creating a [Cloudflare tunnel](https://www.cloudflare.com/en-gb/products/tunnel/). This tunnel is used to safely share the ads-b Grafana dashboard with the outside world. Please change the docker network containing the service you exposed in the Zero-Trusts interface.

## Installation

First, install [Docker](https://docs.docker.com/get-docker/). Then, add the `CLOUDFLARE_TOKEN` of your Cloudflare tunnel to the `.env` file. Lastly, run the `docker compose up -d` command to start the Cloudflare container.
