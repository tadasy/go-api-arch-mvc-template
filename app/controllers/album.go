package controllers

import (
	"go-api-arch-mvc-template/api"
	"go-api-arch-mvc-template/app/models"
	"go-api-arch-mvc-template/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct{}

func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
	var requestBody api.CreateAlbumJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Warn(err.Error())
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}
	createdAlbum, err := models.CreateAlbum(
		requestBody.Title,
		requestBody.ReleaseDate.Time,
		string(requestBody.Category.Name))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdAlbum)
}

func (h *AlbumHandler) GetAlbumById(c *gin.Context, id int) {
	album, err := models.GetAlbum(id)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	if album == nil {
		c.JSON(http.StatusNotFound, api.ErrorResponse{Message: "Album not found"})
		return
	}
	c.JSON(http.StatusOK, album)
}

func (h *AlbumHandler) UpdateAlbumById(c *gin.Context, id int) {
	var requestBody api.UpdateAlbumByIdJSONRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		logger.Warn(err.Error())
		c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
		return
	}
	album, err := models.GetAlbum(id)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}

	if requestBody.Category != nil {
		album.Category.Name = string(requestBody.Category.Name)
	}
	if requestBody.Title != nil {
		album.Title = *requestBody.Title
	}
	if err := album.Save(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

func (h *AlbumHandler) DeleteAlbumById(c *gin.Context, id int) {
	album := &models.Album{ID: id}
	if err := album.Delete(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
