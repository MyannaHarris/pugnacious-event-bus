package actions

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"pugnacious-event-bus/globalVars"
	"pugnacious-event-bus/models"
)

func AddSubscription(subscriptionParams models.SubscriptionParams) (string, error) {
	if (strings.TrimSpace(subscriptionParams.SqsQueue) == "" && strings.TrimSpace(subscriptionParams.ApiUrl) == "") ||
		(strings.TrimSpace(subscriptionParams.EventKey) == "") {
		return "", globalVars.MissingSubscriptionParamsErr
	}

	if strings.TrimSpace(subscriptionParams.SqsQueue) != "" && strings.TrimSpace(subscriptionParams.ApiUrl) != "" {
		return "", globalVars.TooManySubscriptionParamsErr
	}

	var id = uuid.New().String()
	subscription := models.Subscription{
		Id:       id,
		SqsQueue: subscriptionParams.SqsQueue,
		ApiUrl:   subscriptionParams.ApiUrl,
		EventKey: subscriptionParams.EventKey,
		Context:  subscriptionParams.Context,
	}

	fmt.Println(subscription.Id)

	// Store subscription in database

	return id, nil
}

func AddEvent(eventParams models.EventParams) (string, error) {
	if strings.TrimSpace(eventParams.EventKey) == "" {
		return "", globalVars.MissingEventParamsErr
	}

	var id = uuid.New().String()
	event := models.Event{
		Id:       id,
		EventKey: eventParams.EventKey,
		Context:  eventParams.Context,
	}

	fmt.Println(event.Id)

	// Store event in database?

	// Alert subscribed parties

	return id, nil
}
