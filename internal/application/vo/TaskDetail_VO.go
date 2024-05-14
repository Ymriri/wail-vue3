// Package vo
// @Description
// @Author  Ymri  2024/5/11 14:51
// @Update 2024/5/11 14:51
package vo

type TaskDetailVO struct {
	ID uint `json:"id"`
	// 用户信息
	User UserVo `json:"user"`
	// 任务ID
	TaskID int `json:"taskID"`
	// 任务状态
	TaskStatus int `json:"taskStatus"`
	// 生成的文件名
	PreFileName string `json:"preFileName"`
	// 上传的文件名
	FileName string `json:"fileName"`
	// 用户ID
	UserId uint `json:"userId"`
	// 分组名称
	TaskConfigName string `json:"taskConfigName"`
}
