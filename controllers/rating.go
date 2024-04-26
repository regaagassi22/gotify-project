package controllers

import (
	"gotify-project/database"
	"gotify-project/repository"
	"gotify-project/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRating(c *gin.Context) {
	var (
		result gin.H
	)

	ratings, err := repository.GetAllRating(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": ratings,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertRating(c *gin.Context) {
	var rating structs.Rating

	err := c.ShouldBindJSON(&rating)
	if err != nil {
		panic(err)
	}

	err = repository.InsertRating(database.DbConnection, rating)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Rating",
	})
}

func UpdateRating(c *gin.Context) {
	var rating structs.Rating
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&rating)
	if err != nil {
		panic(err)
	}

	rating.ID = int(id)

	err = repository.UpdateRating(database.DbConnection, rating)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Rating",
	})
}

func DeleteRating(c *gin.Context) {
	var rating structs.Rating
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	rating.ID = int(id)

	err = repository.DeleteRating(database.DbConnection, rating)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Rating",
	})
}
