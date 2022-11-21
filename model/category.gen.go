// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCategory = "category"

// Category mapped from table <category>
type Category struct {
	ID               int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Code             string         `gorm:"column:code" json:"code"`
	Name             string         `gorm:"column:name" json:"name"`
	Parent           int32          `gorm:"column:parent" json:"parent"`
	Slug             string         `gorm:"column:slug" json:"slug"`
	DateAdded        int32          `gorm:"column:date_added" json:"date_added"`
	LastModified     int32          `gorm:"column:last_modified" json:"last_modified"`
	FontAwesomeClass string         `gorm:"column:font_awesome_class" json:"font_awesome_class"`
	Thumbnail        string         `gorm:"column:thumbnail" json:"thumbnail"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Courses          []Course       `gorm:"foreignKey:CategoryID" json:"courses"`
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}