package elasticpkg

import (
	"context"

	"github.com/olivere/elastic/v7"
)

// ProviderElastic
// elastic provider entity
type ProviderElastic struct {
	ctx    context.Context
	client *elastic.Client
}
