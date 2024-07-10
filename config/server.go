package config

type ServerConfig struct {
	Enable bool   `yaml:"enable"`
	Mode   string `yaml:"mode"`
	Port   int    `yaml:"port"`
}

func (s *ServerConfig) SetDefault() {
	s.Enable = true
	s.Mode = "release"
	s.Port = 8119
}
