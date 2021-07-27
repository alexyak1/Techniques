FROM golang:1.16-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go get -u github.com/gin-gonic/gin
RUN go mod vendor
RUN go build -o /techniques

# EXPOSE 8787

CMD [ "/techniques" ]