// Package dto
// @Description
// @Author  Ymri  2024/5/8 21:35
// @Update 2024/5/8 21:35
package dto

import "time"

type TasksDTO struct {
	ID               uint
	TaskName         string
	TaskDescription  string
	TaskStartTime    time.Time
	TaskEndTime      time.Time
	TaskDeadline     time.Time
	TaskStatus       int // 0:未开始 1:进行中 2:已结束
	AccessPath       string
	TaskFolder       string
	ResubmissionPath string
	IsResubmission   int
}
