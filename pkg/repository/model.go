package repository

import (
	"gopkg.in/reform.v1"
)

type Repository struct {
	DB *reform.DB
}

type News struct {
	Id         int    `json:"Id",db:"ID"`
	Title      string `json:"Title",db:"Title"`
	Content    string `json:"Content",db:"Content""`
	Categories []int  `json:"Categories"`
}
