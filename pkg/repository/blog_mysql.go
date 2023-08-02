package repository

import "database/sql"

type BlogMysql struct {
	db *sql.DB
}

func NewBlogMysql(db *sql.DB) *BlogMysql {
	return &BlogMysql{db: db}
}

func (r *BlogMysql) Create(blogId int, title string, content string, img string) {

}

func (r *BlogMysql) GetAll() {

}

func (r *BlogMysql) GetById(blogId int) {

}

func (r *BlogMysql) Update(blogId int, title string, content string, img string) {

}

func (r *BlogMysql) Delete(blogId int) {

}
