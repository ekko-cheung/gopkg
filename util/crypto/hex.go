package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/veerdone/gopkg/util"
)

func HexEncodeSha256(s string) string {
	b := HexEncodeSha256B(util.StringToSliceByte(s))

	return util.SliceByteToString(b)
}

func BHexEncodeSha256(b []byte) string {
	bytes := HexEncodeSha256B(b)

	return util.SliceByteToString(bytes)
}

func HexEncodeSha256B(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	sum := h.Sum(nil)

	dst := make([]byte, hex.EncodedLen(len(sum)))
	hex.Encode(dst, sum)

	return dst
}

func HmacSha256(msg, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write(msg)

	return h.Sum(nil)
}
