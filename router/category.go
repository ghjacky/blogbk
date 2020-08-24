package router

import (
	"blogbk/model"
	"blogbk/module/category"
	"github.com/gin-gonic/gin"
	"strconv"
)

func categoryRouter() {
	cr := Router.Group(api("/categories"))
	{
		cr.GET("", fetchCategories)
		cr.POST("", addCategory)
		cr.PUT("", updateCategory)
		cr.GET("/:id", getCategory)
		cr.DELETE("/:id", deleteCategory)
	}
}

func fetchCategories(ctx *gin.Context) {
	tree, err := strconv.Atoi(ctx.Query("tree"))
	if err != nil || (tree != 0 && tree != 1) {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	var categories = new(model.SCategories)
	if err := category.Fetch(categories, tree); err != nil {
		ctx.JSON(500, newHttpResp("分类查询出错", nil))
		return
	}
	ctx.JSON(200, newHttpResp("分类查询成功", categories))
	return
}

func addCategory(ctx *gin.Context) {
	var cat = new(model.SCategory)
	if err := ctx.BindJSON(cat); err != nil {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	if err := category.Add(cat); err != nil {
		ctx.JSON(500, newHttpResp("分类添加出错", nil))
		return
	}
	ctx.JSON(200, newHttpResp("分类添加成功", cat))
	return
}

func updateCategory(ctx *gin.Context) {
	var cat = new(model.SCategory)
	if err := ctx.BindJSON(cat); err != nil {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	if err := category.Update(cat); err != nil {
		ctx.JSON(500, newHttpResp("分类更新错误", nil))
		return
	}
	ctx.JSON(200, newHttpResp("分类更新成功", cat))
	return
}

func deleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	var cat = new(model.SCategory)
	cat.ID = uint(id)
	if err := category.Delete(cat); err != nil {
		ctx.JSON(500, newHttpResp("分类删除失败", nil))
		return
	}
	ctx.JSON(200, newHttpResp("分类删除成功", cat))
	return
}

func getCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		ctx.JSON(400, newHttpResp("参数错误", nil))
		return
	}
	var cat = new(model.SCategory)
	cat.ID = uint(id)
	if err := category.Get(cat); err != nil {
		ctx.JSON(500, newHttpResp("获取分类信息失败", nil))
		return
	}
	ctx.JSON(200, newHttpResp("获取分类信息成功", cat))
	return
}
