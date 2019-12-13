module c_program_language

go 1.13

replace (
	c_program_language/patterns/pool => ./patterns/pool
	c_program_language/patterns/runner => ./patterns/runner
	c_program_language/patterns/work => ./patterns/work
)

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/gorilla/websocket v1.4.1
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20191210023423-ac6580df4449 // indirect
	gopkg.in/go-playground/validator.v9 v9.30.2 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)
