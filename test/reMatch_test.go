// Package test
// @Description
// @Author  Ymri  2024/5/14 23:56
// @Update 2024/5/14 23:56
package test

import (
	"goods-system/internal/application/appmanager"
	"testing"
)

var fileManager = appmanager.GetTaskDetailManagerInstance()

// TestReMatch 测试正则匹配情况
func TestReMatch(t *testing.T) {
	preFile := "2023-2024学年第2学期垃圾综合实践课程方案{_**}" // 预定义文件名
	fileName := "2023-2024学年第2学期综合实践课程方案_《成语大王》_CS1234张三.zip"
	flag := fileManager.FileCHECK(preFile, fileName)
	if flag != false {
		t.Errorf("TestReMatch() failed")
	}

}

func TestReMath2(t *testing.T) {
	preFile := "2023-2024学年第2学期综合实践课程方案_{**}_CS1234张三" // 预定义文件名
	fileName := "2023-2024学年第2学期综合实践课程方案_《成语大王》_CS1234张三.zip"
	flag := fileManager.FileCHECK(preFile, fileName)
	if flag != true {
		t.Errorf("TestReMatch() failed")
	}
}

func TestReMath3(t *testing.T) {
	preFile := "刘玉洁-音乐-wor.zip" // 预定义文件名
	fileName := "刘玉洁-音乐-wor.zip"
	flag := fileManager.FileCHECK(preFile, fileName)
	if flag != true {
		t.Errorf("TestReMatch() failed")
	}
}
