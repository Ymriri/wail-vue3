// Package appmanager
// @Description
// @Author  Ymri  2024/5/14 09:11
// @Update 2024/5/14 09:11
package appmanager

import (
	"goods-system/internal/application/vo"
	"goods-system/internal/service/impl"
)

var taskDetailManagerInstance *TaskDetailManager

type TaskDetailManager struct{}

var taskDetailService = impl.GetTaskDetailServiceInstance()

func GetTaskDetailManagerInstance() *TaskDetailManager {
	return taskDetailManagerInstance
}

// GetTaskDetailByTask 查询任务下所有的子任务
func (t *TaskDetailManager) GetTaskDetailByTask(param vo.TaskDetailVO) []vo.TaskDetailVO {
	return taskDetailService.GetTaskDetailByTaskId(param)
}

// Save 保存任务
func (t *TaskDetailManager) Save(param vo.TaskDetailVO) {
	taskDetailService.Save(param)
}

// Delete 删除任务
func (t *TaskDetailManager) Delete(id uint) {
	taskDetailService.Delete(id)
}

// BatchInsert 批量插入
func (t *TaskDetailManager) BatchInsert(param []vo.TaskDetailVO) {
	taskDetailService.BatchInsert(param)
}

// DeleteByTaskId 根据任务id删除
func (t *TaskDetailManager) DeleteByTaskId(id uint64) {
	taskDetailService.DeleteByTaskId(id)
}

// Update 更新任务
func (t *TaskDetailManager) Update(param vo.TaskDetailVO) {
	taskDetailService.Update(param)
}
