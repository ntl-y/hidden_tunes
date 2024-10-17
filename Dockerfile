FROM golang:1.23-alpine

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o hidden-tunes cmd/main.go

CMD ["./hidden-tunes"]
