package controllers

import (
	"gotify-project/database"
	"gotify-project/repository"
	"gotify-project/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMusic(c *gin.Context) {
	var (
		result gin.H
	)

	musics, err := repository.GetAllMusic(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": musics,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertMusic(c *gin.Context) {
	var music structs.Music

	err := c.ShouldBindJSON(&music)
	if err != nil {
		panic(err)
	}

	err = repository.InsertMusic(database.DbConnection, music)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Music",
	})
}

func UpdateMusic(c *gin.Context) {
	var music structs.Music
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&music)
	if err != nil {
		panic(err)
	}

	music.ID = int(id)

	err = repository.UpdateMusic(database.DbConnection, music)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Music",
	})
}

func DeleteMusic(c *gin.Context) {
	var music structs.Music
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	music.ID = int(id)

	err = repository.DeleteMusic(database.DbConnection, music)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Music",
	})
}
