FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0
RUN go build -o main .

EXPOSE 8080/tcp

CMD ["./main"]