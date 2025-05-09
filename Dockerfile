### Builder
FROM golang:1.24-alpine AS builder
WORKDIR /app

# For Cross-compile (Cloud Run x86)
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

# Dependancy 먼저 For Caching
COPY go.mod go.sum ./
RUN go mod download

# 소스 복사
COPY . .

# 빌드
RUN go build -o run ./cmd/api-server

### Runtime
FROM alpine:latest AS runtime
WORKDIR /app

COPY --from=builder /app/run .

COPY resources ./resources

# Secret
COPY .env ./
COPY ./firebase-application-credentials.json ./
COPY ./google-tts-application-credentials.json ./

EXPOSE 8090

ENTRYPOINT ["./run"]