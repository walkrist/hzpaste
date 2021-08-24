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
        "contact": {
            "name": "API Support",
            "url": "http://ep4sh.cc",
            "email": "ep4sh2k@gm[a]il.com"
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
        "/gc": {
            "get": {
                "description": "returns count of the removed data",
                "produces": [
                    "application/json"
                ],
                "summary": "triggers obsolete data collection and removes it",
                "responses": {
                    "200": {
                        "description": "gc",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        },
        "/pastes": {
            "get": {
                "description": "get pastes list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "shows all pastes",
                "responses": {
                    "200": {
                        "description": "pastes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/paste.Paste"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "add new paste",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "add a paste",
                "responses": {
                    "200": {
                        "description": "new_paste",
                        "schema": {
                            "$ref": "#/definitions/paste.Paste"
                        }
                    }
                }
            },
            "delete": {
                "description": "returns empty slice for Pastes",
                "produces": [
                    "application/json"
                ],
                "summary": "purges all pastes",
                "responses": {
                    "200": {
                        "description": "pastes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/paste.Paste"
                            }
                        }
                    }
                }
            }
        },
        "/pastes/{id}": {
            "get": {
                "description": "get paste by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "finds a single paste identified by an ID in the request URL.",
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Paste ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "paste",
                        "schema": {
                            "$ref": "#/definitions/paste.Paste"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "provides health checks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "provides health checks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "paste.Paste": {
            "type": "object",
            "required": [
                "body",
                "name"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
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
