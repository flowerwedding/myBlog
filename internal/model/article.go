/**
 * @Title  article
 * @Description  文章article
 * @Author  沈来
 * @Update  2020/8/3 15:50
 **/
package model

type Article struct{
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string{
	return "blog_article"
}

type ArticleSwagger struct {
	List  []*Article
	//Pager *app.Pager
}