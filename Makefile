.PHONY: build clean deploy test lint format generate-mock

build:
	GOOS=linux GOARCH=amd64 go build -o bin/haltrds ./cmd/lambda/haltrds/main.go

clean:
	rm -rf ./bin

deploy: clean build
	npm run deploy

remove:
	npm run remove

test:
	go clean -testcache
	go test -p 1 -v $$(go list ./... | grep -v /node_modules/)

lint:
	go vet ./...
	golangci-lint run ./...

format:
	gofmt -l -s -w .
	goimports -w -l ./

generate-mock:
	mockgen -source=infrastructure/rds_client.go -destination mock/rds_client.go -package mock
