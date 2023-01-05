FROM golang:1.20rc1-alpine
RUN apk add --no-cache gcc musl-dev bash curl
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.44.0

WORKDIR /app
COPY . .

CMD ["bash", "./scripts/run_tests.sh"]
