// Package entity
// @Description
// @Author  Ymri  2024/5/11 14:29
// @Update 2024/5/11 14:29
package entity

import "gorm.io/gorm"

type TaskAndDetail struct {
	gorm.Model
	// 主键自增
	ID        uint `gorm:"primaryKey;autoIncrement"`
	TaskID    int  `gorm:"not null"`
	SubTaskID int  `gorm:"not null"`
}
