# pgx-prometheus example

This example shows how to use pgx-prometheus to instrument a PostgreSQL server using `pgxpool`.

## Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)

## Running the example

1. Start a PostgreSQL server using Docker:

```bash
docker run --rm --name pgx-prometheus-example \
  -p 5432:5432 \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=password \
  postgres:15-alpine
```

2. Run the example:

```bash
go run example.go
```

3. Open http://localhost:8080/metrics in your browser to see the metrics.
