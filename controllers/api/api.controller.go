package api

import (
  "github.com/misaliperver/golish/helpers"
  "github.com/misaliperver/golish/constants"
  
  "fmt"
  "log"
  "strings"
  "github.com/gofiber/fiber/v2"
  "github.com/google/uuid"
)

// Hello hanlde api status
func ApiMain(c *fiber.Ctx) error {
	return c.JSON(helpers.ResponseTransform("api routers", false))
}

func DynamicNameHelloApi(c *fiber.Ctx) error {
  name := c.Params("name")
  message := "Merhaba " + name + ". Hoş geldiniz"
	return c.JSON(helpers.ResponseTransform(message, false))
}

func UploadFile(c *fiber.Ctx) error {
  file, err := c.FormFile("image")

  if err != nil {
    log.Println("image save error --> ", err)
    return c.JSON(helpers.ResponseTransform("Server Error", true))
  }

  // generate new uuid for image name
  // remove "- from imageName"
  // extract image extension from original file filename
  // generate image from filename and extension
  uniqueId := uuid.New()
  filename := strings.Replace(uniqueId.String(), "-", "", -1)
  fileExt := strings.Split(file.Filename, ".")[1]
  imageName := fmt.Sprintf("%s.%s", filename, fileExt)
  
  // save image to ./images dir 
  err = c.SaveFile(file, fmt.Sprintf("./tmp/%s", imageName))

  if err != nil {
    log.Println("image save error --> ", err)
    return c.JSON(helpers.ResponseTransform("Server Error", true))
  }

  imageUrl := fmt.Sprintf("%s/images/%s", constants.ServerUrl, imageName)

  // create meta data and send to client
  data := map[string]interface{}{
    "imageName": imageName,
    "imageUrl":  imageUrl,
    "header":    file.Header,
    "size":      file.Size,
  }
  
return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

