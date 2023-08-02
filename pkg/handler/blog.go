package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createBlog(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": "test",
	})
}

func (h *Handler) getAllItems(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": "test",
	})
}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
