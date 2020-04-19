package models

// Structs used by the API
type SubscriptionParams struct {
	SqsQueue string `json:"sqsqueue"`
	EventKey string `json:"eventkey"`
	Context  string `json:"context"`
}

type EventParams struct {
	EventKey string `json:"eventkey"`
	Context  string `json:"context"`
}

// Structs used by the database model
type Subscription struct {
	Id       string
	SqsQueue string
	EventKey string
	Context  string
}

type Event struct {
	Id       string
	EventKey string
	Context  string
}

type EventMessage struct {
	Id       string
	EventKey string
	Context  string
	Source   string
}
