// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameLanguage = "language"

// Language mapped from table <language>
type Language struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Phrase    string         `gorm:"column:phrase" json:"phrase"`
	English   string         `gorm:"column:english" json:"english"`
	Bengali   string         `gorm:"column:Bengali" json:"Bengali"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Language's table name
func (*Language) TableName() string {
	return TableNameLanguage
}