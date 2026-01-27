package routes

import (
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	// to be added later
	// userId := context.GetInt64("userId")
	// eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
	// 	return
	// }

	// event, err := models.GetEventById(eventId)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"message": "could not process request"})
	// 	return
	// }

}

func cancelRegistrationForEvent(context *gin.Context) {

}
