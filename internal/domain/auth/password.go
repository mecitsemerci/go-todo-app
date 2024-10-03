package auth

import "golang.org/x/crypto/bcrypt"

type PasswordHash string

func (p PasswordHash) String() string {
	return string(p)
}

func (p PasswordHash) Equals(other PasswordHash) bool {
	return p.String() == other.String()
}

func NewPasswordHash(password string) (PasswordHash, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return PasswordHash(pwdHash), nil
}

func (p PasswordHash) Verify(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(password))
}
