/**
* @Description: mysql查询语句生成
* @Author: jinyidong
* @Date: 2021/6/19
* @Version V1.0
 */
package orm

import (
	"gorm.io/gorm"
)

func MakeMysqlCondition(q interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		condition := &gormCondition{
			gormPublic: gormPublic{},
			Join:       make([]*gormJoin, 0),
		}
		resolveSearchQuery(Mysql, q, condition)
		for _, join := range condition.Join {
			if join == nil {
				continue
			}
			db = db.Joins(join.JoinOn)
			for k, v := range join.Where {
				db = db.Where(k, v...)
			}
			for k, v := range join.Or {
				db = db.Or(k, v...)
			}
			for _, o := range join.Order {
				db = db.Order(o)
			}
		}
		for k, v := range condition.Where {
			db = db.Where(k, v...)
		}
		for k, v := range condition.Or {
			db = db.Or(k, v...)
		}
		for _, o := range condition.Order {
			db = db.Order(o)
		}
		return db
	}
}
