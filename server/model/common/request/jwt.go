package request

import (
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}
type BaseClaims struct {
	ID          string
	Username    string
	NickName    string
	AuthorityId string
}
