// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRegisteredCourse = "registered_courses"

// RegisteredCourse mapped from table <registered_courses>
type RegisteredCourse struct {
	ID          string                `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID      string                `gorm:"column:user_id;type:uuid;not null;uniqueIndex:registered_courses_user_id_course_id_key,priority:1;index:registered_courses_user_id_year_idx,priority:1" json:"user_id"`
	Year        int16                 `gorm:"column:year;type:smallint;not null;index:registered_courses_user_id_year_idx,priority:2" json:"year"`
	CourseID    *string               `gorm:"column:course_id;type:uuid;uniqueIndex:registered_courses_user_id_course_id_key,priority:2" json:"course_id"`
	Name        *string               `gorm:"column:name;type:text" json:"name"`
	Instructors *string               `gorm:"column:instructors;type:text" json:"instructors"`
	Credit      *float64              `gorm:"column:credit;type:numeric(4,1)" json:"credit"`
	Methods     *string               `gorm:"column:methods;type:text[]" json:"methods"`
	Schedules   *string               `gorm:"column:schedules;type:jsonb" json:"schedules"`
	Memo        string                `gorm:"column:memo;type:text;not null" json:"memo"`
	Attendance  int16                 `gorm:"column:attendance;type:smallint;not null" json:"attendance"`
	Absence     int16                 `gorm:"column:absence;type:smallint;not null" json:"absence"`
	Late        int16                 `gorm:"column:late;type:smallint;not null" json:"late"`
	Tags        []RegisteredCourseTag `gorm:"foreignKey:RegisteredCourse;references:ID" json:"tags"`
}

// TableName RegisteredCourse's table name
func (*RegisteredCourse) TableName() string {
	return TableNameRegisteredCourse
}
