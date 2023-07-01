package user

import (
	"strings"
	"unicode"
)

const MIN_NAME_LENGTH = 3
const MAX_NAME_LENGTH = 30
const MIN_PASSWORD_LENGTH = 8
const MAX_PASSWORD_LENGTH = 100

func IsNameValid(name string) bool {
	name = strings.TrimSpace(name)
	nameLen := len(name)
	if MIN_NAME_LENGTH > nameLen || nameLen > MAX_NAME_LENGTH {
		return false
	}
	first := true
	for _, r := range name {
		if first && !unicode.IsLetter(r) {
			return false
		}
		first = false
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func IsPasswordValid(password string) bool {
	password = strings.TrimSpace(password)
	passwordLen := len(password)
	if passwordLen < MIN_PASSWORD_LENGTH || passwordLen > MAX_PASSWORD_LENGTH {
		return false
	}

	hasLowerLetter := false
	hasUpperLetter := false
	hasDigit := false

	for _, r := range password {
		if !hasLowerLetter && unicode.IsLower(r) {
			hasLowerLetter = true
		} else if !hasUpperLetter && unicode.IsUpper(r) {
			hasUpperLetter = true
		} else if !hasDigit && unicode.IsDigit(r) {
			hasDigit = true
		}

		if hasLowerLetter && hasUpperLetter && hasDigit {
			return true
		}
	}

	return false
}
