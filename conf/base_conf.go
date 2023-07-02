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

type BaseConf struct {
	Db        Db        `yaml:"db" json:"db" toml:"db" properties:"db"`
	Redis     Redis     `yaml:"redis" json:"redis" toml:"redis" properties:"redis"`
	Log       Log       `yaml:"log" json:"log" toml:"log" properties:"log"`
	Etcd      Etcd      `yaml:"etcd" json:"etcd" toml:"etcd" properties:"etcd"`
	Memcached Memcached `yaml:"memcached" json:"memcached" toml:"memcached" properties:"memcached"`
}

type Db struct {
	Username string `yaml:"username" json:"username" toml:"username" properties:"username"`
	Pass     string `yaml:"pass" json:"pass" toml:"pass" properties:"pass"`
	Url      string `yaml:"url" json:"url" toml:"url" properties:"url"`
	Database string `yaml:"database" json:"database" toml:"database" properties:"database"`
}

type Redis struct {
	Addr     string `yaml:"addr" json:"addr" toml:"addr" properties:"addr"`
	Username string `yaml:"username" json:"username" toml:"username" properties:"username"`
	Pass     string `yaml:"pass" json:"pass" toml:"pass" properties:"pass"`
	Db       int    `yaml:"db" json:"db" toml:"db" properties:"db"`
}

type Log struct {
	Level  string   `yaml:"level" json:"level" toml:"level" properties:"level"`
	Output []string `yaml:"output" json:"output" toml:"output" properties:"output"`
}

type Etcd struct {
	Endpoints []string `yaml:"endpoints" json:"endpoints" toml:"endpoints" properties:"endpoints"`
}

type Memcached struct {
	Address []string `json:"address" yaml:"address" toml:"address" properties:"address"`
}
