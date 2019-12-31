package common

import "os"

var JwtSecretKey []byte

func InitConfig() {
	JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}

func GetJwtKey() []byte {
	return JwtSecretKey
}
