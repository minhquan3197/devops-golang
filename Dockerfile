# create image from the official Go image
FROM golang:alpine AS builder

# golang specific variables
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# current working directory is /build in the container
WORKDIR /build

# Create binary directory, install glide and fresh
RUN go get github.com/pilu/fresh

# copy over go.mod and go.sum (module dependencies and checksum)
# over to working directory
COPY . .
# download the dependencies
RUN go mod download

# environment variables for the application
ENV GOLANG_ENV=develop

# copy our application code into the container
COPY . .

# expose the port to run the application on
EXPOSE 8200

# serve the app
ENTRYPOINT ["fresh"]