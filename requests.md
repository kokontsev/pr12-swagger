# Фактический прогон curl

```text
$ curl -sS -i --max-time 15 -X POST http://127.0.0.1:18012/api/v1/notes -H Content-Type: application/json -d {"title": "Первая заметка", "content": "Go REST"}
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Location: /api/v1/notes/1
Date: Sat, 31 Oct 2026 08:44:22 GMT
Content-Length: 151

{"id":1,"title":"Первая заметка","content":"Go REST","createdAt":"2026-10-31T08:44:22.486869Z","updatedAt":"2026-10-31T08:44:22.486869Z"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18012/api/v1/notes/1
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:22 GMT
Content-Length: 151

{"id":1,"title":"Первая заметка","content":"Go REST","createdAt":"2026-10-31T08:44:22.486869Z","updatedAt":"2026-10-31T08:44:22.486869Z"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18012/api/v1/notes?q=Go
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:22 GMT
Content-Length: 3

[]


$ curl -sS -i --max-time 15 -X PATCH http://127.0.0.1:18012/api/v1/notes/1 -H Content-Type: application/json -d {"title": "Go заметка"}
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:22 GMT
Content-Length: 141

{"id":1,"title":"Go заметка","content":"Go REST","createdAt":"2026-10-31T08:44:22.486869Z","updatedAt":"2026-10-31T08:44:22.520404Z"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18012/api/v1/notes?q=Go&limit=1&page=1
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:22 GMT
Content-Length: 143

[{"id":1,"title":"Go заметка","content":"Go REST","createdAt":"2026-10-31T08:44:22.486869Z","updatedAt":"2026-10-31T08:44:22.520404Z"}]


$ curl -sS -i --max-time 15 -X POST http://127.0.0.1:18012/api/v1/notes -H Content-Type: application/json -d {"title": " "}
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:22 GMT
Content-Length: 26

{"error":"invalid_input"}


$ curl -sS -i --max-time 15 -X DELETE http://127.0.0.1:18012/api/v1/notes/1
HTTP/1.1 204 No Content
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:22 GMT



$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18012/api/v1/notes/1
HTTP/1.1 404 Not Found
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:22 GMT
Content-Length: 22

{"error":"not_found"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18012/docs/index.html
HTTP/1.1 200 OK
Content-Type: text/html; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:22 GMT
Transfer-Encoding: chunked


<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Swagger UI</title>
  <link rel="stylesheet" type="text/css" href="./swagger-ui.css" >
  <link rel="icon" type="image/png" href="./favicon-32x32.png" sizes="32x32" />
  <link rel="icon" type="image/png" href="./favicon-16x16.png" sizes="16x16" />
  <style>
    html
    {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
    }
    *,
    *:before,
    *:after
    {
        box-sizing: inherit;
    }

    body {
      margin:0;
      background: #fafafa;
    }
  </style>
</head>

<body>

<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" style="position:absolute;width:0;height:0">
  <defs>
    <symbol viewBox="0 0 20 20" id="unlocked">
      <path d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1 7.334 1 6 2.703 6 5.6V6h2v-.801C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387C5.672 18.861 6.55 19 7.1 19h5.8c.549 0 1.428-.139 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8z"></path>
    </symbol>

    <symbol viewBox="0 0 20 20" id="locked">
      <path d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1 7.334 1 6 2.703 6 5.6V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387C5.672 18.861 6.55 19 7.1 19h5.8c.549 0 1.428-.139 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8zM12 8H8V5.199C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="close">
      <path d="M14.348 14.849c-.469.469-1.229.469-1.697 0L10 11.819l-2.651 3.029c-.469.469-1.229.469-1.697 0-.469-.469-.469-1.229 0-1.697l2.758-3.15-2.759-3.152c-.469-.469-.469-1.228 0-1.697.469-.469 1.228-.469 1.697 0L10 8.183l2.651-3.031c.469-.469 1.228-.469 1.697 0 .469.469.469 1.229 0 1.697l-2.758 3.152 2.758 3.15c.469.469.469 1.229 0 1.698z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="large-arrow">
      <path d="M13.25 10L6.109 2.58c-.268-.27-.268-.707 0-.979.268-.27.701-.27.969 0l7.83 7.908c.268.271.268.709 0 .979l-7.83 7.908c-.268.271-.701.27-.969 0-.268-.269-.268-.707 0-.979L13.25 10z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="large-arrow-down">
      <path d="M17.418 6.109c.272-.268.709-.268.979 0s.271.701 0 .969l-7.908 7.83c-.27.268-.707.268-.979 0l-7.908-7.83c-.27-.268-.27-.701 0-.969.271-.268.709-.268.979 0L10 13.25l7.418-7.141z"/>
    </symbol>

    <symbol viewBox="0 0 24 24" id="jump-to">
      <path d="M19 7v4H5.83l3.58-3.59L8 6l-6 6 6 6 1.41-1.41L5.83 13H21V7z"/>
    </symbol>

    <symbol viewBox="0 0 24 24" id="expand">
      <path d="M10 18h4v-2h-4v2zM3 6v2h18V6H3zm3 7h12v-2H6v2z"/>
    </symbol>
  </defs>
</svg>

<div id="swagger-ui"></div>

<script src="./swagger-ui-bundle.js"> </script>
<script src="./swagger-ui-standalone-preset.js"> </script>
<script>
window.onload = function() {
  
  const ui = SwaggerUIBundle({
    url: "doc.json",
    deepLinking:  true ,
    docExpansion: "list",
    dom_id: "#swagger-ui",
    persistAuthorization:  false ,
    validatorUrl: null,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  })

  window.ui = ui
}
</script>
</body>

</html>


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18012/docs/doc.json
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:22 GMT
Transfer-Encoding: chunked

{
    "schemes": [],
    "swagger": "2.0",
    "info": {
        "description": "REST API for notes",
        "title": "Notes API",
        "contact": {},
        "version": "1.0"
    },
    "host": "",
    "basePath": "/api/v1",
    "paths": {
        "/notes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notes"
                ],
                "summary": "List notes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title search",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "default": 20,
                        "description": "Page size",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/core.Note"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notes"
                ],
                "summary": "Create note",
                "parameters": [
                    {
                        "description": "Note",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.Create"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/core.Note"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/notes/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notes"
                ],
                "summary": "Get note",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Note"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "notes"
                ],
                "summary": "Delete note",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notes"
                ],
                "summary": "Patch note",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Changed fields",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.Patch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Note"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "core.Create": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "core.Note": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "core.Patch": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}
```
