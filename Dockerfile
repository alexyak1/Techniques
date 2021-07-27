## Specify the base image for go app
FROM golang:1.12.0-alpine3.9

## Create /app directory with image that will hold app soursce files
RUN mkdir /app

## Copy everything in the root dir
ADD . /app

## Specify that now execute any commands inside /app
WORKDIR /app
RUN apk add --no-cache git
## go mod command for pull in any dependencies
COPY go.mod ./
COPY go.sum ./
RUN go get -u github.com/gorilla/mux

RUN go mod download


## Run go build to compile the binary
RUN go build  -o main .

## newly created binary which executable
CMD ["/app/main"]
