package actions

import (
	"regexp"
	"testing"

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
				ApiUrl:   "",
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

		// Check the value stored in the database
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
				ApiUrl:   "",
				EventKey: "test-eventkey",
				Context:  "test-context",
			}},
		{"Missing eventKey",
			models.SubscriptionParams{
				SqsQueue: "test-queue",
				ApiUrl:   "",
				EventKey: "",
				Context:  "test-context",
			}},
		{"Has both SqsQueue and ApiUrl",
			models.SubscriptionParams{
				SqsQueue: "test-queue",
				ApiUrl:   "test-apiurl",
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
	var tests = []struct {
		name              string
		subscriptionParam models.EventParams
	}{
		{"Valid Event",
			models.EventParams{
				EventKey: "test-eventkey",
				Context:  "test-context",
			}},
	}

	for _, test := range tests {
		id, err := AddEvent(test.subscriptionParam)

		if err != nil {
			t.Errorf("AddEvent call failed for: %v", test)
		}

		if !IsValidUUID(id) {
			t.Errorf("The UUID from AddEvent is invalid for: %v", test)
		}

		// Check the value stored in the database?

		// Check the subscribed parties are alerted?
	}
}

func TestAddEvent_InvalidInput(t *testing.T) {
	var tests = []struct {
		name              string
		subscriptionParam models.EventParams
	}{
		{"Missing EventKey",
			models.EventParams{
				EventKey: "",
				Context:  "test-context",
			}},
	}

	for _, test := range tests {
		_, err := AddEvent(test.subscriptionParam)

		if err == nil {
			t.Errorf("AddSubscription call succeeded but should have failed for: %v", test)
		}
	}
}
