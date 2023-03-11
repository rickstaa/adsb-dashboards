# Cron

This folder contains some cronjobs I use with the dashboards.

-   **prom_snapshot.sh**: A cronjob I use to automatically create prometheus snapshots (see <https://prometheus.io/docs/prometheus/latest/querying/api/#snapshot>). This can be replaced using any external storage provider (see <https://prometheus.io/docs/operating/integrations/>).

## How to use

You can checkout out [this excellent guide](https://www.cyberciti.biz/faq/how-do-i-add-jobs-to-cron-under-linux-or-unix-oses/) on how to use cronjobs on Linux. I currently use the following cronjob:

```bash
@monthly /usr/bin/bash $HOME/adsb-dashboards/cron/prom_snapshot.sh
```
