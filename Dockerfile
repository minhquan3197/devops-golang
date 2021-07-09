# create image from the official Go image
FROM golang:1.14

# golang specific variables
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Current working directory is /build in the container
RUN mkdir -p /projects

# Set the Current Working Directory inside the container
WORKDIR /projects

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port to run the application on
EXPOSE 8200

# Command to run the executable
CMD ["./main"]