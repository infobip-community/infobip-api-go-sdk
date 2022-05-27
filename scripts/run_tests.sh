#!/bin/bash

go test --cover $(go list ./... | grep -v infobip-api-go-sdk/examples) #-coverprofile=coverage.out
#go tool cover --html=coverage.out
