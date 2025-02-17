FROM golang:alpine3.17 AS builder

WORKDIR /go/src/app
COPY  . . 
RUN go mod download

RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/receipt_processor ./main.go

FROM alpine:3.17
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=builder /go/src/app/bin /go/bin
EXPOSE 8080
CMD ["/go/bin/receipt_processor", "--port", "8080"]
