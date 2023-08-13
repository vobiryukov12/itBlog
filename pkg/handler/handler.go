package handler

import (
	"github.com/Zemavong/itBlog/pkg/service"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.HTMLRender = ginview.Default()

	auth := router.Group("/blogs")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{

		blogs := api.Group("/blogs")
		{
			//blogs.POST("/", h.createList)
			blogs.GET("/", h.createBlog)
			//blogs.GET("/:id", h.getItemById)
			//blogs.PUT("/:id", h.updateList)
			//blogs.DELETE("/:id", h.deleteList)
		}
	}

	return router
}

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hi" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method is not supported.", http.StatusNotFound)
// 		return
// 	}

// 	fmt.Fprintf(w, "Hello!")
// }
