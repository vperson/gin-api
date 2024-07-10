package config

type LogConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

func (l *LogConfig) SetDefault() {
	l.Level = "info"
	l.Format = "text"
}
