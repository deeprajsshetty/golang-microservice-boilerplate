package repository

import "errors"

// ErrNotFound is return when a requested record is not found.
var ErrNotFound = errors.New("not found")
