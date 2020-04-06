FROM golang:latest

# Create an /app directory within our image that will hold our application
# source files.
RUN mkdir /app
# Copy everything in the root directory into our /app directory.
ADD . /app
# Cpecify that we now wish to execute any further commands inside our /app directory.
WORKDIR /app
# Copy dependency manifests to container.
COPY go.mod .
COPY go.sum .
# Get dependencies.
RUN go mod download
# Get Goose binary so that the DB migrate fns can be run in the container.
RUN go get -u github.com/pressly/goose/cmd/goose
# Copy everything into the Docker container.
COPY . .

# Compile main binary used for starting the server.
RUN go build -o bin/server ./cmd/server/main.go
# Run server.
CMD ["/app/server"]
