package elasticpkg

import (
	"context"

	"github.com/olivere/elastic/v7"
)

// ProviderElasticInterface
// main elastic provider interface
type ProviderElasticInterface interface {
	Migrate(index, mapping string) error
	GetConn() *elastic.Client
	Close(ctx context.Context) error
	GetByID(index, id string) (*elastic.GetResult, error)
	GetByQuery(index string, sortFields map[string]bool, query elastic.Query, from, size int) (*elastic.SearchResult, error)
	GetAllByQuery(index, sortField string, sortASC bool, query elastic.Query) (*elastic.SearchResult, error)
	GetByQueryFields(index string, sortFields map[string]bool, query elastic.Query, from, size int, returnFields []string) (*elastic.SearchResult, error)
	GetByQueryGeoSorted(index string, sortField string, sortPoint elastic.GeoPoint, query elastic.Query, from, size int) (*elastic.SearchResult, error)
	StoreByID(index, id string, payload interface{}) error
	DeleteByID(index, id string) error
	DeleteByQuery(index string, query elastic.Query) error

	GetCount(index string, query elastic.Query) (int64, error)
	GetAggregatedCount(index string, query elastic.Query, aggrField string) (map[interface{}]int64, error)
}
