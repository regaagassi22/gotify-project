package controllers

import (
	"gotify-project/database"
	"gotify-project/repository"
	"gotify-project/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllGenre(c *gin.Context) {
	var (
		result gin.H
	)

	genres, err := repository.GetAllGenre(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": genres,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertGenre(c *gin.Context) {
	var genre structs.Genre

	err := c.ShouldBindJSON(&genre)
	if err != nil {
		panic(err)
	}

	err = repository.InsertGenre(database.DbConnection, genre)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Genre",
	})
}

func UpdateGenre(c *gin.Context) {
	var genre structs.Genre
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&genre)
	if err != nil {
		panic(err)
	}

	genre.ID = int(id)

	err = repository.UpdateGenre(database.DbConnection, genre)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Genre",
	})
}

func DeleteGenre(c *gin.Context) {
	var genre structs.Genre
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	genre.ID = int(id)

	err = repository.DeleteGenre(database.DbConnection, genre)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Genre",
	})
}
