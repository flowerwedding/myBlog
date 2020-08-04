/**
 * @Title  db
 * @Description  数据库连接
 * @Author  沈来
 * @Update  2020/8/3 21:19
 **/
package global

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"

var (
	DBEngine *gorm.DB
)