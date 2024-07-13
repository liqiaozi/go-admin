package component

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"lixuefei.com/go-admin/global"
	"time"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token:")
)

type jwtService struct {
}

var JwtService = new(jwtService)

type JwtUser interface {
	GetUid() string
}

type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType = "bearer"
	Issuer    = "app"
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// 生成Token
func (jwtService *jwtService) CreateToken(issuer string, user JwtUser) (tokenOutPut TokenOutPut, err error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + global.App.Server.Jwt.JwtTtl,
				Id:        user.GetUid(),
				Issuer:    issuer, // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(global.App.Server.Jwt.SignKey))
	tokenData := TokenOutPut{
		tokenStr,
		int(global.App.Server.Jwt.JwtTtl),
		TokenType,
	}
	return tokenData, nil
}

// 解析Token
func (jwtService *jwtService) ParseToken(tokenStr string) (*CustomClaims, error) {
	// Token 解析校验
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.App.Server.Jwt.SignKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		} else {
			return nil, TokenInvalid
		}
	} else {
		return nil, TokenInvalid
	}
}
