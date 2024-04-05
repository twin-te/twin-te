package main

import (
	"log"

	"github.com/twin-te/twin-te/back/appenv"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./db/gen/query",
		ModelPkgPath:      "model",
		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	db, err := gorm.Open(postgres.Open(appenv.DB_URL))
	if err != nil {
		log.Fatalln(err)
	}
	g.UseDB(db)

	g.ApplyBasic(g.GenerateAllTable()...)

	g.ApplyBasic(
		g.GenerateModel(
			"users",
			gen.FieldRelate(field.HasMany, "UserAuthentications", g.GenerateModel("user_authentications"), nil),
		),
		g.GenerateModel(
			"registered_courses",
			gen.FieldRelate(field.HasMany, "Tags", g.GenerateModel("registered_course_tags"), &field.RelateConfig{
				GORMTag: field.GormTag{
					"foreignKey": []string{"RegisteredCourse"},
					"references": []string{"ID"},
				},
			}),
		),
		g.GenerateModel(
			"courses",
			gen.FieldRelate(field.HasMany, "RecommendedGrades", g.GenerateModel("course_recommended_grades"), nil),
			gen.FieldRelate(field.HasMany, "Methods", g.GenerateModel("course_methods"), nil),
			gen.FieldRelate(field.HasMany, "Schedules", g.GenerateModel("course_schedules"), nil),
		),
	)

	g.Execute()
}
