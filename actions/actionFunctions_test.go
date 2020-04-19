package actions

import (
	"regexp"
	"testing"

	"pugnacious-event-bus/globalVars"
	"pugnacious-event-bus/models"
)

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func TestAddSubscription_ValidInput(t *testing.T) {
	var tests = []struct {
		name              string
		subscriptionParam models.SubscriptionParams
	}{
		{"Valid SQSQueue",
			models.SubscriptionParams{
				SqsQueue: "test-queue",
				EventKey: "test-eventkey",
				Context:  "test-context",
			}},
	}

	for _, test := range tests {
		id, err := AddSubscription(test.subscriptionParam)

		if err != nil {
			t.Errorf("AddSubscription call failed for: %v", test)
		}

		if !IsValidUUID(id) {
			t.Errorf("The UUID from AddSubscription is invalid for: %v", test)
		}

		subParamExpected := test.subscriptionParam

		// Check the value stored in the database
		if val, ok := globalVars.SubscriptionsMap[id]; ok {
			if val.Id != id || val.SqsQueue != subParamExpected.SqsQueue || val.EventKey != subParamExpected.EventKey ||
				val.Context != subParamExpected.Context {
				t.Errorf("The subscription object in the database doesn't match the expected subscription object for: %v", test)
			}
		}
	}
}

func TestAddSubscription_InvalidInput(t *testing.T) {
	var tests = []struct {
		name              string
		subscriptionParam models.SubscriptionParams
	}{
		{"Missing destination",
			models.SubscriptionParams{
				SqsQueue: "",
				EventKey: "test-eventkey",
				Context:  "test-context",
			}},
		{"Missing eventKey",
			models.SubscriptionParams{
				SqsQueue: "test-queue",
				EventKey: "",
				Context:  "test-context",
			}},
	}

	for _, test := range tests {
		_, err := AddSubscription(test.subscriptionParam)

		if err == nil {
			t.Errorf("AddSubscription call succeeded but should have failed for: %v", test)
		}
	}
}

func TestAddEvent_ValidInput(t *testing.T) {
	// Override SQS calls
	orginalAlertSubscribersToEvent := AlertSubscribersToEvent
	defer func() { AlertSubscribersToEvent = orginalAlertSubscribersToEvent }()
	AlertSubscribersToEvent = func(event models.Event) error { return nil }

	var tests = []struct {
		name       string
		eventParam models.EventParams
	}{
		{"Valid Event",
			models.EventParams{
				EventKey: "test-eventkey",
				Context:  "test-context",
			}},
	}

	for _, test := range tests {
		id, err := AddEvent(test.eventParam)

		if err != nil {
			t.Errorf("AddEvent call failed for: %v", test)
		}

		if !IsValidUUID(id) {
			t.Errorf("The UUID from AddEvent is invalid for: %v", test)
		}

		eventParamExpected := test.eventParam

		// Check the value stored in the database
		if val, ok := globalVars.EventsMap[id]; ok {
			if val.Id != id || val.EventKey != eventParamExpected.EventKey ||
				val.Context != eventParamExpected.Context {
				t.Errorf("The subscription object in the database doesn't match the expected subscription object for: %v", test)
			}
		}
	}
}

func TestAddEvent_InvalidInput(t *testing.T) {
	// Override SQS calls
	orginalAlertSubscribersToEvent := AlertSubscribersToEvent
	defer func() { AlertSubscribersToEvent = orginalAlertSubscribersToEvent }()
	AlertSubscribersToEvent = func(event models.Event) error { return nil }

	var tests = []struct {
		name       string
		eventParam models.EventParams
	}{
		{"Missing EventKey",
			models.EventParams{
				EventKey: "",
				Context:  "test-context",
			}},
	}

	for _, test := range tests {
		_, err := AddEvent(test.eventParam)

		if err == nil {
			t.Errorf("AddSubscription call succeeded but should have failed for: %v", test)
		}
	}
}
