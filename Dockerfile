FROM golang:latest

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /go/src/app/cmd/api

COPY . .

RUN go build -o app

EXPOSE 7777

CMD ["./app"]
