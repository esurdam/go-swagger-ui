package main

import (
	"fmt"
	"net/http"

	swaggerui "github.com/esurdam/go-swagger-ui"
)

var DefaultAssetFn = func(s string) ([]byte, error) {
	return []byte(`{
		"swagger": "2.0",
		"info": {
		  "title": "helloworld api",
		  "version": "1.0",
		  "description": "This is a sample server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key special-key to test the authorization filters.",
		  "contact": {
			"email": "esurdam@gmail.com"
		  },
		  "license": {
			"name": "Apache 2.0",
			"url": "http://www.apache.org/licenses/LICENSE-2.0.html"
		  }
		},
		"tags": [
		  {
			"name": "Greeter"
		  }
		],
		"consumes": [
		  "application/json"
		],
		"produces": [
		  "application/json"
		],
		"paths": {
		  "/v1/greeter": {
			"post": {
			  "summary": "Sends a greeting",
			  "operationId": "Greeter_SayHello",
			  "responses": {
				"200": {
				  "description": "A successful response.",
				  "schema": {
					"$ref": "#/definitions/helloworldHelloReply"
				  }
				},
				"default": {
				  "description": "An unexpected error response.",
				  "schema": {
					"$ref": "#/definitions/rpcStatus"
				  }
				}
			  },
			  "parameters": [
				{
				  "name": "body",
				  "in": "body",
				  "required": true,
				  "schema": {
					"$ref": "#/definitions/helloworldHelloRequest"
				  }
				}
			  ],
			  "tags": [
				"Greeter"
			  ]
			}
		  }
		},
		"definitions": {
		  "helloworldHelloReply": {
			"type": "object",
			"properties": {
			  "message": {
				"type": "string"
			  }
			},
			"title": "The response message containing the greetings"
		  },
		  "helloworldHelloRequest": {
			"type": "object",
			"properties": {
			  "name": {
				"type": "string"
			  }
			},
			"description": "The request message containing the user's name."
		  },
		  "protobufAny": {
			"type": "object",
			"properties": {
			  "@type": {
				"type": "string"
			  }
			},
			"additionalProperties": {}
		  },
		  "rpcStatus": {
			"type": "object",
			"properties": {
			  "code": {
				"type": "integer",
				"format": "int32"
			  },
			  "message": {
				"type": "string"
			  },
			  "details": {
				"type": "array",
				"items": {
				  "$ref": "#/definitions/protobufAny"
				}
			  }
			}
		  }
		}
	  }`), nil
}

func main() {
	mux := swaggerui.NewServeMux(DefaultAssetFn, "swagger.json")
	fmt.Println("* listening on :8080")
	fmt.Println("* visit http://localhost:8080/swagger-ui/")
	fmt.Println("* visit http://localhost:8080/swagger-ui/?url=https://petstore.swagger.io/v2/swagger.json")
	http.ListenAndServe(":8080", mux)
}
