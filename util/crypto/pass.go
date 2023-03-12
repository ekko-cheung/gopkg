package crypto

import (
	"github.com/veerdone/gopkg/log"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func EncoderPass(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Error("encoder pass error", zap.Error(err))
	}
	return string(hash)
}

func ComparePass(hash, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
