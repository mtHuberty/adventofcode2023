#!/bin/bash

set -e

cd ./generate/new-day
go run new-day.go $@

cd ../.. && ./generate-main.sh