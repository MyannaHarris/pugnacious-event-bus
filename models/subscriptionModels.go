package models

// Structs used by the API
type SubscriptionParams struct {
	SqsQueue string `json:"sqsqueue"`
	ApiUrl   string `json:"apiurl"`
	EventKey string `json:"eventkey"`
	Context  string `json:"context"`
}

type EventParams struct {
	EventKey string `json:"eventkey"`
	Context  string `json:"context"`
}

// Structs used by the DDB model
type Subscription struct {
	Id       string
	SqsQueue string
	ApiUrl   string
	EventKey string
	Context  string
}

type Event struct {
	Id       string
	EventKey string
	Context  string
}
