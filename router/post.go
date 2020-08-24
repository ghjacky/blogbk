package router

import (
	"blogbk/model"
	"blogbk/module/post"
	"github.com/gin-gonic/gin"
	"strconv"
)

func postRouter() {
	pr := Router.Group(api("/posts"))
	{
		pr.GET("", fetchList)
		pr.POST("", addPost)
		pr.PUT("", updatePost)
		pr.GET("/:id", getPost)
		pr.DELETE("/:id", deletePost)
	}
}

func fetchList(ctx *gin.Context) {
	var query = new(model.SQuery)
	if err := ctx.BindQuery(query); err != nil {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	var ps = new(model.SPosts)
	if err := post.Fetch(ps, *query); err != nil {
		ctx.JSON(500, newHttpResp("获取文章列表失败", nil))
		return
	}
	ctx.JSON(200, newHttpResp("获取文章列表成功", ps))
	return
}

func addPost(ctx *gin.Context) {
	var p = new(model.SPost)
	if err := ctx.BindJSON(p); err != nil {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	if err := post.Add(p); err != nil {
		ctx.JSON(500, newHttpResp("新增文章失败", nil))
		return
	}
	ctx.JSON(200, newHttpResp("新增文章成功", p))
	return
}

func updatePost(ctx *gin.Context) {
	var p = new(model.SPost)
	if err := ctx.BindJSON(p); err != nil {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	if err := post.Update(p); err != nil {
		ctx.JSON(500, newHttpResp("文章更新失败", nil))
		return
	}
	ctx.JSON(200, newHttpResp("文章更新成功", p))
	return
}

func getPost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	var p = new(model.SPost)
	p.ID = uint(id)
	if err := post.Get(p); err != nil {
		ctx.JSON(500, newHttpResp("文章获取失败", nil))
		return
	}
	ctx.JSON(200, newHttpResp("文章获取成功", p))
	return
}

func deletePost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	var p = new(model.SPost)
	p.ID = uint(id)
	if err := post.Delete(p); err != nil {
		ctx.JSON(500, newHttpResp("文章删除失败", nil))
		return
	}
	ctx.JSON(200, newHttpResp("文章删除成功", p))
	return
}
