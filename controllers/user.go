package controllers

import (
	"gotify-project/database"
	"gotify-project/repository"
	"gotify-project/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	var (
		result gin.H
	)

	users, err := repository.GetAllUser(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertUser(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert User",
	})
}

func UpdateUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	user.ID = int(id)

	err = repository.UpdateUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update User",
	})
}

func DeleteUser(c *gin.Context) {
	var user structs.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	user.ID = int(id)

	err = repository.DeleteUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete User",
	})
}
