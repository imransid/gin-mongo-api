package models

import (
    "time"
   // "github.com/dgrijalva/jwt-go"
    "go.mongodb.org/mongo-driver/bson/primitive"
)


type Auth struct {
	Username 	    string			`json:"username,omitempty" validate:"required"`
    Id       primitive.ObjectID     `json:"id,omitempty"`
    Name     string                 `json:"name,omitempty"`
    Location string                 `json:"location,omitempty"`
    Title    string                 `json:"title,omitempty"`
    Password string                 `json:"password,omitempty"  validate:"required"`
    Email    string                 `json:"email,omitempty"` 
    Phone    string                 `json:"phone"`
    Token    string                 `json:"token"`
    Refresh_token string            `json:"refresh_token"`
    Created_at    time.Time         `json:"created_at"`
    Updated_at    time.Time         `json:"updated_at"`
    User_id       string            `json:"user_id"`
    Role          string            `json:"role" validate:"required"`
}