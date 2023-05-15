package database

import (
	"github.com/eufelipemateus/store-blog-posts/interfaces"
	"gorm.io/gen"
)

func GenerateDB() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./database/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(DB)

	g.ApplyBasic(
		interfaces.Post{},
	)

	g.Execute()

	DB.AutoMigrate(
		interfaces.Post{},
	)
}
