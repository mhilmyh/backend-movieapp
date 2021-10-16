package server

import (
	"app/db"
	"app/entity"
	"app/helper"
	"app/repository"
	"net/http"
	"strconv"
	"time"

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
		helper.SendResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func (mh *MovieHandler) CreateMovie(c *gin.Context) {
	r, err := strconv.Atoi(c.PostForm("rating"))
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	m := entity.Movie{
		Title: c.PostForm("title"),
		Desc: c.PostForm("desc"),
		Rating: int8(r),
		Author: c.PostForm("author"),
	}

	err = mh.movieRepository.Create(&m)
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie": m})
}

func (mh *MovieHandler) UpdateMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	r, err := strconv.Atoi(c.PostForm("rating"))
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	var m entity.Movie
	db.Instance.First(&m, id)

	m.Title = c.PostForm("title")
	m.Desc = c.PostForm("desc")
	m.Rating = int8(r)
	m.Author = c.PostForm("author")

	err = mh.movieRepository.Update(&m)
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie": m})
}

func (mh *MovieHandler) DeleteMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	err = mh.movieRepository.Delete(id)
	if err != nil {
		helper.SendResponseError(c, err)
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
		helper.SendResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"playings": playings})
}

func (ph *PlayingHandler) CreatePlaying(c *gin.Context) {
	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	movieID, err := strconv.Atoi(c.PostForm("movie_id"))
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	start, err := time.Parse(time.RFC3339, c.PostForm("start"))
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	end, err := time.Parse(time.RFC3339, c.PostForm("end"))
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	p := entity.Playing{
		Price: int64(price),
		Start: start,
		End: end,
		MovieID: movieID,
	}

	err = ph.playingRepository.Create(&p)
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"playing": p})
}

func (ph *PlayingHandler) DeletePlaying(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	err = ph.playingRepository.Delete(id)
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}



type ViewerHandler struct {
	viewerRepository repository.ViewerRepository
}

func (vh *ViewerHandler) GetViewers(c *gin.Context) {
	viewers, err := vh.viewerRepository.Get()
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"viewers": viewers})
}

func (vh *ViewerHandler) CreateViewer(c *gin.Context)  {
	p, err := strconv.Atoi(c.PostForm("playing_id"))
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	v := entity.Viewer{
		Name: c.PostForm("name"),
		PlayingID: p,
	}
	err = vh.viewerRepository.Create(&v)
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"viewers": v})
}

func (vh *ViewerHandler) DeleteViewer(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	err = vh.viewerRepository.Delete(id)
	if err != nil {
		helper.SendResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}