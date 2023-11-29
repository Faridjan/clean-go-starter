package pkg

import (
	"context"
	"fmt"
	"net/http"

	kitendpoint "github.com/go-kit/kit/endpoint"
)

const (
	GET_METHOD     = "GET"
	POST_METHOD    = "POST"
	PUT_METHOD     = "PUT"
	DELETE_METHOD  = "DELETE"
	OPTIONS_METHOD = "OPTIONS"
)

type (
	Middleware  func(next http.Handler) http.Handler
	DecoderType func(ctx context.Context, r *http.Request) (interface{}, error)
)

func HandlerJSON(endpoint kitendpoint.Endpoint, midlewares []Middleware, additionalDecoders ...DecoderType) http.HandlerFunc {
	// 1 TODO сделать автоматическое заполнение Request структуры эндпоинта парсингом http.Request-a (склеив с additionalDecoders)
	// 2 TODO смержить все мидлвари и запустить перед ендпоинтом
	// 3 TODO сделать автоматическое заполнение Response структуры эндпоинта для JSON

	return nil
}

func HandlerXML(endpoint kitendpoint.Endpoint, midlewares []Middleware, additionalDecoders ...DecoderType) http.HandlerFunc {
	// 1 TODO сделать автоматическое заполнение Request структуры эндпоинта парсингом http.Request-a  (склеив с additionalDecoders)
	// 2 TODO смержить все мидлвари и запустить перед ендпоинтом
	// 3 TODO сделать автоматическое заполнение Response структуры эндпоинта для XML

	return nil
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == OPTIONS_METHOD {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("[]"))
			if err != nil {
				fmt.Print("getDefaultOptionsMiddleware err:", err)
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}
