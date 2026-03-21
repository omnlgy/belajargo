package router

import (
	"example.com/rest/controller"
	"example.com/rest/helpers"
	"example.com/rest/models"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", controller.GetEvents)
	router.POST("/events", helpers.BodyValidation[models.Event](), controller.CreateEvent)
	router.PUT("/events/:id", helpers.BodyValidation[models.Event](), controller.UpdateEvent)
	router.DELETE("/events/:id", controller.DeleteEvent)
}
