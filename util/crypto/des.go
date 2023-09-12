package crypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"github.com/veerdone/gopkg/util"
)

func UnPadPwd(dst []byte) []byte {
	unpadding := dst[len(dst)-1]                // 获取二进制数组最后一个数值
	result := dst[:(len(dst) - int(unpadding))] // 截取开始至总长度减去填充值之间的有效数据

	return result
}

func PadPwd(srcByte []byte, blockSize int) []byte {
	padding := blockSize - len(srcByte)%blockSize // 要填充的值和个数
	slice1 := []byte{byte(padding)}               // 要填充的单个二进制值
	slice2 := bytes.Repeat(slice1, padding)       // 要填充的二进制数组

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
