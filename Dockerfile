FROM golang:latest

WORKDIR /pugnacious

COPY go.mod go.sum ./

ENV GOPROXY direct

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
