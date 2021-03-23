# Import go image
FROM golang:1.13-alpine
# Label for maintainer
LABEL maintainer="Jack Maarek"
# Set the working directory inside the container
WORKDIR /go/src
# Copy the full project to currennt directory
COPY . .
COPY .docker/consumer/entrypoint.sh ./consumer

# Run command to nstall the dependencies
RUN go mod download

RUN go get -u github.com/cosmtrek/air

WORKDIR consumer

ENTRYPOINT ["./entrypoint.sh"]