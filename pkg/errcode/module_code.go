/**
 * @Title  module_code
 * @description  业务错误码
 * @Author  沈来
 * @Update  2020/8/5 16:09
 **/
package errcode

var (
	ErrorGetTagListFail = NewError(10010, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(10011, "创建标签失败")
	ErrorUpdateTagFail  = NewError(10012, "更新标签失败")
	ErrorDeleteTagFail  = NewError(10013, "删除标签失败")
	ErrorCountTagFail   = NewError(10014, "统计标签失败")
)

var (
	ErrorGetArticleFail    = NewError(10020, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(10021, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(10022, "创建文章失败")
	ErrorUpdateArticleFail = NewError(10023, "更新文章失败")
	ErrorDeleteArticleFail = NewError(10024, "删除文章失败")
)

var (
	ERROR_UPLOAD_FILE_FAIL = NewError(10031, "上传文件失败")
)
