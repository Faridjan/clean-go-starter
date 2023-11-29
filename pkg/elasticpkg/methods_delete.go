package elasticpkg

import (
	"github.com/olivere/elastic/v7"

	"git.centerhome.kz/bcc/backend/toolchain/common-libs/errors"
)

// DeleteByID
// delete from index by ID
func (p *ProviderElastic) DeleteByID(index, id string) error {
	_, err := p.client.Delete().
		Index(index).
		Id(id).
		Do(p.ctx)
	if err != nil {
		return errors.InternalServerError(p.ctx).
			SetDevMessage("DeleteByID.p.client.Delete: " + err.Error())
	}

	return nil
}

// DeleteByQuery
// delete from index by query
func (p *ProviderElastic) DeleteByQuery(index string, query elastic.Query) error {
	_, err := p.client.DeleteByQuery().
		Index(index).
		Query(query).
		Do(p.ctx)
	if err != nil {
		return errors.InternalServerError(p.ctx).
			SetDevMessage("DeleteByQuery.p.client.DeleteByQuery: " + err.Error())
	}

	return nil
}
