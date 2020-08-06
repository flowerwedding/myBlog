/**
 * @Title  md5
 * @description  上传文件格式化
 * @Author  沈来
 * @Update  2020/8/6 19:13
 **/
package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
