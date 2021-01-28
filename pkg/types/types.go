package types

import (
	"time"

)
const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Note struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}