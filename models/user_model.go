package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
    Id       primitive.ObjectID     `json:"id,omitempty"`
    Name          string                 `json:"name,omitempty" validate:"required"`
    Location      string                 `json:"location,omitempty" validate:"required"`
    FirstName     string                 `json:"firstname,omitempty" validate:"required"`
    LastName      string                 `json:"lastname,omitempty" validate:"required"`
    Title         string                 `json:"title,omitempty" validate:"required"`
    Password      string                 `json:"password,omitempty" validate:"required"`
    Email         string                 `json:"email,omitempty" validate:"required"` 
    Phone         string                 `json:"phone" validate:"required"`
    Token         string                 `json:"token"`
    Refresh_token string                `json:"refresh_token"`
    Created_at    time.Time             `json:"created_at"`
    Updated_at    time.Time             `json:"updated_at"`
    User_id       string                `json:"user_id"`
    Role          string                `json:"role" validate:"required"`
}

