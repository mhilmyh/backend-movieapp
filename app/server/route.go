package server

import "github.com/gin-gonic/gin"


func NewRoute(r *gin.RouterGroup)  {
	c := NewDBConnection()
	h := NewHandler(c)

	r.GET("/movies", h.Movie.GetMovies)
	r.POST("/movies", h.Movie.CreateMovie)
	r.PUT("/movies", h.Movie.UpdateMovie)
	r.DELETE("/movies", h.Movie.DeleteMovie)

	r.GET("/playings/:playing/viewers", h.Playing.GetPlayingViewers)
	r.GET("/playings", h.Playing.GetPlayings)
	r.POST("/playings/:playing/viewers", h.Playing.CreatePlayingViewer)
	r.POST("/playings", h.Playing.CreatePlaying)
	r.DELETE("/playings/:playing/viewers/:viewer", h.Playing.DeletePlayingViewer)
	r.DELETE("/playings/:playing", h.Playing.DeletePlaying)

	r.GET("/viewers", h.Viewer.GetViewers)
}