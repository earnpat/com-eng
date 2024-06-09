FROM golang:1.20 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /grpc-api

EXPOSE 9001
CMD ["/grpc-api"]
