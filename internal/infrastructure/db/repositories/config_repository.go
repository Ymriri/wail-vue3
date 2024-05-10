// Package repositories
// @Description
// @Author  Ymri  2024/5/10 10:34
// @Update 2024/5/10 10:34
package repositories

import (
	"goods-system/internal/infrastructure/common/global"
	"goods-system/internal/infrastructure/db/entity"
)

var configInstance *ConfigRepository

type ConfigRepository struct{}

func GetConfigRepositoryInstance() *ConfigRepository {
	if configInstance == nil {
		configInstance = &ConfigRepository{}
	}
	return configInstance
}

// FindAllByParentId 根据父Id 查询所有配置
func (c *ConfigRepository) FindAllByParentId(parentId uint) []entity.ConfigTree {
	var configList []entity.ConfigTree
	global.DB.Where("parent_id = ?", parentId).Find(&configList)
	return configList
}

// FindById 根据Id 查询配置
func (c *ConfigRepository) FindById(id uint) entity.ConfigTree {
	var config entity.ConfigTree
	global.DB.First(&config, id)
	return config
}

// Insert 保存
func (c *ConfigRepository) Insert(config *entity.ConfigTree) {
	global.DB.Create(config)
}

// Update 更新
func (c *ConfigRepository) Update(config *entity.ConfigTree) {
	global.DB.Save(config)
}

// Delete 删除
func (c *ConfigRepository) Delete(id uint) {
	global.DB.Delete(&entity.ConfigTree{}, id)
}
