package controller

import (
	"demo/entity"
	"demo/service"
	"demo/validators"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
	validate *validator.Validate
}

func New(service service.VideoService) VideoController {
	validate := validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{service: service, validate: validate}
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = c.validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	video.ID = id

	err = c.validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Update(video)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error  {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	video.ID = id
	c.service.Delete(video)
	return nil
}



func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
