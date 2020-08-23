/**
 * @Title  tag
 * @Description  标签tag
 * @Author  沈来
 * @Update  2020/8/3 15:45
 **/
package model

import (
	"github.com/jinzhu/gorm"
	"myBlog/pkg/app"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		//别忘 ? 不然sql语句会拼接错误
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffSet, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffSet >= 0 && pageSize > 0 {
		db = db.Offset(pageOffSet).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) ListByIDs(db *gorm.DB, ids []uint32) ([]*Tag, error) {
	var tags []*Tag
	err := db.Where("state = ? AND is_del = ?", t.State, 0).Where("id IN (?)", ids).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? AND state = ?", t.ID, 0, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}

func (t Tag) GetByName(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("name = ? AND is_del = ? AND state = ?", t.Name, 0, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(t).Updates(values).Where("id = ? AND is_del = ?", t.ID, 0).Error; err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
