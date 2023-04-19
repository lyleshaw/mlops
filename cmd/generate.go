package main

import (
	"github.com/lyleshaw/mlops/biz/dal/mysql"
	genModel "github.com/lyleshaw/mlops/biz/model/orm_gen"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./biz/model/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	// reuse your gorm orm_gen
	mysql.Init()
	g.UseDB(mysql.DB)

	// Generate struct based on table
	//g.GenerateAllTable()

	// Generate basic type-safe DAO API for struct `orm_gen.User` following conventions
	g.ApplyBasic(genModel.User{})
	g.ApplyBasic(genModel.Application{})

	// Generate the code
	g.Execute()
}
