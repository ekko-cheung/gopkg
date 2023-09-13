package main

import (
	"flag"
	"github.com/spf13/viper"
	"github.com/veerdone/gopkg/util/crypto"
	"log"
)

var (
	key      string
	filePath string
)

func main() {
	flag.StringVar(&key, "key", "", "crypto key")
	flag.StringVar(&filePath, "file", "", "config file path")
	flag.Parse()
	if key == "" || filePath == "" {
		log.Fatalln("Required parameter missing: key or file")
	}
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("viper read config fail: %s", err)
	}

	if err := encodeViper(key); err != nil {
		log.Fatalf("crypto config fail: %s", err)
	}

	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("viper write config fail: %s", err)
	}
}

func encodeViper(cryptoKey string) error {
	keys := viper.AllKeys()
	for _, key := range keys {
		value := viper.Get(key)
		if s, ok := value.(string); ok {
			desStr, err := crypto.DesEncoding(s, cryptoKey)
			if err != nil {
				return err
			}
			viper.Set(key, desStr)
		}
	}

	return nil
}
