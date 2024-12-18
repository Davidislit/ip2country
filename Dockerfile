FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ip2country

EXPOSE 8080

CMD ["./ip2country"]
