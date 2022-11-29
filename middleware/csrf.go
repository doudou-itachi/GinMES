package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
	"net/http"
)

var csrfMd func(http.Handler) http.Handler

func CSRF() gin.HandlerFunc {
	csrfMd = csrf.Protect(
		[]byte("32-byte-long-auth-key"),
		csrf.Secure(false),
		csrf.HttpOnly(true),
		csrf.ErrorHandler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusForbidden)
			writer.Write([]byte(`{"message":"Forbidden-CSRF token invalid"}`))
		})),
	)
	return adapter.Wrap(csrfMd)
}
