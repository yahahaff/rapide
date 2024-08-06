package cloudflare

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/backend/internal/controllers/api"
	"github.com/yahahaff/rapide/backend/internal/response"
	"github.com/yahahaff/rapide/backend/pkg/cloudflare"
	"mime/multipart"
)

type R2Controller struct {
	api.BaseAPIController
}

// R2Config holds the configuration for accessing Cloudflare R2
type R2Config struct {
	Endpoint    string
	AccessKeyID string
	SecretKey   string
	BucketName  string
}

func (d *R2Controller) GetR2List(c *gin.Context) {

	result, err := cloudflare.R2.ListObjects(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	response.OK(c, result)
}

func (d *R2Controller) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.Abort400(c, "Failed to get file")
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	key := c.PostForm("key")
	if key == "" {
		key = header.Filename
	}

	err = cloudflare.R2.UploadFile(context.TODO(), file, key)
	if err != nil {
		response.Abort400(c, "Failed to upload file")
		return
	}

	response.OK(c, "File uploaded successfully")
}

func (d *R2Controller) DeleteFile(c *gin.Context) {
	key := c.Param("key")
	err := cloudflare.R2.DeleteObject(context.TODO(), key)
	if err != nil {
		response.Abort400(c, "Failed to delete file")
		return
	}

	response.OK(c, "File deleted successfully")
}

func (d *R2Controller) GetFileMetadata(c *gin.Context) {
	key := c.Param("key")
	metadata, err := cloudflare.R2.GetObjectMetadata(context.TODO(), key)
	if err != nil {
		response.Abort400(c, "Failed to get file metadata")
		return
	}

	response.OK(c, metadata)
}
