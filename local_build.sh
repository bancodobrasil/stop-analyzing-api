#!/bin/bash

# enabling githooks directory to host triggered scripts
git config core.hooksPath githooks
docker-compose up -d postgres pgadmin

echo Waiting for postgreSQL startup...
# wait a few seconds to be sure postgreSQL is running
sleep 10

echo "READY !"

cd ./scripts/prisma
sh generate.sh
cd ../../
go mod download
export CGO_ENABLED=0
export GOOS=linux
go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ./stop-analyzing-api .
