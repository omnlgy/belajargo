package controller

import (
	"errors"
	"fmt"
	"net/http"

	"example.com/rest/helpers"
	"example.com/rest/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func validateParsedToken(ctx *gin.Context) *helpers.PayloadToken {
	claim, claimExists := ctx.Get("claim")

	if !claimExists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get claim from validation middleware",
		})
		return nil
	}

	tokenClaim, ok := claim.(*helpers.PayloadToken)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unexpected claim type",
		})
		return nil
	}

	return tokenClaim
}

func validateParsedDataEvent(ctx *gin.Context) *models.Event {
	data, exists := ctx.Get("data")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get data from validation middleware",
		})
		return nil
	}

	event, ok := data.(models.Event)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unexpected event type",
		})
		return nil
	}

	return &event
}

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
	tokenClaim := validateParsedToken(ctx)
	event := validateParsedDataEvent(ctx)

	if tokenClaim == nil || event == nil {
		return
	}

	event.ID = uuid.New()
	event.UserId = uuid.MustParse(tokenClaim.UserId)

	err := event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create event",
			"error":   err.Error(),
		})
		return
	}
	cookieValue, err := ctx.Cookie("tes")
	fmt.Println("cookie", cookieValue, err)
	ctx.JSON(http.StatusOK, gin.H{
		"data":    event,
		"message": "Event created successfully",
	})
}

func UpdateEvent(ctx *gin.Context) {
	tokenClaim := validateParsedToken(ctx)
	event := validateParsedDataEvent(ctx)

	if tokenClaim == nil || event == nil {
		return
	}

	eventId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Event with ID " + ctx.Param("id") + " not found",
		})
		return
	}

	eventFromDb, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Event with ID " + eventId.String() + " not found",
		})
		return
	}

	if eventFromDb.UserId != uuid.MustParse(tokenClaim.UserId) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "You are not authorized to update this event",
		})
		return
	}

	err = event.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update event",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    event,
		"message": "Event updated successfully",
	})
}

func DeleteEvent(ctx *gin.Context) {
	tokenClaim := validateParsedToken(ctx)
	if tokenClaim == nil {
		return
	}

	eventId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Event with ID " + ctx.Param("id") + " not found",
		})
		return
	}

	eventFromDb, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Event with ID " + eventId.String() + " not found",
		})
		return
	}

	if eventFromDb.UserId != uuid.MustParse(tokenClaim.UserId) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "You are not authorized to delete this event",
		})
		return
	}

	err = models.DeleteEvent(eventId)
	if err != nil {
		if errors.Is(err, models.ErrNoRowsAffected) {
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
