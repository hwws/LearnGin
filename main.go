package main

import (
	formats "LearnGin/Formats"
	uploadfile "LearnGin/UploadFile"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 上传文件
	r.POST("/uploadfile", uploadfile.UploadFile)

	// 结构体转XML
	formats.FormatFile()

	err := r.Run(":8080")
	if err != nil {
		log.Panicln(fmt.Sprintf("Router Run Error:%v\n", err))
	}
}
