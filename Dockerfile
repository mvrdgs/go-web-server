FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-web-server cmd/server/main.go

EXPOSE 8080

CMD ["/app/go-web-server"]

# docker build --tag web-server .