package middleware

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/metrics"
	"github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/requests"
)


type ResponseMiddleware struct {
	Response *requests.Response
	Metrix   *metrics.Metrics
}

func (mw *ResponseMiddleware) GetResponse(next http.Handler, lg *slog.Logger) http.Handler {
	start := time.Now()
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(w, r)
        requests.SendResponse(w, *mw.Response, lg)
		end := time.Since(start)
		mw.Metrix.Time.WithLabelValues(strconv.Itoa(mw.Response.Status), r.URL.Path).Observe(end.Seconds())

        mw.Metrix.Hits.WithLabelValues(strconv.Itoa(mw.Response.Status), r.URL.Path).Inc()
    })
}