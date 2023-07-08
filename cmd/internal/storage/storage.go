package storage

import "errors"

var (
	ErrURLNotFount = errors.New("URL Not Found")
	ErrURLExists = errors.New("URL exists")
)