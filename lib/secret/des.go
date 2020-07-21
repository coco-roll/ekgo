package secret

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"golang.org/x/crypto/bcrypt"
)

func padding(src []byte, blocksize int) []byte {
	n := len(src)
	padnum := blocksize - n%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	dst := append(src, pad...)
	return dst
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	dst := src[:n-unpadnum]
	return dst
}

func EncryptDES(src []byte) []byte {
	key := []byte("qimiao66")
	block, _ := des.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func DecryptDES(src []byte) []byte {
	key := []byte("qimiao66")
	block, _ := des.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src

}

//bcrypt加密
func NewBcrypt(str string) string {
	hp, _ := bcrypt.GenerateFromPassword([]byte(str), 10)
	return string(hp)
}

//bcrypt验证
func CheckBcrypt(hasAndPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hasAndPassword), []byte(password))
	return err
}
