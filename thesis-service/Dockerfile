FROM golang:alpine as builder

ENV GO111MODULE=on
WORKDIR /home/app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build ./cmd/main.go

######## Start a new stage from scratch #######
FROM alpine:latest as app

RUN apk --no-cache add ca-certificates

RUN adduser -D myuser
USER myuser

WORKDIR /home/root

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /home/app/main .

# Command to run the executable
CMD ["./main"]

