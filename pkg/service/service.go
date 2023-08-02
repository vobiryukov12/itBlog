package service

import "github.com/Zemavong/itBlog/pkg/repository"

type Blog interface {
	Create(blogId int, title string, content string, img string)
	GetAll()
	GetById(blogId int)
	Delete(blogId int)
	Update(blogId int, title string, content string, img string)
}

type Service struct {
	Blog
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Blog: NewBlogService(repos.Blog),
	}
}
