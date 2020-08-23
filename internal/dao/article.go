/**
 * @Title  article
 * @description  文章article
 * @Author  沈来
 * @Update  2020/8/6 9:52
 **/
package dao

import (
	"errors"
	"myBlog/internal/model"
	"myBlog/pkg/app"
)

func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
	}

	aa, _ := article.GetByTiTle(d.engine)
	if aa.ID != 0 {
		return nil, errors.New("this tag is exist")
	}

	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}

	return article.Update(d.engine, values)
}

func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}, State: state}
	return article.Get(d.engine)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}

func (d *Dao) CountArticleListByTagID(id uint32, state uint8) (int, error) {
	article := model.Article{State: state}
	return article.CountByTagID(d.engine, id)
}

func (d *Dao) GetArticleListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{State: state}
	return article.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}
