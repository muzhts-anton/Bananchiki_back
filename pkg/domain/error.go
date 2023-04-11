package domain

import "errors"

var (
	ErrDatabaseRequest = errors.New("Bad database request")
	ErrDatabaseRange   = errors.New("Index out of database")
	ErrUrlParameter    = errors.New("Error while parsing url parameter")
	ErrInternalServer  = errors.New("Internal server error")
	ErrCodeNotFound    = errors.New("No presentation with requested code")

	ErrPermissionDenied = errors.New("Permission Denied")

	ErrFinishSession   = errors.New("User is already not logged in while trying log out")
	ErrUserNotLoggedIn = errors.New("User is not logged in")
	ErrSessionCast     = errors.New("Incorrect id info in session store")

	ErrInvalidEmail       = errors.New("Invalid email")
	ErrInvalidUsername    = errors.New("Invalid username")
	ErrInvalidPassword    = errors.New("Invalid password")
	ErrEmptyField         = errors.New("Empty field")
	ErrUnmatchedPasswords = errors.New("Unmatched passwords")
	ErrEmailExists        = errors.New("Email not unique")
	ErrBadPassword        = errors.New("Wrong password")
	ErrNoUser             = errors.New("No user found")

	ErrInvalidPresName = errors.New("Presentation name is too long")
)
