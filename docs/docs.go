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
        "/api/v1/greeting": {
            "get": {
                "description": "hello, world endpoint",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "greeting"
                ],
                "summary": "Greeting",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Greeting"
                        }
                    }
                }
            }
        },
        "/api/v1/oauth/google": {
            "get": {
                "description": "Redirect to this endpoint to login with Google OAuth2. When logging in is completed, the redirection to /register in client will occur.",
                "tags": [
                    "google login"
                ],
                "summary": "Login with Google",
                "responses": {
                    "307": {
                        "description": "Temporary Redirect"
                    }
                }
            }
        },
        "/api/v1/property/:id": {
            "get": {
                "description": "Get property by its id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "property"
                ],
                "summary": "Get property by id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Property"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "description": "Get all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Users"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a user by prasing the body",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    }
                }
            }
        },
        "/api/v1/users/:userId": {
            "get": {
                "description": "Get a user by its id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a user with the given id by parsing the body",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user by id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user by its id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user by id",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apperror.AppError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "name": {
                    "type": "string",
                    "example": "internal-server-error"
                }
            }
        },
        "models.BankName": {
            "type": "string",
            "enum": [
                "KASIKORN BANK",
                "BANGKOK BANK",
                "KRUNG THAI BANK",
                "BANK OF AYUDHYA",
                "CIMB THAI BANK",
                "TMBTHANACHART BANK",
                "SIAM COMMERCIAL BANK",
                "GOVERNMENT SAVINGS BANK"
            ],
            "x-enum-varnames": [
                "KBANK",
                "BBL",
                "KTB",
                "BAY",
                "CIMB",
                "TTB",
                "SCB",
                "GSB"
            ]
        },
        "models.Greeting": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Hello, World"
                }
            }
        },
        "models.Property": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "123/4"
                },
                "alley": {
                    "type": "string",
                    "example": "Pattaya Nua 78"
                },
                "country": {
                    "type": "string",
                    "example": "Thailand"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "example": "Et sequi dolor praes"
                },
                "district": {
                    "type": "string",
                    "example": "Bang Phli"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PropertyImage"
                    }
                },
                "owner_id": {
                    "description": "foreign key",
                    "type": "string",
                    "example": "123e4567-e89b-12d3-a456-426614174000"
                },
                "postal_code": {
                    "type": "string",
                    "example": "69096"
                },
                "project_name": {
                    "type": "string",
                    "example": "Supalai"
                },
                "propertyId": {
                    "type": "string"
                },
                "province": {
                    "type": "string",
                    "example": "Pattaya"
                },
                "renting": {
                    "$ref": "#/definitions/models.RentingProperty"
                },
                "residential_type": {
                    "type": "string",
                    "example": "Condo"
                },
                "selling": {
                    "$ref": "#/definitions/models.SellingProperty"
                },
                "street": {
                    "type": "string",
                    "example": "Pattaya"
                },
                "sub_district": {
                    "type": "string",
                    "example": "Bang Bon"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.PropertyImage": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string",
                    "example": "https://image_url.com/abcd"
                }
            }
        },
        "models.RegisteredType": {
            "type": "string",
            "enum": [
                "EMAIL",
                "GOOGLE"
            ],
            "x-enum-varnames": [
                "EMAIL",
                "GOOGLE"
            ]
        },
        "models.RentingProperty": {
            "type": "object",
            "properties": {
                "is_occupied": {
                    "type": "boolean",
                    "example": true
                },
                "price_per_month": {
                    "type": "number",
                    "example": 12345.67
                }
            }
        },
        "models.SellingProperty": {
            "type": "object",
            "properties": {
                "is_sold": {
                    "type": "boolean",
                    "example": true
                },
                "price": {
                    "type": "number",
                    "example": 12345.67
                }
            }
        },
        "models.Users": {
            "type": "object",
            "properties": {
                "bank_account_number": {
                    "type": "string",
                    "example": "1234567890"
                },
                "bank_name": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.BankName"
                        }
                    ],
                    "example": "KBANK"
                },
                "citizen_id": {
                    "type": "string",
                    "example": "1234567890123"
                },
                "citizen_image_url": {
                    "type": "string",
                    "example": "https://image_url.com/abcd"
                },
                "createdAt": {
                    "type": "string"
                },
                "credit_card_cvv": {
                    "type": "string",
                    "example": "123"
                },
                "credit_card_expiration_month": {
                    "type": "string",
                    "example": "12"
                },
                "credit_card_expiration_year": {
                    "type": "string",
                    "example": "2023"
                },
                "credit_card_number": {
                    "type": "string",
                    "example": "1234567890123456"
                },
                "credit_cardholder_name": {
                    "type": "string",
                    "example": "JOHN DOE"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "example": "email@email.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "is_verified": {
                    "type": "boolean",
                    "example": false
                },
                "last_name": {
                    "type": "string",
                    "example": "Doe"
                },
                "password": {
                    "type": "string",
                    "example": "password1234"
                },
                "phone_number": {
                    "type": "string",
                    "example": "0812345678"
                },
                "profile_image_url": {
                    "type": "string",
                    "example": "https://image_url.com/abcd"
                },
                "registered_type": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.RegisteredType"
                        }
                    ],
                    "example": "EMAIL"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Bangkok Property Matchmaking Platform",
	Description:      "Bangkok Property Matchmaking Platform API docs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
