package service

import (
	"tiny-template/src/api/page"
	"tiny-template/src/api/page/middleware"
	"tiny-template/src/repository"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type service struct {
	mainStore   repository.MainPageRepositoryInterface
	searchStore repository.SearchPageRepositoryInterface
}

func NewService(
	mainStore repository.MainPageRepositoryInterface,
	searchStore repository.SearchPageRepositoryInterface,
) page.ServiceInterface {
	serviceStruct := &service{
		mainStore:   mainStore,
		searchStore: searchStore,
	}

	svc := middleware.NewValidationMiddleware()(serviceStruct)
	svc = middleware.NewLoggingMiddleware()(svc)
	svc = middleware.NewInstrumentingMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "page_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "page_service",
			Name:      "error_count",
			Help:      "Number of error requests received.",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "page_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}),
	)(svc)

	return svc
}
