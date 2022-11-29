package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func CsrfToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("X-CSRF-Token", csrf.Token(context.Request))
	}
}
