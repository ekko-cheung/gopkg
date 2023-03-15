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
