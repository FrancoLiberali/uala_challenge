module github.com/FrancoLiberali/uala_challenge/test_integration

go 1.18

require (
	github.com/FrancoLiberali/uala_challenge/app v0.0.1
	github.com/redis/go-redis/v9 v9.4.0
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/FrancoLiberali/uala_challenge/app => ./../app
