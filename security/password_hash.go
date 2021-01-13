package security

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

/*
HashPassword hashes a password with bcrypt.
*/
func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Printf("Error creating Hash = %v", err)
		return "", err
	}
	return string(hash), nil
}

/*
ComparePasswords compares if two plain password matches with hashed password.
*/
func ComparePasswords(hashedPwd, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}
