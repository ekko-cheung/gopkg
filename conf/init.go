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

package conf

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func Init(conf interface{}) {
	var (
		cfgPath  string
		fileName string
	)
	flag.StringVar(&cfgPath, "path", ".", "config file path, default in current path")
	flag.StringVar(&fileName, "file", "config.yaml", "config file name, supports yaml, toml, json, properties")
	flag.Parse()

	sp := strings.Split(fileName, ".")
	if len(sp) < 2 {
		fmt.Printf("config file: '%s' type error", fileName)
		os.Exit(1)
	}
	viper.SetConfigType(sp[1])
	viper.SetConfigName(sp[0])

	viper.AddConfigPath(cfgPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("read config fail: %s", err))
	}

	if err := viper.Unmarshal(conf); err != nil {
		panic(fmt.Sprintf("unmarshal config fail: %s", err))
	}
}

func Parse(key string, conf interface{}) {
	if err := viper.UnmarshalKey(key, conf); err != nil {
		panic(fmt.Sprintf("unmarshal config fail: %s", err))
	}
}
