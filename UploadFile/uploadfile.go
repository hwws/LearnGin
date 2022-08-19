package uploadfile

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UploadFile(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// 上传文件到指定路径
	err := c.SaveUploadedFile(file, "C:\\Users\\13616\\Pictures\\Camera Roll\\b.jpg")
	if err != nil {
		log.Panic("SaveUploadedFileErr:", err)
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
