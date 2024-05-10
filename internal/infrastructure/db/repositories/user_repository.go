// Package repositories
// @Description
// @Author  Ymri  2024/5/9 23:14
// @Update 2024/5/9 23:14
package repositories

import (
	"goods-system/internal/infrastructure/common/api"
	"goods-system/internal/infrastructure/common/global"
	"goods-system/internal/infrastructure/db/entity"
	"time"
)

var userInstance *UserRepository

type UserRepository struct{}

func GetUserRepositoryInstance() *UserRepository {
	return userInstance
}

// page分页查询

func (p *UserRepository) Page(user *entity.User, page, size int) api.Page[entity.User] {
	var total int64
	var userList []entity.User
	tx := global.DB.Model(&entity.User{})

	if user.Name != "" {
		// gorm 模糊查询
		tx.Where("name like ?", "%"+user.Name+"%")
	}
	// -1 表示查询所有
	// 不能为空
	if user.ConfigFileID != -1 {
		tx.Where("config_file_id = ?", user.ConfigFileID)
	}
	// 预加载
	tx.Preload("ConfigTree").
		Count(&total).
		Limit(size).
		Offset((page - 1) * size).
		Find(&userList)
	return api.Page[entity.User]{
		Page:  page,
		Size:  size,
		Total: total,
		Data:  userList,
	}
}

// Insert 保存
func (u *UserRepository) Insert(user *entity.User) {
	user.UpdatedAt = time.Now()
	global.DB.Create(user)
}

// SelectById 查询当个用户
func (u *UserRepository) SelectById(id uint) entity.User {
	var user entity.User
	global.DB.First(&user, id)
	return user
}

// Update 更新用户
func (u *UserRepository) Update(user *entity.User) {

	old := u.SelectById(user.ID)
	updates := make(map[string]interface{})
	flag := false
	if old.Name != user.Name {
		updates["name"] = user.Name
		flag = true
	}
	if old.Grade != user.Grade {
		updates["grade"] = user.Grade
		flag = true

	}
	if old.Class != user.Class {
		updates["class"] = user.Class
		flag = true

	}
	if old.Department != user.Department {
		updates["department"] = user.Department
		flag = true
	}
	if old.Group != user.Group {
		updates["group"] = user.Group
		flag = true
	}
	if old.EmployeeNumber != user.EmployeeNumber {
		updates["employee_number"] = user.EmployeeNumber
		flag = true
	}
	if old.ConfigFileID != user.ConfigFileID {
		updates["config_file_id"] = user.ConfigFileID
		flag = true
	}
	if flag {
		updates["updated_at"] = time.Now()
		global.DB.Model(&entity.User{}).Where("id = ?", user.ID).Updates(updates)
	}
}

// DeleteById  删除用户
func (u *UserRepository) DeleteById(id uint) {
	var user entity.User
	global.DB.Delete(&user, id)
}
