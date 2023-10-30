package toolkit

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
)

var ErrorViLenInvalid = errors.New("length of iv invalid")
var ErrorKeyLenInvalid = errors.New("length of key invalid")

// AEC/CBC/PKCS7Padding

// Encrypt 加密
//
// plainText: 加密目标字符串
// key: 加密Key, 16||24||32位
// iv: 加密iv(AES时固定为16位)
func Encrypt(plainText string, key string, iv string) (string, error) {
	//校验iv length
	if len(iv) != 16 {
		return "", errors.Wrap(
			ErrorViLenInvalid,
			fmt.Sprintf("the length of iv is %d, which excepted to be 16 instead",
				len(iv)))
	}

	//校验key length
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.Wrap(
			ErrorKeyLenInvalid,
			fmt.Sprintf("the length of key is %d, which excepted to be 16,24,or 32 instead",
				len(key)))
	}

	data, err := aesCBCEncrypt([]byte(plainText), []byte(key), []byte(iv))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

// Decrypt 解密
//
// cipherText: 解密目标字符串
// key: 加密Key, 16||24||32位
// iv: 加密iv(AES时固定为16位)
func Decrypt(cipherText string, key string, iv string) (string, error) {
	//校验iv length
	if len(iv) != 16 {
		return "", errors.Wrap(
			ErrorViLenInvalid,
			fmt.Sprintf("the length of iv is %d, which excepted to be 16 instead",
				len(iv)))
	}

	//校验key length
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.Wrap(
			ErrorKeyLenInvalid,
			fmt.Sprintf("the length of key is %d, which excepted to be 16,24,or 32 instead",
				len(key)))
	}

	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	dnData, err := aesCBCDecrypt(data, []byte(key), []byte(iv))
	if err != nil {
		return "", err
	}

	return string(dnData), nil
}

// aesCBCEncrypt AES/CBC/PKCS7Padding 加密
func aesCBCEncrypt(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	// AES
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// PKCS7 填充
	plaintext = paddingPKCS7(plaintext, aes.BlockSize)

	// CBC 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(plaintext, plaintext)

	return plaintext, nil
}

// aesCBCDecrypt AES/CBC/PKCS7Padding 解密
func aesCBCDecrypt(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	// AES
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	// CBC 解密
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// PKCS7 反填充
	result := unPaddingPKCS7(ciphertext)
	return result, nil
}

// PKCS7 填充
func paddingPKCS7(plaintext []byte, blockSize int) []byte {
	paddingSize := blockSize - len(plaintext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(plaintext, paddingText...)
}

// PKCS7 反填充
func unPaddingPKCS7(s []byte) []byte {
	length := len(s)
	if length == 0 {
		return s
	}
	unPadding := int(s[length-1])
	return s[:(length - unPadding)]
}
