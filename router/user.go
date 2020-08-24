package router

import (
	"blogbk/model"
	"blogbk/module/user"
	"github.com/gin-gonic/gin"
)

func userRouter() {
	ur := Router.Group(api("/users"))
	{
		ur.GET("", getUser)
		ur.POST("/auth", authUser)
		ur.POST("/auth/logout", logout)
	}
}

func getUser(ctx *gin.Context) {

}

func authUser(ctx *gin.Context) {
	u := new(model.SUser)
	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	if ok, err := user.Auth(u); err != nil || !ok {
		ctx.JSON(500, newHttpResp("认证失败", nil))
		return
	}
	ss := model.SetSession(*u, ctx.Request.UserAgent())
	ctx.SetCookie("ss_id", ss.Username+"_"+ss.ID.String(), 24*60*60, "/", "", false, false)
	ctx.JSON(200, newHttpResp("认证通过", u.Username))
	return
}

func logout(ctx *gin.Context) {
	u := new(model.SUser)
	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	_, _ = user.Logout(u)
	ctx.JSON(200, newHttpResp("退出登陆", nil))
	return
}
