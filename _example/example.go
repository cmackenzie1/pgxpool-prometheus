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
	defer pool.Close()

	prometheus.MustRegister(pgxpool_prometheus.NewPgxPoolStatsCollector(pool, "database"))

	if err := pool.QueryRow(context.Background(), "SELECT 1").Scan(nil); err != nil {
		panic(err)
	}

	log.Fatalf("Error: %v", http.ListenAndServe(":8080", promhttp.Handler()))
}
