package main

import (
	"github.com/fireyun/gorm-gen-demo/gendemo/dal/model"

	"gorm.io/gen"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery,
	})

	g.ApplyBasic(
		model.App{},
		model.Passport{},
		model.User{},
		model.TemplateSets{})

	// 根据接口定义为模型生成自定义的方法（根据接口方法的sql注释生成对应逻辑）
	g.ApplyInterface(func(model.Method) {}, model.User{})
	g.ApplyInterface(func(model.UserMethod) {}, model.User{})

	g.Execute()
}
