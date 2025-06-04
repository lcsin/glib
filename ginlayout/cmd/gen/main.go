package main

import (
	"github.com/lcsin/glib/ginlayout/ioc"
	"gorm.io/gen"
)

func main() {
	ioc.InitConfig()

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../internal/repository/gen/query",
		// gen.WithoutContext：禁用WithContext模式
		// gen.WithDefaultQuery：生成一个全局Query对象Q
		// gen.WithQueryInterface：生成Query接口
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(ioc.InitDB())
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
