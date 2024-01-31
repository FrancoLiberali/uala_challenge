module github.com/FrancoLiberali/uala_challenge/aws_lambda

go 1.18

require (
	github.com/FrancoLiberali/uala_challenge/app v0.0.1
	github.com/aws/aws-lambda-go v1.45.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/redis/go-redis/v9 v9.4.0 // indirect
)

replace github.com/FrancoLiberali/uala_challenge/app => ./../app
