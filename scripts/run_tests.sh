#!/bin/bash

go test --cover $(go list ./... | grep -v infobip-api-go-sdk/v3/examples) #-coverprofile=coverage.out
#go tool cover --html=coverage.out
