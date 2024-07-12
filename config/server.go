package config

type ServerConfig struct {
	Enable        bool   `yaml:"enable"`
	MetricsEnable bool   `json:"metricsEnable"`
	Mode          string `yaml:"mode"`
	Port          int    `yaml:"port"`
}

func (s *ServerConfig) SetDefault() {
	s.Enable = true
	s.Mode = "release"
	s.MetricsEnable = true
	s.Port = 8119
}
