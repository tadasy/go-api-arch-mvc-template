package controllers

import (
	"encoding/json"
	"go-api-arch-mvc-template/api"
	"go-api-arch-mvc-template/app/models"
	"go-api-arch-mvc-template/pkg/tester"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type AlbumControllerSuite struct {
	tester.DBSQLiteSuite
	albumHandler AlbumHandler
	originalDB   *gorm.DB
}

func TestAlbumControllersTestSuite(t *testing.T) {
	suite.Run(t, new(AlbumControllerSuite))
}

func (suite *AlbumControllerSuite) SetupSuite() {
	suite.DBSQLiteSuite.SetupSuite()
	suite.originalDB = models.DB
	suite.albumHandler = AlbumHandler{}
}

func (suite *AlbumControllerSuite) MockDB() sqlmock.Sqlmock {
	mock, mockGormDB := tester.MockDB()
	models.DB = mockGormDB
	return mock
}

func (suite *AlbumControllerSuite) AfterTest(suiteName, testName string) {
	models.DB = suite.originalDB
}

func (suite *AlbumControllerSuite) TestCreate() {
	request, _ := api.NewCreateAlbumRequest("/api/v1", api.CreateAlbumJSONRequestBody{
		Title:       "test",
		ReleaseDate: api.ReleaseDate{Time: time.Now()},
		Category:    api.Category{Name: "sports"},
	})
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)
	ginContext.Request = request

	suite.albumHandler.CreateAlbum(ginContext)

	suite.Assert().Equal(http.StatusCreated, w.Code)
	bodyBytes, _ := io.ReadAll(w.Body)
	var albumGetResponse api.AlbumResponse
	err := json.Unmarshal(bodyBytes, &albumGetResponse)
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusCreated, w.Code)
	suite.Assert().Equal("test", albumGetResponse.Title)
	suite.Assert().Equal("sports", string(albumGetResponse.Category.Name))
	suite.Assert().NotNil(albumGetResponse.ReleaseDate)
}

func (suite *AlbumControllerSuite) TestCreateRequestBodyFailure() {
	w := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("POST", "/api/v1/album", nil)
	req.Header.Set("Content-Type", "application/json")
	ginContext.Request = req

	suite.albumHandler.CreateAlbum(ginContext)
	suite.Assert().Equal(http.StatusBadRequest, w.Code)
	suite.Assert().JSONEq(
		`{"message":"invalid request"}`,
		w.Body.String(),
	)
}
