// Package appmanager
// @Description
// @Author  Ymri  2024/5/8 21:46
// @Update 2024/5/8 21:46
package appmanager

import (
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/service/impl"
	"goods-system/internal/service/param"
)

var tasksService = impl.GetTaskServiceInstance()

type TasksManager struct {
}

var tasksManager *TasksManager

func GetTasksManagerInstance() *TasksManager {
	return tasksManager
}

// Page 分页
func (p *TasksManager) Page(request request.TasksPageRequest) api.Page[vo.TasksVo] {
	param := convert.ToTasksPageParam(request)
	//分页查询
	page := tasksService.Page(param)
	data := page.Data
	var TaskList []vo.TasksVo
	for _, datum := range data {
		goodsVO := convert.ToTaskVO(datum)
		TaskList = append(TaskList, goodsVO)
	}
	return api.Page[vo.TasksVo]{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Data:  TaskList,
	}
}

// Save 保存任务
func (p *TasksManager) Save(req request.TasksSaveRequest) {
	param := convert.ToTaskSaveParam(req)
	tasksService.Save(param)
}

// Delete 删除商品
func (p *TasksManager) Delete(id string) {
	tasksService.Delete(utils.ToUInt64(id))
}

// All 获得所有商品
func (p *TasksManager) All() []vo.TasksVo {
	var tasksVOS []vo.TasksVo
	taskDTOS := tasksService.List(param.TasksParam{})
	for _, item := range taskDTOS {
		tasksVO := convert.ToTaskVO(item)
		tasksVOS = append(tasksVOS, tasksVO)
	}
	return tasksVOS
}

// GetById 根据id获得商品
func (p *TasksManager) GetById(id string) vo.TasksVo {
	taskDto := tasksService.GetById(utils.ToUInt64(id))
	return convert.ToTaskVO(taskDto)
}
