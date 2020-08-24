package router

import (
	"blogbk/model"
	"blogbk/module/tag"
	"github.com/gin-gonic/gin"
)

func tagRouter() {
	tr := Router.Group(api("/tags"))
	{
		tr.GET("/:name", getTag)
		tr.GET("", fetchTags)
	}
}

func getTag(ctx *gin.Context) {
	name := ctx.Param("name")
	if len(name) == 0 {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	var t = &model.STag{Name: name}
	if err := tag.Get(t); err != nil {
		ctx.JSON(500, newHttpResp("tag获取失败", nil))
		return
	}
	ctx.JSON(200, newHttpResp("tag获取成功", t))
	return
}

func fetchTags(ctx *gin.Context) {
	var ts = new(model.STags)
	if err := tag.Fetch(ts); err != nil {
		ctx.JSON(500, newHttpResp("获取tag列表失败", nil))
		return
	}
	ctx.JSON(200, newHttpResp("获取tag列表成功", ts))
	return
}
