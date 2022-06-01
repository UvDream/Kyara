package config

// Config is the configuration for the application.
type Config struct {
	Zap    ZapConfig `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System    `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql     `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
}
