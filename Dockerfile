# create image from the official Go image
FROM golang:1.14

# golang specific variables
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# current working directory is /build in the container
RUN mkdir -p /projects

WORKDIR /projects

RUN ls -la /projects

COPY . /projects

RUN ls -la /projects

RUN go get github.com/pilu/fresh

# download the dependencies
RUN go mod download

VOLUME /projects

WORKDIR /projects/cmd/app

RUN ls -la

# expose the port to run the application on
EXPOSE 8200

# serve the app
ENTRYPOINT [ "fresh" ]