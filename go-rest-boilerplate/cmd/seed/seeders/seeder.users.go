package seeders

import (
	"github.com/andreas-sujono/go-basic-projects/go-rest-boilerplate/internal/constants"
	"github.com/andreas-sujono/go-basic-projects/go-rest-boilerplate/internal/datasources/records"
	"github.com/andreas-sujono/go-basic-projects/go-rest-boilerplate/pkg/helpers"
	"github.com/andreas-sujono/go-basic-projects/go-rest-boilerplate/pkg/logger"
	"github.com/sirupsen/logrus"
)

var pass string
var UserData []records.Users

func init() {
	var err error
	pass, err = helpers.GenerateHash("12345")
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
	}

	UserData = []records.Users{
		{
			Username: "patrick star 7",
			Email:    "patrick@gmail.com",
			Password: pass,
			Active:   true,
			RoleId:   1,
		},
		{
			Username: "john doe",
			Email:    "johndoe@gmail.com",
			Password: pass,
			Active:   false,
			RoleId:   2,
		},
	}
}
