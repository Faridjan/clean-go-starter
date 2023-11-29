package elasticpkg

import (
	"context"
)

// NewElasticProvider
// elastic provider constructor
func NewElasticProvider(ctx context.Context, nodesURLs []string, setSniff bool, username, password string) (ProviderElasticInterface, error) {
	provider := new(ProviderElastic)

	client, err := initConnect(ctx, nodesURLs, setSniff, username, password)
	if err != nil {
		return nil, err
	}

	provider.client = client
	provider.ctx = ctx

	return provider, nil
}
