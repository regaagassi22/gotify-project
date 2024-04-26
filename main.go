package main

import (
	"database/sql"
	"fmt"
	"gotify-project/controllers"
	"gotify-project/database"
	"gotify-project/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file env")
	} else {
		fmt.Println("SUCCESS read file env")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Failed connected to database")
		panic(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	//Router GIN
	router := gin.Default()
	router.GET("/users", middleware.BasicAuth, middleware.Privilage, controllers.GetAllUser)
	router.POST("/users", controllers.InsertUser)
	router.PUT("/users/:id", middleware.BasicAuth, middleware.Privilage, controllers.UpdateUser)
	router.DELETE("/users/:id", middleware.BasicAuth, middleware.Privilage, controllers.DeleteUser)

	router.GET("/roles", middleware.BasicAuth, middleware.Privilage, controllers.GetAllRole)
	router.POST("/roles", middleware.BasicAuth, middleware.Privilage, controllers.InsertRole)
	router.PUT("/roles/:id", middleware.BasicAuth, middleware.Privilage, controllers.UpdateRole)
	router.DELETE("/roles/:id", middleware.BasicAuth, middleware.Privilage, controllers.DeleteRole)

	router.GET("/musics", middleware.BasicAuth, controllers.GetAllMusic)
	router.POST("/musics", middleware.BasicAuth, middleware.Privilage, controllers.InsertMusic)
	router.PUT("/musics/:id", middleware.BasicAuth, middleware.Privilage, controllers.UpdateMusic)
	router.DELETE("/musics/:id", middleware.BasicAuth, middleware.Privilage, controllers.DeleteMusic)

	router.GET("/genres", middleware.BasicAuth, controllers.GetAllGenre)
	router.POST("/genres", middleware.BasicAuth, middleware.Privilage, controllers.InsertGenre)
	router.PUT("/genres/:id", middleware.BasicAuth, middleware.Privilage, controllers.UpdateGenre)
	router.DELETE("/genres/:id", middleware.BasicAuth, middleware.Privilage, controllers.DeleteGenre)

	router.GET("/ratings", middleware.BasicAuth, controllers.GetAllRating)
	router.POST("/ratings", middleware.BasicAuth, controllers.InsertRating)
	router.PUT("/ratings/:id", middleware.BasicAuth, controllers.UpdateRating)
	router.DELETE("/ratings/:id", middleware.BasicAuth, controllers.DeleteRating)

	router.Run("localhost:8080")

}
