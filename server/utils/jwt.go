package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"server/config"
	"server/global"
	"time"
)

var (
	TokenExpired     = errors.New("令牌已过期")
	TokenNotValidYet = errors.New("令牌尚未激活")
	TokenMalformed   = errors.New("这不是一个token")
	TokenInvalid     = errors.New("无法处理此token")
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Config.JWT.SigningKey),
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims config.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims config.CustomClaims) (string, error) {
	v, err, _ := global.ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*config.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &config.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*config.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// CreateClaims 创建一个claims
func (j *JWT) CreateClaims(baseClaims config.BaseClaims) config.CustomClaims {
	return config.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.Config.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + global.Config.JWT.ExpiresTime,
			Issuer:    global.Config.JWT.Issuer,
		},
	}
}
