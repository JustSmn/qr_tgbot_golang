FROM golang

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o qr-bot .

CMD ["./qr-bot"]