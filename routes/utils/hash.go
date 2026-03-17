package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, HashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(password))
	return err == nil
}
