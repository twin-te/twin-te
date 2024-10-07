// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCourse = "courses"

// Course mapped from table <courses>
type Course struct {
	ID                string                   `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Year              int16                    `gorm:"column:year;type:smallint;not null;uniqueIndex:courses_year_code_key,priority:1" json:"year"`
	Code              string                   `gorm:"column:code;type:text;not null;uniqueIndex:courses_year_code_key,priority:2" json:"code"`
	Name              string                   `gorm:"column:name;type:text;not null" json:"name"`
	Instructors       string                   `gorm:"column:instructors;type:text;not null" json:"instructors"`
	Credit            float64                  `gorm:"column:credit;type:numeric(4,1);not null" json:"credit"`
	Overview          string                   `gorm:"column:overview;type:text;not null" json:"overview"`
	Remarks           string                   `gorm:"column:remarks;type:text;not null" json:"remarks"`
	LastUpdatedAt     time.Time                `gorm:"column:last_updated_at;type:timestamp without time zone;not null" json:"last_updated_at"`
	HasParseError     bool                     `gorm:"column:has_parse_error;type:boolean;not null" json:"has_parse_error"`
	IsAnnual          bool                     `gorm:"column:is_annual;type:boolean;not null" json:"is_annual"`
	RecommendedGrades []CourseRecommendedGrade `json:"recommended_grades"`
	Methods           []CourseMethod           `json:"methods"`
	Schedules         []CourseSchedule         `json:"schedules"`
}

// TableName Course's table name
func (*Course) TableName() string {
	return TableNameCourse
}
