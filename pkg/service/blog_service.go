package service

import "github.com/Zemavong/itBlog/pkg/repository"

type BlogService struct {
	repo repository.Blog
}

func NewBlogService(repo repository.Blog) *BlogService {
	return &BlogService{repo: repo}
}

func (s *BlogService) Create(blogId int, title string, content string, img string) {

}

func (s *BlogService) GetAll() {

}

func (s *BlogService) GetById(blogId int) {

}

func (s *BlogService) Delete(blogId int) {

}

func (s *BlogService) Update(blogId int, title string, content string, img string) {

}
