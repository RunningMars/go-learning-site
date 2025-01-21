package models

import (
    "time"
)

type User struct {
    ID        int64     `json:"id" db:"id"`
    Username  string    `json:"username" db:"username"`
    Phone     string    `json:"phone" db:"phone"`
    Email     string    `json:"email" db:"email"`
    Age       *int      `json:"age" db:"age"`
    Gender    *int      `json:"gender" db:"gender"`
    Avatar    string    `json:"avatar" db:"avatar"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type RegisterRequest struct {
    Username string `json:"username" binding:"required"`
    Phone    string `json:"phone" binding:"required"`
    Password string `json:"password" binding:"required"`
    Code     string `json:"code" binding:"required"`
    Email    string `json:"email"`
    Age      *int   `json:"age"`
    Gender   *int   `json:"gender"`
}

type LoginRequest struct {
    Phone    string `json:"phone" binding:"required"`
    Password string `json:"password" binding:"required"`
}