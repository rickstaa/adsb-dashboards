#!/bin/bash
# Creates a prometheus snapshot.
curl -s -XPOST http://localhost:9090/api/v1/admin/tsdb/snapshot
