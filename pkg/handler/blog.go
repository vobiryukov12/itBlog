package handler

import (
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func (h *Handler) createBlog(c *gin.Context) {
	paths := []string{
		"todo.tmpl",
	}

	t := template.Must(template.New("html-tmpl").ParseFiles(paths...))
	err := t.Execute(os.Stdout, ToDo.User)
	if err != nil {
		panic(err)
	}
}

func (h *Handler) getAllItems(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
