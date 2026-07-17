FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o hello-app main.go

CMD ["./hello-app"]
