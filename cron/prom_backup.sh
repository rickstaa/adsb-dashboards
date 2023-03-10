#!/bin/bash
# Runs a Prometheus snapshot and writes it to external usb.

# Get input from user.
backup_path=$1
if [[ -z $backup_path ]]; then
    echo "Please enter the path to the backup directory."
    exit 1
fi

# Create snapshot and move to external usb.
content=$(curl -s -XPOST http://localhost:9090/api/v1/admin/tsdb/snapshot)
file_name=$(jq -r '.data.name' <<< "$content")
cp -r /var/lib/docker/volumes/adsb-dashboards_prometheus_data/_data/snapshots/$file_name $backup_path
rm -r /var/lib/docker/volumes/adsb-dashboards_prometheus_data/_data/snapshots/$file_name
