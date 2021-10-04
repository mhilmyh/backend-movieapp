package server

import (
	"movieapp/entity"
	"movieapp/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	Movie *MovieHandler
	Playing *PlayingHandler
	Viewer *ViewerHandler
}

func NewHandler(conn *gorm.DB) * Handler {
	return &Handler{
		Movie: &MovieHandler{
			movieRepository: repository.NewMovieRepository(conn),
		},
		Playing: &PlayingHandler{
			playingRepository: repository.NewPlayingRepository(conn),
		},
		Viewer: &ViewerHandler{
			viewerRepository: repository.NewViewerRepository(conn),
		},
	}
}

type MovieHandler struct {
	movieRepository repository.MovieRepository
}

func (mh *MovieHandler) GetMovies(c *gin.Context) {
	movies, err := mh.movieRepository.Get()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func (mh *MovieHandler) CreateMovie(c *gin.Context) {
	err := mh.movieRepository.Create(&entity.Movie{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (mh *MovieHandler) UpdateMovie(c *gin.Context) {
	err := mh.movieRepository.Update(&entity.Movie{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (mh *MovieHandler) DeleteMovie(c *gin.Context) {
	err := mh.movieRepository.Delete(0)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

type PlayingHandler struct {
	playingRepository repository.PlayingRepository
}

func (ph *PlayingHandler) GetPlayings(c *gin.Context) {
	playings, err := ph.playingRepository.Get()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"playings": playings})
}

func (ph *PlayingHandler) GetPlayingViewers(c *gin.Context) {
	viewers, err := ph.playingRepository.Get()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"viewers": viewers})
}

func (ph *PlayingHandler) CreatePlaying(c *gin.Context) {
	
}

func (ph *PlayingHandler) CreatePlayingViewer(c *gin.Context) {
	
}

func (ph *PlayingHandler) DeletePlaying(c *gin.Context) {
	
}

func (ph *PlayingHandler) DeletePlayingViewer(c *gin.Context) {
	
}



type ViewerHandler struct {
	viewerRepository repository.ViewerRepository
}

func (vh *ViewerHandler) GetViewers(c *gin.Context) {
	viewers, err := vh.viewerRepository.Get()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": viewers})
}