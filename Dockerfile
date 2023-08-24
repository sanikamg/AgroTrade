FROM golang:1.20.5-alpine3.18 AS build-stage

# Maintainer info
LABEL maintainer="Sanika M G <sanikamg09@gmail.com>"

WORKDIR /home/app

# Copy go.mod and go.sum files
COPY go.mod go.sum

# Copy the entire project
COPY . .

# Build the application
RUN go build -o /home/build/api ./cmd/api

# Final stage
FROM alpine:3.18

# Maintainer info
LABEL maintainer="Sanika M G <sanikamg9@gmail.com>"

WORKDIR /home/app

# Copy the compiled binary from the build stage
COPY --from=build-stage /home/build/api ./

# Create a config folder and copy the config.json file

COPY /views ./views

COPY .env .

CMD ["./api"]