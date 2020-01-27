# pugnacious-event-bus
The Event Bus for the infamous Pugnacious deployment engine.

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
$ go fmt
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
