// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTimetableModule = "timetable_modules"

// TimetableModule mapped from table <timetable_modules>
type TimetableModule struct {
	Module string `gorm:"column:module;type:text;primaryKey" json:"module"`
}

// TableName TimetableModule's table name
func (*TimetableModule) TableName() string {
	return TableNameTimetableModule
}
