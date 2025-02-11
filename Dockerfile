FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o echo-app

EXPOSE 8080

CMD ["./echo-app"]