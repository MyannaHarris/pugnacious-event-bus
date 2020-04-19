package main

import (
	//"fmt"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"pugnacious-event-bus/globalVars"
)

func performRequest(r http.Handler, method, path string, data string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(data))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestBasicGet(t *testing.T) {
	// Grab our router
	router := SetupRouter()
	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/", "")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
}

// SUbscriptions API
func TestPostSubscriptions_ValidInput(t *testing.T) {
	// Grab our router
	router := SetupRouter()

	var tests = []struct {
		name string
		data string
	}{
		{"Valid sqsqueue and eventkey",
			"{\"sqsqueue\":\"test-queue\",\"eventkey\":\"test-event\"}"},
		{"Valid sqsqueue and eventkey",
			"{\"apiurl\":\"test-apiurl\",\"eventkey\":\"test-event\"}"},
	}

	for _, test := range tests {
		// Perform a POST request with that handler.
		w := performRequest(router, "POST", "/subscriptions", test.data)
		// Assert we encoded correctly,
		// the request gives a 201
		assert.Equal(t, http.StatusCreated, w.Code)
		// Convert the JSON response to a map
		var response map[string]string
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		// Grab the value & whether or not it exists
		value, exists := response["id"]
		// Make some assertions on the correctness of the response.
		assert.Nil(t, err)
		assert.True(t, exists)
		assert.NotNil(t, value)
	}
}

func TestPostSubscriptions_InvalidInput(t *testing.T) {
	// Grab our router
	router := SetupRouter()

	var tests = []struct {
		name        string
		data        string
		expectedErr error
	}{
		{"Missing both sqsqueue and apiurl keys",
			"{\"eventkey\":\"test-event\"}",
			globalVars.MissingSubscriptionParamsErr},
		{"Missing both sqsqueue and apiurl values",
			"{\"sqsqueue\":\"\",\"apiurl\":\"\",\"eventkey\":\"test-event\"}",
			globalVars.MissingSubscriptionParamsErr},
		{"Missing eventkey key",
			"{\"sqsqueue\":\"test-queue\"}",
			globalVars.MissingSubscriptionParamsErr},
		{"Missing eventkey value",
			"{\"sqsqueue\":\"test-queue\",\"eventkey\":\"\"}",
			globalVars.MissingSubscriptionParamsErr},
		{"Missing all keys",
			"{}",
			globalVars.MissingSubscriptionParamsErr},
		{"Missing all values",
			"{\"sqsqueue\":\"\",\"apiurl\":\"\",\"eventkey\":\"\"}",
			globalVars.MissingSubscriptionParamsErr},
		{"Has both sqsqueue and apiurl values",
			"{\"sqsqueue\":\"test-queue\",\"apiurl\":\"test-apiurl\",\"eventkey\":\"test-event\"}",
			globalVars.TooManySubscriptionParamsErr},
	}

	for _, test := range tests {
		// Perform a POST request with that handler.
		w := performRequest(router, "POST", "/subscriptions", test.data)
		// Assert we encoded correctly,
		// the request gives a 500
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		// Convert the JSON response to a map
		var response map[string]string
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		// Grab the value & whether or not it exists
		value, exists := response["error"]
		// Make some assertions on the correctness of the response.
		assert.Nil(t, err)
		assert.True(t, exists)
		assert.NotNil(t, value)
		assert.Equal(t, response["error"], test.expectedErr.Error())
	}
}

// Events API
func TestPostEvents_ValidInput(t *testing.T) {
	// Grab our router
	router := SetupRouter()

	var tests = []struct {
		name string
		data string
	}{
		{"Valid eventkey",
			"{\"eventkey\":\"test-event\"}"},
	}

	for _, test := range tests {
		// Perform a POST request with that handler.
		w := performRequest(router, "POST", "/events", test.data)
		// Assert we encoded correctly,
		// the request gives a 201
		assert.Equal(t, http.StatusCreated, w.Code)
		// Convert the JSON response to a map
		var response map[string]string
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		// Grab the value & whether or not it exists
		value, exists := response["id"]
		// Make some assertions on the correctness of the response.
		assert.Nil(t, err)
		assert.True(t, exists)
		assert.NotNil(t, value)
	}
}

func TestPostEvents_InvalidInput(t *testing.T) {
	// Grab our router
	router := SetupRouter()

	var tests = []struct {
		name        string
		data        string
		expectedErr error
	}{
		{"Missing eventkey key",
			"{}",
			globalVars.MissingEventParamsErr},
		{"Missing eventkey value",
			"{\"eventkey\":\"\"}",
			globalVars.MissingEventParamsErr},
	}

	for _, test := range tests {
		// Perform a POST request with that handler.
		w := performRequest(router, "POST", "/events", test.data)
		// Assert we encoded correctly,
		// the request gives a 500
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		// Convert the JSON response to a map
		var response map[string]string
		err := json.Unmarshal([]byte(w.Body.String()), &response)
		// Grab the value & whether or not it exists
		value, exists := response["error"]
		// Make some assertions on the correctness of the response.
		assert.Nil(t, err)
		assert.True(t, exists)
		assert.NotNil(t, value)
		assert.Equal(t, response["error"], test.expectedErr.Error())
	}
}
