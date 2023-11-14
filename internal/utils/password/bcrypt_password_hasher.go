package password

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHasher struct {
}

func (s *BcryptPasswordHasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), err
}

func (s *BcryptPasswordHasher) Compare(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
