FROM golang:1.23 as builder

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ../ .

RUN go mod download && go mod verify
RUN env GOARCH=arm64 go build -v -o app ./cmd/app

FROM debian:bookworm-slim as deploy

RUN apt-get update 

COPY --from=builder /app/app .

RUN chmod +x app

CMD ["./app"]
