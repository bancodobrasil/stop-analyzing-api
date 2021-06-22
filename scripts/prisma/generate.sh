#!/bin/bash

echo "Running migration tool"

export DATABASE_URL=postgresql://user2020:pass2020@localhost:5432/stop-analyzing-api
go get github.com/prisma/prisma-client-go
go run github.com/prisma/prisma-client-go db push
go run github.com/prisma/prisma-client-go generate