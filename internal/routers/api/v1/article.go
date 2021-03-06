/**
 * @Title  article
 * @Description  文章处理
 * @Author  沈来
 * @Update  2020/8/3 16:17
 **/
package v1

import (
	"github.com/gin-gonic/gin"
	"myBlog/global"
	"myBlog/internal/service"
	"myBlog/pkg/app"
	"myBlog/pkg/convert"
	"myBlog/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

var false struct{}

// @Summary  获取单篇文章
// @Produce  json
// @Param  id path int true "文章ID"
// @Param  state query int false "状态" Enums(0, 1) default(1)
// @Success  200 {object} model.ArticleSwagger "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	param := service.ArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(article)
	return
}

// @Summary  获取多篇文章
// @Produce  json
// @Param  tag_id query uint32 true "标签ID"
// @Param  state query int false "状态" Enums(0, 1) default(1)
// @Param  page query int false "页码"
// @Param  page_size query int false "每页数量"
// @Success  200 {object} model.ArticleSwagger "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articles, totalRows, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}

	response.ToResponseList(articles, totalRows)
	return
}

// @Summary  新增文章
// @Produce  json
// @Param  tag_id query string true "标签id"
// @Param  state query int false "状态" Enums(0, 1) default(1)
// @Param  created_by query string true "创建者" minlength(2) maxlength(100)
// @Param  title query string true "标题" minlength(2) maxlength(100)
// @Param  desc query string true "详情" minlength(2) maxlength(255)
// @Param  content query string true "内容" minlength(2) maxlength(4294967295)
// @Param  cover_image_url query string true "封面URL"
// @Success  200 {object} model.Article "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(nil)
	return
}

// @Summary  更新文章
// @Produce  json
// @Param  id query string true "文章id"
// @Param  tag_id query string true "标签id"
// @Param  state query int false "状态" Enums(0, 1) default(1)
// @Param  modified_by query string true "创建者" minlength(2) maxlength(100)
// @Param  title query string false "标题" minlength(2) maxlength(100)
// @Param  desc query string false "详情" minlength(2) maxlength(255)
// @Param  content query string false "内容" minlength(2) maxlength(4294967295)
// @Param  cover_image_url query  string true "封面URL"
// @Success  200 {object} model.Article "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary  删除文章
// @Produce  json
// @Param  id query int true "文章ID"
// @Success  200 {string} string "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误""
// @Router  /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
