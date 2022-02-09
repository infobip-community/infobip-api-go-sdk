#!/bin/bash

go test --cover $(go list ./... | grep -v infobip-go-client/examples)
