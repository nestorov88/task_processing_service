FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o ./task_proccesing_service ./cmd/app/main.go

ENTRYPOINT ["./task_proccesing_service"]