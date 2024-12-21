package models

import "github.com/google/uuid"

type Todo struct {
    Title       string    `json:"title" validate:"required"`
    IsCompleted *bool      `json:"isCompleted" validate:"required"`
    Id          uuid.UUID `json:"id"`
}