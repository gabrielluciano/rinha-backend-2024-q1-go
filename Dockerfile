FROM golang:1.22.0-bookworm as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /rinha ./cmd/rinha-backend-2024-q1-go 

FROM debian:bookworm

COPY --from=build /rinha /usr/local/bin/rinha

EXPOSE 8080 8081

CMD ["rinha"]
