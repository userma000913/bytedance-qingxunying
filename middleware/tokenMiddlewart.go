package middleware

import (
	"dousheng-backend/common"
	"dousheng-backend/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication(c *gin.Context) {
	//Token解析示例
	_, err := jwt.Parse(c.PostForm("token"), func(token *jwt.Token) (interface{}, error) {
		return []byte("bytedance"), nil
	})
	if err != nil {
		common.Error(c, http.StatusInternalServerError)
		c.Abort()
		return
	}
	userId, _ := util.TokenVerify(c.PostForm("token"))
	c.Set("userId", userId)
	c.Next()
	return
}