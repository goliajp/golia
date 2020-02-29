package libutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func PKCS7Padding(content []byte, blockSize int) []byte {
	padding := blockSize - len(content)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(content, padText...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func AesEncrypt(originContent, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	originContent = PKCS7Padding(originContent, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	cryptContent := make([]byte, len(originContent))
	blockMode.CryptBlocks(cryptContent, originContent)
	return cryptContent, nil
}

func AesDecrypt(cryptContent, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cryptContent))
	blockMode.CryptBlocks(origData, cryptContent)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}
