#!/bin/bash

echo "Running migration tool"

export DATABASE_URL=postgresql://user2020:pass2020@localhost:5432/stop-analyzing-api
go get github.com/prisma/prisma-client-go
go run github.com/prisma/prisma-client-go migrate save --experimental --create-db --name "init"
go run github.com/prisma/prisma-client-go migrate up --experimental
go run github.com/prisma/prisma-client-go generate