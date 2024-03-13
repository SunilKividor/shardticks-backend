package handlers

import (
	"bookmyshow/models"
	"bookmyshow/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShow(c *gin.Context) {
	var req models.CreateShow
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err = postgresql.CreateShowsQuery(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func AddOrganizerNfts(c *gin.Context) {
	var req models.OrganizerNFTs
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err = postgresql.AddOrganizerNftsQuery(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetAllShows(c *gin.Context) {
	shows, err := postgresql.GetAllUserShowsQuery()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, shows)
}
