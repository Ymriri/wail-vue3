// Package impl
// @Description
// @Author  Ymri  2024/5/10 10:36
// @Update 2024/5/10 10:36
package impl

import (
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/common/convert"
	"goods-system/internal/infrastructure/db/entity"
	"goods-system/internal/infrastructure/db/repositories"
)

type ConfigService struct{}

var configService *ConfigService

var configRepo = repositories.GetConfigRepositoryInstance()

func GetConfigServiceInstance() *ConfigService {
	return configService
}

// FindAllByParentId 根据父Id 查询所有配置
func (c *ConfigService) FindAllByParentId(parentId uint) []vo.ConfigTreeVo {
	configList := configRepo.FindAllByParentId(parentId)
	var configTreeVoList []vo.ConfigTreeVo
	for _, config := range configList {
		temp_config := convert.ToConfigVo(config)
		configTreeVoList = append(configTreeVoList, temp_config)
	}
	return configTreeVoList
}

// FindById 根据Id 查询配置
func (c *ConfigService) FindById(id uint) vo.ConfigTreeVo {
	return convert.ToConfigVo(configRepo.FindById(id))
}

// Insert 保存
func (c *ConfigService) Insert(config *entity.ConfigTree) {
	configRepo.Insert(config)
}

// Update 更新
func (c *ConfigService) Update(config *entity.ConfigTree) {
	configRepo.Update(config)
}

// Delete 删除
func (c *ConfigService) Delete(id uint) {
	configRepo.Delete(id)
}
