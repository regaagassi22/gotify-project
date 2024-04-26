package controllers

import (
	"gotify-project/database"
	"gotify-project/repository"
	"gotify-project/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRole(c *gin.Context) {
	var (
		result gin.H
	)

	roles, err := repository.GetAllRole(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": roles,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertRole(c *gin.Context) {
	var role structs.Role

	err := c.ShouldBindJSON(&role)
	if err != nil {
		panic(err)
	}

	err = repository.InsertRole(database.DbConnection, role)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Role",
	})
}

func UpdateRole(c *gin.Context) {
	var role structs.Role
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&role)
	if err != nil {
		panic(err)
	}

	role.ID = int(id)

	err = repository.UpdateRole(database.DbConnection, role)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Role",
	})
}

func DeleteRole(c *gin.Context) {
	var role structs.Role
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	role.ID = int(id)

	err = repository.DeleteRole(database.DbConnection, role)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Role",
	})
}
