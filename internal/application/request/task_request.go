// Package request
// @Description
// @Author  Ymri  2024/5/8 21:52
// @Update 2024/5/8 21:52
package request

type TasksPageRequest struct {
	//商品名称
	TaskName string `json:"taskName"`
	//商品类型
	TasksStatus int `json:"tasksStatus"`
	//当前页
	Page int `json:"page"`
	//页面大小
	Size int `json:"size"`
}

type TasksSaveRequest struct {
	//任务名称
	TaskName string `json:"taskName"`
	//商品类型
	TaskDescription string `json:"taskDescription"`
	TaskStartTime   string `json:"taskStartTime"`
	TaskEndTime     string `json:"taskEndTime"`
	TaskDeadline    string `json:"taskDeadline"`
	TaskStatus      int    `json:"taskStatus"` // 0:未开始 1:进行中 2:已结束
	AccessPath      string `json:"accessPath"`
}
