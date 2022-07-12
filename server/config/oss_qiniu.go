package config

type Qiniu struct {
	AccessKey      string `mapstructure:"access-key" json:"access_key" yaml:"access-key"`
	SecretKey      string `mapstructure:"secret-key" json:"secret_key" yaml:"secret-key"`
	Bucket         string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                            //bucket名称
	Domain         string `mapstructure:"domain" json:"domain" yaml:"domain"`                            //区域 华东:storage.ZoneHuadong/华北:storage.ZoneHuabei/华南:storage.ZoneHuanan/北美:storage.ZoneBeimei /新加坡:storage.ZoneXinjiapo
	Path           string `mapstructure:"path" json:"path" yaml:"path"`                                  //路径
	DomainName     string `mapstructure:"domain-name" json:"domain_name" yaml:"domain-name"`             //	域名
	DomainProtocol string `mapstructure:"domain-protocol" json:"domain_protocol" yaml:"domain-protocol"` //域名协议
}
