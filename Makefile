install_dependencies:
	go install github.com/vektra/mockery/v2@v2.40.1

test_unit:
	go test -v ./app/...

test_integration:
	docker compose -f "docker/docker-compose.yml" up cache -d --wait --wait-timeout 30
	go test -v ./test_integration

generate:
	cd app && go generate ./...

.PHONY: test_integration

