package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// Structs used by the API
type SubscriptionParams struct {
	SqsQueue    string `json:"sqsqueue"`
	ApiUrl      string `json:"apiurl"`
	EventKey   string `json:"eventkey"`
	Context      string `json:"context"`
}

type EventParams struct {
	EventKey   string `json:"eventkey"`
	Context      string `json:"context"`
}

// Structs used by the DDB model
type Subscription struct {
	Id          string
	SqsQueue    string
	ApiUrl      string
	EventKey    string
	Context     string
}

type Event struct {
	Id          string
	EventKey    string
	Context     string
}

func addSubscription(subscriptionParams SubscriptionParams) (string, error) {
	var id = uuid.New().String()
	subscription := Subscription{
		Id:          id,
		SqsQueue:    subscriptionParams.SqsQueue,
		ApiUrl:      subscriptionParams.ApiUrl,
		EventKey:    subscriptionParams.EventKey,
		Context:     subscriptionParams.Context,
	}

	fmt.Println(subscription.Id)

	// Store subscription in database

	return id, nil
}

func addEvent(eventParams EventParams) (string, error) {
	var id = uuid.New().String()
	event := Event{
		Id:          id,
		EventKey:    eventParams.EventKey,
		Context:     eventParams.Context,
	}

	fmt.Println(event.Id)

	// Store event in database?

	// Alert subscribed parties

	return id, nil
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.POST("/subscriptions", func(context *gin.Context) {
		var subscriptionParams SubscriptionParams
		err := context.BindJSON(&subscriptionParams)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Println(subscriptionParams.SqsQueue)

		id, err := addSubscription(subscriptionParams)
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
		var eventParams EventParams
		err := c.BindJSON(&eventParams)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Println(eventParams.EventKey)

		id, err := addEvent(eventParams)
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

	router.Run(":8000")
}
