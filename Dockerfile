# create image from the official Go image
FROM golang:alpine AS builder

# golang specific variables
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# current working directory is /build in the container
RUN mkdir -p /app
WORKDIR /app
COPY . /app

# Create binary directory, install glide and fresh
RUN go get github.com/pilu/fresh

# download the dependencies
RUN go mod download

RUN ls -la /app

# Start app
VOLUME /app

# environment variables for the application
ENV GOLANG_ENV=develop

# expose the port to run the application on
EXPOSE 8200

# serve the app
ENTRYPOINT ["fresh"]