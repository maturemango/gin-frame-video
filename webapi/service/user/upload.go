package user

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"
	"mime/multipart"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	maxFile  int64 = 1024 * 1024 * 10 // 10MB
	maxVideo int64 = 1024 * 1024 * 20 // 20MB
)

// 设置传送视频的大小以及判断是否为视频文件
func UploadUserVideo(c *gin.Context) {
	var data model.UploadUserVideo
	c.ShouldBind(&data)
	data = uploadVideoMessage(data)
	file := c.Request.MultipartForm.File["video"]
	if err := verifyVideoFile(file[0]); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	if c.Request.ContentLength > maxVideo {
		handlers.Base.Fail(c, 413, fmt.Errorf("payload too large"))
		return
	}
	dst := path.Join("./bin/videos", data.VideoNo+"."+strings.Split(file[0].Filename, ".")[1]) // 保存在服务器对应路径下  一般都是保存在云端
	if err := c.SaveUploadedFile(file[0], dst); err != nil {
		handlers.Base.Fail(c, 400, fmt.Errorf("save upload file failed:%v", err))
		return
	}
	if _, err := conn.GetEngine().Insert(&data); err != nil {
		handlers.Base.Fail(c, 500, err)
		return
	}
	handlers.Base.OK(c, "upload success")
}

func verifyVideoFile(file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		return fmt.Errorf("file open failed:%v", err)
	}
	defer f.Close()

	buffer := make([]byte, 512)
	f.Read(buffer)
	contentType := http.DetectContentType(buffer)
	if strings.HasPrefix(contentType, "video") {
		return nil
	} else {
		return fmt.Errorf("file is not video")
	}
}

func uploadVideoMessage(data model.UploadUserVideo) model.UploadUserVideo {
	no := utils.Config.Video.No + fmt.Sprint(time.Now().UnixNano()) + utils.RandomVideoNo(10)
	data.VideoNo = no
	data.UserId = handlers.Identity()
	data.UploadTime = time.Now()
	return data
}
