// Package controller
// @Description
// @Author  Ymri  2024/5/10 10:39
// @Update 2024/5/10 10:39
package controller

import (
	"go/types"
	"goods-system/internal/application/appmanager"
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/api"
)

var configControllerInstance = appmanager.GetConfigManagerInstance()

type ConfigController struct{}

var configController *ConfigController

func GetConfigControllerInstance() *ConfigController {
	return configController
}

// FindAllByParentId 根据父Id 查询所有配置
func (c *ConfigController) FindAllByParentId(parentId uint) api.RespData[[]vo.ConfigTreeVo] {
	data := configControllerInstance.FindAllByParentId(parentId)
	return api.Success[[]vo.ConfigTreeVo](data, "")
}

// FindById 查询单个配置
func (c *ConfigController) FindById(id uint) api.RespData[vo.ConfigTreeVo] {
	return api.Success[vo.ConfigTreeVo](configControllerInstance.FindById(id), "")
}

// Insert 保存
func (c *ConfigController) Insert(config vo.ConfigTreeVo) api.RespData[vo.ConfigTreeVo] {
	data := configControllerInstance.Insert(config)
	return api.Success(data, "")
}

// Update 更新
func (c *ConfigController) Update(config vo.ConfigTreeVo) api.RespData[types.Nil] {
	configControllerInstance.Update(config)
	return api.Success[types.Nil](types.Nil{}, "")
}

// Delete 删除
func (c *ConfigController) Delete(id uint) api.RespData[types.Nil] {
	configControllerInstance.Delete(id)
	return api.Success[types.Nil](types.Nil{}, "")
}
