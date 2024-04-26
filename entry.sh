#!/bin/bash

cd /usr/src/app

echo "running migrations"

go run ./cmd/migration/migrate.go

echo "ending migrations"

app
