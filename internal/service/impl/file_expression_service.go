// Package impl
// @Description
// @Author  Ymri  2024/5/14 10:37
// @Update 2024/5/14 10:37
package impl

import (
	"fmt"
	"goods-system/internal/application/vo"
	"strconv"
)

type FileExpressionService struct{}

var fileExpressionServiceInstance *FileExpressionService

func GetFileExpressionServiceInstance() *FileExpressionService {
	return fileExpressionServiceInstance
}

// ReadFileExpression 读取文件表达式
func (p *FileExpressionService) ReadFileExpression(fileExp string, configId int) []vo.TaskDetailVO {
	fmt.Println(fileExp)
	fmt.Println(configId)
	taskDetailsVoList := []vo.TaskDetailVO{}
	if !p.checkFileExpression(fileExp) {
		return taskDetailsVoList
	}
	// 查询Id为configId的所有用户
	//userService.
	user := vo.UserVo{}
	user.ConfigFileID = configId
	userList := userService.SelectByGroup(user)
	// 检查是否存在{__},忽略__后面所有的字符
	// 字符串分割

	for _, tempuser := range userList {
		preFileString := processString(fileExp, tempuser)

		for _, tempString := range preFileString {
			addTaskDetail := vo.TaskDetailVO{}
			addTaskDetail.User = tempuser
			addTaskDetail.PreFileName = tempString
			addTaskDetail.UserId = tempuser.ID
			taskDetailsVoList = append(taskDetailsVoList, addTaskDetail)
		}
	}
	return taskDetailsVoList

}

// CheckFileExpression 检查文件表达式是否合法
func (p *FileExpressionService) checkFileExpression(fileExp string) bool {
	// 检查左括号数量和右括号数量

	// 检查左括号数量和右括号数量
	leftCount := 0
	rightCount := 0
	for _, v := range fileExp {
		if v == '{' {
			leftCount++
		} else if v == '}' {
			rightCount++
		}
	}
	if leftCount != rightCount {
		return false
	}
	return true

}

func processString(s string, userVo vo.UserVo) []string {
	retString := []string{}
	retString = append(retString, "")
	// 手动遍历
	ss := []rune(s)
	for i := 0; i < len(ss); i++ {
		//检查所有{开头
		if ss[i] == '{' {
			// 开始匹配 单独的数字开头
			// 匹配数字
			if i+1 < len(ss) && ss[i+1] >= '0' && ss[i+1] <= '9' {
				// 字符串转数字

				addTempString := extractUserProperty(int(ss[i+1]-'0'), userVo)
				fmt.Println("内容：", int(ss[i+1]-'0'))
				retString = addString(retString, addTempString)
				// 跳到}
				for j := i + 2; j < len(ss); j++ {
					if ss[j] == '}' {
						i = j
						break
					}
				}
				continue
			}
			// 生成重复次数
			if i+1 < len(ss) && ss[i+1] == 'c' {
				// 需要继续往下重复匹配
				// 从i+2开始匹配
				left := 0
				right := 0
				// 根据-进行分割

				for j := i + 2; j < len(ss) && ss[j] != '-' && ss[j] != '}'; j++ {
					left = left*10 + int(ss[j]-'0')
					i = j
				}
				right = left
				if ss[i+1] == '-' {
					right = 0
					for j := i + 2; j < len(ss) && ss[j] != '}'; j++ {
						right = right*10 + int(ss[j]-'0')
						i = j
					}
				}
				retString = copyAddString(retString, left, right)
				// 跳到}
				for j := i + 1; j < len(ss); j++ {
					if ss[j] == '}' {
						i = j
						break
					}
				}
				continue
			}
			// 保存任意匹配的规则
			retString = addString(retString, string(ss[i]))
		}
		retString = addString(retString, string(ss[i]))
	}
	return retString
}

// s所有都添加
func addString(s []string, add string) []string {
	for i := 0; i < len(s); i++ {
		s[i] += add
	}
	return s
}

// copyAddString 复制并且添加
func copyAddString(s []string, begin int, end int) []string {
	retString := []string{}
	for i := begin; i <= end; i++ {
		//
		// 复制传入的s
		tempString := make([]string, len(s))
		copy(tempString, s)
		// 添加
		tempString = addString(tempString, strconv.Itoa(i))
		retString = append(retString, tempString...)
	}
	return retString
}

func extractUserProperty(num int, userVo vo.UserVo) string {

	switch num {
	case 1:
		return userVo.Name
	case 2:
		return userVo.EmployeeNumber
	case 3:
		return userVo.Class
	case 4:
		return userVo.Grade
	case 5:
		return userVo.Department
	case 6:
		return userVo.EmployeeNumber
	default:
		return ""
	}
}
