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
# Copy everything into the Docker container.
COPY . .

# Compile migrate binary used for running database migrations.
RUN go build -o /app/bin/migrate ./cmd/migrate/main.go
# Compile main binary used for starting the server.
RUN go build -o /app/bin/server ./cmd/server/main.go

# Run server.
CMD ["/app/bin/server"]
