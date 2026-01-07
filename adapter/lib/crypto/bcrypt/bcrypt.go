package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	cost int
}

func New(cost int) *Bcrypt {
	if cost <= bcrypt.DefaultCost {
		cost = bcrypt.DefaultCost
	}
	return &Bcrypt{cost: cost}
}

func (b *Bcrypt) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	return string(bytes), err
}

func (b *Bcrypt) Compare(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
