// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMessage = "message"

// Message mapped from table <message>
type Message struct {
	ID                int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MessageThreadCode string         `gorm:"column:message_thread_code" json:"message_thread_code"`
	Message           string         `gorm:"column:message" json:"message"`
	Sender            string         `gorm:"column:sender" json:"sender"`
	Timestamp         string         `gorm:"column:timestamp" json:"timestamp"`
	ReadStatus        int32          `gorm:"column:read_status" json:"read_status"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}
