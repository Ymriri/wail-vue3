// Package repositories
// @Description
// @Author  Ymri  2024/5/8 20:48
// @Update 2024/5/8 20:48
package repositories

import (
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/global"
	"goods-system/internal/infrastructure/db/entity"
	"time"
)

var taskInstance *TaskRepository

type TaskRepository struct {
}

func GetTaskRespositoryInstance() *TaskRepository {
	return taskInstance
}

// Page 分页
func (p *TaskRepository) Page(tasks *entity.TaskSettings, page, size int) api.Page[entity.TaskSettings] {
	var total int64
	var taskList []entity.TaskSettings
	tx := global.DB.Model(&entity.TaskSettings{})
	if tasks.TaskName != "" {
		// gorm 模糊查询
		tx.Where("task_name like ?", "%"+tasks.TaskName+"%")
	}
	// -1 表示查询所有
	if tasks.TaskStatus != -1 {
		tx.Where("task_status = ?", tasks.TaskStatus)
	}
	tx.Count(&total).
		Limit(size).
		Offset((page - 1) * size).
		Find(&taskList)

	return api.Page[entity.TaskSettings]{
		Page:  page,
		Size:  size,
		Total: total,
		Data:  taskList,
	}
}

// Insert 保存
func (p *TaskRepository) Insert(tasks *entity.TaskSettings) {
	global.DB.Create(tasks)
}

// SelectById 通过id获得商品
func (p *TaskRepository) SelectById(id uint64) entity.TaskSettings {
	var task entity.TaskSettings
	global.DB.First(&task, id)
	return task
}

// DeleteById 根据id删除
func (p *TaskRepository) DeleteById(id uint64) {
	var task entity.TaskSettings
	global.DB.Delete(&task, id)
}

// UpdateById 更新
func (p *TaskRepository) UpdateById(task *entity.TaskSettings) {
	flag := false
	//查询原来的数据
	old := p.SelectById(uint64(task.ID))
	// 创建一个只包含有变化字段的 map
	updates := make(map[string]interface{})
	if task.TaskName != old.TaskName {
		updates["task_name"] = task.TaskName
		flag = true
	}
	if task.TaskStatus != old.TaskStatus {
		updates["task_status"] = task.TaskStatus
		flag = true
	}

	if task.TaskDeadline != old.TaskDeadline {
		updates["task_deadline"] = task.TaskDeadline
		flag = true
	}
	if task.TaskStartTime != old.TaskStartTime {
		updates["task_start_time"] = task.TaskStartTime
		flag = true
	}
	if task.TaskDescription != old.TaskDescription {
		updates["task_description"] = task.TaskDescription
		flag = true
	}
	if task.AccessPath != old.AccessPath {
		updates["access_path"] = task.AccessPath
		flag = true
	}
	if task.TaskEndTime != old.TaskEndTime {
		updates["task_end_time"] = task.TaskEndTime
		flag = true
	}
	if flag {
		updates["updated_at"] = time.Now()
		// goods为新属性
		global.DB.Model(&task).Updates(updates)
	}
}

// Select 获得商品
func (p *TaskRepository) Select(task entity.TaskSettings) []entity.TaskSettings {
	tx := global.DB.Model(&task)
	var taskList []entity.TaskSettings
	if task.TaskName != "" {
		tx.Where("name like ?", "%"+task.TaskName+"%")
	}
	tx.Find(&taskList)
	return taskList
}

// Get 获得一个商品
func (p *TaskRepository) Get(task *entity.TaskSettings) {
	tx := global.DB.Model(task)
	if task.ID != 0 {
		tx.Where("id = ?", task.ID)
	}
	if task.TaskName != "" {
		tx.Where("task_name = ?", task.TaskName)
	}
	tx.First(task)
}

// BatchSave 批量保存
func (p *TaskRepository) BatchSave(goods []entity.Goods) {
	for _, item := range goods {
		global.DB.Save(&item)
	}
}
