package util

import (
	"TikTokLite_v2/common/service"
	"TikTokLite_v2/user_follow/setting"
	"TikTokLite_v2/user_follow/user/dal"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//Claims token claims
type Claims struct {
	Username string `json:"username"`
	UserID   int64  `json:"user_id"`
	jwt.StandardClaims
}

//GetToken 获取对应user的token
func GetToken(u *dal.User) (tokenString string, statusCode int32, statusMsg string, err error) {
	var claims Claims
	claims.Username = u.Name
	claims.UserID = int64(u.ID)
	//token过期时间
	claims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	//token创建时间
	claims.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//加密生成token
	tokenString, err = token.SignedString([]byte(setting.Conf.Token.SecretKey))
	statusCode, statusMsg = service.BuildResponse(err)
	return
}
