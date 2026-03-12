FROM golang:1.25

WORKDIR /app

COPY . .

RUN go mod init devops-security-platform || true
RUN go mod tidy

RUN go build -o auth-service

EXPOSE 9092

CMD ["./auth-service"]