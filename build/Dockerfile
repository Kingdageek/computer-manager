# Build stage
FROM golang:1.23-alpine AS builder

ARG AIR_MODULE=github.com/cosmtrek/air
ARG AIR_VERSION=v1.41.0

RUN apk add --no-cache git \
 && go install ${AIR_MODULE}@${AIR_VERSION}

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin ./main.go

# Run stage for cleaner, smaller image with just compiled binary and minimal dependencies
FROM golang:1.23-alpine AS dev

WORKDIR /app

COPY --from=builder /app/bin/ /app/bin
COPY --from=builder /go/bin/air /usr/local/bin/air
COPY . .

CMD ["air", "-c", "air.toml"]