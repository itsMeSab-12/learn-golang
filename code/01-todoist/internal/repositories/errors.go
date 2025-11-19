package repositories

import "errors"

var ErrNotFound = errors.New("record not found")
var ErrConflict = errors.New("record already exists")
