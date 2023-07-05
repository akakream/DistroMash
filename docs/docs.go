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
        "/": {
            "get": {
                "description": "Get the home page.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "get the home page",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
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
                    "Crdt"
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
            },
            "post": {
                "description": "Post a CRDT key-value pair.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Crdt"
                ],
                "summary": "post crdt key-value pair",
                "parameters": [
                    {
                        "description": "Post Crdt",
                        "name": "crdt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Crdt"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Crdt"
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
                            "$ref": "#/definitions/models.Crdt"
                        }
                    }
                }
            }
        },
        "/api/v1/image": {
            "post": {
                "description": "Upload a multi-platform docker image to ipfs and get the cid.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "upload a multi-platform docker image to ipfs and get the cid",
                "parameters": [
                    {
                        "description": "Post Image",
                        "name": "crdt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Image"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ImageWithCID"
                        }
                    }
                }
            }
        },
        "/api/v1/strategy": {
            "get": {
                "description": "Get all strategies.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Strategy"
                ],
                "summary": "get all strategies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Strategy"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Post the strategy.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Strategy"
                ],
                "summary": "post the strategy",
                "parameters": [
                    {
                        "description": "Name of Strategy",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Strategy"
                        }
                    }
                }
            }
        },
        "/api/v1/strategy/registered": {
            "get": {
                "description": "Get registered strategies.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Strategy"
                ],
                "summary": "get registered strategies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Strategy"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/strategy/{name}": {
            "get": {
                "description": "Get the strategy.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Strategy"
                ],
                "summary": "get the strategy",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of Strategy",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Strategy"
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
        },
        "models.Image": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "models.ImageWithCID": {
            "type": "object",
            "properties": {
                "cid": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Strategy": {
            "type": "object",
            "properties": {
                "name": {
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
	BasePath:         "",
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
