package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPass), err

}

func CheckPasswordHash(plaintext, hashedPass string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plaintext))
	return err == nil
	

}