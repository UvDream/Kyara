package utils

import "golang.org/x/crypto/bcrypt"

//BcryptHash 使用 bcrypt 进行加密
func BcryptHash(data string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	return string(bytes)
}

//BcryptCheck 对比加密后的结果和明文是否相同
func BcryptCheck(data string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
	return err == nil
}
