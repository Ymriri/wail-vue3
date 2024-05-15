// Package impl
// @Description
// @Author  Ymri  2024/5/11 23:29
// @Update 2024/5/11 23:29
package impl

import (
	"fmt"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/infrastructure/db/repositories"
)

type TaskDetailService struct{}

var taskDetailService *TaskDetailService

var taskDetailsRepository = repositories.GetTaskDetailRepository()

var fileService = GetLoadFileService()

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

// GetTaskDetailByName 查询单个任务
func (t *TaskDetailService) GetTaskDetailByName(taskId uint64, fileName string) vo.TaskDetailVO {
	taskDetail := taskDetailsRepository.SelectByTaskIdAndFileName(taskId, fileName)
	return convert.ToTaskDetailVoByDetail(taskDetail)
}

// Save 保存任务
func (t *TaskDetailService) Save(param vo.TaskDetailVO) {
	fmt.Println(param)
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

func (t *TaskDetailService) SaveToExcel(detailVO vo.TaskDetailVO) (string, error) {
	data := t.GetTaskDetailByTaskId(detailVO)
	// 移除异常文件
	for i := 0; i < len(data); i++ {
		if data[i].TaskStatus == 3 {
			data = append(data[:i], data[i+1:]...)
		}
	}
	//查询出task
	taskSetting := tasksRepository.SelectById(utils.ToUInt64(detailVO.TaskID))
	taskVo := convert.ToTaskVO(convert.ToTasksDTO(taskSetting))
	fmt.Println(taskVo)
	fileName, err := fileService.WriteUserExcel(data, taskVo)
	return fileName, err
}
