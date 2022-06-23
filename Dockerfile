FROM golang:1.18-alpine as build

WORKDIR /app

COPY go.mod ./go.mod

COPY go.sum ./go.sum

RUN go mod download

RUN go mod tidy

COPY . .

RUN go build /cmd/main -o main


FROM alpine:3.16.0

WORKDIR /app

EXPOSE 8080

COPY --from=build /app/main /app/main

CMD ["./main"]
