package repository

import (
	"github.com/jmoiron/sqlx"
)

type BlogPostgres struct {
	db *sqlx.DB
}

func NewBlogPostgres(db *sqlx.DB) *BlogPostgres {
	return &BlogPostgres{db: db}
}

func (r *BlogPostgres) Create(blogId int, title string, content string, img string) {

}

func (r *BlogPostgres) GetAll() {

}

func (r *BlogPostgres) GetById(blogId int) {

}

func (r *BlogPostgres) Update(blogId int, title string, content string, img string) {

}

func (r *BlogPostgres) Delete(blogId int) {

}
