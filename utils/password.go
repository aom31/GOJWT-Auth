package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	//bcrypt password
	byteHashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(byteHashedPassword), nil
}

func VerifyPassword(hashedPassword, candidatePassword string) error {
	comparePassword := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	return comparePassword
}
