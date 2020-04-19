package actions

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"

	"pugnacious-event-bus/globalVars"
	"pugnacious-event-bus/models"
)

func AddSubscription(subscriptionParams models.SubscriptionParams) (string, error) {
	if strings.TrimSpace(subscriptionParams.SqsQueue) == "" || strings.TrimSpace(subscriptionParams.EventKey) == "" {
		return "", globalVars.MissingSubscriptionParamsErr
	}

	var id = uuid.New().String()
	subscription := models.Subscription{
		Id:       id,
		SqsQueue: subscriptionParams.SqsQueue,
		EventKey: subscriptionParams.EventKey,
		Context:  subscriptionParams.Context,
	}

	fmt.Println(subscription.Id)

	// Store subscription in database
	globalVars.SubscriptionsMap[subscription.Id] = subscription

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
	globalVars.EventsMap[event.Id] = event

	// Alert subscribed parties
	err := AlertSubscribersToEvent(event)

	return id, err
}

var AlertSubscribersToEvent = func(event models.Event) error {
	message := &models.EventMessage{
		Id:       event.Id,
		EventKey: event.EventKey,
		Context:  event.Context,
		Source:   "pugnacious-event-bus",
	}

	messageString, err := json.Marshal(message)
	if err != nil {
		return err
	}

	for _, subscription := range globalVars.SubscriptionsMap {
		if subscription.EventKey == event.EventKey {
			_, err = globalVars.SqsClient.SendMessage(&sqs.SendMessageInput{
				MessageAttributes: map[string]*sqs.MessageAttributeValue{
					"SourceHostname": &sqs.MessageAttributeValue{
						DataType:    aws.String("String"),
						StringValue: aws.String(globalVars.SOURCE_HOSTNAME),
					},
				},
				MessageBody: aws.String(string(messageString)),
				QueueUrl:    aws.String(subscription.SqsQueue),
			})
		}
	}

	return err
}
