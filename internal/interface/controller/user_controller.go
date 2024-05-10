// Package controller
// @Description
// @Author  Ymri  2024/5/9 23:45
// @Update 2024/5/9 23:45
package controller

import (
	"go/types"
	"goods-system/internal/application/appmanager"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
)

var userManager = appmanager.GetUserManagerInstance()

type UserController struct{}

var userControllerInstance *UserController

func GetUserControllerInstance() *UserController {
	return userControllerInstance
}

// UserPage 分页
func (u *UserController) UserPage(req vo.UserPageVo) api.RespData[api.Page[vo.UserVo]] {
	page := userManager.Page(req)
	return api.Success[api.Page[vo.UserVo]](page, "")
}

// Save 保存
func (u *UserController) Save(req vo.UserVo) api.RespData[types.Nil] {
	userManager.Save(req)
	return api.Success[types.Nil](types.Nil{}, "")
}

// Delete 删除
func (u *UserController) Delete(id uint) api.RespData[types.Nil] {
	userManager.Delete(id)
	return api.Success(types.Nil{}, "")
}

// Update 更新
func (u *UserController) Update(req vo.UserVo) api.RespData[types.Nil] {
	userManager.Update(req)
	return api.Success(types.Nil{}, "")
}
