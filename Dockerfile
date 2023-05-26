FROM golang:1.20.4-alpine3.18

# Creating an api directory
RUN mkdir /api
WORKDIR /api

# Copying the manifests
COPY go.mod go.sum ./

# Downloading dependencies
RUN go mod download

# Copy the files
COPY . .

# Building the application
RUN go build -o main .

# Exposing the port used by Gin
EXPOSE 8081

# Setting entry point
CMD ["./main"]