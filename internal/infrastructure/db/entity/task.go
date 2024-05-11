// Package entity
// @Description
// @Author  Ymri  2024/5/8 20:44
// @Update 2024/5/8 20:44
package entity

import (
	"gorm.io/gorm"
	"time"
)

type TaskSettings struct {
	gorm.Model
	ID               uint      `gorm:"primary_key;AUTO_INCREMENT"`
	TaskName         string    `gorm:"type:text;NOT NULL"`
	TaskDescription  string    `gorm:"type:text"`
	TaskStartTime    time.Time `gorm:"type:datetime"`
	TaskEndTime      time.Time `gorm:"type:datetime"`
	TaskDeadline     time.Time `gorm:"type:datetime"`
	TaskStatus       int       `gorm:"DEFAULT:0"` // 0:未开始 1:进行中 2:已结束
	AccessPath       string    `gorm:"type:text"`
	TaskFolder       string    `gorm:"type:text"`
	ResubmissionPath string    `gorm:"type:text"`
	MathRegulation   string    `gorm:"type:text"`
	ConfigID         int       `gorm:"type:int;"`
	IsResubmission   int       `gorm:"DEFAULT:0"`
}
