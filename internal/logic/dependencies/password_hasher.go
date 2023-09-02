package dependencies

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(password, hash string) bool
}
