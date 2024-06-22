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
        "/detail/delete": {
            "post": {
                "description": "删除订单详情，返回外键检查结果和删除结果，如果有外键依赖则会提示",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "detail"
                ],
                "summary": "删除订单详情",
                "parameters": [
                    {
                        "description": "Detail Request",
                        "name": "detail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Detail"
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
        },
        "/detail/insert": {
            "post": {
                "description": "批量导入订单详情，并且返回成功失败列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "detail"
                ],
                "summary": "批量导入订单详情",
                "parameters": [
                    {
                        "description": "Detail Request",
                        "name": "list",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DetailList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.DetailInsertResponse"
                        }
                    }
                }
            }
        },
        "/detail/select": {
            "post": {
                "description": "请求所有订单详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "detail"
                ],
                "summary": "请求所有订单详情",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.DetailList"
                        }
                    }
                }
            }
        },
        "/detail/update": {
            "post": {
                "description": "更新产品，除了id之外都能改，所有不会有外键冲突",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "detail"
                ],
                "summary": "更新订单详情",
                "parameters": [
                    {
                        "description": "Detail Request",
                        "name": "detail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Detail"
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
        },
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeLoginResponse"
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
        "/order/delete": {
            "post": {
                "description": "删除订单，返回外键检查结果和删除结果，如果有外键依赖则会提示",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "删除订单",
                "parameters": [
                    {
                        "description": "Order Request",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Order"
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
        },
        "/order/insert": {
            "post": {
                "description": "批量导入订单，并且返回成功失败列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "批量导入订单",
                "parameters": [
                    {
                        "description": "Order Request",
                        "name": "list",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OrderList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.OrderInsertResponse"
                        }
                    }
                }
            }
        },
        "/order/select": {
            "post": {
                "description": "请求所有订单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "请求所有订单",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.OrderList"
                        }
                    }
                }
            }
        },
        "/order/update": {
            "post": {
                "description": "更新产品，除了id之外都能改，所有不会有外键冲突",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "更新订单",
                "parameters": [
                    {
                        "description": "Order Request",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Order"
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
                    "producer"
                ],
                "summary": "删除供应商",
                "parameters": [
                    {
                        "description": "Producer Request",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Producer"
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
                    "producer"
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
        "/producer/select": {
            "post": {
                "description": "删除供应商，返回外键检查结果和删除结果，如果有外键依赖则会提示",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producer"
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
                    "producer"
                ],
                "summary": "更新供应商",
                "parameters": [
                    {
                        "description": "Producer Request",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Producer"
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
        },
        "/product/delete": {
            "post": {
                "description": "删除产品，返回外键检查结果和删除结果，如果有外键依赖则会提示",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "删除产品",
                "parameters": [
                    {
                        "description": "Producer Request",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
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
        },
        "/product/insert": {
            "post": {
                "description": "批量导入产品，并且返回成功失败列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "批量导入产品",
                "parameters": [
                    {
                        "description": "Producer Request",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ProductList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProductInsertResponse"
                        }
                    }
                }
            }
        },
        "/product/select": {
            "post": {
                "description": "删除产品，返回外键检查结果和删除结果，如果有外键依赖则会提示",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "请求所有产品",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProductList"
                        }
                    }
                }
            }
        },
        "/product/update": {
            "post": {
                "description": "更新产品，除了id之外都能改，所有不会有外键冲突",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "更新产品",
                "parameters": [
                    {
                        "description": "Producer Request",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
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
        "dto.DetailInsertResponse": {
            "type": "object",
            "properties": {
                "fail_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Detail"
                    }
                },
                "success_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Detail"
                    }
                }
            }
        },
        "dto.DetailList": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Detail"
                    }
                }
            }
        },
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
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Employee"
                    }
                }
            }
        },
        "dto.EmployeeLoginResponse": {
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
                },
                "token": {
                    "type": "string"
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
        "dto.OrderInsertResponse": {
            "type": "object",
            "properties": {
                "fail_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Order"
                    }
                },
                "success_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Order"
                    }
                }
            }
        },
        "dto.OrderList": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Order"
                    }
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
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Producer"
                    }
                }
            }
        },
        "dto.ProductInsertResponse": {
            "type": "object",
            "properties": {
                "fail_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Product"
                    }
                },
                "success_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Product"
                    }
                }
            }
        },
        "dto.ProductList": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Product"
                    }
                }
            }
        },
        "models.Detail": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "list_id": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "number"
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
        "models.Order": {
            "type": "object",
            "properties": {
                "employee_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "time": {
                    "type": "string"
                },
                "total_price": {
                    "type": "number"
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
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
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
                "price": {
                    "type": "number"
                },
                "producer_id": {
                    "type": "integer"
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
