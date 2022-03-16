package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 上传文件处理
func uploadfile(c *gin.Context) {
	// 接收文件
	file, err := c.FormFile("file...")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 储存文件
	log.Println(file.Filename)
	dst := fmt.Sprintf("./%v", file.Filename)
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v uploaded", file.Filename),
	})
}

// 上传多个文件处理
func uploadfiles(c *gin.Context) {
	// 接收文件
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	files := form.File["file..."]
	// 储存文件
	for index, file := range files {
		log.Println(file.Filename)
		dst := fmt.Sprintf("./%v_%d", file.Filename, index)
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d files uploaded", len(files)),
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", uploadfiles)
	r.Run(":9090")
}
