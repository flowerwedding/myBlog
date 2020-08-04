/**
 * @Title  article_tag
 * @Description  文章标签article_tag
 * @Author  沈来
 * @Update  2020/8/3 16:02
 **/
package model

type ArticleTag struct{
	*Model
	TagID   uint32  `json:"tag_id"`
	Article uint32  `json:"article_tag"`
}

func (a ArticleTag) TableName() string{
	return "blog_article_tag"
}