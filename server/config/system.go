package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                  // 环境值
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                               // 端口值
	DbType        string `mapstructure:"db-type" json:"db_type" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"oss_type" yaml:"oss-type"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use_multipoint" yaml:"use-multipoint"` // 多点登录拦截
	//UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                // 使用redis
	LimitCountIP int    `mapstructure:"ip-limit-count" json:"ip_limit_count" yaml:"ip-limit-count"`
	LimitTimeIP  int    `mapstructure:"ip-limit-time" json:"ip_limit_time" yaml:"ip-limit-time"`
	FilePosition string `mapstructure:"file-position" json:"file_position" yaml:"file-position"` // 文件存储位置
	Language     string `mapstructure:"language" json:"language" yaml:"language"`                // 语言
}
