// Package repositories
// @Description
// @Author  Ymri  2024/5/11 14:45
// @Update 2024/5/11 14:45
package repositories

import (
	"fmt"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/common/global"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/infrastructure/db/entity"
	"time"
)

var taskDetailInstance *TaskDetailRepository

type TaskDetailRepository struct{}

func GetTaskDetailRepository() *TaskDetailRepository {
	return taskDetailInstance
}

// GetTaskDetailByTaskID 根据任务ID查询任务详情
func (t *TaskDetailRepository) GetTaskDetailByTaskID(queryTask *vo.TaskDetailVO) []entity.TaskDetail {
	task := taskInstance.SelectById(utils.ToUInt64(queryTask.TaskID))
	var taskDetails []entity.TaskDetail
	tx := global.DB.Model(&entity.TaskDetail{})
	// 根据任务状态查询
	if queryTask.TaskStatus != -1 {
		tx = tx.Where("task_status = ?", queryTask.TaskStatus)
	}
	// 根据任务名称过滤，支持多个任务名称
	if queryTask.PreFileName != "" {
		tx = tx.Where("pre_file_name like ?", "%"+queryTask.PreFileName+"%")
	}
	tx = tx.Where("task_id = ?", task.ID)
	// 预加载User信息
	tx.Preload("User").Find(&taskDetails)
	// for 输出
	for _, taskDetail := range taskDetails {
		fmt.Printf("%s %d ", taskDetail.PreFileName, taskDetail.TaskStatus)

	}
	return taskDetails
}

// SelectById 查找单个任务详情通过ID
func (t *TaskDetailRepository) SelectById(id uint) entity.TaskDetail {
	var taskDetail entity.TaskDetail
	tx := global.DB.Model(&entity.TaskDetail{})
	tx.Where("id = ?", id)
	// Preload("Task") 后续如果需要补充再补
	tx.Preload("User").Find(&taskDetail)
	return taskDetail
}

// SelectByTaskIdAndFileName 通过任务ID和文件名查询
func (t *TaskDetailRepository) SelectByTaskIdAndFileName(taskId uint64, fileName string) entity.TaskDetail {
	var taskDetail entity.TaskDetail
	tx := global.DB.Model(&entity.TaskDetail{})
	tx.Where("task_id = ?", taskId).Where("file_name = ?", fileName)
	tx.Find(&taskDetail)
	return taskDetail
}

// UpdateByTaskDetail 更新单个任务详情 只可以更新状态
func (t *TaskDetailRepository) UpdateByTaskDetail(taskDetail vo.TaskDetailVO) {
	// 只更新有数据的
	updates := make(map[string]interface{})
	flag := false
	tempTaskDetail := convert.ToTaskDetailByVo(taskDetail)
	if taskDetail.ID >= 0 && taskDetail.TaskStatus != 0 {
		updates["task_status"] = tempTaskDetail.TaskStatus
		flag = true
	}
	// 修改预定义文件名名
	if taskDetail.PreFileName != "" {
		updates["pre_file_name"] = tempTaskDetail.PreFileName
		flag = true
	}
	if taskDetail.FileName != "" {
		updates["file_name"] = tempTaskDetail.FileName
		flag = true
	}
	if flag {
		tempTaskDetail.UpdatedAt = time.Now()
		global.DB.Model(&entity.TaskDetail{}).Where("id = ?", tempTaskDetail.ID).Updates(updates)
	}
}

// DeleteById 通过ID删除单条任务
func (t *TaskDetailRepository) DeleteById(id uint) {
	var taskDetails entity.TaskDetail
	global.DB.Delete(&taskDetails, id)
}

// DeleteByTaskId 根据TaskId 批量删除
func (t *TaskDetailRepository) DeleteByTaskId(id uint64) {
	global.DB.Where("tak_id = ?", id).Delete(&entity.TaskDetail{})
}

// InsertOne 插入数据
func (t *TaskDetailRepository) InsertOne(taskDetail vo.TaskDetailVO) {
	taskDetailEntity := convert.ToTaskDetailByVoNoID(taskDetail)
	taskDetailEntity.CreatedAt = time.Now()
	taskDetailEntity.UpdatedAt = time.Now()
	global.DB.Create(&taskDetailEntity)
}

// InsertBatch 批量插入
func (t *TaskDetailRepository) InsertBatch(taskDetails []vo.TaskDetailVO) {
	taskDetailEntities := make([]entity.TaskDetail, 0)
	for _, taskDetail := range taskDetails {
		// 插入的时候忽略ID
		taskDetailEntity := convert.ToTaskDetailByVo(taskDetail)
		taskDetailEntity.CreatedAt = time.Now()
		taskDetailEntity.UpdatedAt = time.Now()
		taskDetailEntities = append(taskDetailEntities, taskDetailEntity)
	}
	global.DB.Create(&taskDetailEntities)
}
