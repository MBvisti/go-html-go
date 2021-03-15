#!/bin/sh
echo "running go.sh..."

echo "running go vet for whole project"
go vet ./...

echo "running go fmt for whole project"
go fmt ./...

echo "sourcing env variables"
source ./.env

if [ ! -d "node_modules" ]
then
    echo "Directory node_modules does not exists."
    exit 9999 # die with error code 9999
fi

echo "go run go"
go run cmd/server/main.go
