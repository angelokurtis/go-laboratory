package metrics

import (
	"database/sql"
	"net/http"

	"github.com/dlmiddlecote/sqlstats"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler http.Handler

func NewHandler(db *sql.DB) Handler {
	collector := sqlstats.NewStatsCollector("test", db)
	prometheus.MustRegister(collector)

	return promhttp.Handler()
}
