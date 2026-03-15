package routes

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/look4suman/events-api/models"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		slog.Error("failed to fetch events", "error", err)
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEventById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to convert id param to int64"})
		slog.Error("failed to convert id param to int64", "error", err)
		return
	}

	eventPtr, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		slog.Error("Could not fetch the event", "error", err)
		return
	}
	ctx.JSON(http.StatusOK, eventPtr)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event data"})
		slog.Error("invalid event data", "error", err)
		return
	}
	e, err := event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		slog.Error("failed to create event", "error", err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created successfully for id: " + strconv.FormatInt(e.ID, 10)})
}
