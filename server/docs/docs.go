// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/public/base/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名, 密码, 验证码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/system.SysUser"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/public/base/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户名, 密码, 验证码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/system.SysUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/system.SysUser"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.LoginRequest": {
            "type": "object",
            "required": [
                "captcha",
                "captcha_id",
                "password",
                "user_name"
            ],
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "captcha_id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "system.SysRole": {
            "type": "object",
            "properties": {
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysRole"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "parent_id": {
                    "description": "父角色ID",
                    "type": "string"
                },
                "role_name": {
                    "description": "角色名",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "system.SysUser": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nick_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "roles": {
                    "description": "关联到角色表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysRole"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample Server pets",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
