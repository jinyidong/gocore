/**
* @Description: 分页组件
* @Author: jinyidong
* @Date: 2021/6/19
* @Version V1.0
 */
package orm

import "gorm.io/gorm"

func Paginate(pageSize, pageIndex int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageIndex - 1) * pageSize
		if offset < 0 {
			offset = 0
		}
		return db.Offset(offset).Limit(pageSize)
	}
}
