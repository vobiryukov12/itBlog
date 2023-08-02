package repository

import "database/sql"

type Blog interface {
	Create(blogId int, title string, content string, img string)
	GetAll()
	GetById(blogId int)
	Delete(blogId int)
	Update(blogId int, title string, content string, img string)
}

type Repository struct {
	Blog
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Blog: NewBlogMysql(db),
	}
}
