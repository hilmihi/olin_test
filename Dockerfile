# STAGE 1
FROM golang:alpine AS builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go clean --modcache
RUN apk add --no-cache make
COPY . .
RUN go build -o main ./app/main.go

# MIGRATE
# RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# RUN make migrate_psql

# STAGE 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE 8912
CMD ["nohup", "./main"]
