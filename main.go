package main

import (
	"be-post/handler"
	"be-post/post"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	// set database connection
	DbUser := os.Getenv("DB_USER")
	DbHost := os.Getenv("DB_HOST")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	DbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword,DbHost, DbPort,DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err!= nil {
        log.Fatal(err.Error())
    }

	// auto migrate database
	db.AutoMigrate(&post.Post{})

	// Post
	postRepository := post.NewPostRepository(db)
	postService := post.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)


	router := gin.Default()
	router.Use(cors.Default())
	
	router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

	// route post
	router.GET("/article", postHandler.Index)
	router.POST("/article", postHandler.Store)
	router.GET("/article/:id", postHandler.Show)
	router.PUT("/article/:id", postHandler.Update)
	router.DELETE("/article/:id", postHandler.Destroy)

	router.Run()
}