# pugnacious-event-bus
The Event Bus for the infamous Pugnacious deployment engine.

## PreReqs

AWS Creds
```
# Use your AWS creds by running something like
$ aws configure
```

AWS SQS Queue
```
# Create a Queue with open permissions 
# or permissions at least for the account you are using the creds from above

# Anyone can send messages
"Statement": [
	{
		"Sid": "<Sid>",
		"Effect": "Allow",
		"Principal": "*",
		"Action": "SQS:SendMessage",
		"Resource": "arn:aws:sqs:us-east-2:<account-that-queue-is-in>:<queue-name>"
	}
]

# Only specific accounts can send messages
"Statement": [
	{
		"Sid": "<Sid>",
		"Effect": "Allow",
		"Principal": "<account-ID-for-above-aws-creds>",
		"Action": "SQS:SendMessage",
		"Resource": "arn:aws:sqs:us-east-2:<account-that-queue-is-in>:<queue-name>"
	}
]
```

Example reulting queue URL:
```
https://sqs.us-east-2.amazonaws.com/<account-ID>/<queue-name> 	
```

## Getting started

You must have Go `1.13` or later installed.

Build the project:

```
$ go build
```

Build and run in one shot:

```
$ go run main.go
```

Once you run the app, you can check it out at `localhost:8000`.

Format your code:

```
$ ./gofmt
```

Run unit tests on server:

```
$ go run main.go
$ go test
$ go test -v ./actions/*
```

Create executable and run it:

```
$ go build -o app
$ ./app
```

## Running in Docker

Build the Docker image:

```
docker build -t pugnacious-event-bus .
```

Run the Docker image:

```
docker run -d -p 8000:8000 pugnacious-event-bus
```

## Call the available APIs

#### GET

* This API is currently just to ping

```
$ go run main.go
```

```
Use Postmate to hit http://localhost:8000

or

curl http://localhost:8000
```

#### POST /subscriptions

* This API is to add a subscription in the event bus

```
$ go run main.go
```

```
Use Postmate to hit http://localhost:8000/subscriptions
Add JSON body with values like "context"

or

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"sqsqueue":"test-queue","eventkey":"test-event"}' \
  http://localhost:8000/subscriptions
```

#### POST /events

* This API is to add an event to the event bus

```
$ go run main.go
```

```
Use Postmate to hit http://localhost:8000/events
Add JSON body with values like "context"

or

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"eventkey":"test-event"}' \
  http://localhost:8000/events
```
