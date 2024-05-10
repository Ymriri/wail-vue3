// Package convert
// @Description
// @Author  Ymri  2024/5/9 23:34
// @Update 2024/5/9 23:34
package convert

import (
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/db/entity"
)

func ToUserByPageVo(param vo.UserPageVo) entity.User {
	return entity.User{
		Name:         param.Name,
		ConfigFileID: param.ConfigFileID,
	}
}

func ToUserByVo(param vo.UserVo) entity.User {
	return entity.User{
		ID:             param.ID,
		Name:           param.Name,
		Grade:          param.Grade,
		Class:          param.Class,
		Department:     param.Department,
		Group:          param.Group,
		EmployeeNumber: param.EmployeeNumber,
		ConfigFileID:   param.ConfigFileID,
	}
}

func ToUserVo(e entity.User) vo.UserVo {
	return vo.UserVo{
		ID:             e.ID,
		Name:           e.Name,
		Grade:          e.Grade,
		Class:          e.Class,
		Department:     e.Department,
		Group:          e.Group,
		EmployeeNumber: e.EmployeeNumber,
		ConfigFileID:   e.ConfigFileID,
	}
}
