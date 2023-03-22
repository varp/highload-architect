#!/usr/bin/env bash

set -eu

for m in /sql/migrations/*.sql; do
  echo "Running migration $m"
  mysql --user=root --password="$MYSQL_ROOT_PASSWORD" < "$m"
done
