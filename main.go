package main

import (
	config "github.com/eufelipemateus/store-blog-posts/configs"
	"github.com/eufelipemateus/store-blog-posts/database"
	"github.com/eufelipemateus/store-blog-posts/functions"
	"github.com/robfig/cron"
)

func main() {

	err := config.Load()
	if err != nil {
		panic("Erro on load config.toml")
	}
	database.OpenConnection()

	if config.GetApp().GenerateDB {
		database.GenerateDB()
	}

	println("Startando wordpress cron...")
	println("Exec.: " + config.GetApp().Cron)

	functions.GetClient()
	functions.GetPosts()

	cron := cron.New()
	cron.AddFunc(config.GetApp().Cron, functions.SaveData)
	cron.Run()
}
