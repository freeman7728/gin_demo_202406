// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/employee/delete": {
            "post": {
                "description": "验证账号密码，然后删除用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "description": "Request",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除成功或失败",
                        "schema": {
                            "type": "msg"
                        }
                    }
                }
            }
        },
        "/employee/insert": {
            "post": {
                "description": "批量导入用户，并且返回成功失败列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "批量导入用户",
                "parameters": [
                    {
                        "description": "Request",
                        "name": "list",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeInsertResult"
                        }
                    }
                }
            }
        },
        "/employee/login": {
            "post": {
                "description": "验证账号密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "Request",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功或失败",
                        "schema": {
                            "type": "msg"
                        }
                    }
                }
            }
        },
        "/employee/update": {
            "post": {
                "description": "验证账号密码，然后把密码改为新密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "用户修改信息",
                "parameters": [
                    {
                        "description": "Request",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "修改成功或失败",
                        "schema": {
                            "type": "msg"
                        }
                    }
                }
            }
        },
        "/producer/delete": {
            "post": {
                "description": "删除供应商，返回外键检查结果和删除结果，如果有外键依赖则会提示",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "请求所有供应商",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProducerList"
                        }
                    }
                }
            }
        },
        "/producer/insert": {
            "post": {
                "description": "批量导入供应商，并且返回成功失败列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "批量导入供应商",
                "parameters": [
                    {
                        "description": "Producer Request",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ProducerList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProducerInsertResponseDto"
                        }
                    }
                }
            }
        },
        "/producer/update": {
            "post": {
                "description": "更新供应商，除了id之外都能改，所有不会有外键冲突",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "更新供应商",
                "parameters": [
                    {
                        "description": "Producer Request",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Producer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "msg",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.EmployeeInsertResult": {
            "type": "object",
            "properties": {
                "fail_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Employee"
                    }
                },
                "success_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Employee"
                    }
                }
            }
        },
        "dto.EmployeeList": {
            "type": "object",
            "properties": {
                "employee_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Employee"
                    }
                }
            }
        },
        "dto.EmployeeUpdateRequest": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "level": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "salary": {
                    "type": "number"
                },
                "tel": {
                    "type": "string"
                }
            }
        },
        "dto.Producer": {
            "type": "object",
            "properties": {
                "producer": {
                    "$ref": "#/definitions/models.Producer"
                }
            }
        },
        "dto.ProducerInsertResponseDto": {
            "type": "object",
            "properties": {
                "fail_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Producer"
                    }
                },
                "success_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Producer"
                    }
                }
            }
        },
        "dto.ProducerList": {
            "type": "object",
            "properties": {
                "producer_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Producer"
                    }
                }
            }
        },
        "models.Employee": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "level": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "salary": {
                    "type": "number"
                },
                "tel": {
                    "type": "string"
                }
            }
        },
        "models.Producer": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "contact_name": {
                    "type": "string"
                },
                "contact_tel": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "short_name": {
                    "type": "string"
                },
                "tel": {
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "超市进销存系统--数据库课设接口文档",
	Description:      "没有描述",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
