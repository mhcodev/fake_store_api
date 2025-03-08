// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/data": {
            "get": {
                "description": "Get data from token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get token data from token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JSONReponseOne"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "get a data of a user logged",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Log in as a user using username and password",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JSONReponseOne"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "Get data from token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Generate a new access token for user",
                "parameters": [
                    {
                        "description": "refresh token credential",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.NewTokenInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JSONReponseOne"
                        }
                    }
                }
            }
        },
        "/category": {
            "get": {
                "description": "Fetch category by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Fetch category by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JSONReponseOne"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Create a category",
                "parameters": [
                    {
                        "description": "Category body request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JSONReponseOne"
                        }
                    }
                }
            }
        },
        "/category/{id}": {
            "put": {
                "description": "Update a category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Update a category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category body request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JSONReponseOne"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Delete a category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JSONReponseOne"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.CreateCategoryRequest": {
            "type": "object",
            "properties": {
                "imageURL": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateCategoryRequest": {
            "type": "object",
            "properties": {
                "imageURL": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.JSONReponseMany": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "count": {
                    "type": "integer"
                },
                "data": {},
                "errors": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.JSONReponseOne": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "success": {
                    "type": "boolean"
                }
            }
        },
        "services.LoginInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "services.NewTokenInput": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Fake Store API",
	Description:      "This is a sample server for a fake store API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
