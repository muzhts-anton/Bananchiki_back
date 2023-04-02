package autusc

import (
	"banana/pkg/domain"

	"net/mail"
	"strings"
	"unicode"
)

const minPasswordLen = 8

func validateEmail(address string) error {
	if _, err := mail.ParseAddress(address); err != nil {
		return domain.ErrInvalidEmail
	}

	return nil
}

func validateUsername(username string) error {
	for _, char := range username {
		if !(unicode.IsLetter(char) || unicode.Is(unicode.Cyrillic, char)) {
			return domain.ErrInvalidUsername
		}
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < minPasswordLen {
		return domain.ErrInvalidPassword
	}

	return nil
}

func trimCredentials(email *string, username *string, password *string, repeatPassword *string) {
	*email = strings.Trim(*email, " ")
	*username = strings.Trim(*username, " ")
	*password = strings.Trim(*password, " ")
	*repeatPassword = strings.Trim(*repeatPassword, " ")
}