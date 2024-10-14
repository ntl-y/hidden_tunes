FROM golang:1.23-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o hidden-tunes ./cmd/main.go

RUN apk add --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate.linux-amd64 /usr/local/bin/migrate

CMD ["sh", "-c", "until pg_isready -h db -p 5432; do sleep 1; done && \
    if ! psql postgres://postgres:govnonasrali@db:5432/postgres -c 'SELECT 1 FROM schema_migrations LIMIT 1;' > /dev/null 2>&1; then \
    migrate -path ./migrations -database \"postgres://postgres:govnonasrali@db:5432/postgres?sslmode=disable\" up; \
    fi && ./hidden-tunes"]
