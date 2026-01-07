package password

type Password interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) (bool, error)
}
