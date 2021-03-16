# Dockerfile References: https://docs.docker.com/engine/reference/builder/

### STAGE 1: BUILD ###
FROM golang:1.16.0-alpine as builder
# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
# Create app directory
RUN mkdir /app
# Set the Current Working Directory inside the container
WORKDIR /app
ADD . /app
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Generate Swagger document
RUN go get github.com/swaggo/swag/cmd/swag && swag init -g ./cmd/api/main.go -o ./docs
# Generate dependencies by wire
RUN go get github.com/google/wire/cmd/wire && wire ./internal/wired/redis.go
# Build the Go api
RUN go build -o ./todoapi ./cmd/api

### STAGE 2: RUN ###
FROM golang:1.16.0-alpine
COPY --from=builder /app/todoapi /go/bin/todoapi
# Expose port 8080 to the outside world
EXPOSE 8080
# Run the executable
CMD ["todoapi"]