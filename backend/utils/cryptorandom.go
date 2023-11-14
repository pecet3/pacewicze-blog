package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
)

func GetRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func HashPassword(password, salt string) (string, error) {
	hasher := sha256.New()

	_, err := hasher.Write([]byte(password + salt))
	if err != nil {
		return "", err
	}
	hashInBytes := hasher.Sum(nil)

	hashString := hex.EncodeToString(hashInBytes)

	return hashString, nil
}

func GenerateSalt(size int) (string, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func VerifyPassword(password, salt string) (bool, error) {
	hashedPassword, err := HashPassword(password, salt)
	if err != nil {
		return false, err
	}
	if hashedPassword != password {
		return true, nil
	}

	return false, errors.New("Passwords are not the same")
}
