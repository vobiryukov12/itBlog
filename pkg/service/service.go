package service

import (
	itblog "github.com/Zemavong/itBlog"
	"github.com/Zemavong/itBlog/pkg/repository"
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
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
	Blog
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Blog:          NewBlogService(repos.Blog),
	}
}
