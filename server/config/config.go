package config

// Config is the configuration for the application.
type Config struct {
	Zap    ZapConfig `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System    `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql     `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Redis  Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
	JWT  JWT  `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	//	验证码
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
