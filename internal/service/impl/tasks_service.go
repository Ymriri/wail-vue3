// Package impl
// @Description
// @Author  Ymri  2024/5/8 21:22
// @Update 2024/5/8 21:22
package impl

import (
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/infrastructure/db/repositories"
	"goods-system/internal/service/dto"
	"goods-system/internal/service/param"
)

type TasksService struct {
}

var tasksService *TasksService
var tasksRepository = repositories.GetTaskRespositoryInstance()

func GetTaskServiceInstance() *TasksService {
	return tasksService
}

// Page 分页
func (p *TasksService) Page(param param.TasksPageParam) api.Page[dto.TasksDTO] {
	tasks := convert.ToTasksByPageParam(param)
	page := tasksRepository.Page(&tasks, param.Page, param.Size)
	//数据库查询任务列表
	taksList := page.Data
	//var goodsDTOList []dto.GoodsDTO
	var tasksDTOList []dto.TasksDTO
	for _, e := range taksList {
		taskDTO := convert.ToTasksDTO(e)
		tasksDTOList = append(tasksDTOList, taskDTO)
	}
	return api.Page[dto.TasksDTO]{
		Page:  page.Page,
		Size:  page.Size,
		Total: page.Total,
		Data:  tasksDTOList,
	}
}

// Save 保存任务
func (p *TasksService) Save(param param.TasksSaveParam) dto.TasksDTO {
	entity := convert.ToTasksSaveParam(param)
	tasksRepository.Insert(&entity)
	return convert.ToTasksDTO(entity)
}

// Delete 删除任务
func (p *TasksService) Delete(id uint64) {
	tasksRepository.DeleteById(id)
}

// Update 更新任务
func (p *TasksService) Update(param param.TasksSaveParam) dto.TasksDTO {
	entity := convert.ToTasksBySaveParam(param)
	tasksRepository.UpdateById(&entity)
	return convert.ToTasksDTO(entity)
}

// List 获得所有商品
func (p *TasksService) List(param param.TasksParam) []dto.TasksDTO {
	var tasksDTOS []dto.TasksDTO
	goods := convert.ToTasksByParam(param)
	goodsList := tasksRepository.Select(goods)
	for _, item := range goodsList {
		goodsDTO := convert.ToTasksDTO(item)
		tasksDTOS = append(tasksDTOS, goodsDTO)
	}
	return tasksDTOS
}

// GetById 根据id获得商品
func (p *TasksService) GetById(id uint64) dto.TasksDTO {
	tasks := entity.TaskSettings{ID: uint(id)}
	tasksRepository.Get(&tasks)
	return convert.ToTasksDTO(tasks)
}
