FROM golang:1.20 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

# COPY *.go ./
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /grpc-api

EXPOSE 9002
CMD ["/grpc-api"]
