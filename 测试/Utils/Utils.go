package Utils

import (
	"demo/Global"
	"demo/Structs"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

//签名密钥
var jwtKey = []byte(Global.Config.JWT.Key)

//分发Token
func ReleaseToken(u Structs.User) (string, error) {

	// token 结构生成
	claims := &Structs.Claims{
		// 使用 ID、Username 作为有效载荷
		UID:      u.UID,
		Account: u.Account,
		Password: u.Password,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + Global.Config.JWT.Expires, // 签名过期时间
			NotBefore: time.Now().Unix() - 1000,                      // 签名生效时间
			Issuer:    Global.Config.JWT.Issuer,                      // 签名发行人
		},
	}

	// 将 Claims 加密存储为 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//解析Token
func ParseToken(tokenString string) (*jwt.Token, *Structs.Claims, error) {
	claims := &Structs.Claims{}
	// 解码
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}

// 判断用户名是否存在
func IsAccountExist(account string) bool {
	var count int64
	Global.DB.Model(Structs.User{}).Where("account = ?", account).Count(&count)
	return count > 0
}

//判断用户邮箱否存在
func IsEmailExist(email string) bool {
	var count int64
	Global.DB.Model(Structs.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

//随即创造六位字符串型验证码
func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
