package config

import "image/color"

type Captcha struct {
	Type            string      `mapstructure:"type" json:"type" yaml:"type"`                                        // 验证码类型 default/audio/math/string/math/chinese
	Language        string      `mapstructure:"language" json:"language" yaml:"language"`                            // 验证码语言 zh/en
	Height          int         `mapstructure:"height" json:"height" yaml:"height"`                                  // 验证码高度
	Width           int         `mapstructure:"width" json:"width" yaml:"width"`                                     // 验证码宽度
	Length          int         `mapstructure:"length" json:"length" yaml:"length"`                                  // 验证码长度
	NoiseCount      int         `mapstructure:"noise-count" json:"noise-count" yaml:"noise-count"`                   // 噪点数量
	ShowLineOptions int         `mapstructure:"show-line-options" json:"show-line-options" yaml:"show-line-options"` // 显示线条选项  2/4/8
	BgColor         *color.RGBA `mapstructure:"bg-color" json:"bg-color" yaml:"bg-color"`                            // 验证码背景颜色
}
