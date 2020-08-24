/**
 * @Title  smms
 * @description  smms图片存储
 * @Author  沈来
 * @Update  2020/8/23 15:20
 **/
package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"myBlog/internal/dao"
	"myBlog/pkg/upload"
	"net/http"
)

type SmmsUploader struct {
}

func (svc *Service) UploadSmms(file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	var (
		resp      *http.Response
		bodyBytes []byte
		ret       dao.SmmsRet
		bodyBuf   = &bytes.Buffer{}
	)
	fileName := upload.GetFileName(fileHeader.Filename)
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("smfile", fileHeader.Filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}
	_ = bodyWriter.Close()

	resp, err = http.Post("https://sm.ms/api/upload", bodyWriter.FormDataContentType(), bodyBuf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &ret)
	if err != nil {
		return nil, err
	}
	if ret.Code == "error" {
		err = errors.New(ret.Data.Msg)
		return nil, err
	}

	_ = svc.dao.UploadSmms(ret)

	url := ret.Data.Url
	return &FileInfo{Name: fileName, AccessUrl: url}, nil
}
