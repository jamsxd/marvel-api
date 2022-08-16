FROM golang:alpine as build
RUN apk add --no-cache git
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o app cmd/main.go

FROM alpine:latest
COPY --from=build /app .

CMD ["/app"]