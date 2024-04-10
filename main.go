package main

import (
	"banner/internal/logger"
	"banner/internal/model"
	psql "banner/internal/repo/Db"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	l := logger.New()
	db, err := psql.InitDb(l)
	if err != nil {
		panic(err)
	}
	db.Insert("test", model.BannerDB{
		Feature:    1,
		Title:      "title",
		Text:       "text",
		Url:        "url",
		Active:     true,
		Created_at: time.Now().Format(time.DateTime),
		Updated_at: time.Now().Format(time.DateTime),
	}, model.Tags{
		1, 2, 3, 4, 5,
	})
	// cmd.Execute()
}
