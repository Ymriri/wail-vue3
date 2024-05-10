// Package appmanager
// @Description
// @Author  Ymri  2024/5/9 23:28
// @Update 2024/5/9 23:28
package appmanager

import (
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/service/impl"
)

var userService = impl.GetUserServiceInstance()

type UserManager struct{}

var userManagerInstance *UserManager

func GetUserManagerInstance() *UserManager {
	return userManagerInstance
}

// Page 分页
func (p *UserManager) Page(req vo.UserPageVo) api.Page[vo.UserVo] {
	return userService.Page(req, req.Page, req.Size)
}

// Save 保存
func (p *UserManager) Save(req vo.UserVo) {
	userService.Save(req)
}

// Delete 删除
func (p *UserManager) Delete(id uint) {
	userService.Delete(id)
}

// Update 更新
func (p *UserManager) Update(req vo.UserVo) {
	userService.Update(req)
}
