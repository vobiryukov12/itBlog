package repository

import (
	itblog "github.com/Zemavong/itBlog"
	"github.com/jmoiron/sqlx"
)

type Blog interface {
	Create(blogId int, title string, content string, img string)
	GetAll()
	GetById(blogId int)
	Delete(blogId int)
	Update(blogId int, title string, content string, img string)
}

type Authorization interface {
	CreateUser(user itblog.User) (int, error)
	GetUser(username, password string) (itblog.User, error)
}

type Repository struct {
	Authorization
	Blog
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Blog:          NewBlogPostgres(db),
	}
}
