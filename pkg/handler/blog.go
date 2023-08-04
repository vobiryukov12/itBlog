package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createBlog(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Main website",
	})
}

func (h *Handler) getAllItems(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
