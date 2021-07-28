package gen

//go:generate yarn --cwd ../../frontend graphql-codegen --config config/1_introspect.yml
//go:generate gqlgenc
//go:generate gofmt -r "Config -> ClientConfig" -w .

//go:generate gqlgen
//go:generate gofmt -r "Config -> ServerConfig" -w .
