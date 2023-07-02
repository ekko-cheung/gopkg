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
