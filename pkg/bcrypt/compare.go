package bcrypt

import "golang.org/x/crypto/bcrypt"

func Compare(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return err
	}

	return nil
}
