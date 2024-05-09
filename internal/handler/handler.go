package handler

import (
	"html/template"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/markovk1n/spoty/internal/service"
)

var token string

type Handler struct {
	tmpl *template.Template

	services     *service.Service
	ClientID     string
	ClientSecret string
}
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func NewHandler(service *service.Service, CID, CS string) *Handler {
	return &Handler{
		tmpl: template.Must(template.ParseGlob("./templates/*.html")),
		// static: template.Must(template.ParseGlob("./static/*.css")),
		services:     service,
		ClientID:     CID,
		ClientSecret: CS,
	}
}
func JSHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/x-javascript")
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.SetHTMLTemplate(h.tmpl)

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", h.home)

	auth := router.Group("/auth")
	{
		auth.GET("/", h.authPage)
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.singIn)
	}

	albums := router.Group("/albums")
	{
		albums.Use(AuthMiddleware())
		albums.GET("/", h.getAllAlbums)
		albums.GET(":id", h.getAlbumById)
		albums.POST("/:id", AuthMiddleware(), h.CreateCommentForAlbum)
	}

	track := router.Group("/track")
	{
		track.GET("/:id", h.getTrack)
		track.POST("/:id", AuthMiddleware(), h.CreateCommentForTrack)
	}

	artists := router.Group("/artists")
	{
		artists.GET("/", h.getArtists)
		artists.GET("/:id", h.getArtistById)
	}

	return router
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session_token")
		if err != nil {
			// Обработка ошибки, если токен отсутствует в cookie
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session token not found"})
			return
		}

		// Добавляем токен в заголовок "Authorization"
		c.Request.Header.Set("Authorization", "Bearer "+cookie)

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}
		c.Request.Header.Set("Authorization", tokenString)
		s := tokenString[7:]
		id, err := service.ParseToken(s)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", id)
		c.Next()
	}
}
