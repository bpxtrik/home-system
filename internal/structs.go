package internal

import (
	"time"

	"github.com/google/uuid"
)

type Motion struct {
	ID        uuid.UUID      
	Timestamp time.Time `json:"timestamp"`
}

type Response struct {
	Detail string `json:"detail"`
}


