package types

import (
	"errors"
	"time"

)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

var (
	//ErrNotFound ...
	ErrNotFound = errors.New("item not found")
	//ErrInternal ...
	ErrInternal = errors.New("internal error")
	//ErrTokenNotFound ...
	ErrTokenNotFound = errors.New("token not found")
	//ErrNoSuchUser ...
	ErrNoSuchUser = errors.New("no such user")
	//ErrInvalidPassword ..
	ErrInvalidPassword = errors.New("invalid password")
	//ErrPhoneUsed ...
	ErrPhoneUsed = errors.New("phone already registered")
	//ErrTokenExpired ...
	ErrTokenExpired = errors.New("token expired")
)
//Note struct 
type Note struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}
