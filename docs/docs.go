// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/task": {
            "post": {
                "description": "create task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create Task",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.POSTTaskResponse"
                        }
                    }
                }
            }
        },
        "/task/{id}": {
            "get": {
                "description": "get task by id",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "task id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GETTaskResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete task by id",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete Task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "task id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.DELETETaskResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "update task by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update Task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "task id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PATCHTaskResponse"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "get all task by param",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Tasks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "used for filter by status",
                        "name": "status_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "used for filter from date",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "used for filter to date",
                        "name": "to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GETAllTaskResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.DELETETaskResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.GETAllTaskResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TaskResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.GETTaskResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.TaskResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.PATCHTaskResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.TaskResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.POSTTaskResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.TaskResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.TaskResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "expected_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/model.TaskStatusResponse"
                },
                "status_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.TaskStatusResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status_name": {
                    "type": "string"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
