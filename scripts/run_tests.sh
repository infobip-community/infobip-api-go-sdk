#!/bin/bash

go test --cover $(go list ./... | grep -v infobip-api-go-sdk/examples)
