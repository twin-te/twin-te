package timetabledomain_test

import (
	"testing"

	timetabledomain "github.com/twin-te/twin-te/back/module/timetable/domain"
)

func TestParseCourseMethod(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    timetabledomain.CourseMethod
		wantErr bool
	}{
		{"online asynchronous", "OnlineAsynchronous", timetabledomain.CourseMethodOnlineAsynchronous, false},
		{"online synchronous", "OnlineSynchronous", timetabledomain.CourseMethodOnlineSynchronous, false},
		{"face to face", "FaceToFace", timetabledomain.CourseMethodFaceToFace, false},
		{"others", "Others", timetabledomain.CourseMethodOthers, false},
		{"invalid", "invalid", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timetabledomain.ParseCourseMethod(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if !tt.wantErr && got.String() != tt.input {
				t.Errorf("String() = %v, want %v", got.String(), tt.input)
			}
		})
	}
}
