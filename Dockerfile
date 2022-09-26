FROM golang:alpine as builder

LABEL maintainer="Daniel Villanueva <villanueva.danielx@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

COPY .env.example .env

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

EXPOSE 8000

CMD ["./main"]