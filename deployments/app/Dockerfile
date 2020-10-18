# create image from the official Go image
FROM golang:1.14

# golang specific variables
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# current working directory is /build in the container
RUN mkdir -p /server

WORKDIR /server

COPY . /server

RUN ls -la /server

RUN go get github.com/pilu/fresh

# download the dependencies
RUN go mod download

VOLUME /server

# expose the port to run the application on
EXPOSE 8200

# serve the app
ENTRYPOINT [ "fresh" ]