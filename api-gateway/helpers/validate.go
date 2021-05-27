package helpers

import "regexp"

type Validation struct {
	Value string
	Valid string
}

func Validate(values []Validation) bool {
	email := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$`)

	for i := 0; i < len(values); i++ {
		if values[i].Valid == "email" {
			if !email.MatchString(values[i].Value) {
				return false
			}
		}
	}

	return true
}