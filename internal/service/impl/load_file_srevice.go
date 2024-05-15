// Package impl
// @Description
// @Author  Ymri  2024/5/10 19:42
// @Update 2024/5/10 19:42
package impl

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"goods-system/internal/application/vo"
	"os"
)

var loadFileService = &LoadFileService{}

type LoadFileService struct{}

func GetLoadFileService() *LoadFileService {
	return loadFileService
}

// ReadUserExcel 读取ReadUserExcel 文件
func (utils *LoadFileService) ReadUserExcel(filePath string) []vo.UserVo {
	// 读取文件
	// 加载excel
	retData := []vo.UserVo{}
	//
	tempMap := make(map[string]int)
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return retData
	}

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//
	// 循环获得sheet
	for _, sheetName := range f.GetSheetMap() {
		fmt.Println(sheetName)
		rows, err := f.GetRows(sheetName)
		if err != nil {
			fmt.Println(err)
		}
		// 锚点
		flag := false
		// 清空map
		tempMap = make(map[string]int)
		for _, row := range rows {
			// 自动搜索姓名所在的行作为锚点
			if !flag {
				for _, cell := range row {
					if cell == "姓名" || cell == "负责人" {
						flag = true
						for index, tempCell := range row {
							if tempCell == "姓名" || tempCell == "负责人" {
								tempMap["name"] = index + 1
							}
							if tempCell == "工号" {
								tempMap["emploteeNumber"] = index + 1
							}
							if tempCell == "所属科组" || tempCell == "科组" {
								tempMap["department"] = index + 1
							}
							if tempCell == "年级" {
								tempMap["grade"] = index + 1
							}
							if tempCell == "班级" {
								tempMap["class"] = index + 1
							}
						}
						// 发现锚点，开始匹配数据
						flag = true
					}

				}
				continue
			}
			user := vo.UserVo{}

			// 如果存在name 数据
			if tempMap["name"] != 0 {
				user.Name = row[tempMap["name"]-1]
			}
			if tempMap["emploteeNumber"] != 0 {
				user.EmployeeNumber = row[tempMap["emploteeNumber"]-1]
			}
			if tempMap["department"] != 0 {
				user.Department = row[tempMap["department"]-1]
			}
			if tempMap["grade"] != 0 {
				user.Grade = row[tempMap["grade"]-1]
			}
			if tempMap["class"] != 0 {
				user.Class = row[tempMap["class"]-1]
			}
			retData = append(retData, user)
		}
		// 目前只支持单个sheet
		break
	}
	return retData

}

// WriteUserExcel 写入excel
func (utils *LoadFileService) WriteUserExcel(data []vo.TaskDetailVO, task vo.TasksVo) (string, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 如果表头不为空则使用表头
	sheetName := "Sheet1"
	//if len(data) > 0 && task.TaskName != "" {
	//	// 写入表头
	//	sheetName = task.TaskName
	//}

	allWriteData := [][]interface{}{}
	allWriteData = append(allWriteData, []interface{}{"序号", "姓名", "工号", "年级", "班级", "科组", "预文件名", "上传到文件名", "状态"})
	// 循环遍历data
	for idx, item := range data {
		// 从第二行开始写入
		row := []interface{}{idx + 1, item.User.Name, item.User.EmployeeNumber, item.User.Grade, item.User.Grade, item.User.Department, item.PreFileName, item.FileName, getStatus(item.TaskStatus)}
		allWriteData = append(allWriteData, row)
	}
	for idx, row := range allWriteData {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		// 如果不存在sheet则创建sheet
		err = f.SetSheetRow(sheetName, cell, &row)
		if err != nil {
			return "", err
		}
	}
	// 先判断文件是否存在，存在则删除
	err := os.Remove(task.TaskName + ".xlsx")
	if err != nil {
		fmt.Println(err)
		// 文件不存在，不管它
	}
	// 自动保存到本地
	if err := f.SaveAs(task.TaskName + ".xlsx"); err != nil {
		fmt.Println(err)
	}

	return task.TaskName + ".xlsx", nil
}

func getStatus(num int) string {
	if num == 1 {
		return "已提交"
	}
	if num == 2 {
		return "未提交"
	}
	return "未提交"
}
