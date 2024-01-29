module github.com/FrancoLiberali/uala_challenge/test_e2e

go 1.18

require (
	github.com/FrancoLiberali/uala_challenge/app v0.0.1
	github.com/cucumber/godog v0.13.0
	github.com/cucumber/messages/go/v21 v21.0.1
	github.com/elliotchance/pie/v2 v2.8.0
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/cucumber/gherkin/go/v26 v26.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-memdb v1.3.4 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/exp v0.0.0-20220321173239-a90fa8a75705 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/FrancoLiberali/uala_challenge/app => ./../app
