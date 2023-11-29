package elasticpkg

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"

	"git.centerhome.kz/bcc/backend/toolchain/common-libs/errors"
)

func (p *ProviderElastic) Close(_ context.Context) error {
	p.client.Stop()
	p.ctx.Done()

	return nil
}

// GetConn
// get elastic client
func (p *ProviderElastic) GetConn() *elastic.Client {
	return p.client
}

// GetByID
// get from index by ID
func (p *ProviderElastic) GetByID(index, id string) (*elastic.GetResult, error) {
	result, err := p.client.Get().Index(index).Id(id).Do(p.ctx)
	if err != nil {
		return nil, errors.InternalServerError(p.ctx).
			SetDevMessage("GetByID.p.client.Get: " + err.Error())
	}

	return result, nil
}

// GetByQuery
// get by query
func (p *ProviderElastic) GetByQuery(index string, sortFields map[string]bool, query elastic.Query, from, size int) (*elastic.SearchResult, error) {
	var fieldsSort []elastic.Sorter

	search := p.client.Search().
		Index(index).
		Query(query).
		From(from).
		Size(size)

	for sortField, sortASC := range sortFields {
		var fieldSort *elastic.FieldSort

		if sortASC {
			fieldSort = elastic.NewFieldSort(sortField).Asc()
		} else {
			fieldSort = elastic.NewFieldSort(sortField).Desc()
		}

		fieldsSort = append(fieldsSort, fieldSort)
	}

	if len(fieldsSort) > 0 {
		search.SortBy(fieldsSort...)
	}

	result, err := search.Do(p.ctx)
	if err != nil {
		return nil, errors.InternalServerError(p.ctx).
			SetDevMessage("GetByQuery.search.Do: " + err.Error())
	}

	return result, nil
}

// GetAllByQuery
// get all by query
func (p *ProviderElastic) GetAllByQuery(index, sortField string, sortASC bool, query elastic.Query) (*elastic.SearchResult, error) {
	search := p.client.Search().
		Index(index).
		Query(query).
		Size(elasticLimit)

	if sortField != "" {
		search.Sort(sortField, sortASC)
	}

	result, err := search.Do(p.ctx)
	if err != nil {
		return nil, errors.InternalServerError(p.ctx).
			SetDevMessage("GetAllByQuery.search.Do: " + err.Error())
	}

	return result, nil
}

// GetByQueryFields
// get only fields by query
func (p *ProviderElastic) GetByQueryFields(index string, sortFields map[string]bool, query elastic.Query, from, size int, returnFields []string) (*elastic.SearchResult, error) {
	var fieldsSort []elastic.Sorter

	fsc := elastic.NewFetchSourceContext(true).Include(returnFields...)
	search := p.client.Search().
		Index(index).
		Query(query).
		FetchSourceContext(fsc).
		From(from).
		Size(size)

	for sortField, sortASC := range sortFields {
		var fieldSort *elastic.FieldSort

		if sortASC {
			fieldSort = elastic.NewFieldSort(sortField).Asc()
		} else {
			fieldSort = elastic.NewFieldSort(sortField).Desc()
		}

		fieldsSort = append(fieldsSort, fieldSort)
	}

	if len(fieldsSort) > 0 {
		search.SortBy(fieldsSort...)
	}

	result, err := search.Do(p.ctx)
	if err != nil {
		return nil, errors.InternalServerError(p.ctx).
			SetDevMessage("GetByQueryFields.search.Do: " + err.Error())
	}

	return result, nil
}

// GetByQueryGeoSorted
// get by query
func (p *ProviderElastic) GetByQueryGeoSorted(index, sortField string, sortPoint elastic.GeoPoint, query elastic.Query, from, size int) (*elastic.SearchResult, error) {
	sort := elastic.NewGeoDistanceSort(sortField).
		Point(sortPoint.Lat, sortPoint.Lon).
		Order(true).
		Unit("m").
		GeoDistance("plane")

	search := p.client.Search().
		Index(index).
		Query(query).
		SortBy(sort).
		From(from).
		Size(size)

	result, err := search.Do(p.ctx)
	if err != nil {
		return nil, errors.InternalServerError(p.ctx).
			SetDevMessage("GetByQueryGeoSorted.search.Do: " + err.Error())
	}

	return result, nil
}

// GetCount
// get count by query
func (p *ProviderElastic) GetCount(index string, query elastic.Query) (int64, error) {
	result, err := p.client.Count(index).Query(query).Do(p.ctx)
	if err != nil {
		return 0, errors.InternalServerError(p.ctx).
			SetDevMessage("GetCount.count.Do: " + err.Error())
	}

	return result, nil
}

// GetAggregatedCount
// get count by query aggregated by field
func (p *ProviderElastic) GetAggregatedCount(index string, query elastic.Query, aggrField string) (map[interface{}]int64, error) {
	aggregationResult := make(map[interface{}]int64)

	aggr := elastic.NewTermsAggregation().Field(aggrField)
	search := p.client.Search().
		Index(index).
		Query(query).
		Size(elasticLimit).
		Aggregation(aggrField, aggr)

	result, err := search.Do(p.ctx)
	if err != nil {
		return nil, errors.InternalServerError(p.ctx).
			SetDevMessage("GetAggregatedCount.search.Do: " + err.Error())
	}

	if aggregation, ok := result.Aggregations[aggrField]; ok {
		var aggregationItem struct {
			Buckets []struct {
				Key      interface{} `json:"key"`
				DocCount int64       `json:"doc_count"`
			} `json:"buckets"`
		}

		err = json.Unmarshal(aggregation, &aggregationItem)
		if err != nil {
			return nil, errors.InternalServerError(p.ctx).
				SetDevMessage("GetAggregatedCount.json.Unmarshal: " + err.Error())
		}

		for _, bucket := range aggregationItem.Buckets {
			aggregationResult[bucket.Key] = bucket.DocCount
		}
	}

	return aggregationResult, nil
}
