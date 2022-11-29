package middleware

import (
	"GinMES/database"
	"GinMES/models"
	"GinMES/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无权限访问，请求未携带token",
			})
			ctx.Abort() //结束后续操作
			return
		}
		log.Print("token:", authHeader)

		//按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "请求头中auth格式有误",
			})
			ctx.Abort()
			return
		}

		//解析token包含的信息
		claims, err := utils.ParseToken(parts[1])
		// 解析到的token如果过期，刷新新的 token设置到 Authorization
		if strings.Contains(err.Error(), "expired") {
			new_token, _ := utils.RefreshToken(parts[1])
			ctx.Header("new_token", new_token)
			ctx.Request.Header.Set("Authorization", parts[0]+" "+new_token)
			ctx.Next()
			return
		}
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无效的Token",
			})
			ctx.Abort()
			return
		}

		/*	if err := CheckUserInfo(claims); err != nil{
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			ctx.Abort()
			return
		}*/

		// 将当前请求的claims信息保存到请求的上下文c上
		ctx.Set("claims", claims)
		ctx.Next() // 后续的处理函数可以用过ctx.Get("claims")来获取当前请求的用户信息

	}
}

// 检查用户名信息
func CheckUserInfo(claims *models.CustomClaims) error {
	username := claims.Username
	password := claims.Password
	//获取数据库用户名及密码
	var user_object models.Users
	database.Db.Model(&user_object).Where(&models.Users{Username: username, Password: password}).First(&user_object)

	if user_object.Username == username && user_object.Password == password {
		return nil
	}
	return errors.New("用户名或密码错误")
}
