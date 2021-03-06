// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/xinliangnote/go-gin-api/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/get": {
            "post": {
                "description": "获取授权信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Demo"
                ],
                "summary": "获取授权信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/demo_handler.authResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/demo/trace": {
            "get": {
                "description": "Trace 示例",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Demo"
                ],
                "summary": "Trace 示例",
                "parameters": [
                    {
                        "type": "string",
                        "description": "签名",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object",
                                "properties": {
                                    "job": {
                                        "description": "工作",
                                        "type": "string"
                                    },
                                    "name": {
                                        "description": "用户名",
                                        "type": "string"
                                    }
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "description": "创建用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "请求信息",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_handler.createRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "签名",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_handler.createResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/user/delete/{id}": {
            "patch": {
                "description": "删除用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "签名",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_handler.deleteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/user/info/{username}": {
            "get": {
                "description": "用户详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "签名",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_handler.detailResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "put": {
                "description": "编辑用户 - 通过用户主键ID更新用户昵称",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "编辑用户 - 通过用户主键ID更新用户昵称",
                "parameters": [
                    {
                        "description": "请求信息",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_handler.updateNickNameByIDRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "签名",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_handler.updateNickNameByIDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/code.Failure"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "code.Failure": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务码",
                    "type": "integer"
                },
                "message": {
                    "description": "描述信息",
                    "type": "string"
                }
            }
        },
        "demo_handler.authResponse": {
            "type": "object",
            "properties": {
                "authorization": {
                    "description": "签名",
                    "type": "string"
                },
                "expire_time": {
                    "description": "过期时间",
                    "type": "integer"
                }
            }
        },
        "user_handler.createRequest": {
            "type": "object",
            "properties": {
                "mobile": {
                    "description": "手机号",
                    "type": "string"
                },
                "nick_name": {
                    "description": "昵称",
                    "type": "string"
                },
                "user_name": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "user_handler.createResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                }
            }
        },
        "user_handler.deleteResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "用户主键ID",
                    "type": "integer"
                }
            }
        },
        "user_handler.detailResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "用户主键ID",
                    "type": "integer"
                },
                "mobile": {
                    "description": "手机号（脱敏）",
                    "type": "string"
                },
                "nick_name": {
                    "description": "昵称",
                    "type": "string"
                },
                "user_name": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "user_handler.updateNickNameByIDRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "用户主键ID",
                    "type": "integer"
                },
                "nick_name": {
                    "description": "昵称",
                    "type": "string"
                }
            }
        },
        "user_handler.updateNickNameByIDResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "用户主键ID",
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "2.0",
	Host:        "127.0.0.1:9999",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "swagger 接口文档",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}