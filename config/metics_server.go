package config

type MetricsServerConfig struct {
	Enable bool   `yaml:"enable"`
	Mode   string `yaml:"mode"`
	Port   int    `yaml:"port"`
}

// SetDefault metrics 的server分开后无法统计业务的request请求数
func (c *MetricsServerConfig) SetDefault() {
	c.Enable = false
	c.Mode = "release"
	c.Port = 9119
}
