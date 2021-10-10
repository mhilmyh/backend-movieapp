package server

import (
	"app/cache"

	"github.com/gin-gonic/gin"
)


func NewRoute(r *gin.RouterGroup)  {
	db := NewDBConnection()
	c := cache.NewRedisConnection()
	h := NewHandler(db, c)

	r.GET("/movies", h.Movie.GetMovies)
	r.POST("/movies", h.Movie.CreateMovie)
	r.PUT("/movies/:id", h.Movie.UpdateMovie)
	r.DELETE("/movies/:id", h.Movie.DeleteMovie)

	r.GET("/playings", h.Playing.GetPlayings)
	r.POST("/playings", h.Playing.CreatePlaying)
	r.DELETE("/playings/:id", h.Playing.DeletePlaying)

	r.GET("/viewers", h.Viewer.GetViewers)
	r.POST("/viewers", h.Viewer.CreateViewer)
	r.DELETE("/viewers/:id", h.Viewer.DeleteViewer)
}