package domain

import "errors"

var (
	ErrDatabaseRequest = errors.New("Bad database request")
	ErrDatabaseRange = errors.New("Index out of database")
)