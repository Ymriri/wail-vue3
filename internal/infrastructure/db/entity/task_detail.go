// Package entity
// @Description
// @Author  Ymri  2024/5/11 14:31
// @Update 2024/5/11 14:31
package entity

import "gorm.io/gorm"

type TaskDetail struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	TaskID      uint64 `gorm:"not null"`  // 科组信息查询一次，然后手动插入
	TaskStatus  int    `gorm:"DEFAULT:0"` // 0:未完成 1:已完成 2: 迟交
	PreFileName string `gorm:"type:text"` // 生成的文件名
	FileName    string `gorm:"type:text"` // 上传的文件名
	UserId      int    `gorm:"type:int"`  // 用户ID
	// 检查时间
	// UserId 外键
	User User `gorm:"foreignKey:UserId;references:ID"`
	// TaskId 外键
	Task TaskSettings `gorm:"foreignKey:TaskID;references:ID"`
}
