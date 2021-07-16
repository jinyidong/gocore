/**
* @Description: 排序
* @Author: jinyidong
* @Date: 2021/6/19
* @Version V1.0
 */
package orm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
默认降序，true or false
*/
func OrderDest(sort string, bl bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(clause.OrderByColumn{Column: clause.Column{Name: sort}, Desc: bl})
	}
}
