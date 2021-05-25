go-swagger-ui
============

This repo provides go handlers for serving the [Swagger UI](https://swagger.io/tools/swagger-ui/) and `swagger.json` file
which is provided during instantiation.

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

1. Add updated assets to //third_party/swagger-ui
2. Run `bash hack/build-ui.sh` which will compile swagger into `swagger/bindata.go`
