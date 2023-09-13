package crypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"github.com/veerdone/gopkg/util"
)

func UnPadPwd(dst []byte) []byte {
	unpadding := dst[len(dst)-1]
	result := dst[:(len(dst) - int(unpadding))]

	return result
}

func PadPwd(srcByte []byte, blockSize int) []byte {
	padding := blockSize - len(srcByte)%blockSize
	slice1 := []byte{byte(padding)}
	slice2 := bytes.Repeat(slice1, padding)

	return append(srcByte, slice2...)
}

// DesEncoding key only supports 8 bytes
func DesEncoding(src, key string) (string, error) {
	keyByte := util.StringToSliceByte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	srcByte := []byte(src)
	paddingBytes := PadPwd(srcByte, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyByte[:blockSize])
	dst := make([]byte, len(paddingBytes))
	blockMode.CryptBlocks(dst, paddingBytes)

	return base64.StdEncoding.EncodeToString(dst), nil
}

// DesDecoding key only supports 8 bytes
func DesDecoding(pwd, key string) (string, error) {
	pwdBytes, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return "", err
	}

	keyByte := util.StringToSliceByte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyByte[:blockSize])

	dst := make([]byte, len(pwdBytes))
	blockMode.CryptBlocks(dst, pwdBytes)
	dst = UnPadPwd(dst)

	return string(dst), nil
}

func DesEncodingMap(m map[string]interface{}, key string) error {
	for k, v := range m {
		if tempMap, ok := v.(map[string]interface{}); ok {
			if err := DesEncodingMap(tempMap, key); err != nil {
				return err
			}
		}
		if s, ok := v.(string); ok {
			desStr, err := DesEncoding(s, key)
			if err != nil {
				return err
			}
			m[k] = desStr
		}
	}

	return nil
}
