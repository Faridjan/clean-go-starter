package elasticpkg

import "git.centerhome.kz/bcc/backend/toolchain/common-libs/errors"

// StoreByID
// create/update by ID
func (p *ProviderElastic) StoreByID(index, id string, payload interface{}) error {
	_, err := p.client.Index().
		Index(index).
		Id(id).
		BodyJson(payload).
		Refresh("true").
		Do(p.ctx)
	if err != nil {
		return errors.InternalServerError(p.ctx).
			SetDevMessage("StoreByID.p.client.Index: " + err.Error())
	}

	return nil
}
