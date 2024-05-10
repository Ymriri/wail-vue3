// Package appmanager
// @Description
// @Author  Ymri  2024/5/10 10:39
// @Update 2024/5/10 10:39
package appmanager

import (
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/service/impl"
	"time"
)

var configService = impl.GetConfigServiceInstance()

type ConfigManager struct{}

var configManagerInstance *ConfigManager

func GetConfigManagerInstance() *ConfigManager {
	return configManagerInstance
}

// FindAllByParentId 根据父Id 查询所有配置
func (c *ConfigManager) FindAllByParentId(parentId uint) []vo.ConfigTreeVo {
	return configService.FindAllByParentId(parentId)
}

// FindById 根据Id 查询配置
func (c *ConfigManager) FindById(id uint) vo.ConfigTreeVo {
	return configService.FindById(id)
}

// Insert 保存
func (c *ConfigManager) Insert(config vo.ConfigTreeVo) {
	tempConfig := convert.ToConfigByVo(config)
	tempConfig.CreatedAt = time.Now()
	tempConfig.UpdatedAt = time.Now()
	configService.Insert(&tempConfig)
}

// Update 更新
func (c *ConfigManager) Update(config vo.ConfigTreeVo) {
	tempConfig := convert.ToConfigByVo(config)
	tempConfig.UpdatedAt = time.Now()
	configService.Update(&tempConfig)
}

// Delete 删除
func (c *ConfigManager) Delete(id uint) {
	configService.Delete(id)
}
