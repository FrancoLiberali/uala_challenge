install_dependencies:
	go install github.com/vektra/mockery/v2@v2.40.1

lint:
	golangci-lint run
	cd app && golangci-lint run --config ../.golangci.yml
	cd test_integration && golangci-lint run --config ../.golangci.yml
	cd test_e2e && golangci-lint run --config ../.golangci.yml
	cd aws_lambda && golangci-lint run --config ../.golangci.yml

test_unit:
	go test -v ./app/...

test_integration:
	docker compose -f "docker/docker-compose.yml" up cache -d --wait --wait-timeout 30
	go test -v ./test_integration

test_e2e:
	docker compose -f "docker/docker-compose.yml" up -d --build --force-recreate --renew-anon-volumes
	go test -v -count=1 ./test_e2e

generate:
	cd app && go generate ./...

aws_build:
	cd aws_lambda && GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap follow/main.go && zip uala-challenge-follow.zip bootstrap
	cd aws_lambda && GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap tweet/main.go && zip uala-challenge-tweet.zip bootstrap
	cd aws_lambda && GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap timeline/main.go && zip uala-challenge-timeline.zip bootstrap

.PHONY: test_integration test_e2e

