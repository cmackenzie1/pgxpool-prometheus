package pgx_prometheus

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
)

// PgxPoolCollector is a Prometheus collector for pgx metrics.
// It implements the prometheus.Collector interface.
type PgxPoolCollector struct {
	db *pgxpool.Pool

	acquireConns            *prometheus.Desc
	canceledAcquireCount    *prometheus.Desc
	constructingConns       *prometheus.Desc
	emptyAcquireCount       *prometheus.Desc
	idleConns               *prometheus.Desc
	maxConns                *prometheus.Desc
	totalConns              *prometheus.Desc
	newConnsCount           *prometheus.Desc
	maxLifetimeDestroyCount *prometheus.Desc
	maxIdleDestroyCount     *prometheus.Desc
}

// NewPgxStatsCollector returns a new pgxCollector.
// The dbName parameter is used to set the "db" label on the metrics.
// The db parameter is the pgxpool.Pool to collect metrics from.
// The db parameter must not be nil.
// The dbName parameter must not be empty.
func NewPgxStatsCollector(db *pgxpool.Pool, dbName string) *PgxPoolCollector {
	fqName := func(name string) string {
		return prometheus.BuildFQName("pgx", "pool", name)
	}
	return &PgxPoolCollector{
		db: db,
		acquireConns: prometheus.NewDesc(
			fqName("acquire_connections"),
			"Number of connections currently in the process of being acquired",
			nil,
			prometheus.Labels{"db": dbName},
		),
		canceledAcquireCount: prometheus.NewDesc(
			fqName("canceled_acquire_count"),
			"Number of times a connection acquire was canceled",
			nil,
			prometheus.Labels{"db": dbName},
		),
		constructingConns: prometheus.NewDesc(
			fqName("constructing_connections"),
			"Number of connections currently in the process of being constructed",
			nil,
			prometheus.Labels{"db": dbName},
		),
		emptyAcquireCount: prometheus.NewDesc(
			fqName("empty_acquire_count"),
			"Number of times a connection acquire was canceled",
			nil,
			prometheus.Labels{"db": dbName},
		),
		idleConns: prometheus.NewDesc(
			fqName("idle_connections"),
			"Number of idle connections in the pool",
			nil,
			prometheus.Labels{"db": dbName},
		),
		maxConns: prometheus.NewDesc(
			fqName("max_connections"),
			"Maximum number of connections allowed in the pool",
			nil,
			prometheus.Labels{"db": dbName},
		),
		totalConns: prometheus.NewDesc(
			fqName("total_connections"),
			"Total number of connections in the pool",
			nil,
			prometheus.Labels{"db": dbName},
		),
		newConnsCount: prometheus.NewDesc(
			fqName("new_connections_count"),
			"Number of new connections created",
			nil,
			prometheus.Labels{"db": dbName},
		),
		maxLifetimeDestroyCount: prometheus.NewDesc(
			fqName("max_lifetime_destroy_count"),
			"Number of connections destroyed due to MaxLifetime",
			nil,
			prometheus.Labels{"db": dbName},
		),
		maxIdleDestroyCount: prometheus.NewDesc(
			fqName("max_idle_destroy_count"),
			"Number of connections destroyed due to MaxIdleTime",
			nil,
			prometheus.Labels{"db": dbName},
		),
	}
}

// Describe implements the prometheus.Collector interface.
func (p PgxPoolCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- p.acquireConns
	descs <- p.canceledAcquireCount
	descs <- p.constructingConns
	descs <- p.emptyAcquireCount
	descs <- p.idleConns
	descs <- p.maxConns
	descs <- p.totalConns
	descs <- p.newConnsCount
	descs <- p.maxLifetimeDestroyCount
	descs <- p.maxIdleDestroyCount
}

// Collect implements the prometheus.Collector interface.
func (p PgxPoolCollector) Collect(metrics chan<- prometheus.Metric) {
	stats := p.db.Stat()

	metrics <- prometheus.MustNewConstMetric(p.acquireConns, prometheus.GaugeValue, float64(stats.AcquiredConns()))
	metrics <- prometheus.MustNewConstMetric(p.canceledAcquireCount, prometheus.CounterValue, float64(stats.CanceledAcquireCount()))
	metrics <- prometheus.MustNewConstMetric(p.constructingConns, prometheus.GaugeValue, float64(stats.ConstructingConns()))
	metrics <- prometheus.MustNewConstMetric(p.emptyAcquireCount, prometheus.CounterValue, float64(stats.EmptyAcquireCount()))
	metrics <- prometheus.MustNewConstMetric(p.idleConns, prometheus.GaugeValue, float64(stats.IdleConns()))
	metrics <- prometheus.MustNewConstMetric(p.maxConns, prometheus.GaugeValue, float64(stats.MaxConns()))
	metrics <- prometheus.MustNewConstMetric(p.totalConns, prometheus.GaugeValue, float64(stats.TotalConns()))
	metrics <- prometheus.MustNewConstMetric(p.newConnsCount, prometheus.CounterValue, float64(stats.NewConnsCount()))
	metrics <- prometheus.MustNewConstMetric(p.maxLifetimeDestroyCount, prometheus.CounterValue, float64(stats.MaxLifetimeDestroyCount()))
	metrics <- prometheus.MustNewConstMetric(p.maxIdleDestroyCount, prometheus.CounterValue, float64(stats.MaxIdleDestroyCount()))
}
