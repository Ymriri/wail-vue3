// Package impl
// @Description
// @Author  Ymri  2024/5/11 23:29
// @Update 2024/5/11 23:29
package impl

import (
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/db/repositories"
)

type TaskDetailService struct{}

var taskDetailService *TaskDetailService

var taskDetailsRepository = repositories.GetTaskDetailRepository()

func GetTaskDetailServiceInstance() *TaskDetailService {
	return taskDetailService
}

// GetTaskDetailByTaskId 查询任务下所有的子任务
func (t *TaskDetailService) GetTaskDetailByTaskId(param vo.TaskDetailVO) []vo.TaskDetailVO {
	taskDetailList := taskDetailsRepository.GetTaskDetailByTaskID(&param)
	var taskDetailVoList []vo.TaskDetailVO
	for _, e := range taskDetailList {
		taskDetailVo := convert.ToTaskDetailVoByDetail(e)
		taskDetailVoList = append(taskDetailVoList, taskDetailVo)
	}
	return taskDetailVoList
}

// Save 保存任务
func (t *TaskDetailService) Save(param vo.TaskDetailVO) {
	taskDetailsRepository.InsertOne(param)
}

// Delete 删除任务
func (t *TaskDetailService) Delete(id uint) {
	taskDetailsRepository.DeleteById(id)
}

// DeleteByTaskId 根据任务id删除
func (t *TaskDetailService) DeleteByTaskId(id uint64) {
	taskDetailsRepository.DeleteByTaskId(id)
}

// BatchInsert 批量插入
func (t *TaskDetailService) BatchInsert(param []vo.TaskDetailVO) {
	taskDetailsRepository.InsertBatch(param)
}

// Update 更新任务
func (t *TaskDetailService) Update(param vo.TaskDetailVO) {
	taskDetailsRepository.UpdateByTaskDetail(param)
}
