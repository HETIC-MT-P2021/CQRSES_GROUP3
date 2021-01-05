# Start from golang base image
FROM golang:1.15-alpine

# ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /go/src/github.com/HETIC-MT-P2021/CQRSES_GROUP3

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download -x
RUN go get -u github.com/cosmtrek/air

#Command to build the Go app with hot reload
ENTRYPOINT ["/go/bin/air"]