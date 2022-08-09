module github.com/blazejsewera/notipie/producer

go 1.19

require (
	github.com/blazejsewera/notipie/core v0.0.0-unpublished
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/blazejsewera/notipie/core v0.0.0-unpublished => ../core // link local notipie core

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
)
