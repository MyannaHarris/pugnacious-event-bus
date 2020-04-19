package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"pugnacious-event-bus/actions"
	"pugnacious-event-bus/models"
)

func main() {
	router := SetupRouter()
	router.Run(":8000")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.POST("/subscriptions", func(context *gin.Context) {
		var subscriptionParams models.SubscriptionParams
		err := context.BindJSON(&subscriptionParams)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := actions.AddSubscription(subscriptionParams)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusCreated, gin.H{
			"id": id,
		})
	})

	router.POST("/events", func(c *gin.Context) {
		var eventParams models.EventParams
		err := c.BindJSON(&eventParams)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := actions.AddEvent(eventParams)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id": id,
		})
	})

	return router
}
