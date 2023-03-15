package log

import (
	"github.com/veerdone/gopkg/conf"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

type fileConf struct {
	FileName   string `yaml:"fileName" json:"fileName" toml:"fileName" properties:"fileName"`
	MaxSize    int    `yaml:"maxSize" json:"maxSize" toml:"maxSize" properties:"maxSize"`
	MaxBackups int    `yaml:"maxBackups" json:"maxBackups" toml:"maxBackups" properties:"maxBackups"`
	MaxAge     int    `yaml:"maxAge" json:"maxAge" toml:"maxAge" properties:"maxAge"`
	Compress   bool   `yaml:"compress" json:"compress" toml:"compress" properties:"compress"`
}

func fileOut() io.WriteCloser {
	c := new(fileConf)
	conf.Parse("log.file", c)

	return &lumberjack.Logger{
		Filename:   c.FileName,
		MaxAge:     c.MaxAge,
		MaxBackups: c.MaxBackups,
		MaxSize:    c.MaxSize,
		Compress:   c.Compress,
	}
}
