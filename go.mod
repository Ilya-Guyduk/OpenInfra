module github.com/Ilya-Guyduk/go-openinfra

go 1.22.5

replace github.com/Ilya-Guyduk/go-openinfra/internal/infra => ./internal/infra

replace github.com/Ilya-Guyduk/go-openinfra/pkg/infra => ./pkg/infra

require (
	github.com/stretchr/testify v1.10.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
)
