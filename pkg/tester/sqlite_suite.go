package tester

import (
	"go-api-arch-mvc-template/app/models"
	"go-api-arch-mvc-template/configs"
	"os"

	"github.com/stretchr/testify/suite"
)

type DBSQLiteSuite struct {
	suite.Suite
}

func (suite *DBSQLiteSuite) SetupSuite() {
	configs.Config.DBName = "unittest.sqlite"
	err := models.SetDatabase(models.InstanceSqlite)
	suite.Assert().Nil(err)

	for _, model := range models.GetModels() {
		err = models.DB.AutoMigrate(model)
		suite.Assert().Nil(err)
	}
}

func (suite *DBSQLiteSuite) TearDownSuite() {
	err := os.Remove(configs.Config.DBName)
	suite.Assert().Nil(err)
}
