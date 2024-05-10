// Package entity
// @Description
// @Author  Ymri  2024/5/9 23:10
// @Update 2024/5/9 23:10
package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string     `gorm:"type:text;not null"`
	Grade          string     `gorm:"type:text"`
	Class          string     `gorm:"type:text"`
	Department     string     `gorm:"type:text"`
	Group          string     `gorm:"column:dgroup;type:text"`
	EmployeeNumber string     `gorm:"type:text"`
	ConfigFileID   int        `gorm:"type:integer"`
	ID             uint       `gorm:"primarykey"`
	ConfigTree     ConfigTree `gorm:"foreignKey:ConfigFileID;references:ID"`
}
