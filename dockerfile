FROM golang:1.22.3 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

COPY --from=build /app/main /main

ENTRYPOINT ["/main"]
