// Package openapi Code generated by swaggo/swag. DO NOT EDIT
package openapi

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
        "/logs": {
            "post": {
                "security": [
                    {
                        "UsersAuth": []
                    }
                ],
                "tags": [
                    "logs"
                ],
                "summary": "insert logs request",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.InsertLogPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.CommonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.CommonResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.CommonResponse"
                        }
                    }
                }
            }
        },
        "/logs/retreive": {
            "post": {
                "security": [
                    {
                        "UsersAuth": []
                    }
                ],
                "tags": [
                    "logs"
                ],
                "summary": "retreive logs request",
                "parameters": [
                    {
                        "description": "filter",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.RetreiveLogsFilter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.RetreiveLogsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.CommonResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.CommonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Action": {
            "type": "object",
            "required": [
                "type"
            ],
            "properties": {
                "details": {
                    "type": "object",
                    "additionalProperties": true
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "domain.CommonResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "domain.InsertLogPayload": {
            "type": "object",
            "required": [
                "action",
                "message",
                "user"
            ],
            "properties": {
                "action": {
                    "$ref": "#/definitions/domain.Action"
                },
                "message": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/domain.User"
                }
            }
        },
        "domain.Log": {
            "type": "object",
            "properties": {
                "action": {
                    "$ref": "#/definitions/domain.Action"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/domain.User"
                }
            }
        },
        "domain.RetreiveLogsFilter": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "query": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "domain.RetreiveLogsResponse": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "logs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Log"
                    }
                },
                "offset": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                },
                "total_pages": {
                    "type": "integer"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "required": [
                "full_name",
                "id",
                "role"
            ],
            "properties": {
                "details": {
                    "type": "object",
                    "additionalProperties": true
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "UsersAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "LOG-MINDER APPLICATION API DOCUMENTATION",
	Description:      "This API contains the source for the LOG-MINDER  app",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
