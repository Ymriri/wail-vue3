// Package convert
// @Description
// @Author  Ymri  2024/5/11 14:54
// @Update 2024/5/11 14:54
package convert

import (
	"goods-system/internal/application/vo"
	"goods-system/internal/infrastructure/db/entity"
)

// ToTaskDetailByVo 转成vo
func ToTaskDetailByVo(param vo.TaskDetailVO) entity.TaskDetail {
	return entity.TaskDetail{
		ID:          param.ID,
		TaskID:      param.TaskID,
		TaskStatus:  param.TaskStatus,
		PreFileName: param.PreFileName,
		FileName:    param.FileName,
		UserId:      param.UserId,
	}
}

// ToTaskDetailVoByDetail 实体转vo 分组名称需要自己补充上去
func ToTaskDetailVoByDetail(param entity.TaskDetail) vo.TaskDetailVO {
	return vo.TaskDetailVO{
		ID:          param.ID,
		TaskID:      param.TaskID,
		TaskStatus:  param.TaskStatus,
		PreFileName: param.PreFileName,
		FileName:    param.FileName,
		UserId:      param.UserId,
		User:        ToUserVo(param.User),
	}
}
