// Package appmanager
// @Description
// @Author  Ymri  2024/5/14 09:11
// @Update 2024/5/14 09:11
package appmanager

import (
	"errors"
	"github.com/labstack/gommon/log"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/utils"
	"goods-system/internal/service/impl"
	"path/filepath"
	"regexp"
	"strings"
)

var taskDetailManagerInstance *TaskDetailManager

type TaskDetailManager struct{}

var taskDetailService = impl.GetTaskDetailServiceInstance()

var fileUtils = impl.GetFtpScannerServiceInstance()

func GetTaskDetailManagerInstance() *TaskDetailManager {
	return taskDetailManagerInstance
}

// SaveToExcel save to excel
func (t *TaskDetailManager) SaveToExcel(param vo.TaskDetailVO) (string, error) {
	return taskDetailService.SaveToExcel(param)
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

// CheckFileList check file
func (t *TaskDetailManager) CheckFileList(taskId int) error {
	// 查询出该任务
	taskDetailVOList := taskDetailService.GetTaskDetailByTaskId(vo.TaskDetailVO{TaskID: taskId})
	taskVo := tasksManager.GetById(utils.ToUInt64(taskId))
	// 任务文件
	// 检查该目录下所有文件
	successTaskVo := make([]vo.TaskDetailVO, 0)
	noTaskVo := make([]vo.TaskDetailVO, 0)
	if &taskVo != nil {
		// scanner ftp filePath
		fileList := fileUtils.ScanFtpFile(taskVo.AccessPath)
		if fileList == nil {
			// 抛出异常
			return errors.New("文件列表为空")
		}
		//fmt.Println(taskVo.AccessPath)
		// 输出文件列表
		//for _, file := range fileList {
		//	log.Printf("File Name: %s\n", file.Name)
		//}
		// 对比任务和查询出来的目录
		for _, taskDetailVO := range taskDetailVOList {
			flag := false
			for _, file := range fileList {
				if t.FileCHECK(taskDetailVO.PreFileName, file.Name) {
					// 匹配成功
					taskDetailVO.FileName = file.Name
					// 修改状态为已提交
					taskDetailVO.TaskStatus = 1
					successTaskVo = append(successTaskVo, taskDetailVO)
					t.Update(taskDetailVO)
					flag = true
					break
				}
			}
			if !flag {
				// 匹配失败
				taskDetailVO.TaskStatus = 2
				noTaskVo = append(noTaskVo, taskDetailVO)
				t.Update(taskDetailVO)
			}
		}
		// 找到未匹配成功的文件名
		for _, file := range fileList {
			flag := false
			for _, taskDetailVO := range successTaskVo {
				if file.Name == taskDetailVO.FileName {
					flag = true
					break
				}
			}
			if flag == false {
				tempTaskDetail := vo.TaskDetailVO{
					FileName:   file.Name,
					TaskStatus: 3,
				}
				// 检查是否保存过异常文件
				taskDetail := taskDetailService.GetTaskDetailByName(taskVo.ID, file.Name)
				if taskDetail.ID == 0 {
					// 保存异常文件
					tempTaskDetail.TaskID = utils.ToInt(taskVo.ID)
					t.Save(tempTaskDetail)
				}
			}
		}
	}
	return nil

}

// FileCHECK 对比文件
func (t *TaskDetailManager) FileCHECK(preFile string, fileName string) bool {
	// filename去拓展名
	ext := filepath.Ext(fileName)                       // 提取文件扩展名
	fileNameNoeExt := strings.TrimSuffix(fileName, ext) // 去除扩展名
	// 检查preFile中是否有相同的拓展名
	if strings.Contains(preFile, ext) {
		// 去除preFile中的拓展名
		preFile = strings.TrimSuffix(preFile, ext)
	}

	// 去除文件名中空格和回车
	fileNameNoeExt = strings.Replace(fileNameNoeExt, " ", "", -1)
	fileNameNoeExt = strings.Replace(fileNameNoeExt, "\n", "", -1)
	preFile = strings.Replace(preFile, " ", "", -1)
	preFile = strings.Replace(preFile, "\n", "", -1)
	// 检查 preFile中是否存在{__}
	// 使用正则进行匹配
	pattern := strings.Replace(preFile, "{", "", -1)
	pattern = strings.Replace(pattern, "}", "", -1)
	pattern = strings.Replace(pattern, "**", ".*", -1)

	// 添加开头和结尾的匹配条件，以防止过度匹配
	pattern = "^" + pattern + "$"

	// 编译正则表达式
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Printf("Pattern compile failed: %v\n", err)
		return false
	}

	// 使用正则表达式进行匹配
	isMatch := re.MatchString(fileNameNoeExt)
	if isMatch {
		//fmt.Println("Matched")
		return true
	} else {
		//fmt.Println("Not matched")
		return false
	}
}
