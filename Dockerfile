## Specify the base image for go app
FROM golang:1.12.0-alpine3.9

ENV GO111MODULE=on
ENV PORT=8787


## Specify that now execute any commands inside /app
WORKDIR /app/server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build
CMD ["./server"]
