/**
 * @Title  auth
 * @description  JWT
 * @Author  沈来
 * @Update  2020/8/6 22:08
 **/
package service

import "errors"

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(
		param.AppKey,
		param.AppSecret,
	)
	if err != nil {
		return nil
	}
	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
