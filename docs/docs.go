// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/request": {
            "post": {
                "description": "This endpoint proxies a request.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "proxy"
                ],
                "summary": "Handles the proxy request",
                "parameters": [
                    {
                        "description": "Proxy request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ProxyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ProxyResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ProxyRequest": {
            "type": "object",
            "properties": {
                "headers": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "method": {
                    "type": "string",
                    "default": "GET"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.ProxyResponse": {
            "type": "object",
            "properties": {
                "headers": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "length": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "https://proxy-server-u78h.onrender.com",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
