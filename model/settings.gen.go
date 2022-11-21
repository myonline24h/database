// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSetting = "settings"

// Setting mapped from table <settings>
type Setting struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Key       string         `gorm:"column:key" json:"key"`
	Value     string         `gorm:"column:value" json:"value"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Setting's table name
func (*Setting) TableName() string {
	return TableNameSetting
}