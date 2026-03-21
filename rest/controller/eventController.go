package controller

import (
	"net/http"

	"example.com/rest/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetEvents(ctx *gin.Context) {
	allEvents, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get events",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": allEvents,
		"meta": gin.H{
			"total": len(allEvents),
			"page":  1,
			"size":  10,
		},
	})
}

func CreateEvent(ctx *gin.Context) {
	data, exists := ctx.Get("data")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get data from validation middleware",
		})
		return
	}

	event := data.(models.Event)
	event.ID = uuid.New()
	event.UserId = uuid.New()

	err := event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create event",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    event,
		"message": "Event created successfully",
	})
}

func UpdateEvent(ctx *gin.Context) {
	data, exists := ctx.Get("data")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get data from validation middleware",
		})
		return
	}

	eventId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Event with ID " + ctx.Param("id") + " not found",
		})
		return
	}

	event := data.(models.Event)
	event.ID = eventId

	err = event.Update()

	if err != nil {
		if err.Error() == "no rows affected" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Event not found",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to update event",
				"error":   err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    event,
		"message": "Event updated successfully",
	})
}

func DeleteEvent(ctx *gin.Context) {
	eventId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Event with ID " + ctx.Param("id") + " not found",
		})
		return
	}

	err = models.DeleteEvent(eventId)
	if err != nil {
		if err.Error() == "no rows affected" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Event not found",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to delete event",
				"error":   err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
