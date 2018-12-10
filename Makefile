SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=
OS=$(shell uname -s)

export GO111MODULE := on

setup:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
	go mod download
.PHONY: setup

test:
	go test $(TEST_OPTIONS) -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m
.PHONY: test

cover: test
	go tool cover -html=coverage.txt
.PHONY: cover

fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
	jsonnet fmt -i ./grafana/*.jsonnet
.PHONY: fmt

lint:
	./bin/golangci-lint run --enable-all ./...
	promtool check rules prometheus/rules/couchbase.rules.yml
	jsonnet fmt --test ./grafana/*.jsonnet
.PHONY: lint

ci: build test lint
.PHONY: ci

build: grafana
	go build
.PHONY: build

todo:
	@grep \
		--exclude-dir=vendor \
		--exclude-dir=bin \
		--exclude-dir=grafana/grafonnet \
		--exclude=Makefile \
		--text \
		--color \
		-nRo -E ' TODO:.*|SkipNow' .
.PHONY: todo

grafana:
	jsonnet -J grafana grafana/dashboard.jsonnet -o grafana/dashboard.json
.PHONY: grafana

.DEFAULT_GOAL := build
