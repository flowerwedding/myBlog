/**
 * @Title  upload
 * @description  文件上传
 * @Author  沈来
 * @Update  2020/8/6 20:15
 **/
package api

import (
	"github.com/gin-gonic/gin"
	"myBlog/global"
	"myBlog/internal/service"
	"myBlog/pkg/app"
	"myBlog/pkg/convert"
	"myBlog/pkg/errcode"
	"myBlog/pkg/upload"
)

// @Summary  文件上传
// @Produce  json
// @Param  file body string true "文件"
// @Param  type body string true "类型"
// @Success  200 {object} app.Response "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /upload/file [post]
func Upload(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ERROR_UPLOAD_FILE_FAIL.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
