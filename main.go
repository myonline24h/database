// You can edit this code!
// Click here and start typing.
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

// generate code
func main() {
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "./models",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/elearning?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)

	g.GenerateAllTable()
	// execute the action of code generation

	course := g.GenerateModel("course")
	section := g.GenerateModel("section")
	lesson := g.GenerateModel("lesson")

	category := g.GenerateModel("category",
		gen.FieldRelate(field.HasMany, "Courses", course,
			&field.RelateConfig{
				//RelateSlice: true,
				GORMTag: "foreignKey:CategoryID",
			}),
	)

	g.GenerateModel("course",
		gen.FieldRelate(field.BelongsTo, "Category", category,
			&field.RelateConfig{
				//RelateSlice: true,
				GORMTag: "foreignKey:CategoryID",
			}),
		gen.FieldRelate(field.HasMany, "Section", section,
			&field.RelateConfig{
				//RelateSlice: true,
				GORMTag: "foreignKey:CourseID",
			}),
	)

	g.GenerateModel("section",
		gen.FieldRelate(field.BelongsTo, "Courses", course,
			&field.RelateConfig{
				//RelateSlice: true,
				GORMTag: "foreignKey:CourseID",
			}),
		gen.FieldRelate(field.HasMany, "Lesson", lesson,
			&field.RelateConfig{
				//RelateSlice: true,
				GORMTag: "foreignKey:SectionID",
			}),
	)

	g.GenerateModel("lesson",
		gen.FieldRelate(field.BelongsTo, "Course", course,
			&field.RelateConfig{
				//RelateSlice: true,
				GORMTag: "foreignKey:CourseID",
			}),
		gen.FieldRelate(field.BelongsTo, "Section", section,
			&field.RelateConfig{
				//RelateSlice: true,
				GORMTag: "foreignKey:SectionID",
			}),
	)
	//g.ApplyBasic(category, course)

	g.Execute()
}
