// Package entity
// @Description
// @Author  Ymri  2024/5/8 20:39
// @Update 2024/5/8 20:39
package entity

import (
	"gorm.io/gorm"
)

type ConfigTree struct {
	gorm.Model
	// 配置主键
	ID uint `gorm:"primary_key;AUTO_INCREMENT"`
	// 父级ID，0为顶级
	ParentID uint `gorm:"type:int(11);NOT NULL"`
	// 配置名称
	ConfigName string `gorm:"type:text;NOT NULL"`
	// 配置值
	ConfigValue string `gorm:"type:text"`
}
