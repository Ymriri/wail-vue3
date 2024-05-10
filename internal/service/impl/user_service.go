// Package impl
// @Description
// @Author  Ymri  2024/5/9 23:30
// @Update 2024/5/9 23:30
package impl

import (
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/db/repositories"
)

type UserService struct{}

var userService *UserService

var userRepo = repositories.GetUserRepositoryInstance()

func GetUserServiceInstance() *UserService {
	return userService
}

// Page 分页查询
func (u *UserService) Page(param vo.UserPageVo, page, size int) api.Page[vo.UserVo] {
	uservo := convert.ToUserByPageVo(param)
	pageData := userRepo.Page(&uservo, page, size)
	userList := pageData.Data
	var userVoList []vo.UserVo
	for _, e := range userList {
		userVo := convert.ToUserVo(e)
		userVoList = append(userVoList, userVo)
	}
	return api.Page[vo.UserVo]{
		Page:  pageData.Page,
		Size:  pageData.Size,
		Total: pageData.Total,
		Data:  userVoList,
	}
}

// Save 保存单个用户
func (u *UserService) Save(param vo.UserVo) {
	entity := convert.ToUserByVo(param)
	userRepo.Insert(&entity)
}

// Delete 删除用户
func (u *UserService) Delete(id uint) {
	userRepo.DeleteById(id)
}

// Update 更新用户
func (u *UserService) Update(param vo.UserVo) {
	entity := convert.ToUserByVo(param)
	userRepo.Update(&entity)
}
