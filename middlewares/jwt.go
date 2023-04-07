package middlewares

import (
	"errors"
	"net/http"
	"time"

	"admin_api/global"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// JWT 鉴权取头部信息 x-token 登录时回返回 token 信息
		// 这里前端需要把 token 存储到 cookie 或者本地 local storage 中
		// 过期时间需约定,可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "请登录"})
			c.Abort()
			return
		}

		// 解析 token 包含的信息
		claims, err := NewJWT().ParseToken(token)
		if err != nil {
			// token 过期
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{"msg": "授权已过期"})
				c.Abort()
				return
			}

			c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登陆"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Set("user_id", claims.UserId)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.ServerConfig.JWTInfo.SigningKey), //可以配置过期时间
	}
}

// Claims PayLoad
type Claims struct {
	UserId string // TODO 完善 PayLoad
	jwt.StandardClaims
}

// CreateToken 创建一个 token
func (j *JWT) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
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
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

// RefreshToken 更新 token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
