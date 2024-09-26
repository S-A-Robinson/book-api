FROM golang:1.23-bullseye AS builder

# Set the current working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=1 GOOS=linux go build -o /books-api

EXPOSE 8080

CMD ["/books-api"]

