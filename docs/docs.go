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
        "/user/changePassword": {
            "post": {
                "description": "Change Password user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "change password",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "changePasswordInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.ChangePasswordUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TokenObject"
                        }
                    }
                }
            }
        },
        "/user/googleAuth": {
            "post": {
                "description": "Google Auth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "google auth",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "googleOAuthInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.GoogleOAuthInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TokenObject"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "loginUserInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TokenObject"
                        }
                    }
                }
            }
        },
        "/user/profile": {
            "get": {
                "description": "Get user profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "get profile",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            }
        },
        "/user/refreshToken": {
            "post": {
                "description": "Refresh JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "refresh token",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd refresh token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User Data",
                        "name": "refreshTokenInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RefreshTokenInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TokenObject"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Adding new user to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "register",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "registerUserInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RegisterUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TokenObject"
                        }
                    }
                }
            }
        },
        "/user/resetPassword": {
            "post": {
                "description": "Send email consist of OTP to user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "reset password",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "resetPasswordInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.ResetPasswordInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TokenObject"
                        }
                    }
                }
            }
        },
        "/user/validateToken": {
            "get": {
                "description": "Validate JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "validate token",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TokenObject"
                        }
                    }
                }
            }
        },
        "/user/verifyOtp": {
            "post": {
                "description": "Verify OTP user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "verify otp",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "verifyOtpInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/otp.VerifyOtpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TokenObject"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.TokenObject": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "otp.VerifyOtpInput": {
            "type": "object",
            "properties": {
                "otp_code": {
                    "type": "string"
                }
            }
        },
        "user.ChangePasswordUserInput": {
            "type": "object",
            "properties": {
                "change_password": {
                    "type": "string"
                },
                "otp_code": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.GoogleOAuthInput": {
            "type": "object",
            "required": [
                "oAuthAccessToken"
            ],
            "properties": {
                "oAuthAccessToken": {
                    "type": "string"
                }
            }
        },
        "user.LoginUserInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.RefreshTokenInput": {
            "type": "object",
            "required": [
                "accessToken"
            ],
            "properties": {
                "accessToken": {
                    "type": "string"
                }
            }
        },
        "user.RegisterUserInput": {
            "type": "object",
            "required": [
                "confirmPassword",
                "email",
                "password",
                "username"
            ],
            "properties": {
                "confirmPassword": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.ResetPasswordInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "last_login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
