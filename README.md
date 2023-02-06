# pgxpool-prometheus

[![Go Reference](https://pkg.go.dev/badge/github.com/cmackenzie1/pgxpool-prometheus.svg)](https://pkg.go.dev/github.com/cmackenzie1/pgxpool-prometheus)

A [pgx](https://github.com/jackc/pgx) [Prometheus](https://prometheus.io/) metrics collector for Go applications
using [pgxpool](https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool).

## Installation

```bash
go get github.com/cmackenzie1/pgxpool-prometheus
```

## Usage

Please see the [example](./_example) directory for a complete example.

```go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/cmackenzie1/pgxpool-prometheus"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a pgxpool.Config
	config, err := pgxpool.ParseConfig("postgres://postgres:password@localhost:5432/?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Create a pgxpool.Pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	prometheus.MustRegister(pgxpool_prometheus.NewPgxPoolStatsCollector(pool, "database"))

	log.Fatalf("Error: %v", http.ListenAndServe(":8080", promhttp.Handler()))
}
```

## Exported Metrics

| Metric                                | Type    | Description                                                         |
|---------------------------------------|---------|---------------------------------------------------------------------|
| `pgx_pool_acquire_connections`        | Gauge   | Number of connections currently in the process of being acquired    |
| `pgx_pool_canceled_acquire_count`     | Counter | Number of times a connection acquire was canceled                   |
| `pgx_pool_constructing_connections`   | Gauge   | Number of connections currently in the process of being constructed |
| `pgx_pool_empty_acquire_count`        | Counter | Number of times a connection acquire was canceled                   |
| `pgx_pool_idle_connections`           | Gauge   | Number of idle connections in the pool                              |
| `pgx_pool_max_connections`            | Gauge   | Maximum number of connections allowed in the pool                   |
| `pgx_pool_max_idle_destroy_count`     | Counter | Number of connections destroyed due to MaxIdleTime                  |
| `pgx_pool_max_lifetime_destroy_count` | Counter | Number of connections destroyed due to MaxLifetime                  |
| `pgx_pool_new_connections_count`      | Counter | Number of new connections created                                   |
| `pgx_pool_total_connections`          | Gauge   | Total number of connections in the pool                             |

## License

pgx-prometheus is licensed under the MIT License. See [LICENSE](./LICENSE) for the full license text.