FROM golang:1.23-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o hidden-tunes ./cmd/main.go

CMD ["./hidden-tunes"]