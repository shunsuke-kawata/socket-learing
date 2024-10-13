FROM golang:1.23.2-alpine3.20

RUN apk update
WORKDIR /app
COPY ./be/go.mod /app/go.mod
COPY ./be/go.sum /app/go.sum

RUN go mod tidy

ENTRYPOINT ["go", "run","main.go"]
