// Package convert
// @Description
// @Author  Ymri  2024/5/10 10:32
// @Update 2024/5/10 10:32
package convert

import (
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/db/entity"
)

func ToConfigByVo(param vo.ConfigTreeVo) entity.ConfigTree {
	return entity.ConfigTree{
		ID:          param.ID,
		ParentID:    param.ParentID,
		ConfigName:  param.ConfigName,
		ConfigValue: param.ConfigValue,
	}
}

func ToConfigVo(e entity.ConfigTree) vo.ConfigTreeVo {
	return vo.ConfigTreeVo{
		ID:          e.ID,
		ParentID:    e.ParentID,
		ConfigName:  e.ConfigName,
		ConfigValue: e.ConfigValue,
	}
}
