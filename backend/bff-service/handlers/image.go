package handlers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"bff-service/clients"
	proto "bff-service/proto/image"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		//log.Println("[UploadImage] ❌ Invalid form file:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image file"})
		return
	}
	defer file.Close()

	filename := header.Filename
	//log.Printf("[UploadImage] 📤 Received file: %s", filename)

	if !strings.HasSuffix(strings.ToLower(filename), ".jpg") {
		//log.Println("[UploadImage] ❌ Only .jpg files are allowed")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only .jpg files are allowed"})
		return
	}

	contentType := header.Header.Get("Content-Type")
	if contentType != "image/jpeg" {
		//log.Println("[UploadImage] ❌ Only JPEG images are allowed")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only JPEG images are allowed"})
		return
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		//log.Println("[UploadImage] ❌ Error reading image data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading image data"})
		return
	}

	entityId := c.PostForm("entity_id")
	entityType := c.PostForm("entity_type")
	//log.Printf("[UploadImage] 🧩 entity_id=%s, entity_type=%s", entityId, entityType)

	res, err := clients.ImageClient.UploadImage(c, &proto.UploadImageRequest{
		EntityId:   entityId,
		EntityType: entityType,
		FileType:   contentType,
		ImageData:  buf.Bytes(),
	})
	if err != nil {
		//log.Println("[UploadImage] ❌ gRPC error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gRPC UploadImage failed: " + err.Error()})
		return
	}

	//log.Printf("[UploadImage] ✅ Upload successful: image_id=%s", res.ImageId)
	c.JSON(http.StatusOK, gin.H{
		"image_id": res.ImageId,
	})
}

func GetImage(c *gin.Context) {
	imageID := c.Param("id")
	//log.Printf("[GetImage] 🔍 Fetching image: image_id=%s", imageID)

	res, err := clients.ImageClient.GetImage(c, &proto.GetImageRequest{ImageId: imageID})
	if err != nil {
		log.Printf("[GetImage] ❌ Image not found: %s", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	//log.Printf("[GetImage] ✅ Image fetched: image_id=%s", imageID)
	c.Data(http.StatusOK, res.FileType, res.ImageData)
}

func DeleteImage(c *gin.Context) {
	imageID := c.Param("id")
	log.Printf("[DeleteImage] 🗑️ Deleting image: image_id=%s", imageID)

	res, err := clients.ImageClient.DeleteImage(c, &proto.DeleteImageRequest{ImageId: imageID})
	if err != nil {
		//log.Printf("[DeleteImage] ❌ Failed to delete image: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
		return
	}

	//log.Printf("[DeleteImage] ✅ Deleted image: image_id=%s", imageID)
	c.JSON(http.StatusOK, gin.H{"status": res.Status})
}
