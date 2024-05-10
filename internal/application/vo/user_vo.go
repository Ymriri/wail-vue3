// Package vo
// @Description
// @Author  Ymri  2024/5/9 23:24
// @Update 2024/5/9 23:24
// 页面显示用户信息
package vo

type UserPageVo struct {
	// 姓名
	Name string `json:"name"`
	// 分组
	ConfigFileID int `json:"configFileID"`
	Page         int `json:"page"`
	Size         int `json:"size"`
}

type UserVo struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Grade          string `json:"grade"`
	Class          string `json:"class"`
	Department     string `json:"department"`
	Group          string `json:"group"`
	EmployeeNumber string `json:"employeeNumber"`
	ConfigFileID   int    `json:"configFileId"`
	ConfigName     string `json:"configName"`
}
