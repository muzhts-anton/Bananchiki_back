package domain

import "errors"

var (
	ErrInternalServer = errors.New("Internal server error")

	ErrDatabaseRequest = errors.New("Bad database request")
	ErrDatabaseRange   = errors.New("Index out of database")
	ErrCodeNotFound    = errors.New("No presentation with requested code")

	ErrUrlParameter = errors.New("Error while parsing url parameter")

	ErrPermissionDenied = errors.New("Permission Denied")

	ErrInvalidPresName = errors.New("Presentation name is too long")

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

	ErrWrongEmotions = errors.New("Wrong Emotions")

	ErrGrpc = errors.New("gRPC error")

	ErrSecondVote          = errors.New("Try to vote twice")
	ErrRunout              = errors.New("Quiz time has run out")
	ErrUnexpectedTimeValue = errors.New("Quiz Answer time is beyond expected interval")
)
