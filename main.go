package main

import (
	"log"

	"demo.com/blog/models"
	"github.com/gin-gonic/gin"
)

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()

	// API Blogs
	router := r.Group("/blogs")
	{
		router.POST("/create", postBlog)
		router.GET("/", readBlog)
		router.POST("/update/:id", updateBlog)
		router.GET("/delete/:id", deleteBlog)
	}

	// By default it serves on :8080
	r.Run()
}

func postBlog(c *gin.Context) {
	c.JSON(200, gin.H{"message": "A new Record Created!"})
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readBlog(c *gin.Context) {
	blogs, err := models.GetBlogs()
	checkErr(err)

	if blogs == nil {
		c.JSON(404, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(200, gin.H{"data": blogs})
	}
}

func updateBlog(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Updated!"})
}
func deleteBlog(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Record Deleted!"})
}
