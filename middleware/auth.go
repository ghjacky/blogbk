package middleware

import (
	"blogbk/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

func Auth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if isNotNeedCheckAuth(ctx.Request.RequestURI, ctx.Request.Method) {
			ctx.Next()
			return
		}
		cki, err := ctx.Cookie("ss_id")
		if err == nil && len(cki) > 0 {
			uss := strings.Split(cki, "_")
			if len(uss) == 2 {
				un := uss[0]
				sid := uuid.MustParse(uss[1])
				var sess = model.SSession{Username: un, ID: sid}
				if sess.Validate() {
					ctx.Next()
					return
				}
			}
		}
		ctx.JSON(403, "请登录")
		ctx.Abort()
		return
	}
}

func isNotNeedCheckAuth(path string, method string) bool {
	return path == "/api/v1/users/auth" ||
		path == "/api/v1/users/auth/logout" ||
		strings.ToLower(method) == "get"
}
