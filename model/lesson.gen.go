// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameLesson = "lesson"

// Lesson mapped from table <lesson>
type Lesson struct {
	ID             int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Title          string         `gorm:"column:title" json:"title"`
	Duration       string         `gorm:"column:duration" json:"duration"`
	CourseID       int32          `gorm:"column:course_id" json:"course_id"`
	SectionID      int32          `gorm:"column:section_id" json:"section_id"`
	VideoType      string         `gorm:"column:video_type" json:"video_type"`
	VideoURL       string         `gorm:"column:video_url" json:"video_url"`
	VideoCallURL   string         `gorm:"column:video_call_url" json:"video_call_url"`
	VideoCallAt    time.Time      `gorm:"column:video_call_at" json:"video_call_at"`
	VideoCallCode  string         `gorm:"column:video_call_code" json:"video_call_code"`
	VideoCallApp   string         `gorm:"column:video_call_app" json:"video_call_app"`
	DateAdded      int32          `gorm:"column:date_added" json:"date_added"`
	LastModified   int32          `gorm:"column:last_modified" json:"last_modified"`
	LastModifiedBy string         `gorm:"column:last_modified_by" json:"last_modified_by"`
	LessonType     string         `gorm:"column:lesson_type" json:"lesson_type"`
	AttachmentURL  string         `gorm:"column:attachment_url" json:"attachment_url"`
	AttachmentType string         `gorm:"column:attachment_type" json:"attachment_type"`
	Summary        string         `gorm:"column:summary" json:"summary"`
	Order          int32          `gorm:"column:order;not null" json:"order"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	CreatedBy      string         `gorm:"column:created_by" json:"created_by"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Course         Course         `gorm:"foreignKey:CourseID" json:"course"`
	Section        Section        `gorm:"foreignKey:SectionID" json:"section"`
}

// TableName Lesson's table name
func (*Lesson) TableName() string {
	return TableNameLesson
}