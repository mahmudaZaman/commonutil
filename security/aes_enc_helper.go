package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/mergermarket/go-pkcs7"
	"io"
)

/*
SimpleAESEncrypt encrypts a string.
*/
func SimpleAESEncrypt(key []byte, unencrypted string) (string, error) {
	plainText := []byte(unencrypted)
	plainText, plainTextErr := pkcs7.Pad(plainText, aes.BlockSize)
	if plainTextErr != nil {
		return "", fmt.Errorf(`plainText: "%s" has error - (pkcs7.Pad failed).Internal error = %v`, plainText, plainTextErr)
	}
	if len(plainText)%aes.BlockSize != 0 {
		return "", fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
	}
	block, blockErr := aes.NewCipher(key)
	if blockErr != nil {
		return "", fmt.Errorf(`failed to create bock - (aes.NewCipher failed), Internal error =%v`, blockErr)
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	_, readErr := io.ReadFull(rand.Reader, iv)
	if readErr != nil {
		return "", readErr
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	data := fmt.Sprintf("%x", cipherText)
	return data, nil
}
