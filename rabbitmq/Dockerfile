# Import go image
FROM golang:1.13-alpine
# Label for maintainer
LABEL maintainer="Jack Maarek"
# Set the working directory inside the container
WORKDIR /go/src
# Copy the full project to currennt directory
COPY . .

RUN ls -al
# Run command to nstall the dependencies
RUN go mod download
RUN go get -u github.com/cosmtrek/air

#Command to build the Go app with hot reload
ENTRYPOINT ["/go/bin/air"]