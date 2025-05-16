package models_test

import (
	"go-api-arch-mvc-template/app/models"
	"go-api-arch-mvc-template/pkg/tester"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CategoryTestSuite struct {
	tester.DBSQLiteSuite
}

func TestCategoryTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryTestSuite))
}

func (suite *CategoryTestSuite) TestCategory() {
	category, err := models.GetOrCreateCategory("sports")
	suite.Assert().Nil(err)
	suite.Assert().NotNil(category.ID)
	suite.Assert().Equal("sports", category.Name)

	category2, err := models.GetOrCreateCategory("sports")
	suite.Assert().Nil(err)
	suite.Assert().Equal(category.ID, category2.ID)
	suite.Assert().Equal("sports", category2.Name)
}
