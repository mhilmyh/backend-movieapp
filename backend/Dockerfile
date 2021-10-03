FROM golang:1.17-alpine

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /backend-movieapp

CMD [ "/backend-movieapp" ]