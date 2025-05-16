package models

import (
	"encoding/json"
	"go-api-arch-mvc-template/api"
	"go-api-arch-mvc-template/pkg"
	"time"
)

type Album struct {
	ID          int
	Title       string
	ReleaseDate time.Time
	CategoryID  int
	Category    *Category
}

func (a *Album) Anniversary(clock pkg.Clock) int {
	now := clock.Now()
	years := now.Year() - a.ReleaseDate.Year()
	releaseDay := pkg.GetAdjustedReleaseDay(a.ReleaseDate, now)
	if now.YearDay() < releaseDay {
		years--
	}
	return years
}

func (a *Album) MarshalJSON() ([]byte, error) {
	return json.Marshal(&api.AlbumResponse{
		Id:          a.ID,
		Title:       a.Title,
		Anniversary: a.Anniversary(pkg.RealClock{}),
		ReleaseDate: api.ReleaseDate{Time: a.ReleaseDate},
		Category: api.Category{
			Id:   &a.Category.ID,
			Name: api.CategoryName(a.Category.Name),
		},
	})
}

func CreateAlbum(title string, releaseDate time.Time, categoryName string) (*Album, error) {
	category, err := GetOrCreateCategory(categoryName)
	if err != nil {
		return nil, err
	}
	album := &Album{
		Title:       title,
		ReleaseDate: releaseDate,
		CategoryID:  category.ID,
		Category:    category,
	}
	tx := DB.Create(album)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return album, nil
}

func GetAlbum(id int) (*Album, error) {
	album := &Album{}
	tx := DB.Preload("Category").First(album, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return album, nil
}

func (a *Album) Save() error {
	category, err := GetOrCreateCategory(a.Category.Name)
	if err != nil {
		return err
	}
	a.Category = category
	a.CategoryID = category.ID
	tx := DB.Save(&a)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (a *Album) Delete() error {
	tx := DB.Where("id = ?", &a.ID).Delete(&a)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
