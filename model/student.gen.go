// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameStudent = "student"

// Student mapped from table <student>
type Student struct {
	ID           int32  `gorm:"column:id;primaryKey" json:"id"`
	ComparyID    int32  `gorm:"column:compary_id;not null" json:"compary_id"`
	UserID       int32  `gorm:"column:user_id;not null" json:"user_id"`
	Relationship string `gorm:"column:relationship" json:"relationship"`
}

// TableName Student's table name
func (*Student) TableName() string {
	return TableNameStudent
}
