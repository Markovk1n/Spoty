package handler

import (
	"html/template"

	"github.com/gin-gonic/gin"

	"github.com/markovk1n/spoty/internal/service"
)

type Handler struct {
	tmpl     *template.Template
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		tmpl:     template.Must(template.ParseGlob("./templates/*.html")),
		services: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.SetHTMLTemplate(h.tmpl)

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", h.home)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.singIn)
	}

	albums := router.Group("/albums")
	{
		albums.GET("/", h.getAllAlbums)
		albums.GET("/:id", h.getAlbumById)

		album := albums.Group(":id/album")
		{
			album.GET("/", h.getAlbumTracks)

			album.POST("/create_review", h.create_album_review)
		}
	}

	track := router.Group(":id/track")
	{
		track.GET("/", h.getTrack)
		track.POST("/create_review", h.create_track_review)
	}

	// artists := router.Group("/artists")
	// {

	// }

	reviews := router.Group("/review")
	{
		reviews.GET("/", h.getAllReviews)

	}

	// singles := router.Group("/singles")
	// {

	// }

	return router
}
