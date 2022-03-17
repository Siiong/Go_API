package service

import (
	"log"
	"main/db"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
	db *db.Storage
}

func New(db *db.Storage) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetSneakers(c *gin.Context) {
	id := c.Param("id")
	snkrs, err := s.db.Sneakers.GetSneakers(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sneakers": snkrs,
	})
}

func (s *Service) GetAllSneakers(c *gin.Context) {
	list, err := s.db.Sneakers.GetAllSneakers()
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sneakers": list,
	})
}

/* func (s *Service) GetAllCollection(c *gin.Context) {
	list, err := s.db.Sneakers.GetAllCollection()
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"collection": list,
	})
} */

func (s *Service) CreateSneakers(c *gin.Context) {
	var item model.Sneakers
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	_, err = s.db.Sneakers.CreateSneakers(&item)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sneakers": item,
	})
}

func (s *Service) DeleteSneakers(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error id": id,
		})
		return
	}
	err := s.db.Sneakers.DeleteSneakers(id)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"delete": id,
	})
}

//TODO Write func name to get no error
