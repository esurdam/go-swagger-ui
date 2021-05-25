package swaggerui_test

import (
	swaggerui "github.com/esurdam/go-swagger-ui"
	"net/http/httptest"
	"testing"
)

var DefaultAssetFn = func(s string) ([]byte, error) {
	return []byte("test"), nil
}

func TestNewServeMux(t *testing.T) {
	type args struct {
		assetFn  swaggerui.AssetFn
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestNewServeMux",
			args: args{
				assetFn: DefaultAssetFn,
				filename: "swagger.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux := swaggerui.NewServeMux(tt.args.assetFn, tt.args.filename)
			if mux == nil {
				t.Error("NewServeMuxWithRoot() got nil")
				return
			}
			req := httptest.NewRequest("GET", "/swagger-ui/?url="+tt.args.filename, nil)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, req)
			if w.Code != 200 {
				t.Errorf("NewServeMuxWithRoot() GET swagger-ui failed, expected 200, got %d", w.Code)
			}

			// Test getting the json
			req = httptest.NewRequest("GET", "/"+tt.args.filename, nil)
			w = httptest.NewRecorder()
			mux.ServeHTTP(w, req)
			if w.Code != 200 {
				t.Errorf("NewServeMuxWithRoot() GET swagger-ui failed, expected 200, got %d", w.Code)
			}
			if res := w.Body.String(); res != "test" {
				t.Errorf("NewServeMuxWithRoot() GET json failed, expected 'test', got %s",res)
			}
		})
	}
}

func TestNewServeMuxWithRoot(t *testing.T) {
	type args struct {
		assetFn  swaggerui.AssetFn
		filename string
		root     string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestNewServeMuxWithRoot",
			args: args{
				assetFn:  DefaultAssetFn,
				filename: "swagger.json",
				root:     "/v1/auth",
			},
		},
		{
			name: "TestNewServeMuxWithRoot2",
			args: args{
				assetFn:  DefaultAssetFn,
				filename: "api.swagger.json",
				root:     "/v1/api",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux := swaggerui.NewServeMuxWithRoot(tt.args.assetFn, tt.args.filename, tt.args.root)
			if mux == nil {
				t.Error("NewServeMuxWithRoot() got nil")
				return
			}
			req := httptest.NewRequest("GET", tt.args.root+"/swagger-ui/?url="+tt.args.filename, nil)
			w := httptest.NewRecorder()
			mux.ServeHTTP(w, req)
			if w.Code != 200 {
				t.Errorf("NewServeMuxWithRoot() GET swagger-ui failed, expected 200, got %d", w.Code)
			}

			// Test getting the json
			req = httptest.NewRequest("GET", tt.args.root+"/"+tt.args.filename, nil)
			w = httptest.NewRecorder()
			mux.ServeHTTP(w, req)
			if w.Code != 200 {
				t.Errorf("NewServeMuxWithRoot() GET json failed, expected 200, got %d", w.Code)
			}
			if res := w.Body.String(); res != "test" {
				t.Errorf("NewServeMuxWithRoot() GET json failed, expected 'test', got %s",res)
			}
		})
	}
}
