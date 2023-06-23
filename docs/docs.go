// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Ahmet Kerem Aksoy",
            "email": "a.aksoy@tu-berlin.de"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/crdt": {
            "get": {
                "description": "Get all CRDT key-value pairs.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Crdts"
                ],
                "summary": "get all crdt key-value pairs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Crdt"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/crdt/{key}": {
            "get": {
                "description": "Get the CRDT value by key.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Crdt"
                ],
                "summary": "get crdt value by given key",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Key of Value",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Crdt"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Crdt": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "DistroMash API",
	Description:      "DistroMash meshes your Docker Distribution",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
