package main

import (
	"github.com/andreas-sujono/go-basic-projects/go-rest-boilerplate/cmd/seed/seeders"
	"github.com/andreas-sujono/go-basic-projects/go-rest-boilerplate/internal/config"
	"github.com/andreas-sujono/go-basic-projects/go-rest-boilerplate/internal/constants"
	"github.com/andreas-sujono/go-basic-projects/go-rest-boilerplate/internal/utils"
	"github.com/andreas-sujono/go-basic-projects/go-rest-boilerplate/pkg/logger"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := config.InitializeAppConfig(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
	logger.Info("configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}

func main() {
	db, err := utils.SetupPostgresConnection()
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
	}
	defer db.Close()

	logger.Info("seeding...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})

	seeder := seeders.NewSeeder(db)
	err = seeder.UserSeeder(seeders.UserData)
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
	}

	logger.Info("seeding success!", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
}
