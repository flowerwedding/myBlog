/**
 * @Title  smms
 * @description  #
 * @Author  沈来
 * @Update  2020/8/23 15:30
 **/
package model

import "github.com/jinzhu/gorm"

type SmmsFile struct {
	*Model
	FileName  string `json:"filename"`
	StoreName string `json:"storename"`
	Size      int    `json:"size"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Hash      string `json:"hash"`
	Delete    string `json:"delete"`
	Url       string `json:"url"`
	Path      string `json:"path"`
}

func (sf SmmsFile) Insert(db *gorm.DB) (err error) {
	return db.Create(&sf).Error
}
