// Package impl
// @Description
// @Author  Ymri  2024/5/10 19:42
// @Update 2024/5/10 19:42
package impl

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"goods-system/internal/application/vo"
)

var loadFileService = &LoadFileService{}

type LoadFileService struct{}

func GetLoadFileService() *LoadFileService {
	return loadFileService
}

// 读取ReadUserExcel 文件
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

	// 定义返回数据
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//
	// 循环获得sheet
	for _, sheetName := range f.GetSheetMap() {
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
						fmt.Println("记录", tempMap)
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
