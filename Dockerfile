FROM golang:1.24-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . . 
RUN go build -o spy-cat-agency ./main.go 

FROM alpine:latest


COPY --from=builder /app/spy-cat-agency .
COPY --from=builder /app/migrations ./migrations


EXPOSE 8080

CMD ["./spy-cat-agency"]