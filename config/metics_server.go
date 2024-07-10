package config

type MetricsServerConfig struct {
	Enable bool   `yaml:"enable"`
	Mode   string `yaml:"mode"`
	Port   int    `yaml:"port"`
}

func (c *MetricsServerConfig) SetDefault() {
	c.Enable = true
	c.Mode = "release"
	c.Port = 9119
}
