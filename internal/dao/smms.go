/**
 * @Title  smms
 * @description  #
 * @Author  沈来
 * @Update  2020/8/23 15:42
 **/
package dao

import (
	"myBlog/internal/model"
)

type SmmsRet struct {
	Code string `json:"code"`
	Data struct {
		FileName  string `json:"filename"`
		StoreName string `json:"storename"`
		Size      int    `json:"size"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Hash      string `json:"hash"`
		Delete    string `json:"delete"`
		Url       string `json:"url"`
		Path      string `json:"path"`
		Msg       string `json:"msg"`
	} `json:"data"`
}

func (d *Dao) UploadSmms(ret SmmsRet) error {
	smmsFile := model.SmmsFile{ //
		FileName:  ret.Data.FileName,
		StoreName: ret.Data.StoreName,
		Size:      ret.Data.Size,
		Width:     ret.Data.Width,
		Height:    ret.Data.Height,
		Hash:      ret.Data.Hash,
		Delete:    ret.Data.Delete,
		Url:       ret.Data.Url,
		Path:      ret.Data.Path,
	}
	err := smmsFile.Insert(d.engine)
	if err != nil {
		return err
	}

	return nil
}
