# Build stage
FROM golang:1.26-bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Run stage
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /app/main .
CMD ["/app/main"]