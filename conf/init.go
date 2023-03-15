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
