FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=build /app/main /app/main

WORKDIR /app

EXPOSE 8080

CMD ["/app/main"]
