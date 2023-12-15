package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/metrics"
)

type Middleware struct {
	metrics *metrics.Metrics
}

func GetMiddleware() *Middleware {
	return &Middleware{
		metrics: metrics.GetMetrics(),
	}
}

func (m *Middleware) GetMetrics(next http.Handler, responseStatus int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record the start time of the request
		startTime := time.Now()

		// Serve the HTTP request by calling the next handler
		next.ServeHTTP(w, r)

		// Calculate the request processing time
		duration := time.Since(startTime).Seconds()

		// Extract relevant information from the request, such as status and path
		status := responseStatus
		path := r.URL.Path

		// Record the request processing time in the histogram
		m.metrics.Time.WithLabelValues(strconv.Itoa(status), path).Observe(duration)

		// Increment the hits counter for the corresponding status and path
		m.metrics.Hits.WithLabelValues(strconv.Itoa(status), path).Inc()
	})
}
