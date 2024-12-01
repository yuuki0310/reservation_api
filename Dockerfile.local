FROM golang:1.23-alpine

WORKDIR /go/src/github.com/yuuki0310/reservation_api


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@v1.27.3

CMD ["air"]
