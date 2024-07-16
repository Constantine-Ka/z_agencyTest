package main

import (
	"zeroagencytest/pkg/config"
	"zeroagencytest/pkg/db"
	"zeroagencytest/pkg/repository"
	"zeroagencytest/pkg/router"
	"zeroagencytest/pkg/utils/logging"
)

func main() {
	logger := logging.GetLogger()
	config.Load(".env")
	database, err := db.New()
	//defer database.Close()
	if err != nil {
		logger.Fatal(err)
		return
	}
	repo := repository.New(database)
	router.Init(repo, config.GetString("APP_PORT"))

}
