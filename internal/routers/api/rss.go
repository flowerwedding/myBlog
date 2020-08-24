/**
 * @Title  rss
 * @description  #
 * @Author  沈来
 * @Update  2020/8/24 10:12
 **/
package api

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	"github.com/microcosm-cc/bluemonday"
	"html/template"
	"myBlog/global"
	"myBlog/internal/service"
	"myBlog/pkg/app"
	"time"
)

// @Summary  文章订阅
// @Produce  json
// @Param  tag_id query uint32 true "标签ID"
// @Param  state query int false "状态" Enums(0, 1) default(1)
// @Param  page query int false "页码"
// @Param  page_size query int false "每页数量"
// @Success  200 {object} string "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /api/v1/rss [get]
func RssGet(c *gin.Context) {
	now := time.Now()
	domain := global.RssSetting.Domain
	feed := &feeds.Feed{
		Title:       "myBlog",
		Link:        &feeds.Link{Href: domain},
		Description: "Wblog,talk about golang and so on.",
		Author:      &feeds.Author{Name: "shenlai", Email: global.EmailSetting.From},
		Created:     now,
	}

	svc := service.New(c.Request.Context())
	feed.Items = make([]*feeds.Item, 0)
	param := service.ArticleListRequest{}
	_, _ = app.BindAndValid(c, &param)
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	posts, _, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		_ = seelog.Error(err)
		return
	}

	for _, post := range posts {
		item := &feeds.Item{
			Id:          fmt.Sprintf("%s/post/%d", domain, post.ID),
			Title:       post.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("%s/post/%d", domain, post.ID)},
			Description: string(Excerpt(post)),
			Created:     now,
		}
		feed.Items = append(feed.Items, item)
	}
	rss, err := feed.ToRss()
	if err != nil {
		_ = seelog.Error(err)
		return
	}

	_, _ = c.Writer.WriteString(rss)
}

func Excerpt(post *service.Article) template.HTML {
	//you can sanitize, cut it down, add images, etc
	policy := bluemonday.StrictPolicy() //remove all html tags
	sanitized := policy.Sanitize(post.Content)
	runes := []rune(sanitized)
	if len(runes) > 300 {
		sanitized = string(runes[:300])
	}
	excerpt := template.HTML(sanitized + "...")
	return excerpt
}
