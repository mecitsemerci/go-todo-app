package auth

import "regexp"

type Email string

func (e Email) String() string {
	return string(e)
}

func (e Email) Equals(other Email) bool {
	return e.String() == other.String()
}

func (e Email) IsValid() bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(e.String())
}
