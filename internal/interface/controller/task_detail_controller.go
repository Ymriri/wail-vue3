// Package controller
// @Description
// @Author  Ymri  2024/5/14 09:15
// @Update 2024/5/14 09:15
package controller

import (
	"github.com/labstack/gommon/log"
	"go/types"
	"goods-system/internal/application/appmanager"
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/utils"
)

type TaskDetailController struct{}

var taskDetailManager = appmanager.GetTaskDetailManagerInstance()
var taskDetail *TaskDetailController

func GetTaskDetailControllerInstance() *TaskDetailController {
	return taskDetail
}

// save to excel
func (t *TaskDetailController) SaveToExcel(param vo.TaskDetailVO) api.RespData[types.Nil] {
	fileName, err := taskDetailManager.SaveToExcel(param)
	if err != nil {
		log.Error(err)
		return api.Fail[types.Nil]("保存到excel失败！")
	}
	return api.Success[types.Nil](types.Nil{}, "保存到本地 "+fileName+" 成功！")
}

// GetTaskDetailByTask 查询任务下所有的子任务
func (t *TaskDetailController) GetTaskDetailByTask(param vo.TaskDetailVO) api.RespData[[]vo.TaskDetailVO] {
	//fmt.Println(param)
	return api.Success[[]vo.TaskDetailVO](taskDetailManager.GetTaskDetailByTask(param), "查询成功！")
}

// ScannerFileSys begin scanner file sys
func (t *TaskDetailController) ScannerFileSys(param request.TasksUpdateRequest) api.RespData[types.Nil] {
	err := taskDetailManager.CheckFileList(utils.ToInt(param.ID))
	// 把任务设置扫描完成
	param.TaskStatus = 2
	taskManager.Update(param)
	if err != nil {
		return api.Fail[types.Nil]("扫描失败！请检查路径是否正确！")
	}
	return api.Success[types.Nil](types.Nil{}, "扫描完成！")
}

// Save 保存任务
func (t *TaskDetailController) Save(param vo.TaskDetailVO) api.RespData[types.Nil] {
	taskDetailManager.Save(param)
	return api.Success[types.Nil](types.Nil{}, "保存成功！")
}

// Delete 删除任务
func (t *TaskDetailController) Delete(id uint) api.RespData[types.Nil] {
	taskDetailManager.Delete(id)
	return api.Success[types.Nil](types.Nil{}, "删除成功！")
}

// BatchInsert 批量插入
func (t *TaskDetailController) BatchInsert(param []vo.TaskDetailVO) api.RespData[types.Nil] {
	taskDetailManager.BatchInsert(param)
	return api.Success[types.Nil](types.Nil{}, "批量插入成功！")
}

// DeleteByTaskId 根据任务id删除
func (t *TaskDetailController) DeleteByTaskId(id uint64) api.RespData[types.Nil] {
	taskDetailManager.DeleteByTaskId(id)
	return api.Success[types.Nil](types.Nil{}, "删除成功！")
}

// Update 更新任务
func (t *TaskDetailController) Update(param vo.TaskDetailVO) api.RespData[types.Nil] {
	taskDetailManager.Update(param)
	return api.Success[types.Nil](types.Nil{}, "更新成功！")
}

func (t *TaskDetailController) FileCheck(taskId int) api.RespData[types.Nil] {

	return api.Success[types.Nil](types.Nil{}, "文件校验成功！")
}
