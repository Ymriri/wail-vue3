// Package param
// @Description
// @Author  Ymri  2024/5/8 21:27
// @Update 2024/5/8 21:27
package param

import "time"

// TaskParam 商品参数
type TasksParam struct {
	//商品名称
	TaskName string
	//商品编号
	ID         string
	TaskStatus int
}

// TaskseParam 删除
type TasksReduceParam struct {
	//id
	Id uint64
}

// TasksageParam 商品分页参数
type TasksPageParam struct {
	//商品名称
	TasksName string
	//商品类型
	TasksStatus int
	//当前页
	Page int
	//页面大小
	Size int
}

// tasksSaveParam 商品保存
type TasksSaveParam struct {
	ID              uint      `json:"id"`
	TaskName        string    `json:"taskName"`
	TaskDescription string    `json:"taskDescription"`
	TaskStartTime   time.Time `json:"taskStartTime"`
	TaskEndTime     time.Time `json:"taskEndTime"`
	TaskDeadline    time.Time `json:"taskDeadline"`
	TaskStatus      int       `json:"taskStatus"` // 0:未开始 1:进行中 2:已结束
	AccessPath      string    `json:"accessPath"`
	MathRegulation  string    `json:"mathRegulation"`
	ConfigID        int       `json:"configID"`
}
