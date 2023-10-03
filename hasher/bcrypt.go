package modules

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct {
}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (u *BcryptHasher) Hash(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, 10)
}

func (u *BcryptHasher) Compare(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (u *BcryptHasher) Cost(hashedPassword []byte) (int, error) {
	return bcrypt.DefaultCost, nil
}
