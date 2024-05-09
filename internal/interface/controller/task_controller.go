package controller

import (
	"fmt"
	"go/types"
	"goods-system/internal/application/appmanager"
	"goods-system/internal/application/request"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
)

var taskManager = appmanager.GetTasksManagerInstance()

type TasksController struct {
}

var taskController *TasksController

func GetTasksController() *TasksController {
	return taskController
}
func GetTaskManagerInstance() *TasksController {
	return taskController
}

// GoodsGetById 根据id获得商品
func (p *TasksController) TaskGetById(id string) api.RespData[vo.TasksVo] {
	taskvo := taskManager.GetById(id)
	return api.Success(taskvo, "")
}

// GoodsAll 所有商品
func (p *TasksController) TaskAll() api.RespData[[]vo.TasksVo] {
	var taskVOS []vo.TasksVo
	taskVOS = taskManager.All()
	return api.Success(taskVOS, "")
}

// GoodsPage 分页
func (p *TasksController) TaskPage(req request.TasksPageRequest) api.RespData[api.Page[vo.TasksVo]] {
	page := taskManager.Page(req)
	return api.Success[api.Page[vo.TasksVo]](page, "")
}

// GoodsSave 保存商品
func (p *TasksController) TaskSave(req request.TasksSaveRequest) api.RespData[types.Nil] {
	if req.TaskName == "" || req.TaskDescription == "" || req.AccessPath == "" {
		return api.Fail[types.Nil]("参数不全，检查输入是否为空")
	}
	taskManager.Save(req)
	return api.Success[types.Nil](types.Nil{}, "")
}

// GoodsSave 保存商品
func (p *TasksController) TaskUpdate(req request.TasksUpdateRequest) api.RespData[types.Nil] {
	fmt.Println(req)
	if req.TaskName == "" || req.TaskDescription == "" || req.AccessPath == "" {
		return api.Fail[types.Nil]("参数不全，检查输入是否为空")
	}
	taskManager.Update(req)
	return api.Success[types.Nil](types.Nil{}, "")
}

// GoodsDelete 删除商品
func (p *TasksController) TaskDelete(id string) api.RespData[types.Nil] {
	if id == "" {
		return api.Fail[types.Nil]("id不能为空")
	}
	taskManager.Delete(id)
	return api.Success[types.Nil](types.Nil{}, "")
}
