/**
 * @Title  service
 * @description  公共service
 * @Author  沈来
 * @Update  2020/8/5 15:52
 **/
package service

import (
	"context"
	"myBlog/global"
	"myBlog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
