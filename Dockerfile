# build stage, discard most of this later
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/api

# copy binary from building stage
FROM gcr.io/distroless/static-debian12
WORKDIR /
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
