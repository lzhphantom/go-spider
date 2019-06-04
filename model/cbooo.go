package model

import "github.com/jinzhu/gorm"

type ChoooRankListInfo struct {
	gorm.Model
	BoxOffice    string `json:"box_office"`
	MovieID      string `json:"movie_id"`
	MovieEnName  string `json:"movie_en_name"`
	MovieName    string `json:"movie_name"`
	Ranking      string `json:"ranking"`
	DefaultImage string `json:"default_image"`
	ReleaseYear  string `json:"release_year"`
	Country      string `json:"country"`
}

func (ChoooRankListInfo) TableName() string {
	return "cbooo_movies"
}
