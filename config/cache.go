package config

type CacheConfig struct {
	Enable   bool   `yaml:"enable"`
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

func (d *CacheConfig) SetDefault() {
	d.Db = 0
	d.Address = "127.0.0.1:6379"
}
