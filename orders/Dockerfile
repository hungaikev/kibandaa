## We specify the base image we need for our
## go application
FROM golang:alpine as builder

# Args
ARG APP_NAME
ARG VCS_REF

ENV APP_NAME ${APP_NAME}

# Add Maintainer Info
LABEL maintainer="Hungai Amuhinda <hungaikevin@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git curl


# Install golang migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate.linux-amd64 /usr/bin/migrate


# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.build=${VCS_REF}" -a -installsuffix cgo -o main ./cmd/orders/*.go


# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/main .
# Copy the migration folder from the previous stage.
COPY --from=builder /app/internal/storage/db/migrations ./internal/storage/db/migrations
# Copy migrate
COPY --from=builder /usr/bin/migrate /usr/bin/migrate
# Copy Wait for it
COPY --from=builder /app/wait-for .

# Expose port 9090 to the outside world
EXPOSE 9090

#Command to run the executable
CMD ["./main"]