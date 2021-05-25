// Package swaggerui provides Handlers for serving the swagger UI.
package swaggerui

import (
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/esurdam/go-swagger-ui/swagger"
)

const defaultPath = "/swagger-ui/"

// AssetFn is a function handler which returns the swagger.json bytes.
// Used to instantiate handler with swagger.json.
type AssetFn func(string) ([]byte, error)

type pathHandler struct {
	root   http.Handler
	prefix string
}

// ServeHTTP handles prefixes for serving the UI.
func (f *pathHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If prefix contains requested path. (Allows for /v1/auth/swagger-ui/ or /v1/auth/swagger-ui
	if strings.HasPrefix(f.prefix, r.URL.Path) && r.URL.Query().Get("url") == "" {
		parts := strings.Split(f.prefix, "/")
		q := r.URL.Query()
		q.Add("url", strings.Join(parts[0:len(parts)-2], "/")+"/swagger.json")
		http.Redirect(w, r, r.URL.String()+"?"+q.Encode(), http.StatusTemporaryRedirect)
		return
	}
	f.root.ServeHTTP(w, r)
}

// AddHandle adds the swagger-ui Handle to the provided mux at the provided root.
func AddHandle(mux *http.ServeMux, root string) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "./third_party/swagger-ui",
	})
	prefix := defaultPath
	if root != "/" {
		prefix = root + defaultPath
		mux.Handle(prefix, &pathHandler{
			root:   http.StripPrefix(prefix, fileServer),
			prefix: prefix,
		})
		return
	}
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

// NewServeMux creates a new http.ServeMux which serves the swagger ui and json with the default root "/".
func NewServeMux(assetFn AssetFn, filename string) *http.ServeMux {
	return NewServeMuxWithRoot(assetFn, filename, "/")
}

// NewServeMuxWithRoot creates a new http.ServeMux which serves the swagger ui and json with the provided root.
// Passing in root = "/v1/auth" will yield a path of "/v1/auth/swagger.json"
func NewServeMuxWithRoot(assetFn AssetFn, filename, root string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(filepath.Join(root, filename), func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		data, err := assetFn(filename)
		if err != nil {
			panic("Unable to load swagger")
		}
		res.Write(data)
	})
	AddHandle(mux, root)
	return mux
}
