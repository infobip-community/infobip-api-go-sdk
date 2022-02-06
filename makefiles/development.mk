IMAGE=client_tests:local

.PHONY: fmt
fmt: ##@development Runs formatter.
fmt:
	./scripts/gofmt_check.sh

.PHONY: install-lint
install-lint: ##@development Installs linters.
install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.44.0

.PHONY: lint
lint: ##@development Runs linters.
lint:
	golangci-lint run

.PHONY: test
test: ##@development Runs tests.
test:
	go test ./...

.PHONY: all-checks
all-checks: ##@development Runs all checks and tests, as ran by the CI.
all-checks: fmt lint test

.PHONY: build
build: ##@development Build containers. Needed only after dependency changes.
build:
	docker build --progress=plain  -t ${IMAGE} .

.PHONY: docker-fmt
docker-fmt: ##@development Runs formatter inside Docker.
docker-fmt: build
	docker run --rm ${IMAGE} ./scripts/gofmt_check.sh

.PHONY: docker-lint
docker-lint: ##@development Runs linters inside Docker.
docker-lint: build
docker-lint:
	docker run --rm ${IMAGE} golangci-lint run

.PHONY: docker-test
docker-test: ##@development Runs tests inside Docker.
docker-test: build
docker-test:
	docker run --rm ${IMAGE}

.PHONY: docker-all-checks
docker-all-checks: ##@development Runs all checks and tests inside Docker, as ran by the CI.
docker-all-checks: build docker-fmt docker-lint docker-test
