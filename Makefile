install_dependencies:
	go install github.com/vektra/mockery/v2@v2.40.1

lint:
	golangci-lint run
	cd app && golangci-lint run --config ../.golangci.yml
	cd test_integration && golangci-lint run --config ../.golangci.yml
	cd test_e2e && golangci-lint run --config ../.golangci.yml

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

.PHONY: test_integration test_e2e

