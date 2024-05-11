// Package vo
// @Description
// @Author  Ymri  2024/5/8 21:53
// @Update 2024/5/8 21:53
package vo

import "time"

type TasksVo struct {
	ID              uint64    `json:"id"`
	TaskName        string    `json:"taskName"`
	TaskDescription string    `json:"taskDescription"`
	TaskStartTime   time.Time `json:"taskStartTime"`
	TaskEndTime     time.Time `json:"taskEndTime"`
	TaskDeadline    time.Time `json:"taskDeadline"`
	TaskStatus      int       `json:"taskStatus"`
	AccessPath      string    `json:"accessPath"`
	MathRegulation  string    `json:"mathRegulation"`
	ConfigID        int       `gorm:"configId"`
}
