package blog

import "github.com/skylee/gin-web-layout/internal/storage"

type Repository interface {
}

type BlogRepo struct {
	db *storage.Storage
}

func NewBlogRepo(storage *storage.Storage) *BlogRepo {
	return &BlogRepo{storage}
}
