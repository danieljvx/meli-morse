FROM golang:alpine as builder

LABEL maintainer="Daniel Villanueva <villanueva.danielx@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

COPY .env.example .env

RUN go mod download

COPY . .

RUN go get -u github.com/swaggo/swag/cmd/swag

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g main.go --output docs

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

EXPOSE 8000

CMD ["./main"]