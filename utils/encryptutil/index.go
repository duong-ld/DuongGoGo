package encryptutil

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	encryptedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(encryptedByte), err
}

func ComparePassword(encryptedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))

	if err != nil {
		return false
	}
	return true
}
