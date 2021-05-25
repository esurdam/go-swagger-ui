go-swagger-ui
============

[![Documentation](https://godoc.org/github.com/esurdam/go-swagger-ui?status.svg)](http://godoc.org/github.com/esurdam/go-swagger-ui)
[![Go Report Card](https://goreportcard.com/badge/github.com/esurdam/go-swagger-ui)](https://goreportcard.com/report/github.com/esurdam/go-swagger-ui)
[![test](https://github.com/esurdam/go-swagger-ui/actions/workflows/go.yml/badge.svg)](https://github.com/esurdam/go-swagger-ui/actions/workflows/go.yml)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/esurdam/go-swagger-ui/blob/main/LICENSE)

This repo provides go handlers for serving `swagger.json` and the [Swagger UI](https://swagger.io/tools/swagger-ui/).

Commonly used with [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) compiled swagger.  

## Usage

```go
import (
	"github.com/esurdam/go-swagger-ui"
)

// Asset represents a AssetFn - compiled bindata swagger file
mux := swaggerui.NewServeMux(Asset, "swagger.json") // add swagger bindata asset

// /swagger.json serves json
// /swagger-ui serves the swagger-ui
```


With custom root:
```go
import (
	"github.com/esurdam/go-swagger-ui"
)

// Asset represents a AssetFn - compiled bindata swagger file
mux := swaggerui.NewServeMuxWithRoot(Asset, "swagger.json", "/v1/auth") // add swagger bindata asset

// v1/auth/swagger.json serves json
// v1/auth/swagger-ui serves the swagger-ui
```

## Updating UI

`swagger` directory contains auto-generated output.

1. Add updated assets to //third_party/swagger-ui
2. Run `make build` which will compile swagger into `swagger/bindata.go`
