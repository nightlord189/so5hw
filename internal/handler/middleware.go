package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/nightlord189/so5hw/internal/model"
	"github.com/nightlord189/so5hw/internal/service"
	"net/http"
	"strings"
)

func CheckAuthMiddleware(jwtAccessSecret string) func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		splitted := strings.Split(authHeader, " ")
		if len(splitted) != 2 {
			c.JSON(http.StatusUnauthorized, model.GenericError(-91, "wrong auth header"))
			c.Abort()
			return
		}
		tokenStr := splitted[1]
		jwtToken, err := service.GetJwtToken(tokenStr, jwtAccessSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.GenericError(-92, err.Error()))
			c.Abort()
			return
		}
		if !jwtToken.Valid {
			c.JSON(http.StatusUnauthorized, model.GenericError(-93, "invalid token"))
			c.Abort()
			return
		}
		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, model.GenericError(-94, "error getting claims"))
			c.Abort()
			return
		}
		c.Set("user_id", claims["user_id"])
		c.Set("username", claims["username"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
