package domain

import "errors"

var (
	ErrDatabaseRequest = errors.New("Bad database request")
	ErrDatabaseRange   = errors.New("Index out of database")
	ErrUrlParameter    = errors.New("Error while parsing url parameter")
	ErrInternalServer  = errors.New("Internal server error")

	ErrFinishSession   = errors.New("User is already not logged in while trying log out")
	ErrUserNotLoggedIn = errors.New("User is not logged in")
	ErrSessionCast     = errors.New("Incorrect id info in session store")
)
