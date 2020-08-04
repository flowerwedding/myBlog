/**
 * @Title  tag
 * @Description  标签tag
 * @Author  沈来
 * @Update  2020/8/3 15:45
 **/
package model

type Tag struct{
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (a Tag) TableName() string{
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	//Pager *app.Pager
}