/**
 * @Title  qiniu
 * @description  七牛云图片存储
 * @Author  沈来
 * @Update  2020/8/23 15:20
 **/
package service

/*
import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"mime/multipart"
	"myBlog/global"
	"myBlog/pkg/upload"
	"os"
)

// 获取文件大小的接口
type Size interface {
	Size() int64
}

// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type QiniuUploader struct {
}

func (svc *Service) UploadQiniu(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	accessKey := global.QiniuSetting.AccessKey
	secretKey := global.QiniuSetting.SecretKey
	bucket := global.QiniuSetting.Bucket

	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	localfile := uploadSavePath + "/" + fileName              //文件原来路径
	key := global.AppSetting.UploadServerUrl + "/" + fileName //目标路径

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localfile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	url := global.QiniuSetting.FileServer + ret.Key
	return &FileInfo{Name: fileName, AccessUrl: url}, nil
}*/
