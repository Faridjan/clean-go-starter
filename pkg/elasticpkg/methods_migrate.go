package elasticpkg

import "git.centerhome.kz/bcc/backend/toolchain/common-libs/errors"

// Migrate
// create index with mapping
func (p *ProviderElastic) Migrate(index, mapping string) error {
	exist, err := p.client.IndexExists(index).Do(p.ctx)
	if err != nil {
		return errors.InternalServerError(p.ctx).
			SetDevMessage("Migrate.p.client.IndexExists: " + err.Error())
	}

	if !exist {
		_, err = p.client.CreateIndex(index).Body(mapping).Do(p.ctx)
		if err != nil {
			return errors.InternalServerError(p.ctx).
				SetDevMessage("Migrate.p.client.CreateIndex: " + err.Error())
		}
	}

	return nil
}
