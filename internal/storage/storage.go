package storage

import "errors"

var (
	ErrURLNotFound         = errors.New("url not found")
	ErrURLNotFoundToDelete = errors.New("url not found to delete")
	ErrURLExists           = errors.New("url exists")
)
