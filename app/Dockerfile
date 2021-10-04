FROM golang:1.17-alpine

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download
RUN go build -o /app main.go

CMD [ "/app" ]