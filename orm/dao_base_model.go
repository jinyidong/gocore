/**
* @Description: 公用model信息
* @Author: jinyidong
* @Date: 2021/6/21
* @Version V1.0
 */
package orm

import (
	"gorm.io/gorm"
	"time"
)

type ModelID struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}
