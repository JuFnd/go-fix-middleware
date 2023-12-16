package middleware

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Vkladyshi/pkg/requests"
)


type ResponseMiddleware struct {
	Response *requests.Response
}

func (mw *ResponseMiddleware) GetResponse(next http.Handler, lg *slog.Logger) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(w, r)
        // Your logic to further modify the response if needed
        requests.SendResponse(w, *mw.Response, lg)
    })
}


