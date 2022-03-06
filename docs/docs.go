// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/movies": {
            "get": {
                "description": "List an array of movies containing the name, opening crawl and comment count\"",
                "produces": [
                    "application/json"
                ],
                "summary": "Route Gets all movies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Movie"
                            }
                        }
                    }
                }
            }
        },
        "/movies/:movieID/comments/": {
            "get": {
                "description": "Endpoint Gets a list of comments for a particular movieID",
                "produces": [
                    "application/json"
                ],
                "summary": "Endpoint Gets a list of comments",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "movie_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Comment"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a new comment to a post with  movieID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Adds a new comment to a post",
                "parameters": [
                    {
                        "description": "Comment",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Comment"
                            }
                        }
                    },
                    {
                        "type": "integer",
                        "description": "MovieId",
                        "name": "movie_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/movies/{movieID}/characters/": {
            "get": {
                "description": "accept sort parameters to sort by one of name, gender or height in ascending or descending order.\"",
                "produces": [
                    "application/json"
                ],
                "summary": "Get characters",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "movie_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Sort by height or name or gender",
                        "name": "sortBy",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "descending or ascending order",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "can be filtered by male or female options",
                        "name": "filterBy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Character"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Character": {
            "type": "object",
            "properties": {
                "birth_year": {
                    "type": "string"
                },
                "eye_color": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "hair_color": {
                    "type": "string"
                },
                "height": {
                    "type": "string"
                },
                "mass": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "skin_color": {
                    "type": "string"
                }
            }
        },
        "domain.Comment": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "movie_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.Movie": {
            "type": "object",
            "properties": {
                "comment_count": {
                    "type": "integer"
                },
                "episode_id": {
                    "type": "integer"
                },
                "opening_crawl": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:8090",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Movie-API Service",
	Description:      "Repo can be found here:https://github.com/Tambarie/movie-api//",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}