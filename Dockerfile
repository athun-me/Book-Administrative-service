# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the dependencies
RUN go mod download

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Set the working directory to the location of the main.go file
WORKDIR /go/src/app/cmd/api

# Copy the current directory contents into the container at /go/src/app/cmd/api
COPY . .

# Compile the application
RUN go build -o app

# Expose the port on which the application will run
EXPOSE 7777

# Set the default command to run when starting the container
CMD ["./app"]
