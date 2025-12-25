package internal

import (
	"time"

	"github.com/google/uuid"
)

type Motion struct {
	ID        uuid.UUID      
	Timestamp time.Time 
}

type Response struct {
	Detail string `json:"detail"`
}

type MotionRequest struct {
	AccessKey string `json:"access_key"`
	Timestamp time.Time `json:"timestamp"`
}
