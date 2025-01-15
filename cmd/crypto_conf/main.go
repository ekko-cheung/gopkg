/*
 * Copyright 2023 veerdone
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"github.com/spf13/viper"
	"github.com/veerdone/gopkg/util/crypto"
	"log"
)

var (
	key     string
	srcFile string
	outFile string
)

func main() {
	flag.StringVar(&key, "key", "", "crypto key")
	flag.StringVar(&srcFile, "file", "", "input config file")
	flag.StringVar(&outFile, "out", "", "out config file")
	flag.Parse()
	if key == "" || srcFile == "" {
		log.Fatalln("Required parameter missing: key or file")
	}
	viper.SetConfigFile(srcFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("viper read config fail: %s", err)
	}

	if err := encodeViper(key); err != nil {
		log.Fatalf("crypto config fail: %s", err)
	}

	if outFile == "" {
		err = viper.WriteConfig()
	} else {
		err = viper.WriteConfigAs(outFile)
	}
	if err != nil {
		log.Fatalf("viper wrire config fail: %s", err)
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
